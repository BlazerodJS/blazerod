package resolver

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var fileExtensions = []string{".ts", ".mjs", ".js"}

// NpmResolver resolves imports as NPM packages
type NpmResolver struct {
	moduleBase string
}

// Resolve attempts to resolve the specified module from NPM
func (r NpmResolver) Resolve(specifier, referrer string) ([]byte, string) {
	return r.resolveModule(specifier, r.moduleBase)
}

func (r NpmResolver) resolveModule(specifier, basedir string) ([]byte, string) {
	relativeToBase, err := filepath.Rel(r.moduleBase, basedir)
	if err != nil {
		return nil, ""
	}

	parts := strings.Split(relativeToBase, string(filepath.Separator))
	for idx := range parts {
		if idx != 0 && parts[idx-1] == "node_modules" {
			continue
		}

		candidatePath := filepath.Join(append(parts[:idx], "node_modules", specifier)...)
		candidatePath = filepath.Join(r.moduleBase, candidatePath)
		bytes, resolvedPath := r.resolvePath(candidatePath)
		if bytes != nil {
			return bytes, resolvedPath
		}
	}

	return nil, ""
}

func (r NpmResolver) resolvePath(candidatePath string) ([]byte, string) {
	bytes, resolvedPath := r.resolveFile(candidatePath)
	if bytes != nil {
		return bytes, resolvedPath
	}

	packageJSON, _ := ioutil.ReadFile(filepath.Join(candidatePath, "package.json"))
	if packageJSON != nil {
		var pkg struct {
			Module string
			Main   string
		}

		err := json.Unmarshal(packageJSON, &pkg)
		if err == nil && (pkg.Module != "" || pkg.Main != "") {
			if pkg.Module != "" {
				bytes, resolvedPath := r.resolveFile(pkg.Module)
				if bytes != nil {
					return bytes, resolvedPath
				}
			}

			if pkg.Main != "" {
				bytes, resolvedPath := r.resolveFile(pkg.Main)
				if bytes != nil {
					return bytes, resolvedPath
				}
			}
		}
	}

	for _, ext := range fileExtensions {
		path := filepath.Join(candidatePath, "index"+ext)
		bytes, err := ioutil.ReadFile(path)
		if err == nil {
			return bytes, path
		}
	}

	return nil, ""
}

func (r NpmResolver) resolveFile(candidatePath string) ([]byte, string) {
	bytes, err := ioutil.ReadFile(candidatePath)
	if err == nil {
		return bytes, candidatePath
	}

	for _, ext := range fileExtensions {
		path := candidatePath + ext
		bytes, err := ioutil.ReadFile(path)
		if err == nil {
			return bytes, path
		}
	}

	return nil, ""
}
