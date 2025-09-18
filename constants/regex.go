package regex

import "regexp"

// Returns a regex pattern that matches a version string
// that starts with a custom prefix text, optionally followed by ":" or whitespace.
// The pattern is case-insensitive and allows for optional whitespace around the version string.
func Version(prefixText string) *regexp.Regexp {
	escapedPrefix := regexp.QuoteMeta(prefixText) // Escape special characters in the prefix
	return regexp.MustCompile(`(?i)` + escapedPrefix + `[\s:]*v?(\d+\.\d+\.\d+)`)
}

// AndroidApk returns a regex that matches app-debug.apk and app-*-androidTest.apk
func AndroidApk() *regexp.Regexp {
	return regexp.MustCompile(`^(app-debug\.apk|app-.*-androidTest\.apk)$`)
}
