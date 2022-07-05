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
	r.NotFoundHandler = http.HandlerFunc(h.NotFoundRedirect)
	// main handler
	sr := r.PathPrefix("/").Subrouter()
	sr.Handle("/test/{testCode}", live.NewHttpHandler(store, h.Test()))
	sr.Handle("/result/{takeID}", live.NewHttpHandler(store, h.Result()))
	sr.Handle("/about", live.NewHttpHandler(store, h.About()))
	sr.Handle("/profile", live.NewHttpHandler(store, h.Profile()))
	sr.Handle("/privacy", live.NewHttpHandler(store, h.Privacy()))
	sr.Handle("/terms", live.NewHttpHandler(store, h.Terms()))
	sr.Handle("/admin", live.NewHttpHandler(store, h.Admin()))
	sr.Handle("/404", live.NewHttpHandler(store, h.NotFound()))
	sr.Handle("/", live.NewHttpHandler(store, h.Home()))
	sr.Use(h.Middleware)

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
