package handler

import (
	"context"
	"errors"
	"html/template"
	"log"
	"math"

	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/google/uuid"

	"github.com/jfyne/live"
)

const (
	// events
	eventUserProfileUpdatePage    = "challenge-list-update-page"
	eventProfileSelectOngoing     = "select-ongoing"
	eventProfileSelectFinished    = "select-finished"
	eventProfileSelectMine        = "select-mine"
	eventProfileCreateNew         = "create-new"
	eventProfileCreateNewSubmit   = "create-new-submit"
	eventProfileCreateNewValidate = "create-new-validate"
	// params
	paramUserProfilePage             = "page"
	paramProfileCreateNewContent     = "content"
	paramProfileCreateNewDescription = "description"
	paramProfileCreateNewStartTime   = "start-time"
	paramProfileCreateNewEndTime     = "end-time"
	// params value
)

type (
	ProfileInstance struct {
		*CommonInstance
		Summary             *domain.SystemSymmary
		UserSummary         *domain.UserSummary
		Challenges          []*domain.Challenge
		ChallengeCount      int
		FilterArgs          domain.FilterChallengesArgs
		Page                int
		MaxPage             int
		ShowMine            bool
		CreateChallengeForm bool
		CreateArgs          domain.CreateChallengeArgs
		CreatedChallenge    *domain.Challenge
		FormError           error
		TimeLayout          string
	}
)

func (ins *ProfileInstance) withError(err error) *ProfileInstance {
	ins.Error = err
	return ins
}

func (ins *ProfileInstance) NextPage() int {
	if ins.Page >= ins.MaxPage {
		return ins.Page
	}
	return ins.Page + 1
}

func (ins *ProfileInstance) PrevPage() int {
	if ins.Page == 1 {
		return ins.Page
	}
	return ins.Page - 1
}

func (ins *ProfileInstance) updateChallenges(ctx context.Context, h *Handler) error {
	ins.FilterArgs.Limit = h.app.Cfg.App.DefaultChallengePageLimit
	ins.FilterArgs.UserID = ins.UserID
	if ins.Page > 0 {
		ins.FilterArgs.Offset = (ins.Page - 1) * h.app.Cfg.App.DefaultChallengePageLimit
	} else {
		ins.FilterArgs.Offset = 0
	}

	var chs []*domain.Challenge
	var count int
	var err error
	if ins.FilterArgs.AuthorID == uuid.Nil {
		chs, count, err = h.app.FilterUserChallenges(ctx, &ins.FilterArgs)
		if err != nil {
			return err
		}
	} else {
		chs, count, err = h.app.GetChallengesByAuthor(
			ctx, ins.FilterArgs.AuthorID,
			ins.FilterArgs.Limit, ins.FilterArgs.Offset,
		)
		if err != nil {
			return err
		}
	}
	ins.Challenges = chs
	ins.ChallengeCount = count
	ins.MaxPage = int(math.Ceil(float64(count) / float64(h.app.Cfg.App.DefaultChallengePageLimit)))

	return nil
}

func (h *Handler) NewProfileInstance(s live.Socket) *ProfileInstance {
	m, ok := s.Assigns().(*ProfileInstance)
	if !ok {
		return &ProfileInstance{
			CommonInstance: h.NewCommon(s, viewProfile),
			Page:           1,
			FilterArgs: domain.FilterChallengesArgs{
				Ongoing:  false,
				Finished: false,
			},
			ShowMine:            false,
			CreateChallengeForm: true,
			FormError:           errors.New("provide challenge details"),
			CreatedChallenge:    nil,
			TimeLayout:          h.app.Cfg.App.DefaultTimeLayout,
		}
	}

	return m
}

func (h *Handler) Profile() live.Handler {
	t, err := template.ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.profile.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewProfileInstance // NB: make sure constructor is correct
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
		instance := h.NewProfileInstance(s)
		instance.fromContext(ctx)

		if instance.User == nil || instance.UserID == uuid.Nil {
			s.Redirect(h.url404())
			return nil, nil
		}

		userSummary, err := h.app.GetUserSummary(ctx, instance.UserID)
		if err != nil {
			return instance.withError(err), nil
		}
		instance.UserSummary = userSummary

		if err := instance.updateChallenges(ctx, h); err != nil {
			return instance.withError(err), nil
		}

		return instance, nil
	})

	lvh.HandleEvent(eventUserProfileUpdatePage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		page := p.Int(paramUserProfilePage)
		instance := h.NewProfileInstance(s)
		instance.Page = page
		err := instance.updateChallenges(ctx, h)
		if err != nil {
			return instance.withError(err), nil
		}
		return instance, nil
	})

	lvh.HandleEvent(eventProfileSelectOngoing, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewProfileInstance(s)
		if instance.FilterArgs.Ongoing {
			return instance, nil
		}

		instance.Page = 1
		instance.ShowMine = false
		instance.FilterArgs.AuthorID = uuid.Nil
		instance.FilterArgs.Ongoing = true
		instance.FilterArgs.Finished = false
		instance.CreateChallengeForm = false
		instance.CreatedChallenge = nil
		instance.FormError = nil

		err := instance.updateChallenges(ctx, h)
		return instance.withError(err), nil
	})

	lvh.HandleEvent(eventProfileSelectMine, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewProfileInstance(s)
		if instance.ShowMine {
			return instance, nil
		}

		instance.Page = 1
		instance.FilterArgs.AuthorID = instance.UserID
		instance.ShowMine = true
		instance.FilterArgs.Ongoing = false
		instance.FilterArgs.Finished = false
		instance.CreateChallengeForm = false
		instance.CreatedChallenge = nil
		instance.FormError = nil

		err := instance.updateChallenges(ctx, h)
		return instance.withError(err), nil
	})

	lvh.HandleEvent(eventProfileSelectFinished, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewProfileInstance(s)
		if instance.FilterArgs.Finished {
			return instance, nil
		}

		instance.Page = 1
		instance.ShowMine = false
		instance.FilterArgs.AuthorID = uuid.Nil
		instance.FilterArgs.Finished = true
		instance.FilterArgs.Ongoing = false
		instance.CreateChallengeForm = false
		instance.CreatedChallenge = nil
		instance.FormError = nil

		err := instance.updateChallenges(ctx, h)
		return instance.withError(err), nil
	})

	lvh.HandleEvent(eventProfileCreateNew, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewProfileInstance(s)

		instance.ShowMine = false
		instance.FilterArgs.AuthorID = uuid.Nil
		instance.FilterArgs.Finished = false
		instance.FilterArgs.Ongoing = false
		instance.CreateChallengeForm = true
		instance.FormError = errors.New("provide challenge details")

		return instance, nil
	})

	lvh.HandleEvent(eventProfileCreateNewSubmit, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewProfileInstance(s)

		instance.CreateArgs = profileCreateArgsFromParams(p, h.app.Cfg.App.DefaultTimeLayout, instance.UserID)
		instance.FormError = instance.CreateArgs.Validate()
		if instance.FormError != nil {
			return instance, nil
		}

		challenge, err := h.app.CreateChallengeFromArgs(ctx, instance.CreateArgs, true)
		if err != nil {
			return instance.withError(err), nil
		}

		instance.CreatedChallenge = challenge
		instance.CreateArgs = domain.CreateChallengeArgs{}

		return instance, nil
	})

	lvh.HandleEvent(eventProfileCreateNewValidate, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewProfileInstance(s)

		instance.CreateArgs = profileCreateArgsFromParams(p, h.app.Cfg.App.DefaultTimeLayout, instance.UserID)
		instance.FormError = instance.CreateArgs.Validate()

		return instance, nil
	})

	return lvh
}

func profileCreateArgsFromParams(p live.Params, layout string, userID uuid.UUID) domain.CreateChallengeArgs {
	return domain.CreateChallengeArgs{
		Type:        domain.ChallengeTypeBool,
		Outcome:     nil,
		Content:     p.String(paramProfileCreateNewContent),
		Description: p.String(paramProfileCreateNewDescription),
		StartTime:   p.String(paramProfileCreateNewStartTime),
		EndTime:     p.String(paramProfileCreateNewEndTime),
		Published:   false,
		TimeLayout:  layout,
		AuthorID:    userID,
	}
}
