package comment

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
)

var (
	ErrFetchingComment = errors.New("could not fetch comment by ID")
	ErrUpdatingComment = errors.New("could not update comment")
	ErrNoCommentFound  = errors.New("no comment found")
	ErrDeletingComment = errors.New("could not delete comment")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - defines our comment structure
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// CommentStore - defines the interface we need our comment storage
// layer to implement
type CommentStore interface {
	GetComment(context.Context, string) (Comment, error)
	Ping(context.Context) error
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
func (s *Service) GetComment(ctx context.Context, ID string) (Comment, error) {
	// calls store passing in the context
	cmt, err := s.Store.GetComment(ctx, ID)
	if err != nil {
		log.Errorf("an error occured fetching the comment: %s", err.Error())
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

// GetCommentsBySlug - retrieves all comments by slug (path - /article/name/)
func (s *Service) GetCommentsBySlug(ctx context.Context, slug string) ([]Comment, error) {
	return []Comment{}, ErrNotImplemented
}

// PostComment - adds a new comment to the database
func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}

// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(ctx context.Context, ID string, newComment Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}

// DeleteComment - deletes a comment from the database by ID
func (s *Service) DeleteComment(ctx context.Context, ID string) error {
	return ErrNotImplemented
}

// GetAllComments - retrieves all comments from the database
func (s *Service) GetAllComments(ctx context.Context) ([]Comment, error) {
	return []Comment{}, ErrNotImplemented
}

// ReadyCheck - a function that tests we are functionally ready to serve requests
func (s *Service) ReadyCheck(ctx context.Context) error {
	log.Info("Checking readiness")
	return s.Store.Ping(ctx)
}
