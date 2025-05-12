package migration

import "embed"

// Reserving this file to deal with .sql files (while building)

//go:embed *.sql
var FS embed.FS
