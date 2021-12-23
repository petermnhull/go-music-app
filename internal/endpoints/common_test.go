package endpoints_test

import (
	"testing"

	"github.com/petermnhull/go-music-app/internal/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestNewFailedAPIResponse(t *testing.T) {
	t.Run("new failed API response works", func(t *testing.T) {
		actual := endpoints.NewFailedAPIResponse(500, "database connection failure")
		expected := endpoints.APIResponse{500, "failed", map[string]string{"error": "database connection failure"}}
		assert.Equal(t, &expected, actual)
	})
}
