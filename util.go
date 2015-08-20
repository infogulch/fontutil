// Package fontutil is a small collection of utility functions that make working with fonts easier.
package fontutil

import (
	"image"

	"golang.org/x/image/math/fixed"

	"github.com/golang/freetype"
)

// DrawWidth simulates drawing s using Context c and returns the raster.Point
// where X is the width of the drawn string. The context's src and dst are
// unused, but the font should already be set. Clip is set to an empty Rectangle
// and should be reset afterwards.
func DrawWidth(c *freetype.Context, s string) (fixed.Point26_6, error) {
	// nil rectangle is always empty so draw is never called
	c.SetClip(image.Rectangle{})
	p, err := c.DrawString(s, fixed.Point26_6{}) // 0,0
	return p, err
}

// CenterX returns the Point at which string s will be centered within rect r
// along the X coordinate when passed to c.DrawString. The Y coordinate will be 0.
// Clip is set to r before returning.
func CenterX(c *freetype.Context, s string, r image.Rectangle) (fixed.Point26_6, error) {
	w, err := DrawWidth(c, s)
	c.SetClip(r)
	if err != nil {
		return fixed.Point26_6{}, err
	}
	half := Int26_6(0.5)
	xmin := freetype.Pt(r.Min.X, 0)
	return freetype.Pt(r.Dx(), 0).Mul(half).Sub(w.Mul(half)).Add(xmin), nil
}

// Int26_6 returns a fixed.Int26_6 representation of f.
func Int26_6(f float64) fixed.Int26_6 {
	return fixed.Int26_6(f * float64(1<<6))
}
