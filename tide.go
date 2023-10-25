package stormglass

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ExtremesPointStation represents a tide station.
type ExtremesPointStation struct {
	Distance float64 `json:"distance,omitempty"`
	Lat      float64 `json:"lat,omitempty"`
	Lan      float64 `json:"lan,omitempty"`
	Name     string  `json:"name,omitempty"`
	Source   string  `json:"source,omitempty"`
}

// ExtremesPointMeta represents the meta data from the extremes point request.
type ExtremesPointMeta struct {
	Meta
	Station ExtremesPointStation `json:"station,omitempty"`
}

// ExtremesPoint represents an extreme point.
type ExtremesPoint struct {
	Height float64   `json:"height,omitempty"`
	Time   time.Time `json:"time,omitempty"`
	Type   string    `json:"type,omitempty"`
}

// ExtremesPoints represents the extremes points request response.
type ExtremesPoints struct {
	Data []ExtremesPoint   `json:"data,omitempty"`
	Meta ExtremesPointMeta `json:"meta,omitempty"`
}

// ExtremesPointsRequestOptions represents the options for the extremes points request.
type ExtremesPointsRequestOptions struct {
	CommonRequestOptions
	Datum ExtremesPointsDatumOption `json:"datum,omitempty"`
}

// ExtremesPointsDatumOption represents the datum option for the extremes points request.
type ExtremesPointsDatumOption string

// Datum options for the extremes points request.
const (
	MLLW ExtremesPointsDatumOption = "mllw"
	MSL  ExtremesPointsDatumOption = "msl"
)

// GetExtremesPoint send an extreme point request: https://docs.stormglass.io/#/tide?id=extremes-point-request
func (c *Client) GetExtremesPoint(ctx context.Context, options ExtremesPointsRequestOptions) (*ExtremesPoints, error) {
	path, err := url.JoinPath(c.BaseURL, "tide", "extremes", "point")
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	values := u.Query()
	values.Add("lat", fmt.Sprintf("%f", options.Lat))
	values.Add("lng", fmt.Sprintf("%f", options.Lng))

	if options.Start != nil {
		values.Add("start", fmt.Sprintf("%d", options.Start.Unix()))
	}

	if options.End != nil {
		values.Add("end", fmt.Sprintf("%d", options.End.Unix()))
	}

	if options.Datum != "" {
		values.Add("datum", string(options.Datum))
	}

	u.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", u.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	res := ExtremesPoints{}

	if err = c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
