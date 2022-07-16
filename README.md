# Stormglass.io V2 GO client

![tests](https://github.com/yawlhead91/stormglassgo/actions/workflows/integration-test/badge.svg)

Go client for the stormglass.io v2 REST API see [documentation](https://docs.stormglass.io/#/) for more information.

This library is currently a work in progress. See the [status](#status) section for more information and what endpoints are implemented.

## Installation

```
go get github.com/yawlhead91/stormglassgo
```

## Usage
```go
package main

import (
	"context"
	"log"
	"time"

	stormglass "github.com/yawlhead91/stormglassgo"
)

const (
	lat = 58.7984
	lng = 17.8081
	key = "..."
)

func main() {

	start := time.Now()
	end := time.Now().Add(time.Hour)

	client := stormglass.NewClient(key)

	ctx := context.Background()
	points, err := client.GetPoint(ctx, stormglass.PointsRequestOptions{
		Lat: lat,
		Lng: lng,
		Params: stormglass.ParamsOptions{
			AirTemperature: true,
		},
		Start: &start,
		End:   &end,
		Source: stormglass.SourcesOptions{
			ICON: true,
		},
	})
	if err != nil {
		log.Fatalf("get weather points: %v", err)
	}

	for _, p := range points.Hours {
		log.Printf("%+v", p)
	}
}
```

## Status

#### Weather

- [x] GET	/weather/point

#### Bio

- [ ] GET	/bio/point

#### Tide

- [ ] GET	/tide/extremes/point
- [ ] GET	/tide/sea-level/point
- [ ] GET	/tide/sea-level/stations
- [ ] GET	/tide/sea-level/stations/area

### Astronomy

- [ ] GET	/tide/astronomy/point

### Solar

- [ ] GET	/tide/solar/point

### Elevation

- [ ] GET	/tide/elevation/point


## Contributing

All contributions are welcome following the existing code style and conventions and submit a PR.





