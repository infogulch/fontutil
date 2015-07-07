package fontutil

import (
	"io/ioutil"
	"testing"

	"code.google.com/p/freetype-go/freetype"
)

func TestFind(t *testing.T) {
	filename, err := Find("Times New Roman")
	if err != nil {
		t.Fatalf("cannot find font path: %s", err)
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("cannot read font file: %s", err)
	}
	font, err := freetype.ParseFont(file)
	if err != nil {
		t.Fatalf("cannot parse font: %s", err)
	}
	_ = font
}
