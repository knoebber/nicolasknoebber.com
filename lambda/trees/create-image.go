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

func drawTree(c *gg.Context,lineWidth, x0, y0, length, degrees float64, param treeParam) {
  if lineWidth < 1 || x0 < 1 || y0 < 1 || length < 1 {
    return
  }
  var x1 float64
  var y1 float64

  c.SetLineWidth(lineWidth)
  lineWidth -= 2
  x1, y1 = polarLine(c,x0,y0,length,degrees)
  drawTree(c,lineWidth,x1,y1,length-param.leftLengthOffset, degrees-param.leftAngleOffset,param)
  drawTree(c,lineWidth,x1,y1,length-param.rightLengthOffset, degrees + param.rightAngleOffset,param)
}

type treeParam struct {
  leftLengthOffset float64
  leftAngleOffset float64
  rightLengthOffset float64
  rightAngleOffset float64
}

func draw() (buffer *bytes.Buffer, err error) {
	c := gg.NewContext(width, height)
	c.SetRGB(0, 0, 0)
  param := treeParam {
    leftAngleOffset: 15,
    leftLengthOffset:1,
    rightLengthOffset:1,
    rightAngleOffset:45,
  }

  drawTree(c,10,width/2, height,100,270,param)
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
