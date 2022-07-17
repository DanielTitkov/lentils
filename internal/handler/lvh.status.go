package handler

import (
	"context"
	"html/template"

	"github.com/DanielTitkov/orrery/internal/domain"

	"github.com/jfyne/live"
)

type (
	StatusInstance struct {
		*CommonInstance
		Errors  []error
		Events  []domain.Event
		Version string
		Summary *domain.SystemSummary
	}
)

func (ins *StatusInstance) withError(err error) *StatusInstance {
	ins.Error = err
	return ins
}

// must be present in all instances
func (ins *StatusInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	return nil
}

func (h *Handler) NewStatusInstance(s live.Socket) *StatusInstance {
	m, ok := s.Assigns().(*StatusInstance)
	if !ok {
		return &StatusInstance{
			CommonInstance: h.NewCommon(s, viewStatus),
			Errors:         nil,
			Events:         nil,
		}
	}

	return m
}

func (h *Handler) Status() live.Handler {
	t := template.Must(template.New("base.layout.html").Funcs(funcMap).ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.status.html",
		h.t+"part.system_summary.html",
	))

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewStatusInstance // NB: make sure constructor is correct
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
		instance := h.NewStatusInstance(s)
		instance.fromContext(ctx)

		instance.Errors = h.app.Errors
		instance.Events = h.app.Events
		instance.Version = h.app.Cfg.App.Version

		// summary
		instance.Summary, err = h.app.GetSystemSummary(ctx)
		if err != nil {
			return instance.withError(err), nil
		}

		return instance, nil
	})

	return lvh
}
