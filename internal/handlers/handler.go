package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Handler struct {
	clerkClient clerk.Client
}

func New(clerkClient clerk.Client) *Handler {
	return &Handler{
		clerkClient: clerkClient,
	}
}

func (h *Handler) LoadRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/health"))
	r.Use(clerk.WithSessionV2(h.clerkClient))

	r.Get("/", returnActiveSession)

	return r
}

func returnActiveSession(w http.ResponseWriter, r *http.Request) {
	sessionClaims, ok := clerk.SessionFromContext(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	jsonResp, _ := json.Marshal(sessionClaims)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonResp))
}
