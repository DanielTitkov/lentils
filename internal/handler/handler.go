package handler

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/jfyne/live"

	"github.com/google/uuid"

	"github.com/DanielTitkov/orrery/internal/app"
	"github.com/DanielTitkov/orrery/internal/domain"
	"github.com/DanielTitkov/orrery/internal/logger"
)

const (
	// views
	view404     = "404"
	viewAbout   = "about"
	viewAdmin   = "admin"
	viewTest    = "test"
	viewResult  = "result"
	viewHome    = "home"
	viewProfile = "profile"
	viewPrivacy = "privacy"
	viewTerms   = "terms"
	viewStatus  = "status"
	// events (common)
	eventCloseAuthModals = "close-auth-modals"
	eventOpenLogoutModal = "open-logout-modal"
	eventOpenLoginModal  = "open-login-modal"
	eventCloseError      = "close-error-notification"
	eventCloseMessage    = "close-message-notification"
	eventToggleDark      = "toggle-dark"
	// context
	userCtxKeyValue   = "user"
	localeCtxKeyValue = "locale"
)

type (
	Handler struct {
		app *app.App
		log *logger.Logger
		t   string // template path
		ui  map[string]*UITranslation
	}

	CommonInstance struct {
		Env             string
		Session         string
		Error           error
		Message         *string
		User            *domain.User
		UserID          uuid.UUID
		ShowLoginModal  bool
		ShowLogoutModal bool
		CurrentView     string
		Version         string
		Dark            bool
		UI              *UITranslation
		ui              map[string]*UITranslation
		locale          string
	}

	contextKey struct {
		name string
	}
)

var userCtxKey = &contextKey{userCtxKeyValue}
var localeCtxKey = &contextKey{localeCtxKeyValue}

func NewHandler(
	app *app.App,
	logger *logger.Logger,
	t string,
) *Handler {
	return &Handler{
		app: app,
		log: logger,
		t:   t,
		ui:  initTraslationMap(),
	}
}

func initTraslationMap() map[string]*UITranslation {
	m := make(map[string]*UITranslation)
	for _, l := range domain.Locales() {
		m[l] = newUITranslation(l)
	}

	return m
}

func (h *Handler) NewCommon(s live.Socket, currentView string) *CommonInstance {
	c := &CommonInstance{
		Env:             h.app.Cfg.Env,
		Session:         fmt.Sprint(s.Session()),
		Error:           nil,
		Message:         nil,
		ShowLoginModal:  false,
		ShowLogoutModal: false,
		CurrentView:     currentView,
		Version:         h.app.Cfg.App.Version,
		ui:              h.ui,
		locale:          domain.DefaultLocale(), // it's private because changing requesed additional logic
	}

	c.SetLocale(c.locale)
	return c
}

func (h *Handler) url404() *url.URL {
	u, _ := url.Parse("/404")
	return u
}

func (h *Handler) NotFoundRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, h.url404().String(), http.StatusTemporaryRedirect)
}

func (c *CommonInstance) getTranslation(locale string) *UITranslation {
	trans, ok := c.ui[locale]
	if !ok {
		return c.ui[domain.DefaultLocale()]
	}

	return trans
}

func (c *CommonInstance) CloseAuthModals() {
	c.ShowLoginModal = false
	c.ShowLogoutModal = false
}

func (c *CommonInstance) OpenLoginModal() {
	c.ShowLoginModal = true
}

func (c *CommonInstance) OpenLogoutModal() {
	c.ShowLogoutModal = true
}

func (c *CommonInstance) CloseError() {
	c.Error = nil
}

func (c *CommonInstance) CloseMessage() {
	c.Message = nil
}

func (c *CommonInstance) SetLocale(l string) {
	if !domain.IsValidLocale(l) {
		l = domain.DefaultLocale()
	}
	c.locale = l
	c.UI = c.getTranslation(l)
}

func (c *CommonInstance) Locale() string {
	return c.locale
}

func UserFromCtx(ctx context.Context) (*domain.User, uuid.UUID) {
	user, ok := ctx.Value(userCtxKey).(*domain.User)
	if !ok {
		return nil, uuid.Nil
	}
	if user == nil {
		return nil, uuid.Nil
	}
	return user, user.ID
}

func localeFromCtx(ctx context.Context) string {
	locale, ok := ctx.Value(localeCtxKey).(string)
	if !ok {
		return domain.LocaleEn
	}
	return locale
}

func (c *CommonInstance) fromContext(ctx context.Context) {
	c.locale = localeFromCtx(ctx)
	user, userID := UserFromCtx(ctx)
	c.User = user
	c.UserID = userID
}
