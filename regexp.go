package tools

import (
	"regexp"
)

var (
	FileExtension = regexp.MustCompile(`\.[a-zA-Z0-9]+$`)
)
