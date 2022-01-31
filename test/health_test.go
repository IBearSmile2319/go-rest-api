//go:build e2e
// +build e2e

package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestHealth - tests our health endpoint
func TestHealthEndPint(t *testing.T) {
	fmt.Println("Running E2E test for health endpoint")
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/health")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}
