package fontutil_test

import (
	"fmt"
	"hash/fnv"
	"image"
	"image/color"
	"image/png"

	"github.com/infogulch/fontutil"

	"code.google.com/p/freetype-go/freetype"
)

var font = fontutil.Must(fontutil.FindAndParse("Arial"))

func Example() {
	// As usual in examples, this ignores all errors. Don't do this in your program.

	// setup and find start point for centering
	s := "Hello, World!"
	size := image.Rect(0, 0, 120, 20)
	dst := image.NewRGBA(size)
	c := freetype.NewContext()
	c.SetFont(font)
	c.SetFontSize(14.0)
	c.SetSrc(image.NewUniform(color.Black))
	c.SetDst(dst)
	start, _ := fontutil.CenterX(c, s, size) // CenterX calls c.SetClip(size)

	// perform draw at start.X + y 16
	c.DrawString(s, start.Add(freetype.Pt(0, 16)))

	// write the image out to a file
	// out, _ := os.Create("helloworld.png")
	// defer out.Close()

	// write image to hash for testing purposes
	out := fnv.New64()
	_ = png.Encode(out, dst)
	fmt.Printf("Hash of compressed image: %x", out.Sum64())
	// Output: Hash of compressed image: 2bbf3be4499fb96c
}
