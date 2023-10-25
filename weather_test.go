package stormglass

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"sort"
	"testing"
	"unicode"

	"github.com/jinzhu/now"

	"github.com/stretchr/testify/assert"
)

func TestParamOptionsToList(t *testing.T) {
	t.Run("with no options", func(t *testing.T) {
		var options WeatherParamsOptions
		list := options.toList()

		assert.Len(t, list, 0)
	})

	t.Run("with all options", func(t *testing.T) {
		options := WeatherParamsOptions{}

		v := reflect.ValueOf(&options)
		e := v.Elem()

		for i := 0; i < e.NumField(); i++ {
			e.Field(i).SetBool(true)
		}

		list := options.toList()

		assertion := assert.New(t)
		assertion.Len(list, e.NumField())

		var expected []string
		// this assumes naming conventions for fields names
		// as camel case matches of expected values
		for i := 0; i < e.NumField(); i++ {
			expected = append(expected, lcFirstLetter(e.Type().Field(i).Name))
		}

		s := func(a []string) {
			sort.Slice(a, func(i, j int) bool {
				return a[i] < a[j]
			})
		}

		s(expected)
		s(list)

		assertion.Equal(expected, list)
	})
}

func TestSourceOptionsToList(t *testing.T) {
	t.Run("with no options", func(t *testing.T) {
		var options WeatherSourcesOptions
		list := options.toList()

		assert.Len(t, list, 0)
	})

	t.Run("with all options", func(t *testing.T) {
		options := WeatherSourcesOptions{
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
		}

		list := options.toList()

		v := reflect.ValueOf(options)

		assertion := assert.New(t)
		assertion.Len(list, v.NumField())

		expected := []string{
			"icon",
			"dwt",
			"noaa",
			"meteo",
			"meto",
			"fcoo",
			"fmi",
			"yr",
			"smhi",
			"sg",
		}

		assertion.Equal(expected, list)
	})
}

func TestClient_GetPoint(t *testing.T) {
	var (
		start        = now.BeginningOfDay()
		end          = now.EndOfDay()
		testKey      = "testkey123"
		endpointPath = "/weather/point"
		lat          = 58.7984
		lng          = 17.8081
	)

	t.Run("test full url composition", func(t *testing.T) {
		assertion := assert.New(t)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertion.NotNil(r.URL)
			assertion.Equal(endpointPath, r.URL.Path)

			expectedValues := url.Values{}
			expectedValues.Set("lat", fmt.Sprintf("%f", lat))
			expectedValues.Set("lng", fmt.Sprintf("%f", lng))
			expectedValues.Set("start", fmt.Sprintf("%d", start.Unix()))
			expectedValues.Set("end", fmt.Sprintf("%d", end.Unix()))
			expectedValues.Set("params", "airTemperature,waveDirection")
			expectedValues.Set("source", "fcoo,fmi")

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
		res, err := c.GetPoint(ctx, PointsRequestOptions{
			CommonRequestOptions: CommonRequestOptions{
				Lat:   lat,
				Lng:   lng,
				Start: &start,
				End:   &end,
			},
			Params: WeatherParamsOptions{
				AirTemperature: true,
				WaveDirection:  true,
			},
			Source: WeatherSourcesOptions{
				FCOO: true,
				FMI:  true,
			},
		})

		assertion.NoError(err, "expecting nil err")
		assertion.NotNil(res, "expecting non-nil response")
	})
	t.Run("test url composition with no params", func(t *testing.T) {
		assertion := assert.New(t)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertion.NotNil(r.URL)
			assertion.Equal(endpointPath, r.URL.Path)

			expectedValues := url.Values{}
			expectedValues.Set("lat", fmt.Sprintf("%f", lat))
			expectedValues.Set("lng", fmt.Sprintf("%f", lng))
			expectedValues.Set("start", fmt.Sprintf("%d", start.Unix()))
			expectedValues.Set("end", fmt.Sprintf("%d", end.Unix()))
			expectedValues.Set("source", "fcoo,fmi")

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
		res, err := c.GetPoint(ctx, PointsRequestOptions{
			CommonRequestOptions: CommonRequestOptions{
				Lat:   lat,
				Lng:   lng,
				Start: &start,
				End:   &end,
			},
			Source: WeatherSourcesOptions{
				FCOO: true,
				FMI:  true,
			},
		})

		assertion.Nil(err, "expecting nil err")
		assertion.NotNil(res, "expecting non-nil response")
	})
	t.Run("test url composition with no source", func(t *testing.T) {
		assertion := assert.New(t)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertion.NotNil(r.URL)
			assertion.Equal(endpointPath, r.URL.Path)

			expectedValues := url.Values{}
			expectedValues.Set("lat", fmt.Sprintf("%f", lat))
			expectedValues.Set("lng", fmt.Sprintf("%f", lng))
			expectedValues.Set("start", fmt.Sprintf("%d", start.Unix()))
			expectedValues.Set("end", fmt.Sprintf("%d", end.Unix()))

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
		res, err := c.GetPoint(ctx, PointsRequestOptions{
			CommonRequestOptions: CommonRequestOptions{
				Lat:   lat,
				Lng:   lng,
				Start: &start,
				End:   &end,
			},
		})

		assertion.Nil(err, "expecting nil err")
		assertion.NotNil(res, "expecting non-nil response")
	})
	t.Run("test url composition with no start date", func(t *testing.T) {
		assertion := assert.New(t)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertion.NotNil(r.URL)
			assertion.Equal(endpointPath, r.URL.Path)

			expectedValues := url.Values{}
			expectedValues.Set("lat", fmt.Sprintf("%f", lat))
			expectedValues.Set("lng", fmt.Sprintf("%f", lng))
			expectedValues.Set("end", fmt.Sprintf("%d", end.Unix()))

			assertion.Equal(
				expectedValues.Encode(),
				r.URL.RawQuery,
			)
			_, _ = fmt.Fprintln(w, http.NoBody)
		}))
		defer ts.Close()

		c := NewClient(testKey)
		c.BaseURL = ts.URL
		c.HTTPClient = ts.Client()

		ctx := context.Background()
		res, err := c.GetPoint(ctx, PointsRequestOptions{
			CommonRequestOptions: CommonRequestOptions{
				Lat: lat,
				Lng: lng,
				End: &end,
			},
		})

		assertion.Nil(err)
		assertion.NotNil(res)
	})

	t.Run("with response error", func(t *testing.T) {
		assertion := assert.New(t)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			want := `{"errors:{"key", "Unauthorized – Your API key is invalid."}}`
			w.WriteHeader(http.StatusUnauthorized)
			if _, err := w.Write([]byte(want)); err != nil {
				t.Fatal(err)
			}
		}))
		defer ts.Close()

		c := NewClient(testKey)
		c.BaseURL = ts.URL
		c.HTTPClient = ts.Client()

		ctx := context.Background()
		res, err := c.GetPoint(ctx, PointsRequestOptions{
			CommonRequestOptions: CommonRequestOptions{
				Lat: lat,
				Lng: lng,
				End: &end,
			},
		})

		assertion.Nil(res)
		assertion.NotNil(err)
	})
}

// lower case first letter helper func.
func lcFirstLetter(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}

	return ""
}
