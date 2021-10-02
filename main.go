package gpx

import (
	"encoding/xml"
)

func Unmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

func Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(v, prefix, indent)
}

func CreateGpx(m Metadata) Gpx {
	g := Gpx{}
	g.Creator = "github.com/cconcannon/gpx"
	g.Version = "1.1"
	return g
}
