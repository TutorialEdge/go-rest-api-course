package database

import (
	"context"

	"github.com/TutorialEdge/go-rest-api-course/internal/models"
)

// CommentRow - models how our comments look in the database
type CommentRow struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// GetComment - retrieves a comment from the database by ID
func (d *Database) GetComment(ctx context.Context, uuid string) (models.Comment, error) {
	// fetch CommentRow from the database and then convert to models.Comment

	var cmt models.Comment
	// sqlx with context
	return cmt, nil
}

// GetCommentsBySlug -
func (d *Database) GetCommentsBySlug(ctx context.Context, slug string) ([]models.Comment, error) {
	return []models.Comment{}, nil
}

func (d *Database) PostComment(ctx context.Context) (models.Comment, error) {
	return models.Comment{}, nil
}

func (d *Database) UpdateComment(ctx context.Context) (models.Comment, error) {
	return models.Comment{}, nil
}

func (d *Database) DeleteComment(ctx context.Context) error {
	return nil
}
