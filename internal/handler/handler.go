package handler

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"time"

	"github.com/bradfitz/iter"
	"github.com/jfyne/live"

	"github.com/google/uuid"

	"github.com/tinygodsdev/orrery/internal/app"
	"github.com/tinygodsdev/orrery/internal/domain"
	"github.com/tinygodsdev/orrery/internal/logger"
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
	// params (common)
	paramLocale = "locale"
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
		Domain          string
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

	Constants struct {
		// to have constants in templates
		IntroStatus     string
		QuestionsStatus string
		FinishStatus    string
		ResultStatus    string
		MethodSten      string
		MethodPerc      string
		MethodMean      string
		MethodSum       string
	}

	contextKey struct {
		name string
	}
)

var userCtxKey = &contextKey{userCtxKeyValue}
var localeCtxKey = &contextKey{localeCtxKeyValue}
var funcMap = template.FuncMap{
	"N":     iter.N,
	"Plus1": func(i int) int { return i + 1 },
	"Sum": func(data ...float64) float64 {
		var res float64
		for _, n := range data {
			res += n
		}
		return res
	},
	"Sub": func(f1, f2 float64) float64 {
		return f1 - f2
	},
	"Mean": func(data ...float64) float64 {
		if len(data) == 0 {
			return 0
		}
		var sum float64
		for _, n := range data {
			sum += n
		}
		return sum / float64(len(data))
	},
	"LocaleIcon": domain.LocaleIcon,
	"Perc": func(min, max, v float64) float64 {
		if max == min {
			return 0
		}
		return (v - min) / (max - min)
	},
	"DerefInt": func(i *int) int {
		if i == nil {
			return 0
		}
		return *i
	},
	"DisplayTime": func(t time.Time) string {
		return t.Format(domain.DefaultDisplayTime)
	},
	"CodeInTags": func(code string, tags []*domain.Tag) bool {
		for _, t := range tags {
			if t.Code == code {
				return true
			}
		}
		return false
	},
	"DisplayTechTime": func(t time.Time) string {
		return t.Format("2006-01-02 15:04:05.000 MST")
	},
	"Since": func(t time.Time) time.Duration {
		return time.Since(t)
	},
	"UILocales": func() []string {
		return domain.UILocales()
	},
	"LocaleParam": func(locale string) string {
		if locale == domain.DefaultLocale() {
			return "/"
		}
		return fmt.Sprintf("?locale=%s", locale)
	},
	"FormatDuration": func(d time.Duration) string {
		z := time.Unix(0, 0).UTC()
		return z.Add(d).Format("4:05")
	},
}

func NewHandler(
	app *app.App,
	logger *logger.Logger,
	t string,
) *Handler {
	return &Handler{
		app: app,
		log: logger,
		t:   t,
		ui:  initTranslationMap(),
	}
}

func initTranslationMap() map[string]*UITranslation {
	m := make(map[string]*UITranslation)
	for _, l := range domain.Locales() {
		m[l] = newUITranslation(l)
	}

	return m
}

func (h *Handler) NewConstants() *Constants {
	return &Constants{
		IntroStatus:     domain.TestStepIntro,
		QuestionsStatus: domain.TestStepQuestions,
		FinishStatus:    domain.TestStepFinish,
		ResultStatus:    domain.TestStepResult,
		MethodSten:      domain.ScaleTypeSten,
		MethodPerc:      domain.ScaleTypePerc,
		MethodMean:      domain.ScaleTypeMean,
		MethodSum:       domain.ScaleTypeSum,
	}
}

func (h *Handler) NewCommon(s live.Socket, currentView string) *CommonInstance {
	c := &CommonInstance{
		Env:             h.app.Cfg.Env,
		Domain:          h.app.Cfg.Server.Domain,
		Session:         fmt.Sprint(s.Session()),
		Error:           nil,
		Message:         nil,
		ShowLoginModal:  false,
		ShowLogoutModal: false,
		CurrentView:     currentView,
		Version:         h.app.Cfg.App.Version,
		ui:              h.ui,
		locale:          domain.DefaultLocale(), // it's private because changing requires additional logic
	}

	c.SetLocale(c.locale)
	return c
}

func (h *Handler) url404() *url.URL {
	u, _ := url.Parse("/404")
	return u
}

func (h *Handler) HandleError(ctx context.Context, err error) {
	h.log.Error("got bad request", err)
	w := live.Writer(ctx)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("bad request: " + err.Error()))
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
