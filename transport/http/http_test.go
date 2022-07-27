package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttp_Call(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/test-url", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"data":{"test":"test"}}`))
		}))
		defer server.Close()

		testHTTP := New(server.URL, "", "")
		testURL, err := url.Parse("/test-url")
		assert.NoError(t, err)

		res, err := testHTTP.UnSignedRequest(context.Background(), testURL, http.MethodGet, []byte{}, map[string]string{})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, `{"data":{"test":"test"}}`, string(res))
	})

	t.Run("unauthorized error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/test-error-url", r.URL.String())
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"errors":[{"error":"unauthorized"}]}`))
		}))
		defer server.Close()

		testHTTP := New(server.URL, "", "")
		testURL, err := url.Parse("/test-error-url")
		assert.NoError(t, err)

		res, err := testHTTP.UnSignedRequest(context.Background(), testURL, http.MethodGet, []byte{}, map[string]string{})
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "HTTP 401: ", err.Error())
	})

}
