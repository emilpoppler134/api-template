package utils

import (
	"fmt"
	"strconv"
)

func ParseInt(str string) (int, error) {
	result, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse the string %s to an integer", str)
	}
	return result, nil
}
