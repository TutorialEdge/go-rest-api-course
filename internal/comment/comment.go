package comment

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Service - our comment service
type Service struct {
	DB *gorm.DB
}

// Comment -
type Comment struct {
	ID      string
	Slug    string
	Body    string
	Author  string
	Created time.Time
}

// CommentService -
type CommentService interface {
	GetComment(ID string) (Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(comment Comment) (Comment, error)
	DeleteComment(comment Comment) error
	GetAllComments() ([]Comment, error)
}

// NewService - returns a new comments service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetComment -
func (s *Service) GetComment(id string) (Comment, error) {
	return Comment{}, nil
}

// PostComment -
func (s *Service) PostComment(comment Comment) (Comment, error) {
	return Comment{}, nil
}

// UpdateComment -
func (s *Service) UpdateComment(comment Comment) (Comment, error) {
	return Comment{}, nil
}

// DeleteComment -
func (s *Service) DeleteComment(comment Comment) (Comment, error) {
	return Comment{}, nil
}

// GetAllComments -
func (s *Service) GetAllComments() ([]Comment, error) {
	return []Comment{}, nil
}
