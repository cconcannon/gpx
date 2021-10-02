# github.com/cconcannon/gpx

A Go module for interfacing with `.gpx` files.

GPX stands for the "GPS Exchange Format" and the [current `.gpx` file spec can be found here](https://www.topografix.com/gpx.asp).

## Use

`go get github.com/cconcannon/gpx`

```go
include (
	...

	"github.com/cconcannon/gpx"
)
```

### Read a GPX file:

```go
gpxFile, err := os.Open(file)

if err != nil {
	fmt.Println(err)
}

defer gpxFile.Close()

data, err := ioutil.ReadAll(gpxFile)

var g gpx.Gpx

err = gpx.Unmarshal(data, &g)

if err != nil {
	fmt.Println(err)
}

// access your data via g.Tracks...
```

## Get Points from a Track

```go
tracks := g.Tracks
segments := (*tracks)[0].TrackSegments
trackPoints := (*segments)[0].TrackPoints

for i, point := range *trackPoints {
	fmt.Printf("\nPoint %v:\nLatitude: %v\nLongitude: %v\n", i, point.Lat, point.Lon)
}
```

### Create a GPX file:

```go
m := gpx.Metadata{
	Name: "data identifier"
	Author: "your name"
	...
}

// Create empty GPX data structure with metadata
g := gpx.CreateGpx(m)

// Add track data
g.Tracks = ...

```