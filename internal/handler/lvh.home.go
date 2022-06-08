package handler

import (
	"context"
	"html/template"
	"log"

	"github.com/DanielTitkov/lentils/internal/domain"

	"github.com/jfyne/live"
)

type (
	HomeInstance struct {
		*CommonInstance
		Summary *domain.SystemSymmary
	}
)

func (h *Handler) NewHomeInstance(s live.Socket) *HomeInstance {
	m, ok := s.Assigns().(*HomeInstance)
	if !ok {
		return &HomeInstance{
			CommonInstance: h.NewCommon(s, viewHome),
		}
	}

	return m
}

func (h *Handler) Home() live.Handler {
	t, err := template.ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.home.html",
		h.t+"part.system_summary.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewHomeInstance // NB: make sure constructor is correct
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

	// Set the mount function for this handler.
	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		instance := h.NewHomeInstance(s)
		instance.fromContext(ctx)

		// summary
		summary, err := h.app.GetSystemSummary(ctx)
		if err != nil {
			instance.Error = err
			return instance, nil
		}
		instance.Summary = summary

		return instance, nil
	})

	return lvh
}
