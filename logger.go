package dht

import (
	"os"
	"strings"

	logger "github.com/d2r2/go-logger"
)

var lg logger.PackageLog

func init() {
	var level logger.LogLevel

	switch strings.ToLower(os.Getenv("DHT_LOG_LEVEL")) {
	case "fatal":
		level = logger.FatalLevel
	case "error":
		level = logger.ErrorLevel
	case "warn":
		level = logger.WarnLevel
	case "info":
		level = logger.InfoLevel
	case "debug":
		level = logger.DebugLevel
	default:
		level = logger.ErrorLevel
	}

	lg = logger.NewPackageLogger("dht", level)
}
