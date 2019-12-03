package resolver

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func resolveByFile(specifier, referrer string) ([]byte, string) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, ""
	}

	path := specifier
	if !filepath.IsAbs(path) {
		path = filepath.Join(dir, path)
	}

	if filepath.Ext(path) == "" {
		_, err = os.Stat(path + ".js")
		switch {
		case os.IsNotExist(err):
			path = filepath.Join(path, "index.js")
		case err != nil:
			return nil, ""
		default:
			path = path + ".js"
		}
	}

	if _, err := os.Stat(path); err != nil {
		return nil, ""
	}
	fmt.Printf("Resolving with %s\n", path)
	code, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return code, path
}
