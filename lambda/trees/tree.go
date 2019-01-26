package main

import (
  "bytes"
  "github.com/fogleman/gg"
  "fmt"
)

func draw() (buffer *bytes.Buffer, err error) {
	c := gg.NewContext(1000, 1000)

  // Create the image.
	c.DrawCircle(500, 500, 400)
	c.SetRGB(0, 0, 0)
	c.Fill()

  // Write the bytes from the image in the context to a buffer.
  buffer = new(bytes.Buffer)
  if err = c.EncodePNG(buffer); err != nil {
    fmt.Printf("failed to encode png %s",err.Error())
  }
  return
}
