// +build integration
package comment_test

import (
	"fmt"
	"testing"

	"github.com/TutorialEdge/go-rest-api-course/internal/comment"
	"github.com/TutorialEdge/go-rest-api-course/internal/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// CommentTestSuite -
type CommentTestSuite struct {
	suite.Suite
	service *comment.Service
}

// TestCommentTestSuite - main entry point for our test suite
func TestCommentTestSuite(t *testing.T) {
	suite.Run(t, new(CommentTestSuite))
}

// SetupTest -
func (s *CommentTestSuite) SetupTest() {
	fmt.Println("This is run once at the start of a test suite")
	db, _ := database.NewDatabase()
	s.service = comment.NewService(db)
}

// TestGetComment - tests to see if the comment service can
// retrieve comments from the database
func (s *CommentTestSuite) TestGetComment() {
	s.T().Run("Test Get Comment", func(t *testing.T) {
		comment, err := s.service.GetComment(1)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "/testget", comment.Slug)
	})

	s.T().Run("Test No Comment", func(t *testing.T) {
		comment, err := s.service.GetComment(9999)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "/testget", comment.Slug)
	})
}

// TestGetAllComments - tests to see if the comment service can
// retrieve all comments from the database
func (s *CommentTestSuite) TestGetAllComments() {
	s.T().Run("Test Get All Comments", func(t *testing.T) {
		comments, err := s.service.GetAllComments()
		assert.NoError(s.T(), err)
		assert.NotEmpty(s.T(), comments)
	})
}

// TestPostComment - tests to see if the comment service can post a comment
// to the database
func (s *CommentTestSuite) TestPostComment() {

}

// TestDeleteComment - tests to see if the comment service can delete a
// comment from the database
func (s *CommentTestSuite) TestDeleteComment() {

}
