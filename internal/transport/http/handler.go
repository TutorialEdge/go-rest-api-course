package http

import (
	"fmt"
	"net/http"

	"github.com/TutorialEdge/go-rest-api-course/internal/comment"
	"github.com/gorilla/mux"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id:[0-9]+}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id:[0-9]+}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id:[0-9]+}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})
}

// GetComment -
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.GetComment(1)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, "%+v", comment)
}

// GetAllComments -
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, "%+v", comments)
}

// PostComment -
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, "%+v", comment)
}

// UpdateComment -
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.UpdateComment(3, comment.Comment{
		Slug: "/new",
	})
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, "%+v", comment)
}

// DeleteComment -
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.DeleteComment(1)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, "%+v", comment)
}
