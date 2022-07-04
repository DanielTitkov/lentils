package handler

import (
	"context"
	"html/template"
	"log"

	"github.com/DanielTitkov/orrery/internal/domain"

	"github.com/jfyne/live"
)

type (
	AboutInstance struct {
		*CommonInstance
		Summary *domain.SystemSymmary
	}
)

func (ins *AboutInstance) withError(err error) *AboutInstance {
	ins.Error = err
	return ins
}

func (h *Handler) NewAboutInstance(s live.Socket) *AboutInstance {
	m, ok := s.Assigns().(*AboutInstance)
	if !ok {
		return &AboutInstance{
			CommonInstance: h.NewCommon(s, viewAbout),
		}
	}

	return m
}

func (h *Handler) About() live.Handler {
	t, err := template.ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.about.html",
		h.t+"part.system_summary.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewAboutInstance // NB: make sure constructor is correct
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
		instance := h.NewAboutInstance(s)
		instance.fromContext(ctx)

		summary, err := h.app.GetSystemSummary(ctx)
		if err != nil {
			return instance.withError(err), nil
		}
		instance.Summary = summary

		return instance, nil
	})

	return lvh
}
