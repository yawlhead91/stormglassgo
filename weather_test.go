package stormglass

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"unicode"

	"github.com/jinzhu/now"

	"github.com/stretchr/testify/assert"
)

const (
	testKey      = "testkey123"
	endpointPath = "/weather/point"
	lat          = 58.7984
	lng          = 17.8081
)

func TestParamOptionsToList(t *testing.T) {
	t.Run("with no options", func(t *testing.T) {
		var options ParamsOptions
		list := options.toList()

		assert.Len(t, list, 0)
	})

	t.Run("with all options", func(t *testing.T) {
		options := ParamsOptions{
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

		var expected []string
		// this assumes naming conventions for fields names
		// as camel case matches of expected values
		for i := 0; i < v.NumField(); i++ {
			expected = append(expected, lcFirstLetter(v.Type().Field(i).Name))
		}

		assertion.Equal(expected, list)
	})
}

func TestSourceOptionsToList(t *testing.T) {
	t.Run("with no options", func(t *testing.T) {
		var options SourcesOptions
		list := options.toList()

		assert.Len(t, list, 0)
	})

	t.Run("with all options", func(t *testing.T) {
		options := SourcesOptions{
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

func TestClientGetPoint(t *testing.T) {
	var (
		start = now.BeginningOfDay()
		end   = now.EndOfDay()
	)

	t.Run("test full url composition", func(t *testing.T) {
		assertion := assert.New(t)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertion.NotNil(r.URL)
			assertion.Equal(r.URL.Path, endpointPath)
			assertion.Equal(
				r.URL.RawQuery,
				fmt.Sprintf(
					"lat=%f&lng=%f&params=%s,%s&start=%d&end=%d&source=%s,%s",
					lat,
					lng,
					"airTemperature",
					"waveDirection",
					start.Unix(),
					end.Unix(),
					"fcoo",
					"fmi",
				),
			)
			fmt.Fprintln(w, "{}")
		}))
		defer ts.Close()

		c := NewClient(testKey)
		c.BaseURL = ts.URL
		c.HTTPClient = ts.Client()

		ctx := context.Background()
		res, err := c.GetPoint(ctx, PointsRequestOptions{
			Lat: lat,
			Lng: lng,
			Params: ParamsOptions{
				AirTemperature: true,
				WaveDirection:  true,
			},
			Source: SourcesOptions{
				FCOO: true,
				FMI:  true,
			},
			Start: &start,
			End:   &end,
		})

		assertion.Nil(err, "expecting nil err")
		assertion.NotNil(res, "expecting non-nil response")
	})
	t.Run("test url composition with no params", func(t *testing.T) {
		assertion := assert.New(t)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertion.NotNil(r.URL)
			assertion.Equal(r.URL.Path, endpointPath)
			assertion.Equal(
				r.URL.RawQuery,
				fmt.Sprintf(
					"lat=%f&lng=%f&start=%d&end=%d&source=%s,%s",
					lat,
					lng,
					start.Unix(),
					end.Unix(),
					"fcoo",
					"fmi",
				),
			)
			fmt.Fprintln(w, "{}")
		}))
		defer ts.Close()

		c := NewClient(testKey)
		c.BaseURL = ts.URL
		c.HTTPClient = ts.Client()

		ctx := context.Background()
		res, err := c.GetPoint(ctx, PointsRequestOptions{
			Lat: lat,
			Lng: lng,
			Source: SourcesOptions{
				FCOO: true,
				FMI:  true,
			},
			Start: &start,
			End:   &end,
		})

		assertion.Nil(err, "expecting nil err")
		assertion.NotNil(res, "expecting non-nil response")
	})
	t.Run("test url composition with no source", func(t *testing.T) {
		assertion := assert.New(t)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertion.NotNil(r.URL)
			assertion.Equal(r.URL.Path, endpointPath)
			assertion.Equal(
				r.URL.RawQuery,
				fmt.Sprintf(
					"lat=%f&lng=%f&start=%d&end=%d",
					lat,
					lng,
					start.Unix(),
					end.Unix(),
				),
			)
			fmt.Fprintln(w, "{}")
		}))
		defer ts.Close()

		c := NewClient(testKey)
		c.BaseURL = ts.URL
		c.HTTPClient = ts.Client()

		ctx := context.Background()
		res, err := c.GetPoint(ctx, PointsRequestOptions{
			Lat:   lat,
			Lng:   lng,
			Start: &start,
			End:   &end,
		})

		assertion.Nil(err, "expecting nil err")
		assertion.NotNil(res, "expecting non-nil response")
	})
	t.Run("test url composition with no start date", func(t *testing.T) {
		assertion := assert.New(t)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assertion.NotNil(r.URL)
			assertion.Equal(r.URL.Path, endpointPath)
			assertion.Equal(
				r.URL.RawQuery,
				fmt.Sprintf(
					"lat=%f&lng=%f&end=%d",
					lat,
					lng,
					end.Unix(),
				),
			)
			fmt.Fprintln(w, http.NoBody)
		}))
		defer ts.Close()

		c := NewClient(testKey)
		c.BaseURL = ts.URL
		c.HTTPClient = ts.Client()

		ctx := context.Background()
		res, err := c.GetPoint(ctx, PointsRequestOptions{
			Lat: lat,
			Lng: lng,
			End: &end,
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
			Lat: lat,
			Lng: lng,
			End: &end,
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
