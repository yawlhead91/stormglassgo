//go:build integration
// +build integration

package integration_tests_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jinzhu/now"

	"github.com/stretchr/testify/assert"
	stormglass "github.com/yawlhead91/stormglassgo"
)

const (
	// Banzai Pipeline, Oahu, Hawaii
	lat = 58.7984
	lng = 17.8081
)

func TestGetPoint(t *testing.T) {
	t.Run(" return api error", func(t *testing.T) {
		var start = now.BeginningOfDay()

		c := stormglass.NewClient("")

		ctx := context.Background()
		res, err := c.GetPoint(ctx, stormglass.PointsRequestOptions{
			Lat: lat,
			Lng: lng,
			Params: stormglass.ParamsOptions{
				AirTemperature: true,
			},
			Start: &start,
		})

		assert.Nil(t, res, "expecting nil response")
		assert.NotNil(t, err, "expecting non-nil error")
		assert.Equal(t, "403: key:API key is invalid", err.Error(), "unexpected error")
	})
	t.Run("success: return 24 hour points", func(t *testing.T) {
		tme := time.Now().In(time.UTC)
		var start = now.New(tme).BeginningOfDay()
		var end = now.New(tme).BeginningOfDay().Add(time.Hour * 23) // api returns hours ahead

		loc := start.Location()
		log.Println(loc)

		c := stormglass.NewClient(os.Getenv("STORMGLASS_API_KEY"))

		params := stormglass.ParamsOptions{
			AirTemperature:   true,
			WaveDirection:    true,
			WaterTemperature: true,
		}

		sources := stormglass.SourcesOptions{
			ICON:        true,
			UKMetOffice: true,
			StormGlass:  true,
		}

		ctx := context.Background()
		res, err := c.GetPoint(ctx, stormglass.PointsRequestOptions{
			Lat:    lat,
			Lng:    lng,
			Params: params,
			Start:  &start,
			End:    &end,
			Source: sources,
		})

		assert.Nil(t, err, "expecting nil response")
		assert.NotNil(t, res, "expecting non-nil response")
		assert.Equal(t, 24, len(res.Hours), "unexpected hour point count")
		assert.NotNil(t, res.Meta, "expecting non-nil response")
		assert.Equal(t, res.Meta.Cost, 1, "unexpected meta value for cost")
		assert.Equal(t, start.Format("2006-01-02 15:04"), res.Meta.Start, "unexpected meta value for start")
		assert.Equal(t, end.Format("2006-01-02 15:04"), res.Meta.End, "unexpected meta value for end")
		assert.Equal(t, lat, res.Meta.Lat, "unexpected meta value for latitude")
		assert.Equal(t, lng, res.Meta.Lng, "unexpected meta value for longitude")
		for _, h := range res.Hours {
			assert.NotNil(t, h.AirTemperature, "expected air temperature to be set")
			assert.NotNil(t, h.WaveDirection, "expected wave direction to be set")
			assert.NotNil(t, h.WaterTemperature, "expected wave temperature to be set")
		}
	})
}
