package fontutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func find(fontname string) (filename string, err error) {
	keyname := "Software\\Microsoft\\Windows NT\\CurrentVersion\\Fonts"
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, keyname, registry.READ)
	if err != nil {
		return "", fmt.Errorf("cannot open fonts registry key: %s", err)
	}
	valuenames, err := key.ReadValueNames(0) // 0 means all
	if err != nil {
		return "", fmt.Errorf("cannot read font names from registry: %s", err)
	}
	for _, valuename := range valuenames {
		name := strings.TrimSpace(strings.TrimSuffix(valuename, "(TrueType)"))
		if name != fontname {
			continue
		}
		value, valuetype, err := key.GetStringValue(valuename)
		if err != nil {
			return "", fmt.Errorf("cannot read font key value (%s of type %d): %s", valuename, valuetype, err)
		}
		if !filepath.IsAbs(value) {
			value = filepath.Join(os.Getenv("windir"), "Fonts", value)
		}
		return value, nil
	}
	return "", fmt.Errorf("cannot find font name: %s", fontname)
}
