package endpoints_test

import (
	"testing"

	"github.com/petermnhull/go-music-app/internal/endpoints"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAPIResponseSuccess(t *testing.T) {
	t.Run("new success API response works", func(t *testing.T) {
		data := map[string]string{}
		actual := endpoints.NewAPIResponseSuccess(200, data)
		expected := endpoints.APIResponse{200, "success", map[string]string{}}
		assert.Equal(t, &expected, actual)
	})
}

func TestNewAPIResponseFailed(t *testing.T) {
	t.Run("new failed API response works", func(t *testing.T) {
		actual := endpoints.NewAPIResponseFailed(500, "database connection failure")
		expected := endpoints.APIResponse{500, "failed", map[string]string{"error": "database connection failure"}}
		assert.Equal(t, &expected, actual)
	})
}

func TestAPIResponseToOutput(t *testing.T) {
	t.Run("api response renders output", func(t *testing.T) {
		data := map[string]any{}
		r := endpoints.APIResponse{200, "success", data}
		actual := r.ToOutput()
		expected := `{"code": 200, "status": "success", "data": {}}`
		require.JSONEq(t, expected, actual)
	})
}
