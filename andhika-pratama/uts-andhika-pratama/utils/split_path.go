package utils

import (
	"errors"
	"strings"
)

func SplitPath(path string) ([]string, error){
	parts := strings.Split(path, "/")
	if len(parts) > 4 {
		return []string{} , errors.New("request is not valid")
	}
	return parts, nil
}