package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/TutorialEdge/go-rest-api-course/internal/comment"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/go-playground/validator/v10"
)

type CommentService interface {
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, newCmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
	ReadyCheck(ctx context.Context) error
}

// GetComment - retrieve a comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		if errors.Is(err, comment.ErrFetchingComment) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

// PostCommentRequest
type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func commentFromPostCommentRequest(u PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   u.Slug,
		Author: u.Author,
		Body:   u.Body,
	}
}

// PostComment - adds a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var postCmtReq PostCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&postCmtReq); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(postCmtReq)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt := commentFromPostCommentRequest(postCmtReq)
	cmt, err = h.Service.PostComment(r.Context(), cmt)
	if err != nil {
		log.Error(err)
		return
	}
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

// UpdateCommentRequest -
type UpdateCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

// convert the validated struct into something that the service layer understands
// this is a little verbose, but it allows us to remove tight coupling between our components
func commentFromUpdateCommentRequest(u UpdateCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   u.Slug,
		Author: u.Author,
		Body:   u.Body,
	}
}

// UpdateComment - updates a comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID := vars["id"]

	var updateCmtRequest UpdateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&updateCmtRequest); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(updateCmtRequest)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt := commentFromUpdateCommentRequest(updateCmtRequest)

	cmt, err = h.Service.UpdateComment(r.Context(), commentID, cmt)
	if err != nil {
		log.Error(err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

// DeleteComment - deletes a comment by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID := vars["id"]

	if commentID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteComment(r.Context(), commentID)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully Deleted"}); err != nil {
		panic(err)
	}
}
