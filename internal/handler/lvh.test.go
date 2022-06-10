package handler

import (
	"context"
	"errors"
	"html/template"
	"log"

	"github.com/DanielTitkov/lentils/internal/domain"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

const (
	// events
	eventBeginTest = "begin-test"
	// params
	paramTestCode = "testCode"
	// params values
)

type (
	TestInstance struct {
		*CommonInstance
		Test     *domain.Test
		TestStep string
		Page     int
	}
)

func (ins *TestInstance) withError(err error) *TestInstance {
	ins.Error = err
	return ins
}

func (h *Handler) NewTestInstance(s live.Socket) *TestInstance {
	m, ok := s.Assigns().(*TestInstance)
	if !ok {
		return &TestInstance{
			CommonInstance: h.NewCommon(s, viewTest),
			TestStep:       domain.TestStepIntro,
		}
	}

	return m
}

func (h *Handler) Test() live.Handler {
	t, err := template.ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.test.html",
	)
	if err != nil {
		log.Fatal(err)
	}

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
		// SAFE TO COPY END
	}
	// COMMON BLOCK END

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		r := live.Request(ctx)
		testCode, ok := mux.Vars(r)[paramTestCode]
		if !ok {
			return nil, errors.New("test code is required")
		}

		instance := h.NewTestInstance(s)
		instance.fromContext(ctx)

		test, err := h.app.PrepareTest(ctx, testCode, "en", nil) // FIXME
		if err != nil {
			return instance.withError(err), nil
		}
		instance.Test = test

		return instance, nil
	})

	return lvh
}
