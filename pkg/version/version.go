package version

import (
	"fmt"

	"github.com/BlazerodJS/blazerod/pkg/v8engine"
)

var blazerodVersion = "dev"

// Version returns the version string, including dependencies
func Version() string {
	return fmt.Sprintf("blazerod %s (V8 %s)", blazerodVersion, v8engine.Version())
}
