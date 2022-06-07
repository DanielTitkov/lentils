package handler

import (
	"context"
	"net/http"
)

func (ah *Handler) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := ah.app.GetUserBySession(r)
		if err != nil {
			ah.log.Error("failed to get user from session", err)
		}
		ctx := context.WithValue(r.Context(), userCtxKey, user)

		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
