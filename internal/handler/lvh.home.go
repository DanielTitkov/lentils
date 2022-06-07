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
		Summary                       *domain.SystemSymmary
		RandomFinishedChallenges      []*domain.Challenge
		RandomFinishedChallengesCount int
		RandomOngoingChallenges       []*domain.Challenge
		RandomOngoingChallengesCount  int
		RandomPendingChallenges       []*domain.Challenge
		RandomPendingChallengesCount  int
		ClosingChallenges             []*domain.Challenge
		ClosingChallengesCount        int
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

		// TODO: make single func for all challenge types
		// random finished
		randomFinishedChallenges, err := h.app.GetRandomFinishedChallenges(ctx)
		if err != nil {
			instance.Error = err
			return instance, nil
		}
		instance.RandomFinishedChallenges = randomFinishedChallenges
		instance.RandomFinishedChallengesCount = len(randomFinishedChallenges)

		// random ongoing
		randomOngoingChallenges, err := h.app.GetRandomOngoingChallenges(ctx, instance.UserID, h.app.Cfg.App.HomeChallengePageLimit)
		if err != nil {
			instance.Error = err
			return instance, nil
		}
		instance.RandomOngoingChallenges = randomOngoingChallenges
		instance.RandomOngoingChallengesCount = len(randomOngoingChallenges)

		// random pending
		randomPendingChallenges, err := h.app.GetRandomPendingChallenges(ctx)
		if err != nil {
			instance.Error = err
			return instance, nil
		}
		instance.RandomPendingChallenges = randomPendingChallenges
		instance.RandomPendingChallengesCount = len(randomPendingChallenges)

		// closing
		closingChallenges, err := h.app.GetClosingChallenges(ctx)
		if err != nil {
			instance.Error = err
			return instance, nil
		}
		instance.ClosingChallenges = closingChallenges
		instance.ClosingChallengesCount = len(closingChallenges)

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
