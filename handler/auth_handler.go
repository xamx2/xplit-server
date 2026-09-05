package handler

import (
	"net/http"

	"github.com/xamx2/xplit-server/contexts"
)

type AuthHandler struct {
	Next http.Handler
}

var _ http.Handler = (*AuthHandler)(nil)

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := contexts.WithAuth(r.Context(), r)
	h.Next.ServeHTTP(w, r.WithContext(ctx))
}
