package resolver

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BlazerodJS/blazerod/pkg/v8engine"
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

// Resolver implements module resolution (stdlib, file, NPM)
type Resolver struct {
	engine  *v8engine.Engine
	basedir string
}

// NewResolver creates a new Resolver
func NewResolver(engine *v8engine.Engine, basedir string) *Resolver {
	return &Resolver{
		engine:  engine,
		basedir: basedir,
	}
}

// ResolveModule imports the requested module
func (r *Resolver) ResolveModule(specifier, referrer string) (string, int) {
	fmt.Printf("Requested %s (%s)\n", specifier, referrer)
	if specifier == "test" {
		return "test", 0
	}

	code, path := resolveByFile(specifier, referrer)

	if path != "" {
		resolver := r
		resolver.basedir = filepath.Dir(path)
		fmt.Printf("Loading %s\n", path)
		if ret := r.engine.LoadModule(string(code), path, resolver.ResolveModule); ret != 0 {
			return "", 1
		}
		return path, 0
	}

	return "", 1
}
