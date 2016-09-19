package shapes

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/ajstarks/svgo"
)

func toXml(s interface{}) string {
	var b bytes.Buffer
	enc := xml.NewEncoder(&b)
	enc.Encode(s)
	return b.String()
}

type Properties struct {
	Anchor image.Point `xml:"anchor"`
	Colour color.Color `xml:"colour"`
}

type Circle struct {
	Radius int        `xml:"radius,attr"`
	Props  Properties `xml:"properties"`
}

func (c *Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}

func (c *Circle) Xml() string {
	return toXml(c)
}

func (c *Circle) Draw(s *svg.SVG) {
	s.Circle(c.Props.Anchor.X, c.Props.Anchor.Y, c.Radius, fmt.Sprintf("fill:%s;stroke:none", colorToHex(c.Props.Colour)))
}

func colorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("%02x%02x%02x", r>>8, g>>8, b>>8)
}

func NewCircle(radius int, anchor image.Point, colour color.Color) *Circle {
	return &Circle{radius, Properties{anchor, colour}}
}

type Rectangle struct {
	Width  int        `xml:"width,attr"`
	Height int        `xml:"height,attr"`
	Props  Properties `xml:"properties"`
}

func (r *Rectangle) Area() float64 {
	return float64(r.Width * r.Height)
}

func (r *Rectangle) Xml() string {
	return toXml(r)
}

func (r *Rectangle) Draw(s *svg.SVG) {
	s.Rect(r.Props.Anchor.X, r.Props.Anchor.Y, r.Width, r.Height, fmt.Sprintf("fill:%s;stroke:none", colorToHex(r.Props.Colour)))
}

func NewRectangle(width int, height int, anchor image.Point, colour color.Color) *Rectangle {
	return &Rectangle{width, height, Properties{anchor, colour}}
}

func TotalArea(shapes ...Shape) (total float64) {
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}
