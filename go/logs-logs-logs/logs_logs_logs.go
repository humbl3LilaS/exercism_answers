package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, val := range log {
		switch val {
		case '❗':
			return "recommendation"
		case '🔍' :
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	original := fmt.Sprintf("%c", oldRune)
	replacement := fmt.Sprintf("%c", newRune)
	return strings.ReplaceAll(log, original, replacement)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) > limit {
		return false
	}
	return true
}
