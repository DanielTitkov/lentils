package handler

import (
	"context"
	"fmt"
	"net/http"
)

func (ah *Handler) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("MIDDLEWARE", r) // FIXME
		user, err := ah.app.GetUserBySession(r)
		if err != nil {
			ah.log.Error("failed to get user from session", err)
		}
		if user != nil {
			fmt.Println("MIDDLEWARE WRITING TO CONTEXT") // FIXME
			ctx := context.WithValue(r.Context(), userCtxKey, user)
			r = r.WithContext(ctx)
		} else {
			fmt.Println("MIDDLEWARE ___NOT___ WRITING TO CONTEXT") // FIXME
		}

		h.ServeHTTP(w, r)
	})
}
