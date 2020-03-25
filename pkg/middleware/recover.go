package middleware

import (
	"github.com/gorilla/mux"
	"github.com/imjuanleonard/palu-covid/pkg/logger"
	"net/http"
	"runtime/debug"
)

func Recover() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Errorf("Recovered from panic: %+v\n%s", err, string(debug.Stack()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
