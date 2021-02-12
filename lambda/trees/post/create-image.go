package main

import (
	"bytes"
	"fmt"
	"github.com/fogleman/gg"
	"math"
)

const (
	width  = 400
	height = 400
)

func polarLine(c *gg.Context, x0, y0, length, degrees float64) (x1, y1 float64) {
	theta := gg.Radians(degrees)
	x1 = length*(math.Cos(theta)) + x0
	y1 = length*(math.Sin(theta)) + y0
	c.DrawLine(x0, y0, x1, y1)
	c.Stroke()
	return
}

func tree(c *gg.Context, lineWidth, x0, y0, length, degrees float64, p TreeParam) {
	if lineWidth < 1 || x0 < 1 || y0 < 1 || x0 > width || y0 > height || length < 1 {
		return
	}

	c.SetLineWidth(lineWidth)
	lineWidth -= 2
	x1, y1 := polarLine(c, x0, y0, length, degrees)
	tree(c, lineWidth, x1, y1, length-p.LeftLength, degrees-p.LeftAngle, p)
	tree(c, lineWidth, x1, y1, length-p.RightLength, degrees+p.RightAngle, p)
}

func createTree(p TreeParam) (buffer *bytes.Buffer, err error) {

	c := gg.NewContext(width, height)
	c.SetRGB(0, 0, 0)
	tree(c, 15, width/2, height, 100, 270, p)
	// Save the image to the disk if testing locally.
	if dev {
		c.SavePNG("test.png")
	}

	// Write the bytes from the image in the context to a buffer.
	buffer = new(bytes.Buffer)
	if err = c.EncodePNG(buffer); err != nil {
		fmt.Printf("failed to encode png %s", err.Error())
	}
	return
}
