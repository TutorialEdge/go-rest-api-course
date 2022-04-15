// +build e2e

package tests

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckEndpoint(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/alive")
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}
