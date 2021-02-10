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
	gorm.Model
	Slug    string
	Body    string
	Author  string
	Created time.Time
}

// CommentService -
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

// NewService - returns a new comments service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetComment -
func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// PostComment -
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// UpdateComment -
func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}

	if result := s.DB.Model(&comment).Updates(newComment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// DeleteComment -
func (s *Service) DeleteComment(ID uint) (Comment, error) {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return Comment{}, nil
}

// GetAllComments -
func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}
	return comments, nil
}
