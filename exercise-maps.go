package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	for v := range strings.Fields(s) {
		var key = strings.Fields(s)[v]
		elem, ok := m[key]
		if ok {
			m[strings.Fields(s)[v]] = elem + 1
		} else {
			m[strings.Fields(s)[v]] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
