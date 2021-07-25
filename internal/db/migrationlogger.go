package db

import (
	"log"
)

// MigrationLogger logs activity during the migration process
type MigrationLogger struct {
	verbose bool
}

func (ml *MigrationLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}

func (ml *MigrationLogger) Verbose() bool {
	return ml.verbose
}
