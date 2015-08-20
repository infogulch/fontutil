package fontutil

import (
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

// Find is a platform-agnostic way to locate font files installed on the system by name.
func Find(fontname string) (filename string, err error) {
	// call to platform dependent function
	return find(fontname)
}

// FindAndParse Finds a font file named by fontname, reads it, and parses it.
func FindAndParse(fontname string) (*truetype.Font, error) {
	filepath, err := Find(fontname)
	if err != nil {
		return nil, err
	}
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return freetype.ParseFont(file)
}

// Must is a helper that wraps a call to FindAndParse and panics if the err is
// non-nil. It is intended for use in package variable initializations.
func Must(t *truetype.Font, e error) *truetype.Font {
	if e != nil {
		panic(e)
	}
	return t
}
