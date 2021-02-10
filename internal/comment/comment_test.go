// +build integration
package comment_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// CommentTestSuite -
type CommentTestSuite struct {
	suite.Suite
}

// TestCommentTestSuite - main entry point for our test suite
func TestCommentTestSuite(t *testing.T) {
	suite.Run(t, new(CommentTestSuite))
}

// SetupTest -
func (s *CommentTestSuite) SetupTest() {

}

// TestGetComment - tests to see if the comment service can
// retrieve comments from the database
func (s *CommentTestSuite) TestGetComment() {

}

// TestGetAllComments - tests to see if the comment service can
// retrieve all comments from the database
func (s *CommentTestSuite) TestGetAllComments() {

}

// TestPostComment - tests to see if the comment service can post a comment
// to the database
func (s *CommentTestSuite) TestPostComment() {

}

// TestDeleteComment - tests to see if the comment service can delete a
// comment from the database
func (s *CommentTestSuite) TestDeleteComment() {

}
