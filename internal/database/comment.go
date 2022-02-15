package database

import (
	"context"
	"errors"

	"github.com/TutorialEdge/go-rest-api-course/internal/comment"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

// CommentRow - models how our comments look in the database
type CommentRow struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// GetComment - retrieves a comment from the database by ID
func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	// fetch CommentRow from the database and then convert to comment.Comment

	var cmt comment.Comment
	// sqlx with context to ensure context cancelation is honoured
	return cmt, ErrNotImplemented
}

// GetCommentsBySlug -
func (d *Database) GetCommentsBySlug(ctx context.Context, slug string) ([]comment.Comment, error) {
	return []comment.Comment{}, ErrNotImplemented
}

func (d *Database) PostComment(ctx context.Context) (comment.Comment, error) {
	return comment.Comment{}, ErrNotImplemented
}

func (d *Database) UpdateComment(ctx context.Context) (comment.Comment, error) {
	return comment.Comment{}, ErrNotImplemented
}

func (d *Database) DeleteComment(ctx context.Context) error {
	return ErrNotImplemented
}
