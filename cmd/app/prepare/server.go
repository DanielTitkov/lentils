package prepare

import (
	"net/http"
	"time"

	"github.com/DanielTitkov/orrery/internal/configs"
	"github.com/DanielTitkov/orrery/internal/handler"
	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

func Mux(cfg configs.Config, store live.HttpSessionStore, h *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	r.Use(h.Middleware)
	r.NotFoundHandler = http.HandlerFunc(h.NotFoundRedirect)
	// main handler
	r.Handle("/test/{testCode}", live.NewHttpHandler(store, h.Test()))
	r.Handle("/result/{takeID}", live.NewHttpHandler(store, h.Result()))
	r.Handle("/about", live.NewHttpHandler(store, h.About()))
	r.Handle("/profile", live.NewHttpHandler(store, h.Profile()))
	r.Handle("/privacy", live.NewHttpHandler(store, h.Privacy()))
	r.Handle("/terms", live.NewHttpHandler(store, h.Terms()))
	r.Handle("/admin", live.NewHttpHandler(store, h.Admin()))
	r.Handle("/404", live.NewHttpHandler(store, h.NotFound()))
	r.Handle("/", live.NewHttpHandler(store, h.Home()))

	// live scripts
	r.Handle("/live.js", live.Javascript{})
	r.Handle("/auto.js.map", live.JavascriptMap{})

	// auth
	r.HandleFunc("/auth/logout", h.Logout)
	r.HandleFunc("/auth/{provider}", h.BeginOAuth)
	r.HandleFunc("/auth/{provider}/callback", h.CompleteOAuth)

	// static
	r.HandleFunc("/favicon.ico", faviconHandler)
	r.HandleFunc("/static/css/styles.css", stylesHandler)

	return r
}

func Server(cfg configs.Config, handler *mux.Router) *http.Server {
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
	}

	return server
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/media/favicon.ico")
}

func stylesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/dist/css/styles.css")
}
