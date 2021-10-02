package gpx

import (
	"encoding/xml"
)

// Gpx is the root element in the XML file.
type Gpx struct {
	XMLName    xml.Name    `xml:"gpx"`
	Version    string      `xml:"version,attr"`
	Creator    string      `xml:"creator,attr"`
	Metadata   *Metadata   `xml:"metadata,omitempty"`
	Waypoints  *[]Wpt      `xml:"wpt,omitempty"`
	Routes     *[]Rte      `xml:"rte,omitempty"`
	Tracks     *[]Trk      `xml:"trk,omitempty"`
	Extensions *Extensions `xml:"extensions,omitempty"`
}

// Metadata contains information about the GPX file, author, and copyright restrictions goes in the
// metadata section. Providing rich, meaningful information about your GPX files allows others to
// search for and use your GPS data.
type Metadata struct {
	XMLName    xml.Name    `xml:"metadata"`
	Name       string      `xml:"nameomitempty"`
	Desc       string      `xml:"desc,omitempty"`
	Author     *Person     `xml:"author,omitempty"`
	Copyright  *Copyright  `xml:"copyright,omitempty"`
	Links      *[]Link     `xml:"link,omitempty"`
	Time       string      `xml:"time,omitempty"`
	Keywords   string      `xml:"keywords,omitempty"`
	Bounds     *Bounds     `xml:"bounds,omitempty"`
	Extensions *Extensions `xml:"extensions,omitempty"`
}

// Wpt represents a waypoint, point of interest, or named feature on a map.

type Wpt struct {
	XMLName       xml.Name    `xml:"trkpt"`
	Elevation     float64     `xml:"ele,omitempty"`
	Time          string      `xml:"time,omitempty"`
	Magvar        float64     `xml:"magvar,omitempty"`
	GeoidHeight   float64     `xml:"geoidheight,omitempty"`
	Name          string      `xml:"name,omitempty"`
	Cmt           string      `xml:"cmt,omitempty"`
	Desc          string      `xml:"desc,omitempty"`
	Src           string      `xml:"src,omitempty"`
	Links         *[]Link     `xml:"link,omitempty"`
	Sym           string      `xml:"sym,omitempty"`
	Type          string      `xml:"type,omitempty"`
	Fix           string      `xml:"fix,omitempty"`
	Sat           int         `xml:"sat,omitempty"`
	Hdop          float64     `xml:"hdop,omitempty"`
	Vdop          float64     `xml:"vdop,omitempty"`
	Pdop          float64     `xml:"pdop,omitempty"`
	AgeOfDgpsData float64     `xml:"ageofdgpsdata,omitempty"`
	DgpsId        int         `xml:"dgpsid,omitempty"`
	Extensions    *Extensions `xml:"extensions,omitempty"`
	Lat           float64     `xml:"lat,attr"`
	Lon           float64     `xml:"lon,attr"`
}

// Rte represents a route - an ordered list of waypoints representing a series of turn points
// leading to a destination
type Rte struct {
	XMLName     xml.Name    `xml:"rte"`
	Name        string      `xml:"name,omitempty"`
	Cmt         string      `xml:"cmt,omitempty"`
	Desc        string      `xml:"desc,omitempty"`
	Src         string      `xml:"src,omitempty"`
	Links       *[]Link     `xml:"link,omitempty"`
	Number      int         `xml:"number,omitempty"`
	Type        string      `xml:"type,omitempty"`
	Extensions  *Extensions `xml:"extensions,omitempty"`
	RoutePoints *[]Wpt      `xml:"rtept,omitempty"`
}

// Trk represents a track - an ordered list of points describing a path
type Trk struct {
	XMLName       xml.Name    `xml:"trk"`
	Name          string      `xml:"name,omitempty"`
	Cmt           string      `xml:"cmt,omitempty"`
	Desc          string      `xml:"desc,omitempty"`
	Src           string      `xml:"src,omitempty"`
	Links         *[]Link     `xml:"link,omitempty"`
	Number        int         `xml:"number,omitempty"`
	Type          string      `xml:"type,omitempty"`
	Extensions    *Extensions `xml:"extensions,omitempty"`
	TrackSegments *[]Trkseg   `xml:"trkseg,omitempty"`
}

// TrkSeg represents a track segment, which holds a list of Track Points which are logically
// connected in order. To represent a single GPS track where GPS reception was lost, or the GPS
// receiver was turned off, start a new Track Segment for each continuous span of track data.
type Trkseg struct {
	XMLName     xml.Name `xml:"trkseg"`
	TrackPoints *[]Wpt   `xml:"trkpt,omitempty"`
}

// Extensions can be used to extend the GPX format
type Extensions interface{}

// Person represents a person or organization
type Person struct {
	XMLName xml.Name `xml:"author"`
	Name    string   `xml:"name,omitempty"`
	Email   *Email   `xml:"email,omitempty"`
	Link    *Link    `xml:"link,omitempty"`
}

// Link represents a link to an external resource (Web page, digital photo, video clip, etc) with
// additional information
type Link struct {
	XMLName xml.Name `xml:"link"`
	Text    string   `xml:"text,omitempty"`
	Type    string   `xml:"type,omitempty"`
	Href    string   `xml:"href"`
}

// Email is an email address. Broken into two parts (id, domain) to help prevent email harvesting.
type Email struct {
	XMLName xml.Name `xml:"email"`
	Id      string   `xml:"id,attr"`
	Domain  string   `xml:"domain,attr"`
}

// Bounds describes two lat/lon pairs defining the extent of an element.
type Bounds struct {
	XMLName xml.Name `xml:"bounds"`
	MinLat  float64  `xml:"minlat,attr"`
	MinLon  float64  `xml:"minlon,attr"`
	MaxLat  float64  `xml:"maxlat,attr"`
	MaxLon  float64  `xml:"maxlon,attr"`
}

// Copyright contains information about the copyright holder and any license governing use of this
// file. By linking to an appropriate license, you may place your data into the public domain or
// grant additional usage rights.
type Copyright struct {
	XMLName xml.Name `xml:"copyright"`
	Year    int      `xml:"year,omitempty"`
	License string   `xml:"license,omitempty"`
	Author  string   `xml:"author"`
}

// Pt is a geographic point with optional elevation and time. Available for use by other schemas.
type Pt struct {
	XMLName   xml.Name `xml:"pt"`
	Elevation float64  `xml:"ele,omitempty"`
	Time      string   `xml:"time,omitempty"`
	Lat       float64  `xml:"lat,attr"`
	Lon       float64  `xml:"lon,attr"`
}

// PtSeg is an ordered sequence of points (for polygons or polylines, etc.)
type PtSeg struct {
	XMLName xml.Name `xml:"ptseg"`
	Points  *[]Pt    `xml:"pt,omitempty"`
}
