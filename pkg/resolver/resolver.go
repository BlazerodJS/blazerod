package resolver

import "fmt"

// ResolveModule imports the requested module
func ResolveModule(specifier, referrer string) (string, int) {
	fmt.Printf("Requested %s (%s)\n", specifier, referrer)
	if specifier == "test" {
		return "test", 0
	}

	return "", 1
}
