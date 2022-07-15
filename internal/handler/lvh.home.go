package handler

import (
	"context"
	"html/template"

	"github.com/DanielTitkov/orrery/internal/domain"

	"github.com/jfyne/live"
)

const (
	// events
	eventHomeToggleTag        = "toggle-tag"
	eventHomeToggleFilterMode = "toggle-filter-mode"
	// params
	paramHomeTag = "tag"
)

type (
	HomeInstance struct {
		*CommonInstance
		Tests         []*domain.Test
		Summary       *domain.SystemSymmary
		Tags          []*domain.Tag
		ActiveTags    []*domain.Tag
		FilterModeAny bool
	}
)

func (h *Handler) NewHomeInstance(s live.Socket) *HomeInstance {
	m, ok := s.Assigns().(*HomeInstance)
	if !ok {
		return &HomeInstance{
			CommonInstance: h.NewCommon(s, viewHome),
			FilterModeAny:  true,
		}
	}

	return m
}

func (ins *HomeInstance) withError(err error) *HomeInstance {
	ins.Error = err
	return ins
}

// must be present in all instances
func (ins *HomeInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	var err error
	// tags
	ins.Tags, err = h.app.GetTags(ctx, ins.Locale())
	if err != nil {
		ins.withError(err)
	}

	// tests
	err = ins.updateTests(ctx, h)
	if err != nil {
		ins.withError(err)
	}

	return nil
}

func (ins *HomeInstance) updateTests(ctx context.Context, h *Handler) (err error) {
	// update tests
	ins.Tests, err = h.app.GetTestsForLocale(ctx, &domain.QueryTestsArgs{
		Locale:        ins.Locale(),
		Tags:          ins.ActiveTags,
		FilterModeAny: ins.FilterModeAny,
	})
	if err != nil {
		return err
	}
	return nil
}

func (ins *HomeInstance) isTagActive(code string) bool {
	for _, t := range ins.ActiveTags {
		if t.Code == code {
			return true
		}
	}

	return false
}

func (ins *HomeInstance) tagToActive(code string) {
	for _, t := range ins.Tags {
		if t.Code == code {
			ins.ActiveTags = append(ins.ActiveTags, t)
		}
	}
}

func (ins *HomeInstance) tagToInactive(code string) {
	for i, t := range ins.ActiveTags {
		if t.Code == code {
			ins.ActiveTags = append(ins.ActiveTags[:i], ins.ActiveTags[i+1:]...)
		}
	}
}

func (ins *HomeInstance) toggleTag(code string) error {
	if ins.isTagActive(code) {
		ins.tagToInactive(code)
	} else {
		ins.tagToActive(code)
	}

	return nil
}

func (h *Handler) Home() live.Handler {
	t := template.Must(template.New("base.layout.html").Funcs(funcMap).ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.home.html",
		h.t+"part.system_summary.html",
	))

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

	// Set the mount function for this handler.
	lvh.HandleMount(func(ctx context.Context, s live.Socket) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		instance.fromContext(ctx)

		// summary
		instance.Summary, err = h.app.GetSystemSummary(ctx)
		if err != nil {
			return instance.withError(err), nil
		}

		instance.updateForLocale(ctx, s, h)

		return instance.withError(err), nil
	})

	lvh.HandleEvent(eventHomeToggleFilterMode, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		instance.FilterModeAny = !instance.FilterModeAny
		return instance.withError(instance.updateTests(ctx, h)), nil
	})

	lvh.HandleEvent(eventHomeToggleTag, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		tagCode := p.String(paramHomeTag)
		instance.toggleTag(tagCode)
		// update tests
		return instance.withError(instance.updateTests(ctx, h)), nil
	})

	return lvh
}
