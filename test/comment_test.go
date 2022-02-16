// +build e2e

package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

const (
	BASE_URL = "http://localhost:8080"
)

func TestGetComments(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/v1/comment")
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())

	// resp, err := client.R().Get(BASE_URL + "/api/v1/comment")
	// if err != nil {
	// 	t.Fail()
	// }

	// assert.Equal(t, 200, resp.StatusCode())
}

func TestPostComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/v1/comment")
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}
