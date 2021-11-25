package main

import (
	"strings"
)

type FileFilter func(buf []byte) []byte

func NewProtoFilter() FileFilter {
	return func(buf []byte) []byte {
		bufStr := string(buf)

		result := make([]string, 0)
		for _, ln := range strings.Split(bufStr, "\n") {
			if strings.HasPrefix(ln, "syntax") {
				result = append(result, ln)
				continue
			}
			if strings.HasPrefix(ln, "package") {
				result = append(result, ln)
				continue
			}
			if strings.HasPrefix(ln, "option") {
				result = append(result, ln)
				continue
			}
		}

		resultStr := strings.Join(result, "\n")
		return []byte(resultStr)
	}

}
