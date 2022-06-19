package stormglass

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Points struct {
	Hours []Hours `json:"hours,omitempty"`
	Meta  Meta    `json:"meta,omitempty"`
}
type AirTemperature struct {
	Dwd  float64 `json:"dwd,omitempty"`
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type AirTemperature1000Hpa struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type AirTemperature100M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type AirTemperature200Hpa struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type AirTemperature500Hpa struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type AirTemperature800Hpa struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type AirTemperature80M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type CloudCover struct {
	Dwd  float64 `json:"dwd,omitempty"`
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type CurrentDirection struct {
	Meto float64 `json:"meto,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type CurrentSpeed struct {
	Meto float64 `json:"meto,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type Gust struct {
	Dwd  float64 `json:"dwd,omitempty"`
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type Humidity struct {
	Dwd  float64 `json:"dwd,omitempty"`
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type IceCover struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type Precipitation struct {
	Dwd  float64 `json:"dwd,omitempty"`
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type Pressure struct {
	Dwd  float64 `json:"dwd,omitempty"`
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type SeaLevel struct {
	Meto float64 `json:"meto,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type SecondarySwellDirection struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type SecondarySwellHeight struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type SecondarySwellPeriod struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type SnowDepth struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type SwellDirection struct {
	Dwd   float64 `json:"dwd,omitempty"`
	Icon  float64 `json:"icon,omitempty"`
	Meteo float64 `json:"meteo,omitempty"`
	Noaa  float64 `json:"noaa,omitempty"`
	Sg    float64 `json:"sg,omitempty"`
}
type SwellHeight struct {
	Dwd   float64 `json:"dwd,omitempty"`
	Icon  float64 `json:"icon,omitempty"`
	Meteo float64 `json:"meteo,omitempty"`
	Noaa  float64 `json:"noaa,omitempty"`
	Sg    float64 `json:"sg,omitempty"`
}
type SwellPeriod struct {
	Dwd   float64 `json:"dwd,omitempty"`
	Icon  float64 `json:"icon,omitempty"`
	Meteo float64 `json:"meteo,omitempty"`
	Noaa  float64 `json:"noaa,omitempty"`
	Sg    float64 `json:"sg,omitempty"`
}
type Visibility struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WaterTemperature struct {
	Meto float64 `json:"meto,omitempty"`
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WaveDirection struct {
	Icon  float64 `json:"icon,omitempty"`
	Meteo float64 `json:"meteo,omitempty"`
	Noaa  float64 `json:"noaa,omitempty"`
	Sg    float64 `json:"sg,omitempty"`
}
type WaveHeight struct {
	Dwd   float64 `json:"dwd,omitempty"`
	Icon  float64 `json:"icon,omitempty"`
	Meteo float64 `json:"meteo,omitempty"`
	Noaa  float64 `json:"noaa,omitempty"`
	Sg    float64 `json:"sg,omitempty"`
}
type WavePeriod struct {
	Icon  float64 `json:"icon,omitempty"`
	Meteo float64 `json:"meteo,omitempty"`
	Noaa  float64 `json:"noaa,omitempty"`
	Sg    float64 `json:"sg,omitempty"`
}
type WindDirection struct {
	Icon float64 `json:"icon,omitempty"`
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindDirection1000Hpa struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindDirection100M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindDirection200Hpa struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindDirection20M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindDirection30M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindDirection40M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindDirection500Hpa struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindDirection50M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindDirection800Hpa struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindDirection80M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindSpeed struct {
	Icon float64 `json:"icon,omitempty"`
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindSpeed1000Hpa struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindSpeed100M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindSpeed200Hpa struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindSpeed20M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindSpeed30M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindSpeed40M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindSpeed500Hpa struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindSpeed50M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindSpeed800Hpa struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindSpeed80M struct {
	Noaa float64 `json:"noaa,omitempty"`
	Sg   float64 `json:"sg,omitempty"`
}
type WindWaveDirection struct {
	Dwd   float64 `json:"dwd,omitempty"`
	Icon  float64 `json:"icon,omitempty"`
	Meteo float64 `json:"meteo,omitempty"`
	Noaa  float64 `json:"noaa,omitempty"`
	Sg    float64 `json:"sg,omitempty"`
}
type WindWaveHeight struct {
	Dwd   float64 `json:"dwd,omitempty"`
	Icon  float64 `json:"icon,omitempty"`
	Meteo float64 `json:"meteo,omitempty"`
	Noaa  float64 `json:"noaa,omitempty"`
	Sg    float64 `json:"sg,omitempty"`
}
type WindWavePeriod struct {
	Dwd   float64 `json:"dwd,omitempty"`
	Icon  float64 `json:"icon,omitempty"`
	Meteo float64 `json:"meteo,omitempty"`
	Noaa  float64 `json:"noaa,omitempty"`
	Sg    float64 `json:"sg,omitempty"`
}
type Hours struct {
	AirTemperature          *AirTemperature          `json:"airTemperature,omitempty"`
	AirTemperature1000Hpa   *AirTemperature1000Hpa   `json:"airTemperature1000hpa,omitempty"`
	AirTemperature100M      *AirTemperature100M      `json:"airTemperature100m,omitempty"`
	AirTemperature200Hpa    *AirTemperature200Hpa    `json:"airTemperature200hpa,omitempty"`
	AirTemperature500Hpa    *AirTemperature500Hpa    `json:"airTemperature500hpa,omitempty"`
	AirTemperature800Hpa    *AirTemperature800Hpa    `json:"airTemperature800hpa,omitempty"`
	AirTemperature80M       *AirTemperature80M       `json:"airTemperature80m,omitempty"`
	CloudCover              *CloudCover              `json:"cloudCover,omitempty"`
	CurrentDirection        *CurrentDirection        `json:"currentDirection,omitempty"`
	CurrentSpeed            *CurrentSpeed            `json:"currentSpeed,omitempty"`
	Gust                    *Gust                    `json:"gust,omitempty"`
	Humidity                *Humidity                `json:"humidity,omitempty"`
	IceCover                *IceCover                `json:"iceCover,omitempty"`
	Precipitation           *Precipitation           `json:"precipitation,omitempty"`
	Pressure                *Pressure                `json:"pressure,omitempty"`
	SeaLevel                *SeaLevel                `json:"seaLevel,omitempty"`
	SecondarySwellDirection *SecondarySwellDirection `json:"secondarySwellDirection,omitempty"`
	SecondarySwellHeight    *SecondarySwellHeight    `json:"secondarySwellHeight,omitempty"`
	SecondarySwellPeriod    *SecondarySwellPeriod    `json:"secondarySwellPeriod,omitempty"`
	SnowDepth               *SnowDepth               `json:"snowDepth,omitempty"`
	SwellDirection          *SwellDirection          `json:"swellDirection,omitempty"`
	SwellHeight             *SwellHeight             `json:"swellHeight,omitempty"`
	SwellPeriod             *SwellPeriod             `json:"swellPeriod,omitempty"`
	Time                    *time.Time               `json:"time,omitempty"`
	Visibility              *Visibility              `json:"visibility,omitempty"`
	WaterTemperature        *WaterTemperature        `json:"waterTemperature,omitempty"`
	WaveDirection           *WaveDirection           `json:"waveDirection,omitempty"`
	WaveHeight              *WaveHeight              `json:"waveHeight,omitempty"`
	WavePeriod              *WavePeriod              `json:"wavePeriod,omitempty"`
	WindDirection           *WindDirection           `json:"windDirection,omitempty"`
	WindDirection1000Hpa    *WindDirection1000Hpa    `json:"windDirection1000hpa,omitempty"`
	WindDirection100M       *WindDirection100M       `json:"windDirection100m,omitempty"`
	WindDirection200Hpa     *WindDirection200Hpa     `json:"windDirection200hpa,omitempty"`
	WindDirection20M        *WindDirection20M        `json:"windDirection20m,omitempty"`
	WindDirection30M        *WindDirection30M        `json:"windDirection30m,omitempty"`
	WindDirection40M        *WindDirection40M        `json:"windDirection40m,omitempty"`
	WindDirection500Hpa     *WindDirection500Hpa     `json:"windDirection500hpa,omitempty"`
	WindDirection50M        *WindDirection50M        `json:"windDirection50m,omitempty"`
	WindDirection800Hpa     *WindDirection800Hpa     `json:"windDirection800hpa,omitempty"`
	WindDirection80M        *WindDirection80M        `json:"windDirection80m,omitempty"`
	WindSpeed               *WindSpeed               `json:"windSpeed,omitempty"`
	WindSpeed1000Hpa        *WindSpeed1000Hpa        `json:"windSpeed1000hpa,omitempty"`
	WindSpeed100M           *WindSpeed100M           `json:"windSpeed100m,omitempty"`
	WindSpeed200Hpa         *WindSpeed200Hpa         `json:"windSpeed200hpa,omitempty"`
	WindSpeed20M            *WindSpeed20M            `json:"windSpeed20m,omitempty"`
	WindSpeed30M            *WindSpeed30M            `json:"windSpeed30m,omitempty"`
	WindSpeed40M            *WindSpeed40M            `json:"windSpeed40m,omitempty"`
	WindSpeed500Hpa         *WindSpeed500Hpa         `json:"windSpeed500hpa,omitempty"`
	WindSpeed50M            *WindSpeed50M            `json:"windSpeed50m,omitempty"`
	WindSpeed800Hpa         *WindSpeed800Hpa         `json:"windSpeed800hpa,omitempty"`
	WindSpeed80M            *WindSpeed80M            `json:"windSpeed80m,omitempty"`
	WindWaveDirection       *WindWaveDirection       `json:"windWaveDirection,omitempty"`
	WindWaveHeight          *WindWaveHeight          `json:"windWaveHeight,omitempty"`
	WindWavePeriod          *WindWavePeriod          `json:"windWavePeriod,omitempty"`
}

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
}

func (p ParamsOptions) toList() []string {
	var params []string

	if p.Time {
		params = append(params, "time")
	}
	if p.AirTemperature {
		params = append(params, "airTemperature")
	}
	if p.AirTemperature80m {
		params = append(params, "airTemperature80m")
	}
	if p.AirTemperature100m {
		params = append(params, "airTemperature100m")
	}
	if p.AirTemperature1000hpa {
		params = append(params, "airTemperature1000hpa")
	}
	if p.AirTemperature800hpa {
		params = append(params, "airTemperature800hpa")
	}
	if p.AirTemperature500hpa {
		params = append(params, "airTemperature500hpa")
	}
	if p.AirTemperature200hpa {
		params = append(params, "airTemperature200hpa")
	}
	if p.Pressure {
		params = append(params, "pressure")
	}
	if p.CloudCover {
		params = append(params, "cloudCover")
	}
	if p.CurrentDirection {
		params = append(params, "currentDirection")
	}
	if p.CurrentSpeed {
		params = append(params, "currentSpeed")
	}
	if p.Gust {
		params = append(params, "gust")
	}
	if p.Humidity {
		params = append(params, "humidity")
	}
	if p.IceCover {
		params = append(params, "iceCover")
	}
	if p.Precipitation {
		params = append(params, "precipitation")
	}
	if p.SnowDepth {
		params = append(params, "snowDepth")
	}
	if p.SeaLevel {
		params = append(params, "seaLevel")
	}
	if p.SwellDirection {
		params = append(params, "swellDirection")
	}
	if p.SwellHeight {
		params = append(params, "swellHeight")
	}
	if p.SwellPeriod {
		params = append(params, "swellPeriod")
	}
	if p.SecondarySwellPeriod {
		params = append(params, "secondarySwellPeriod")
	}
	if p.SecondarySwellDirection {
		params = append(params, "secondarySwellDirection")
	}
	if p.SecondarySwellHeight {
		params = append(params, "secondarySwellHeight")
	}
	if p.Visibility {
		params = append(params, "visibility")
	}
	if p.WaterTemperature {
		params = append(params, "waterTemperature")
	}
	if p.WaveDirection {
		params = append(params, "waveDirection")
	}

	return params
}

// SourcesOptions : https://docs.stormglass.io/#/sources?id=available-sources
type SourcesOptions struct {
	ICON        bool
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

type PointsRequestOptions struct {
	Lng    float64        `json:"lng,omitempty"`
	Lat    float64        `json:"lat,omitempty"`
	Params ParamsOptions  `json:"params,omitempty"`
	Start  *time.Time     `json:"start,omitempty"`
	End    *time.Time     `json:"end,omitempty"`
	Source SourcesOptions `json:"sources,omitempty"`
}

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

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	res := Points{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
