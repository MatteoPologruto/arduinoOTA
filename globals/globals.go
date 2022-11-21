package globals

import (
	"os"
	"path/filepath"

	"github.com/arduino/arduinoOTA/version"
)

var (
	// VersionInfo contains all info injected during build
	VersionInfo = version.NewInfo(filepath.Base(os.Args[0]))
)
