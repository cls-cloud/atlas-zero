package utils

import (
	"strconv"
	"strings"
)

func StrAtoi(str string) int64 {
	atoi, _ := strconv.ParseInt(str, 10, 64)
	return atoi
}

func SplitToInt64(str, sep string) ([]int64, error) {
	parts := strings.Split(str, sep)
	result := make([]int64, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		value, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}
