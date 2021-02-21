// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestCommentsEndpoint(t *testing.T) {
	fmt.Println("Running E2E Tests For Comments API")

	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comment")
	if err != nil {
		t.Fail()
	}

	fmt.Println(resp.StatusCode())
}
