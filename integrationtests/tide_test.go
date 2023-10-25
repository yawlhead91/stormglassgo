//go:build integration

package integrationtests

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/jinzhu/now"
	"github.com/stretchr/testify/assert"
	stormglass "github.com/yawlhead91/stormglassgo"
)

func Test_GetExtremesPoint(t *testing.T) {
	t.Run(" return api error", func(t *testing.T) {
		var start = now.BeginningOfDay()

		c := stormglass.NewClient("")

		ctx := context.Background()
		res, err := c.GetExtremesPoint(ctx, stormglass.ExtremesPointsRequestOptions{
			CommonRequestOptions: stormglass.CommonRequestOptions{
				Lat:   lat,
				Lng:   lng,
				Start: &start,
			},
			Datum: stormglass.MSL,
		})
		assert.Nil(t, res, "expecting nil response")
		assert.NotNil(t, err, "expecting non-nil error")
		assert.Equal(t, "403: key:API key is invalid", err.Error(), "unexpected error")
	})

	t.Run("success: return 24 hour points", func(t *testing.T) {
		tme := time.Now().In(time.UTC)
		var start = now.New(tme).BeginningOfDay()
		var end = now.New(tme).BeginningOfDay().Add(time.Hour * 24)

		c := stormglass.NewClient(os.Getenv("STORMGLASS_API_KEY"))

		ctx := context.Background()
		res, err := c.GetExtremesPoint(ctx, stormglass.ExtremesPointsRequestOptions{
			CommonRequestOptions: stormglass.CommonRequestOptions{
				Lat:   lat,
				Lng:   lng,
				Start: &start,
				End:   &end,
			},
			Datum: stormglass.MSL,
		})

		assert.Nil(t, err, "expecting nil response")
		require.NotNil(t, res, "expecting non-nil response")
		require.NotNil(t, res.Data, "expecting non-nil response hours")
		assert.NotEqual(t, 0, len(res.Data), "unexpected data count")
		assert.NotNil(t, res.Meta, "expecting non-nil response")
		assert.Equal(t, res.Meta.Cost, 1, "unexpected meta value for cost")
		assert.Equal(t, start.Format("2006-01-02 15:04"), res.Meta.Start, "unexpected meta value for start")
		assert.Equal(t, end.Format("2006-01-02 15:04"), res.Meta.End, "unexpected meta value for end")
		assert.Equal(t, lat, res.Meta.Lat, "unexpected meta value for latitude")
		assert.Equal(t, lng, res.Meta.Lng, "unexpected meta value for longitude")
		assert.Equal(t, stormglass.MSL, res.Meta.Datum, "unexpected meta value for longitude")
		assert.NotEqual(t, 0, res.Meta.Station.Distance, "expected some value for station distance")
		assert.NotEqual(t, "", res.Meta.Station.Name, "expected some value for station name")
		assert.NotEqual(t, "", res.Meta.Station.Source, "expected some value for station source")
		assert.NotEqual(t, 0, res.Meta.Station.Lat, "expected some value for station lat")
		assert.NotEqual(t, 0, res.Meta.Station.Lng, "expected some value for station lat")

		for _, h := range res.Data {
			assert.NotNil(t, h.Height, "expected height to be set")
			assert.NotNil(t, h.Time, "expected time to be set")
			assert.NotNil(t, h.Type, "expected type to be set")
		}
	})

}
