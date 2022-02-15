package comment

import (
	"context"
	"errors"

	"github.com/TutorialEdge/go-rest-api-course/internal/models"
)

var (
	ErrFetchingComment = errors.New("could not fetch comment by ID")
	ErrUpdatingComment = errors.New("could not update comment")
	ErrNoCommentFound  = errors.New("no comment found")
	ErrDeletingComment = errors.New("could not delete comment")
	ErrNotImplemented  = errors.New("not implemented")
)

// CommentStore - defines the interface we need our comment storage
// layer to implement
type CommentStore interface {
	GetComment(context.Context, string) (models.Comment, error)
}

// Service - the struct for our comment service
type Service struct {
	Store CommentStore
}

// NewService - returns a new comment service
func NewService(store CommentStore) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment - retrieves comments by their ID from the database
func (s *Service) GetComment(ctx context.Context, ID string) (models.Comment, error) {
	return models.Comment{}, ErrNotImplemented
}

// GetCommentsBySlug - retrieves all comments by slug (path - /article/name/)
func (s *Service) GetCommentsBySlug(ctx context.Context, slug string) ([]models.Comment, error) {
	return []models.Comment{}, ErrNotImplemented
}

// PostComment - adds a new comment to the database
func (s *Service) PostComment(ctx context.Context, cmt models.Comment) (models.Comment, error) {
	return models.Comment{}, ErrNotImplemented
}

// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(ctx context.Context, ID string, newComment models.Comment) (models.Comment, error) {
	return models.Comment{}, ErrNotImplemented
}

// DeleteComment - deletes a comment from the database by ID
func (s *Service) DeleteComment(ctx context.Context, ID string) error {
	return ErrNotImplemented
}

// GetAllComments - retrieves all comments from the database
func (s *Service) GetAllComments(ctx context.Context) ([]models.Comment, error) {
	return []models.Comment{}, ErrNotImplemented
}
