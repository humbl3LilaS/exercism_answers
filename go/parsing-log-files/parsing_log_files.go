package parsinglogfiles

import (
	"fmt"
	"regexp"
	"slices"
)

var validTags = []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FIL]"}

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[[A-Z]{3}\]`)
	if !re.MatchString(text) {
		return false
	}

	tag := re.FindString(text)

	if !slices.Contains(validTags, tag) {
		return false
	}
	return true
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)

	total := 0
	for _,line := range lines {
		if re.MatchString(line) {
			total += 1
		}
	}

	return total
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]*`)
	return string(re.ReplaceAllLiteralString(text, ""))
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+[A-Za-z]+[A-Za-z0-9]*`)
	result := []string{}

	for _, line := range lines {
		if re.MatchString(line) {
			user := re.FindString(line)
			reUser := regexp.MustCompile(`User\s+`)
			newLine := fmt.Sprintf("[USR] %s %s", reUser.ReplaceAllLiteralString(user, ""), line)
			result = append(result, newLine)
		} else {
			result = append(result, line)
		}
	}
	return result
}
