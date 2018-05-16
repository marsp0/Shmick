package main

import (
	"strings"
)

func licenseKeyFormatting(S string, K int) string {
	// https://leetcode.com/problems/license-key-formatting/description/
	var (
		letterList = strings.Split(S, "")
		resultList = []string{}
		counter    = 0
	)
	for i := len(S) - 1; i >= 0; i-- {
		if counter == K {
			counter = 0
			resultList = append(resultList, "-")
		}
		if letterList[i] != "-" {
			resultList = append(resultList, strings.ToUpper(letterList[i]))
			counter++
		}
	}
	var (
		i = 0
		j = len(resultList) - 1
	)
	for i < j {
		resultList[i], resultList[j] = resultList[j], resultList[i]
		i++
		j--
	}
	if len(resultList) > 0 && resultList[0] == "-" {
		resultList = resultList[1:]
	}
	return strings.Join(resultList, "")
}
