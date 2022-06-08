package handler

import (
	"context"
	"html/template"
	"log"

	"github.com/google/uuid"

	"github.com/jfyne/live"
)

const (
	// events
	eventAdminUpdatePage        = "challenge-list-update-page"
	eventAdminSelectPending     = "select-pending"
	eventAdminSelectUnpublished = "select-unpublished"
	eventAdminCreateNew         = "create-new"
	eventAdminCreateNewSubmit   = "create-new-submit"
	eventAdminCreateNewValidate = "create-new-validate"
	// params
	paramAdminPage                 = "page"
	paramAdminCreateNewContent     = "content"
	paramAdminCreateNewDescription = "description"
	paramAdminCreateNewStartTime   = "start-time"
	paramAdminCreateNewEndTime     = "end-time"
	paramAdminCreateNewPublished   = "published"
)

type (
	AdminInstance struct {
		*CommonInstance
		Page       int
		MaxPage    int
		TimeLayout string
	}
)

func (ins *AdminInstance) withError(err error) *AdminInstance {
	ins.Error = err
	return ins
}

func (ins *AdminInstance) NextPage() int {
	if ins.Page >= ins.MaxPage {
		return ins.Page
	}
	return ins.Page + 1
}

func (ins *AdminInstance) PrevPage() int {
	if ins.Page == 1 {
		return ins.Page
	}
	return ins.Page - 1
}

func (h *Handler) NewAdminInstance(s live.Socket) *AdminInstance {
	m, ok := s.Assigns().(*AdminInstance)
	if !ok {
		return &AdminInstance{
			CommonInstance: h.NewCommon(s, viewAdmin),
			Page:           1,
			TimeLayout:     h.app.Cfg.App.DefaultTimeLayout,
		}
	}

	return m
}

func (h *Handler) Admin() live.Handler {
	t, err := template.ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.admin.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewAdminInstance // NB: make sure constructor is correct
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
		instance := h.NewAdminInstance(s)
		instance.fromContext(ctx)

		if instance.User == nil || instance.UserID == uuid.Nil || !instance.User.Admin {
			s.Redirect(h.url404())
			return nil, nil
		}

		return instance, nil
	})

	lvh.HandleEvent(eventAdminUpdatePage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		page := p.Int(paramAdminPage)
		instance := h.NewAdminInstance(s)
		instance.Page = page

		return instance, nil
	})

	return lvh
}