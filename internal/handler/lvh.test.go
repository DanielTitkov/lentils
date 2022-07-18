package handler

import (
	"context"
	"errors"
	"fmt"
	"html/template"

	"github.com/tinygodsdev/orrery/internal/util"

	"github.com/tinygodsdev/orrery/internal/domain"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

const (
	// events
	eventBeginTest              = "begin-test"
	eventEndTest                = "end-test"
	eventNextPage               = "next-page"
	eventPrevPage               = "prev-page"
	eventResponseUpdate         = "response-update"
	eventSetLocale              = "set-locale"
	eventToggleAutoNext         = "toggle-auto-next"
	eventToggleShowDetails      = "toggle-show-details"
	eventToggleShowInstruction  = "toggle-show-instruction"
	eventTestToggleShowAdvanced = "toggle-show-advanced"
	eventTestMarkUpdate         = "test-mark-update"
	eventTestSetMethod          = "set-method"
	// params
	paramTestCode      = "testCode"
	paramTestItem      = "item"
	paramTestItemValue = "val"
	paramTestLocale    = "locale"
	paramTestMarkValue = "val"
	paramTestMethod    = "method"
	// params values
)

type (
	TestInstance struct {
		*CommonInstance
		*Constants
		Test                 *domain.Test
		CurrentQuestions     []*domain.Question
		TestCode             string
		TestStep             string
		Page                 int
		AutoNext             bool
		ShowDetails          bool
		ShowInstruction      bool
		ShowAdvancedSettings bool
		OverrideMethod       string
	}
)

// must be present in all instances
func (ins *TestInstance) withError(err error) *TestInstance {
	ins.Error = err
	return ins
}

// must be present in all instances
func (ins *TestInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	r := live.Request(ctx)
	var err error

	if ins.User == nil {
		ins.User, err = h.app.GetUserBySession(r, s.Session())
		if err != nil || ins.User == nil {
			return fmt.Errorf("user is nil, sid: %s, error: %s", s.Session(), err)
		}
	}

	ins.Test, err = h.app.PrepareTest(ctx, ins.TestCode, ins.Locale(), &domain.PrepareTestArgs{
		UserID:  ins.User.ID,
		Session: ins.Session,
	})
	if err != nil {
		ins.withError(err)
	}
	return nil
}

func (ins *TestInstance) nextPage() int {
	if ins.Page >= ins.Test.PageCount() {
		return ins.Page
	}
	return ins.Page + 1
}

func (ins *TestInstance) prevPage() int {
	if ins.Page == 1 {
		return ins.Page
	}
	return ins.Page - 1
}

func (h *Handler) NewTestInstance(s live.Socket) *TestInstance {
	m, ok := s.Assigns().(*TestInstance)
	if !ok {
		return &TestInstance{
			CommonInstance:  h.NewCommon(s, viewTest),
			Constants:       h.NewConstants(),
			TestStep:        domain.TestStepIntro,
			AutoNext:        false,
			ShowDetails:     false,
			ShowInstruction: false,
			OverrideMethod:  "",
		}
	}

	return m
}

func (h *Handler) Test() live.Handler {
	t := template.Must(template.New("base.layout.html").Funcs(funcMap).ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.test.html",
	))

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewTestInstance // NB: make sure constructor is correct
		// SAFE TO COPY
		lvh.HandleEvent(eventCloseAuthModals, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.CloseAuthModals()
			return instance, nil
		})

		lvh.HandleEvent(eventOpenLogoutModal, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.OpenLogoutModal()
			return instance, nil
		})

		lvh.HandleEvent(eventOpenLoginModal, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.OpenLoginModal()
			return instance, nil
		})

		lvh.HandleEvent(eventCloseError, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.CloseError()
			return instance, nil
		})

		lvh.HandleEvent(eventCloseMessage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.CloseMessage()
			return instance, nil
		})

		lvh.HandleEvent(eventToggleDark, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
			instance := constructor(s)
			if instance.User != nil {
				instance.User.UseDarkTheme = !instance.User.UseDarkTheme
			}
			instance.User, err = h.app.UpdateUser(ctx, instance.User)
			if err != nil {
				return instance.withError(err), nil
			}
			return instance, nil
		})

		// update locale logic
		lvh.HandleParams(func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.SetLocale(p.String(paramLocale))
			err := instance.updateForLocale(ctx, s, h)
			if err != nil {
				return nil, err
			}
			return instance, nil
		})

		lvh.HandleError(func(ctx context.Context, err error) {
			h.HandleError(ctx, err)
		})
		// SAFE TO COPY END
	}
	// COMMON BLOCK END

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (i interface{}, err error) {
		r := live.Request(ctx)
		testCode, ok := mux.Vars(r)[paramTestCode]
		if !ok {
			return nil, errors.New("test code is required")
		}

		instance := h.NewTestInstance(s)
		instance.fromContext(ctx)
		instance.TestCode = testCode
		instance.updateForLocale(ctx, s, h)

		return instance, nil
	})

	lvh.HandleEvent(eventBeginTest, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewTestInstance(s)

		var err error
		instance.Test, err = h.app.BeginTest(ctx, instance.Test)
		if err != nil {
			return instance.withError(err), nil
		}
		instance.Page = 1
		instance.CurrentQuestions = instance.Test.QuestionsForPage(instance.Page)
		instance.ShowDetails = false

		return instance, nil
	})

	lvh.HandleEvent(eventEndTest, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewTestInstance(s)

		var err error
		// update test status
		instance.Test, err = h.app.EndTest(ctx, instance.Test)
		if err != nil {
			return instance.withError(err), nil
		}

		// load all test data from the db and calculate result
		instance.Test, err = h.app.PrepareTestResult(ctx, instance.Test, instance.Locale(), instance.OverrideMethod)
		if err != nil {
			return instance.withError(err), nil
		}

		return instance, nil
	})

	lvh.HandleEvent(eventToggleAutoNext, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewTestInstance(s)
		instance.AutoNext = !instance.AutoNext
		return instance, nil
	})

	lvh.HandleEvent(eventToggleShowDetails, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewTestInstance(s)
		instance.ShowDetails = !instance.ShowDetails
		return instance, nil
	})

	lvh.HandleEvent(eventToggleShowInstruction, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewTestInstance(s)
		instance.ShowInstruction = !instance.ShowInstruction
		return instance, nil
	})

	lvh.HandleEvent(eventTestToggleShowAdvanced, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewTestInstance(s)
		instance.ShowAdvancedSettings = !instance.ShowAdvancedSettings
		return instance, nil
	})

	lvh.HandleEvent(eventNextPage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewTestInstance(s)

		instance.Page = instance.nextPage()
		instance.CurrentQuestions = instance.Test.QuestionsForPage(instance.Page)

		return instance, nil
	})

	lvh.HandleSelf(eventNextPage, func(ctx context.Context, s live.Socket, data interface{}) (interface{}, error) {
		instance := h.NewTestInstance(s)

		instance.Page = instance.nextPage()
		instance.CurrentQuestions = instance.Test.QuestionsForPage(instance.Page)

		return instance, nil
	})

	lvh.HandleEvent(eventPrevPage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewTestInstance(s)

		instance.Page = instance.prevPage()
		instance.CurrentQuestions = instance.Test.QuestionsForPage(instance.Page)

		return instance, nil
	})

	lvh.HandleEvent(eventResponseUpdate, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewTestInstance(s)

		itemCode := p.String(paramTestItem)
		if itemCode == "" {
			return instance.withError(errors.New("item code is empty")), nil
		}

		value := p.Int(paramTestItemValue)

		meta := util.NewMeta()
		meta["session"] = instance.Session
		item := instance.Test.GetItem(itemCode)
		err := item.AddResponse(instance.Test.Take.ID, value, meta)
		if err != nil {
			return instance.withError(err), nil
		}

		// save response to db
		instance.Test.Take.Page = instance.Page
		instance.Test.Take, item.Response, err = h.app.AddResponse(ctx, instance.Test.Take, item)
		if err != nil {
			return instance.withError(err), nil
		}

		if instance.AutoNext {
			if !instance.Test.IsPageNotDone(instance.Page) {
				err := s.Self(ctx, eventNextPage, nil)
				if err != nil {
					return instance.withError(err), nil
				}
			}
		}

		return instance, nil
	})

	lvh.HandleEvent(eventTestMarkUpdate, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewTestInstance(s)
		mark := p.Int(paramTestMarkValue)
		if mark >= domain.TakeMinMark && mark <= domain.TakeMaxMark {
			instance.Test.Take.Mark = &mark
			err := h.app.UpdateTakeMark(ctx, instance.Test.Take.ID, mark)
			if err != nil {
				return instance.withError(err), nil
			}
		}
		return instance, nil
	})

	lvh.HandleEvent(eventTestSetMethod, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewTestInstance(s)
		instance.OverrideMethod = p.String(paramTestMethod)
		// load all test data from the db and calculate result
		instance.Test, err = h.app.PrepareTestResult(ctx, instance.Test, instance.Locale(), instance.OverrideMethod)
		if err != nil {
			return instance.withError(err), nil
		}
		return instance, nil
	})

	return lvh
}
