package main

import (
	"regexp"
	"strings"
)

type byNumDesc []*movie

func (m byNumDesc) Len() int {
	return len(m)
}

func (m byNumDesc) Less(i, j int) bool {
	return m[i].num > m[j].num
}

func (m byNumDesc) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func extractYear(s string) string {
	re := regexp.MustCompile(`\(.*(\d\d\d\d)\)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}

func extractDate(s string) string {
	return strings.Split(s, ",")[0]
}
