package main

import (
  "bytes"
  "github.com/fogleman/gg"
  "fmt"
  "math"
)

const (
  width = 400
  height = 400
)

func polarLine(c *gg.Context, x0 ,y0, length, degrees float64)(x1,y1 float64) {
  theta := gg.Radians(degrees)
  x1 = length*(math.Cos(theta))+x0
  y1 = length*(math.Sin(theta))+y0
  // fmt.Printf("(%f, %f),(%f,%f)\n",x0,y0,x1,y1)
  c.DrawLine(x0,y0,x1,y1)
  c.Stroke()
  return
}

func drawTree(c *gg.Context,lineWidth, x0, y0, length, degrees float64) {
  if lineWidth < 1 || x0 < 1 || y0 < 1 || length < 1 {
    return
  }
  numBranches := 2
  var x1 float64
  var y1 float64

  for i:= 0; i< numBranches; i += 1 {
    fmt.Printf("setting line width %f\n", lineWidth)
    c.SetLineWidth(lineWidth)
    x1, y1 = polarLine(c,x0,y0,length,degrees)
    if i % 2 == 0 {
      drawTree(c,lineWidth -2,x1,y1,length - 20, degrees - 20)
    } else {
      drawTree(c,lineWidth -2,x1,y1,length - 10, degrees + 10)
      }
  }
}

func draw() (buffer *bytes.Buffer, err error) {
	c := gg.NewContext(width, height)
	c.SetRGB(0, 0, 0)
  drawTree(c,10,200, height,100,270)
  // If testing locally, save the image to the disk.
  if dev {
    c.SavePNG("test.png")
  }

  // Write the bytes from the image in the context to a buffer.
  buffer = new(bytes.Buffer)
  if err = c.EncodePNG(buffer); err != nil {
    fmt.Printf("failed to encode png %s",err.Error())
  }
  return
}
