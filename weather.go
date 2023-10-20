package stormglass

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Points represents a Point request response.
type Points struct {
	Hours []Hour `json:"hours,omitempty"`
	Meta  Meta   `json:"meta,omitempty"`
}

// SourcesOptions : https://docs.stormglass.io/#/sources?id=available-sources
type SourcesOptions struct {
	ICON        bool
	DWD         bool
	NOAA        bool
	MeteoFrance bool
	UKMetOffice bool
	FCOO        bool
	FMI         bool
	YR          bool
	SMHI        bool
	StormGlass  bool
}

func (s SourcesOptions) toList() []string {
	var sources []string
	if s.ICON {
		sources = append(sources, "icon")
	}

	if s.DWD {
		sources = append(sources, "dwt")
	}

	if s.NOAA {
		sources = append(sources, "noaa")
	}

	if s.MeteoFrance {
		sources = append(sources, "meteo")
	}

	if s.UKMetOffice {
		sources = append(sources, "meto")
	}

	if s.FCOO {
		sources = append(sources, "fcoo")
	}

	if s.FMI {
		sources = append(sources, "fmi")
	}

	if s.YR {
		sources = append(sources, "yr")
	}

	if s.SMHI {
		sources = append(sources, "smhi")
	}

	if s.StormGlass {
		sources = append(sources, "sg")
	}

	return sources
}

// SourceValues represents potential source values response.
type SourceValues struct {
	ICON        *float64 `json:"icon,omitempty"`
	DWD         *float64 `json:"dwd,omitempty"`
	NOAA        *float64 `json:"noaa,omitempty"`
	MeteoFrance *float64 `json:"meteo,omitempty"`
	UKMetOffice *float64 `json:"meto,omitempty"`
	FCOO        *float64 `json:"fcoo,omitempty"`
	FMI         *float64 `json:"fmi,omitempty"`
	YR          *float64 `json:"yr,omitempty"`
	SMHI        *float64 `json:"smhi,omitempty"`
	StormGlass  *float64 `json:"sg,omitempty"`
}

// Hour represents an hour data point response.
type Hour struct {
	AirTemperature          *SourceValues `json:"airTemperature,omitempty"`
	AirTemperature1000Hpa   *SourceValues `json:"airTemperature1000hpa,omitempty"`
	AirTemperature100M      *SourceValues `json:"airTemperature100m,omitempty"`
	AirTemperature200Hpa    *SourceValues `json:"airTemperature200hpa,omitempty"`
	AirTemperature500Hpa    *SourceValues `json:"airTemperature500hpa,omitempty"`
	AirTemperature800Hpa    *SourceValues `json:"airTemperature800hpa,omitempty"`
	AirTemperature80M       *SourceValues `json:"airTemperature80m,omitempty"`
	CloudCover              *SourceValues `json:"cloudCover,omitempty"`
	CurrentDirection        *SourceValues `json:"currentDirection,omitempty"`
	CurrentSpeed            *SourceValues `json:"currentSpeed,omitempty"`
	Gust                    *SourceValues `json:"gust,omitempty"`
	Humidity                *SourceValues `json:"humidity,omitempty"`
	IceCover                *SourceValues `json:"iceCover,omitempty"`
	Precipitation           *SourceValues `json:"precipitation,omitempty"`
	Pressure                *SourceValues `json:"pressure,omitempty"`
	SeaLevel                *SourceValues `json:"seaLevel,omitempty"`
	SecondarySwellDirection *SourceValues `json:"secondarySwellDirection,omitempty"`
	SecondarySwellHeight    *SourceValues `json:"secondarySwellHeight,omitempty"`
	SecondarySwellPeriod    *SourceValues `json:"secondarySwellPeriod,omitempty"`
	SnowDepth               *SourceValues `json:"snowDepth,omitempty"`
	SwellDirection          *SourceValues `json:"swellDirection,omitempty"`
	SwellHeight             *SourceValues `json:"swellHeight,omitempty"`
	SwellPeriod             *SourceValues `json:"swellPeriod,omitempty"`
	Time                    *time.Time    `json:"time,omitempty"`
	Visibility              *SourceValues `json:"visibility,omitempty"`
	WaterTemperature        *SourceValues `json:"waterTemperature,omitempty"`
	WaveDirection           *SourceValues `json:"waveDirection,omitempty"`
	WaveHeight              *SourceValues `json:"waveHeight,omitempty"`
	WavePeriod              *SourceValues `json:"wavePeriod,omitempty"`
	WindDirection           *SourceValues `json:"windDirection,omitempty"`
	WindDirection1000Hpa    *SourceValues `json:"windDirection1000hpa,omitempty"`
	WindDirection100M       *SourceValues `json:"windDirection100m,omitempty"`
	WindDirection200Hpa     *SourceValues `json:"windDirection200hpa,omitempty"`
	WindDirection20M        *SourceValues `json:"windDirection20m,omitempty"`
	WindDirection30M        *SourceValues `json:"windDirection30m,omitempty"`
	WindDirection40M        *SourceValues `json:"windDirection40m,omitempty"`
	WindDirection500Hpa     *SourceValues `json:"windDirection500hpa,omitempty"`
	WindDirection50M        *SourceValues `json:"windDirection50m,omitempty"`
	WindDirection800Hpa     *SourceValues `json:"windDirection800hpa,omitempty"`
	WindDirection80M        *SourceValues `json:"windDirection80m,omitempty"`
	WindSpeed               *SourceValues `json:"windSpeed,omitempty"`
	WindSpeed1000Hpa        *SourceValues `json:"windSpeed1000hpa,omitempty"`
	WindSpeed100M           *SourceValues `json:"windSpeed100m,omitempty"`
	WindSpeed200Hpa         *SourceValues `json:"windSpeed200hpa,omitempty"`
	WindSpeed20M            *SourceValues `json:"windSpeed20m,omitempty"`
	WindSpeed30M            *SourceValues `json:"windSpeed30m,omitempty"`
	WindSpeed40M            *SourceValues `json:"windSpeed40m,omitempty"`
	WindSpeed500Hpa         *SourceValues `json:"windSpeed500hpa,omitempty"`
	WindSpeed50M            *SourceValues `json:"windSpeed50m,omitempty"`
	WindSpeed800Hpa         *SourceValues `json:"windSpeed800hpa,omitempty"`
	WindSpeed80M            *SourceValues `json:"windSpeed80m,omitempty"`
	WindWaveDirection       *SourceValues `json:"windWaveDirection,omitempty"`
	WindWaveHeight          *SourceValues `json:"windWaveHeight,omitempty"`
	WindWavePeriod          *SourceValues `json:"windWavePeriod,omitempty"`
}

// Meta data from Point request.
type Meta struct {
	Cost         int      `json:"cost,omitempty"`
	DailyQuota   int      `json:"dailyQuota,omitempty"`
	End          string   `json:"end,omitempty"`
	Lat          float64  `json:"lat,omitempty"`
	Lng          float64  `json:"lng,omitempty"`
	Params       []string `json:"params,omitempty"`
	RequestCount int      `json:"requestCount,omitempty"`
	Start        string   `json:"start,omitempty"`
}

// ParamsOptions holds optional parameters.
type ParamsOptions struct {
	Time                    bool
	AirTemperature          bool
	AirTemperature80m       bool
	AirTemperature100m      bool
	AirTemperature1000hpa   bool
	AirTemperature800hpa    bool
	AirTemperature500hpa    bool
	AirTemperature200hpa    bool
	Pressure                bool
	CloudCover              bool
	CurrentDirection        bool
	CurrentSpeed            bool
	Gust                    bool
	Humidity                bool
	IceCover                bool
	Precipitation           bool
	SnowDepth               bool
	SeaLevel                bool
	SwellDirection          bool
	SwellHeight             bool
	SwellPeriod             bool
	SecondarySwellPeriod    bool
	SecondarySwellDirection bool
	SecondarySwellHeight    bool
	Visibility              bool
	WaterTemperature        bool
	WaveDirection           bool
	WaveHeight              bool
	WavePeriod              bool
	WindDirection           bool
	WindDirection1000Hpa    bool
	WindDirection100M       bool
	WindDirection200Hpa     bool
	WindDirection20M        bool
	WindDirection30M        bool
	WindDirection40M        bool
	WindDirection500Hpa     bool
	WindDirection50M        bool
	WindDirection800Hpa     bool
	WindDirection80M        bool
	WindSpeed               bool
	WindSpeed1000Hpa        bool
	WindSpeed100M           bool
	WindSpeed200Hpa         bool
	WindSpeed20M            bool
	WindSpeed30M            bool
	WindSpeed40M            bool
	WindSpeed500Hpa         bool
	WindSpeed50M            bool
	WindSpeed800Hpa         bool
	WindSpeed80M            bool
	WindWaveDirection       bool
	WindWaveHeight          bool
	WindWavePeriod          bool
}

func (p ParamsOptions) toList() []string {
	var params []string

	b, _ := json.Marshal(&p)
	var m map[string]bool
	_ = json.Unmarshal(b, &m)

	for k, v := range m {
		if v {
			params = append(params, k)
		}
	}

	return params
}

// PointsRequestOptions for available query paramters.
type PointsRequestOptions struct {
	Lng    float64        `json:"lng,omitempty"`
	Lat    float64        `json:"lat,omitempty"`
	Params ParamsOptions  `json:"params,omitempty"`
	Start  *time.Time     `json:"start,omitempty"`
	End    *time.Time     `json:"end,omitempty"`
	Source SourcesOptions `json:"sources,omitempty"`
}

// GetPoint sends a Point request https://docs.stormglass.io/#/weather?id=point-request.
func (c *Client) GetPoint(ctx context.Context, options PointsRequestOptions) (*Points, error) {
	endpoint := fmt.Sprintf("%s/weather/point?lat=%f&lng=%f", c.BaseURL, options.Lat, options.Lng)

	params := options.Params.toList()
	if len(params) > 0 {
		endpoint = fmt.Sprintf("%s&params=%s", endpoint, strings.Join(params, ","))
	}

	if options.Start != nil {
		endpoint = fmt.Sprintf("%s&start=%d", endpoint, options.Start.Unix())
	}

	if options.End != nil {
		endpoint = fmt.Sprintf("%s&end=%d", endpoint, options.End.Unix())
	}

	sources := options.Source.toList()
	if len(sources) > 0 {
		endpoint = fmt.Sprintf("%s&source=%s", endpoint, strings.Join(sources, ","))
	}

	req, err := http.NewRequest("GET", endpoint, http.NoBody)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	res := Points{}

	if err = c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
