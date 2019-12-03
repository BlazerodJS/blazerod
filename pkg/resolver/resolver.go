package resolver

import (
	"fmt"
	"path/filepath"

	"github.com/BlazerodJS/blazerod/pkg/v8engine"
)

// Resolver implements module resolution (stdlib, file, NPM)
type Resolver struct {
	engine          *v8engine.Engine
	basedir         string
	originalBasedir string
}

// NewResolver creates a new Resolver
func NewResolver(engine *v8engine.Engine, basedir string) *Resolver {
	return &Resolver{
		engine:          engine,
		basedir:         basedir,
		originalBasedir: basedir,
	}
}

// ResolveModule imports the requested module
func (r *Resolver) ResolveModule(specifier, referrer string) (string, int) {
	fmt.Printf("Requested %s (%s)\n", specifier, referrer)
	if specifier == "test" {
		return "test", 0
	}

	code, path := resolveByFile(specifier, referrer)

	if code == nil {
		n := &NpmResolver{r.originalBasedir}
		code, path = n.Resolve(specifier, referrer)
	}

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
