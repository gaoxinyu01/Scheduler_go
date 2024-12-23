package middleware

import "net/http"

type CheckAuthMiddleware struct {
}

func NewCheckAuthMiddleware() *CheckAuthMiddleware {
	return &CheckAuthMiddleware{}
}

func (m *CheckAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
