package utillib

import (
	"fmt"
	"strconv"
	"strings"
)

func UnicodeToCN(s string) string{
	sUnicodev := strings.Split(s, "\\u")
	var context string
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		context += fmt.Sprintf("%c", temp)
	}
	return context
}
