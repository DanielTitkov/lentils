package handler

import (
	"context"
	"errors"
	"html/template"
	"net/http"

	"github.com/google/uuid"

	"github.com/DanielTitkov/orrery/internal/domain"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

const (
	// events
	eventResultSetLocale = "set-locale"
	// params
	paramTakeID = "takeID"
	// params values
)

type (
	ResultInstance struct {
		*CommonInstance
		Test     *domain.Test
		TestStep string
		// to have constants in templates
		IntroStatus     string
		QuestionsStatus string
		FinishStatus    string
		ResultStatus    string
		ShowDetails     bool
		ShowInstruction bool
	}
)

// must be present in all instances
func (ins *ResultInstance) withError(err error) *ResultInstance {
	ins.Error = err
	return ins
}

// must be present in all instances
func (ins *ResultInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	r := live.Request(ctx)
	takeIDStr, ok := mux.Vars(r)[paramTakeID]
	if !ok {
		return errors.New("take id is required")
	}

	takeID, err := uuid.Parse(takeIDStr)
	if err != nil {
		ins.withError(err)
	}

	take, err := h.app.GetTake(ctx, takeID)
	if err != nil {
		ins.withError(err)
	}

	ins.Test, err = h.app.PrepareTestResult(ctx, &domain.Test{Take: take}, ins.Locale())
	if err != nil {
		ins.withError(err)
	}
	return nil
}

func (h *Handler) NewResultInstance(s live.Socket) *ResultInstance {
	m, ok := s.Assigns().(*ResultInstance)
	if !ok {
		return &ResultInstance{
			CommonInstance:  h.NewCommon(s, viewResult),
			TestStep:        domain.TestStepIntro,
			IntroStatus:     domain.TestStepIntro,
			QuestionsStatus: domain.TestStepQuestions,
			FinishStatus:    domain.TestStepFinish,
			ResultStatus:    domain.TestStepResult,
			ShowDetails:     false,
			ShowInstruction: false,
		}
	}

	return m
}

func (h *Handler) Result() live.Handler {
	t := template.Must(template.New("base.layout.html").Funcs(funcMap).ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.test.html",
	))

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewResultInstance // NB: make sure constructor is correct
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
		// SAFE TO COPY END
	}
	// COMMON BLOCK END

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (i interface{}, err error) {
		instance := h.NewResultInstance(s)
		err = instance.updateForLocale(ctx, s, h)
		if err != nil {
			return nil, err
		}

		return instance, nil
	})

	// lvh.HandleEvent(eventResultSetLocale, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
	// 	instance := h.NewResultInstance(s)

	// 	instance.SetLocale(p.String(paramTestLocale))
	// 	instance.Test, err = h.app.PrepareTestResult(ctx, instance.Test, instance.Locale())
	// 	if err != nil {
	// 		return instance.withError(err), nil
	// 	}

	// 	return instance, nil
	// })

	lvh.HandleEvent(eventToggleShowDetails, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewResultInstance(s)
		instance.ShowDetails = !instance.ShowDetails
		return instance, nil
	})

	lvh.HandleEvent(eventToggleShowInstruction, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewResultInstance(s)
		instance.ShowInstruction = !instance.ShowInstruction
		return instance, nil
	})

	lvh.HandleError(func(ctx context.Context, err error) {
		w := live.Writer(ctx)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request: " + err.Error()))
	})

	return lvh
}
