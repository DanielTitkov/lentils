package handler

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/bradfitz/iter"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

const (
// events
// params
// params values
)

var resultFuncMap = template.FuncMap{
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
}

type (
	ResultInstance struct {
		*CommonInstance
		Test     *domain.Test
		Locale   string
		TestStep string
		// to have constants in templates
		FinishStatus    string
		ResultStatus    string
		ShowDetails     bool
		ShowInstruction bool
	}
)

func (ins *ResultInstance) withError(err error) *ResultInstance {
	ins.Error = err
	return ins
}

func (h *Handler) NewResultInstance(s live.Socket) *ResultInstance {
	m, ok := s.Assigns().(*ResultInstance)
	if !ok {
		return &ResultInstance{
			CommonInstance:  h.NewCommon(s, viewTest),
			TestStep:        domain.TestStepIntro,
			FinishStatus:    domain.TestStepFinish,
			ResultStatus:    domain.TestStepResult,
			Locale:          domain.LocaleEn,
			ShowDetails:     false,
			ShowInstruction: false,
		}
	}

	return m
}

func (h *Handler) Result() live.Handler {
	t := template.Must(template.New("base.layout.html").Funcs(testFuncMap).ParseFiles(
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
		// SAFE TO COPY END
	}
	// COMMON BLOCK END

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (i interface{}, err error) {
		r := live.Request(ctx)
		testCode, ok := mux.Vars(r)[paramTestCode]
		if !ok {
			return nil, errors.New("test code is required")
		}

		instance := h.NewResultInstance(s)
		instance.fromContext(ctx)

		if instance.User == nil {
			return instance.withError(errors.New("user is nil")), nil
		}

		instance.Test, err = h.app.PrepareTest(ctx, testCode, instance.Locale, &domain.PrepareTestArgs{
			UserID:  instance.UserID,
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
		instance := h.NewResultInstance(s)

		locale := p.String(paramTestLocale)
		if !domain.IsValidLocale(locale) {
			return instance.withError(fmt.Errorf("unknown locale: %s", locale)), nil
		}

		instance.Locale = locale
		instance.Test, err = h.app.PrepareTest(ctx, testCode, instance.Locale, &domain.PrepareTestArgs{
			UserID:  instance.UserID,
			Session: instance.Session,
		})
		if err != nil {
			return instance.withError(err), nil
		}

		return instance, nil
	})

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
		w.Write([]byte("this is a bad request: " + err.Error()))
	})

	return lvh
}
