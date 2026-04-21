package luhn

import (
	"regexp"
	"strings"
	"strconv"
)

func Valid(id string) bool {

	normalizedId := strings.ReplaceAll(id, " ", "")

	if len(normalizedId) <= 1 {
		return false
	}


	re := regexp.MustCompile(`^\d+$`)


	if !re.MatchString(normalizedId) {
		return false
	}



	result := 0

	for  i := len(normalizedId); i > 0; i-- {
		idx := len(normalizedId) - i
		val, _ := strconv.Atoi(string(normalizedId[i - 1]))
		if idx % 2 != 0 {
			j := val * 2
			if j > 9 {
				result += (j - 9)
			} else {

				result += j
			}
		} else {
			result += val
		}
	}	
	

    return result % 10 == 0

}
