package http_service

import (
	"bytes"
	"fmt"
	"github.com/carolinaborim/go-qr-code-generator/qrtst"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handler(t *testing.T) {
	t.Run("returning valid qr image for url", func(t *testing.T) {
		url := "https://google.com"
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/?url=%s", url), nil)

		handler(w, req)
		result := w.Result()

		assert.Equal(t, "image/png", result.Header.Get("Content-Type"))
		assert.Equal(t, http.StatusOK, result.StatusCode)

		bodyBytes, err := io.ReadAll(result.Body)
		require.NoError(t, err)
		assert.Equal(t, url, qrtst.Read(t, bytes.NewReader(bodyBytes)))
	})

	t.Run("returning an error if url is missing", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		handler(w, req)
		result := w.Result()

		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		assert.Equal(t, http.StatusBadRequest, result.StatusCode)

		bodyBytes, err := io.ReadAll(result.Body)
		require.NoError(t, err)
		assert.JSONEq(t, `{"error":"missing url parameter"}`, string(bodyBytes))
		t.Log(string(bodyBytes))
	})

}
