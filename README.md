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
```

### Get Points from a Track

```go
tracks := g.GetTracks()
segments := tracks[0].GetSegments()
points := segments[0].GetTrackPoints()

for _, point := range points {
	fmt.Printf("\nPoint %v:\nLatitude: %v\nLongitude: %v\n", i, point.Lat, point.Lon)
}
```