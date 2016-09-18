package main

import (
	"fmt"

	"github.com/ajstarks/svgo"
	"github.com/paulcadman/shapes"
	"image"
	"image/color"
	"log"
	"net/http"
)

func main() {
	circle := shapes.NewCircle(10, image.Pt(100, 100), color.RGBA{255, 0, 0, 0})
	rect := shapes.NewRectangle(50, 50, image.Pt(100, 150), color.Gray{200})
	fmt.Println(shapes.TotalArea(circle, rect))
	fmt.Println(circle.Xml())
	fmt.Println(rect.Xml())

	http.Handle("/draw", http.HandlerFunc(draw(circle, rect)))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func draw(shapes ...shapes.SVGDrawer) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		s := svg.New(w)
		s.Start(500, 500)
		for _, shape := range shapes {
			shape.Draw(s)
		}
		s.End()
	}
}
