package handler

import (
	"context"
	"html/template"
	"log"

	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/google/uuid"

	"github.com/jfyne/live"
)

type (
	AboutInstance struct {
		*CommonInstance
		Summary                 *domain.SystemSymmary
		TrueChallengeExample    *domain.Challenge
		FalseChallengeExample   *domain.Challenge
		OngoingChallengeExample *domain.Challenge
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

		trueExample, err := h.app.GetRandomTrueChallenges(ctx, 1)
		if err != nil {
			return instance.withError(err), nil
		}
		if len(trueExample) == 1 {
			instance.TrueChallengeExample = trueExample[0]
		}

		falseExample, err := h.app.GetRandomFalseChallenges(ctx, 1)
		if err != nil {
			return instance.withError(err), nil
		}
		if len(falseExample) == 1 {
			instance.FalseChallengeExample = falseExample[0]
		}

		ongoingExample, err := h.app.GetRandomOngoingChallenges(ctx, uuid.Nil, 1)
		if err != nil {
			return instance.withError(err), nil
		}
		if len(ongoingExample) == 1 {
			instance.OngoingChallengeExample = ongoingExample[0]
		}

		return instance, nil
	})

	return lvh
}
