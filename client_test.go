package stormglass

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			want := `{"errors":{"key":"API key is invalid"}}`
			w.WriteHeader(http.StatusUnauthorized)
			if _, err := w.Write([]byte(want)); err != nil {
				t.Fatal(err)
			}
		}))
		defer ts.Close()

		c := NewClient(testKey)
		c.BaseURL = ts.URL
		c.HTTPClient = ts.Client()

		req, _ := http.NewRequest("GET", ts.URL, http.NoBody)
		err := c.sendRequest(req, nil)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "API key is invalid")
	})
	t.Run("test decode error", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			want := `{{`
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte(want)); err != nil {
				t.Fatal(err)
			}
		}))
		defer ts.Close()

		c := NewClient(testKey)
		c.BaseURL = ts.URL
		c.HTTPClient = ts.Client()

		req, _ := http.NewRequest("GET", ts.URL, http.NoBody)
		err := c.sendRequest(req, nil)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "decode error")
	})

	t.Run("success", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			want := `{ "name": "test" }`
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte(want)); err != nil {
				t.Fatal(err)
			}
		}))
		defer ts.Close()

		c := NewClient(testKey)
		c.BaseURL = ts.URL
		c.HTTPClient = ts.Client()

		var obj struct {
			Name string `json:"name"`
		}
		req, _ := http.NewRequest("GET", ts.URL, http.NoBody)
		err := c.sendRequest(req, &obj)
		assert.Nil(t, err)
		assert.Equal(t, obj.Name, "test")
	})
}
