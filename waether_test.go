// +build integration

package stormglass_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jinzhu/now"

	"github.com/stretchr/testify/assert"
	stormglass "github.com/yawlhead91/stormglassgo"
)

const (
	lat = 58.7984
	lng = 17.8081
)

func TestGetPoint(t *testing.T) {
	var start = now.BeginningOfDay()
	var end = now.BeginningOfDay().Add(time.Hour * 23)

	c := stormglass.NewClient(os.Getenv("STORMGLASS_API_KEY"))

	ctx := context.Background()
	res, err := c.GetPoint(ctx, stormglass.PointsRequestOptions{
		Lat: lat,
		Lng: lng,
		Params: stormglass.ParamsOptions{
			AirTemperature: true,
		},
		Start: &start,
		End:   &end,
	})

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.Equal(t, len(res.Hours), 24, "hours should be 24 hours long")
	//assert.Equal(t, "integration_face_id", res.Faces[0].FaceID, "expecting correct face_id")
	//assert.NotEmpty(t, res.Faces[0].FaceToken, "expecting non-empty face_token")
	//assert.Greater(t, len(res.Faces[0].FaceImages), 0, "expecting non-empty face_images")
}
