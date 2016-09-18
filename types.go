package shapes

import "github.com/ajstarks/svgo"

type XmlStringer interface {
	Xml() string
}

type Shape interface {
	XmlStringer
	Area() float64
	Properties() Properties
}

type SVGDrawer interface {
	Draw(*svg.SVG)
}
