package stormglass

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/jinzhu/now"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetExtremesPoint(t *testing.T) {
	var (
		start        = now.BeginningOfDay()
		end          = now.EndOfDay()
		testKey      = "testkey123"
		endpointPath = "/tide/extremes/point"
		lat          = 58.7984
		lng          = 17.8081
	)

	t.Run("test full url composition", func(t *testing.T) {
		assertion := assert.New(t)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertion.NotNil(r.URL)
			assertion.Equal(endpointPath, r.URL.Path)

			expectedValues := url.Values{}
			expectedValues.Add("lat", fmt.Sprintf("%f", lat))
			expectedValues.Add("lng", fmt.Sprintf("%f", lng))
			expectedValues.Add("start", fmt.Sprintf("%d", start.Unix()))
			expectedValues.Add("end", fmt.Sprintf("%d", end.Unix()))
			expectedValues.Add("datum", "MLLW")
			assertion.Equal(
				expectedValues.Encode(),
				r.URL.RawQuery,
			)

			_, _ = fmt.Fprintln(w, "{}")
		}))
		defer ts.Close()

		c := NewClient(testKey)
		c.BaseURL = ts.URL
		c.HTTPClient = ts.Client()

		ctx := context.Background()
		res, err := c.GetExtremesPoint(ctx, ExtremesPointsRequestOptions{
			CommonRequestOptions: CommonRequestOptions{
				Lat:   lat,
				Lng:   lng,
				Start: &start,
				End:   &end,
			},
			Datum: MLLW,
		})

		assertion.NoError(err, "expecting nil err")
		assertion.NotNil(res, "expecting non-nil response")
	})

	t.Run("test full with no datum", func(t *testing.T) {
		assertion := assert.New(t)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertion.NotNil(r.URL)
			assertion.Equal(endpointPath, r.URL.Path)

			expectedValues := url.Values{}
			expectedValues.Add("lat", fmt.Sprintf("%f", lat))
			expectedValues.Add("lng", fmt.Sprintf("%f", lng))
			expectedValues.Add("start", fmt.Sprintf("%d", start.Unix()))
			expectedValues.Add("end", fmt.Sprintf("%d", end.Unix()))
			assertion.Equal(
				expectedValues.Encode(),
				r.URL.RawQuery,
			)

			_, _ = fmt.Fprintln(w, "{}")
		}))
		defer ts.Close()

		c := NewClient(testKey)
		c.BaseURL = ts.URL
		c.HTTPClient = ts.Client()

		ctx := context.Background()
		res, err := c.GetExtremesPoint(ctx, ExtremesPointsRequestOptions{
			CommonRequestOptions: CommonRequestOptions{
				Lat:   lat,
				Lng:   lng,
				Start: &start,
				End:   &end,
			},
		})

		assertion.NoError(err, "expecting nil err")
		assertion.NotNil(res, "expecting non-nil response")
	})

	t.Run("test full with no dates", func(t *testing.T) {
		assertion := assert.New(t)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertion.NotNil(r.URL)
			assertion.Equal(endpointPath, r.URL.Path)

			expectedValues := url.Values{}
			expectedValues.Add("lat", fmt.Sprintf("%f", lat))
			expectedValues.Add("lng", fmt.Sprintf("%f", lng))
			assertion.Equal(
				expectedValues.Encode(),
				r.URL.RawQuery,
			)

			_, _ = fmt.Fprintln(w, "{}")
		}))
		defer ts.Close()

		c := NewClient(testKey)
		c.BaseURL = ts.URL
		c.HTTPClient = ts.Client()

		ctx := context.Background()
		res, err := c.GetExtremesPoint(ctx, ExtremesPointsRequestOptions{
			CommonRequestOptions: CommonRequestOptions{
				Lat: lat,
				Lng: lng,
			},
		})

		assertion.NoError(err, "expecting nil err")
		assertion.NotNil(res, "expecting non-nil response")
	})
}
