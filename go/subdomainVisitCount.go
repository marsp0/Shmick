package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) (result []string) {
	// https://leetcode.com/problems/subdomain-visit-count/description/
	var (
		domainCounter = map[string]int{}
	)
	for _, value := range cpdomains {
		var (
			splitStrings = strings.Split(value, " ")
			count, _     = strconv.Atoi(splitStrings[0])
			domain       = splitStrings[1]
		)
		domainCounter[domain] += count
		for index, byteLetter := range domain {
			if string(byteLetter) == "." {
				domainCounter[domain[index+1:]] += count
			}
		}
	}
	for key, value := range domainCounter {
		result = append(result, fmt.Sprintf("%d %s", value, key))
	}
	return result
}
