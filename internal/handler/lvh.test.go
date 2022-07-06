package handler

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/DanielTitkov/orrery/internal/util"

	"github.com/DanielTitkov/orrery/internal/domain"
	"github.com/bradfitz/iter"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

const (
	// events
	eventBeginTest             = "begin-test"
	eventEndTest               = "end-test"
	eventNextPage              = "next-page"
	eventPrevPage              = "prev-page"
	eventResponseUpdate        = "response-update"
	eventSetLocale             = "set-locale"
	eventToggleAutoNext        = "toggle-auto-next"
	eventToggleShowDetails     = "toggle-show-details"
	eventToggleShowInstruction = "toggle-show-instruction"
	// params
	paramTestCode      = "testCode"
	paramTestItem      = "item"
	paramTestItemValue = "val"
	paramTestLocale    = "locale"
	// params values
)

var testFuncMap = template.FuncMap{
	"N":     iter.N,
	"Plus1": func(i int) int { return i + 1 },
	"Sum": func(data ...float64) float64 {
		var res float64
		for _, n := range data {
			res += n
		}
		return res
	},
	"Sub": func(f1, f2 float64) float64 {
		return f1 - f2
	},
	"Mean": func(data ...float64) float64 {
		if len(data) == 0 {
			return 0
		}
		var sum float64
		for _, n := range data {
			sum += n
		}
		return sum / float64(len(data))
	},
	"LocaleIcon": domain.LocaleIcon,
	"Perc": func(min, max, v float64) float64 {
		if max == min {
			return 0
		}
		return (v - min) / (max - min)
	},
}

type (
	TestInstance struct {
		*CommonInstance
		Test             *domain.Test
		CurrentQuestions []*domain.Question
		TestStep         string
		Page             int
		// to have constants in templates
		IntroStatus     string
		QuestionsStatus string
		FinishStatus    string
		ResultStatus    string
		AutoNext        bool
		ShowDetails     bool
		ShowInstruction bool
	}
)

func (ins *TestInstance) withError(err error) *TestInstance {
	ins.Error = err
	return ins
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
			TestStep:        domain.TestStepIntro,
			IntroStatus:     domain.TestStepIntro,
			QuestionsStatus: domain.TestStepQuestions,
			FinishStatus:    domain.TestStepFinish,
			ResultStatus:    domain.TestStepResult,
			AutoNext:        false,
			ShowDetails:     false,
			ShowInstruction: false,
		}
	}

	return m
}

func (h *Handler) Test() live.Handler {
	t := template.Must(template.New("base.layout.html").Funcs(testFuncMap).ParseFiles(
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

		if instance.User == nil {
			instance.User, err = h.app.GetUserBySession(r, s.Session())
			if err != nil || instance.User == nil {
				return nil, fmt.Errorf("user is nil, sid: %s, error: %s", s.Session(), err)
			}
		}

		instance.Test, err = h.app.PrepareTest(ctx, testCode, instance.Locale(), &domain.PrepareTestArgs{
			UserID:  instance.User.ID,
			Session: instance.Session,
		})
		if err != nil {
			return instance.withError(err), nil
		}

		return instance, nil
	})

	lvh.HandleEvent(eventSetLocale, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		r := live.Request(ctx)
		testCode, ok := mux.Vars(r)[paramTestCode]
		if !ok {
			return nil, errors.New("test code is required")
		}
		instance := h.NewTestInstance(s)

		instance.SetLocale(p.String(paramTestLocale))
		instance.Test, err = h.app.PrepareTest(ctx, testCode, instance.Locale(), &domain.PrepareTestArgs{
			UserID:  instance.UserID,
			Session: instance.Session,
		})
		if err != nil {
			return instance.withError(err), nil
		}

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
		instance.Test, err = h.app.PrepareTestResult(ctx, instance.Test, instance.Locale())
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

	lvh.HandleError(func(ctx context.Context, err error) {
		w := live.Writer(ctx)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("this is a bad request: " + err.Error()))
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

	return lvh
}
