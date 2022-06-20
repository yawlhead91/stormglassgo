package stormglass

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("should return default client", func(t *testing.T) {
		client := NewClient(testKey)
		assert.NotNil(t, client)
		assert.Equal(t, client.BaseURL, BaseURLV2)
		assert.NotNil(t, client.HTTPClient)
	})
}

func TestClientSendRequest(t *testing.T) {
	t.Run("test error status code", func(t *testing.T) {
		assert := assert.New(t)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			want := `{"errors:{"key", "Unauthorized – Your API key is invalid."}}`
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(want))
		}))
		defer ts.Close()

		c := NewClient(testKey)
		c.BaseURL = ts.URL
		c.HTTPClient = ts.Client()

	})
	t.Run("test decode error", func(t *testing.T) {

	})
	t.Run("success", func(t *testing.T) {

	})
}
