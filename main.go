package main

import (
	"check_elasticsearch/cmd"
	"fmt"
)

var (
	// These get filled at build time with the proper vaules
	version = "development"
	commit  = "HEAD"
	date    = "latest"
)

func main() {
	cmd.Execute(buildVersion())
}

//goland:noinspection GoBoolExpressions
func buildVersion() string {
	result := version

	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}

	if date != "" {
		result = fmt.Sprintf("%s\ndate: %s", result, date)
	}

	return result
}
