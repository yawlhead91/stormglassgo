package stormglass

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// Points represents a Point request response.
type Points struct {
	Hours []Hour `json:"hours,omitempty"`
	Meta  Meta   `json:"meta,omitempty"`
}

// WeatherSourcesOptions : https://docs.stormglass.io/#/sources?id=available-sources
type WeatherSourcesOptions struct {
	ICON        bool
	DWD         bool
	NOAA        bool
	MeteoFrance bool
	FCOO        bool
	FMI         bool
	YR          bool
	SMHI        bool
	StormGlass  bool
	UKMetOffice bool
}

func (s WeatherSourcesOptions) toList() []string {
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

// WeatherSourceValues represents potential source values response.
type WeatherSourceValues struct {
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
	AirTemperature          *WeatherSourceValues `json:"airTemperature,omitempty"`
	AirTemperature1000Hpa   *WeatherSourceValues `json:"airTemperature1000hpa,omitempty"`
	AirTemperature100M      *WeatherSourceValues `json:"airTemperature100m,omitempty"`
	AirTemperature200Hpa    *WeatherSourceValues `json:"airTemperature200hpa,omitempty"`
	AirTemperature500Hpa    *WeatherSourceValues `json:"airTemperature500hpa,omitempty"`
	AirTemperature800Hpa    *WeatherSourceValues `json:"airTemperature800hpa,omitempty"`
	AirTemperature80M       *WeatherSourceValues `json:"airTemperature80m,omitempty"`
	CloudCover              *WeatherSourceValues `json:"cloudCover,omitempty"`
	CurrentDirection        *WeatherSourceValues `json:"currentDirection,omitempty"`
	CurrentSpeed            *WeatherSourceValues `json:"currentSpeed,omitempty"`
	Gust                    *WeatherSourceValues `json:"gust,omitempty"`
	Humidity                *WeatherSourceValues `json:"humidity,omitempty"`
	IceCover                *WeatherSourceValues `json:"iceCover,omitempty"`
	Precipitation           *WeatherSourceValues `json:"precipitation,omitempty"`
	Pressure                *WeatherSourceValues `json:"pressure,omitempty"`
	SeaLevel                *WeatherSourceValues `json:"seaLevel,omitempty"`
	SecondarySwellDirection *WeatherSourceValues `json:"secondarySwellDirection,omitempty"`
	SecondarySwellHeight    *WeatherSourceValues `json:"secondarySwellHeight,omitempty"`
	SecondarySwellPeriod    *WeatherSourceValues `json:"secondarySwellPeriod,omitempty"`
	SnowDepth               *WeatherSourceValues `json:"snowDepth,omitempty"`
	SwellDirection          *WeatherSourceValues `json:"swellDirection,omitempty"`
	SwellHeight             *WeatherSourceValues `json:"swellHeight,omitempty"`
	SwellPeriod             *WeatherSourceValues `json:"swellPeriod,omitempty"`
	Time                    *time.Time           `json:"time,omitempty"`
	Visibility              *WeatherSourceValues `json:"visibility,omitempty"`
	WaterTemperature        *WeatherSourceValues `json:"waterTemperature,omitempty"`
	WaveDirection           *WeatherSourceValues `json:"waveDirection,omitempty"`
	WaveHeight              *WeatherSourceValues `json:"waveHeight,omitempty"`
	WavePeriod              *WeatherSourceValues `json:"wavePeriod,omitempty"`
	WindDirection           *WeatherSourceValues `json:"windDirection,omitempty"`
	WindDirection1000Hpa    *WeatherSourceValues `json:"windDirection1000hpa,omitempty"`
	WindDirection100M       *WeatherSourceValues `json:"windDirection100m,omitempty"`
	WindDirection200Hpa     *WeatherSourceValues `json:"windDirection200hpa,omitempty"`
	WindDirection20M        *WeatherSourceValues `json:"windDirection20m,omitempty"`
	WindDirection30M        *WeatherSourceValues `json:"windDirection30m,omitempty"`
	WindDirection40M        *WeatherSourceValues `json:"windDirection40m,omitempty"`
	WindDirection500Hpa     *WeatherSourceValues `json:"windDirection500hpa,omitempty"`
	WindDirection50M        *WeatherSourceValues `json:"windDirection50m,omitempty"`
	WindDirection800Hpa     *WeatherSourceValues `json:"windDirection800hpa,omitempty"`
	WindDirection80M        *WeatherSourceValues `json:"windDirection80m,omitempty"`
	WindSpeed               *WeatherSourceValues `json:"windSpeed,omitempty"`
	WindSpeed1000Hpa        *WeatherSourceValues `json:"windSpeed1000hpa,omitempty"`
	WindSpeed100M           *WeatherSourceValues `json:"windSpeed100m,omitempty"`
	WindSpeed200Hpa         *WeatherSourceValues `json:"windSpeed200hpa,omitempty"`
	WindSpeed20M            *WeatherSourceValues `json:"windSpeed20m,omitempty"`
	WindSpeed30M            *WeatherSourceValues `json:"windSpeed30m,omitempty"`
	WindSpeed40M            *WeatherSourceValues `json:"windSpeed40m,omitempty"`
	WindSpeed500Hpa         *WeatherSourceValues `json:"windSpeed500hpa,omitempty"`
	WindSpeed50M            *WeatherSourceValues `json:"windSpeed50m,omitempty"`
	WindSpeed800Hpa         *WeatherSourceValues `json:"windSpeed800hpa,omitempty"`
	WindSpeed80M            *WeatherSourceValues `json:"windSpeed80m,omitempty"`
	WindWaveDirection       *WeatherSourceValues `json:"windWaveDirection,omitempty"`
	WindWaveHeight          *WeatherSourceValues `json:"windWaveHeight,omitempty"`
	WindWavePeriod          *WeatherSourceValues `json:"windWavePeriod,omitempty"`
}

// WeatherParamsOptions holds optional parameters.
type WeatherParamsOptions struct {
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

func (p WeatherParamsOptions) toList() []string {
	var params []string

	b, _ := json.Marshal(&p)
	var m map[string]bool
	_ = json.Unmarshal(b, &m)

	for k, v := range m {
		if v {
			params = append(params, firstToLower(k))
		}
	}

	sort.Slice(params, func(i, j int) bool {
		return params[i] < params[j]
	})

	return params
}

func firstToLower(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}
	lc := unicode.ToLower(r)
	if r == lc {
		return s
	}
	return string(lc) + s[size:]
}

// PointsRequestOptions for available query parameters.
type PointsRequestOptions struct {
	CommonRequestOptions
	Params WeatherParamsOptions  `json:"params,omitempty"`
	Source WeatherSourcesOptions `json:"sources,omitempty"`
}

// GetPoint sends a Point request https://docs.stormglass.io/#/weather?id=point-request.
func (c *Client) GetPoint(ctx context.Context, options PointsRequestOptions) (*Points, error) {
	path, err := url.JoinPath(c.BaseURL, "weather", "point")
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	values := url.Values{}

	values.Add("lat", fmt.Sprintf("%f", options.Lat))
	values.Add("lng", fmt.Sprintf("%f", options.Lng))

	params := options.Params.toList()
	if len(params) > 0 {
		values.Add("params", strings.Join(params, ","))
	}

	if options.Start != nil {
		values.Add("start", fmt.Sprintf("%d", options.Start.Unix()))
	}

	if options.End != nil {
		values.Add("end", fmt.Sprintf("%d", options.End.Unix()))
	}

	sources := options.Source.toList()
	if len(sources) > 0 {
		values.Add("source", strings.Join(sources, ","))
	}

	u.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", u.String(), http.NoBody)
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
