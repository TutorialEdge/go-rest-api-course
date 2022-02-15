package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service CommentService
}

// Response objecgi
type Response struct {
	Message string `json:"message"`
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service CommentService) *Handler {
	var h Handler
	h.Router = mux.NewRouter()
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LoggingMiddleware)

	h.mapRoutes()
	return &Handler{
		Service: service,
	}
}

// mapRoutes - sets up all the routes for our application
func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", JWTAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", JWTAuth(h.UpdateComment)).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", JWTAuth(h.DeleteComment)).Methods("DELETE")
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am Alive!"}); err != nil {
			panic(err)
		}
	})
}
