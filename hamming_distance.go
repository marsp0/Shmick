package main

import "fmt"
import "strconv"
import "strings"

func hammingDistance(x int, y int) int {
	var first = int64(x)
	var second = int64(y) 
	var x_string = strconv.FormatInt(first,2)
	var y_string = strconv.FormatInt(second,2)
	var smaller *string
	var larger *string
	if x < y {
		smaller = &x_string
		larger = &y_string
	} else {
		smaller = &y_string
		larger = &x_string
	}
	var empty_values = strings.Repeat("0",len(*larger)-len(*smaller))
	var counter = 0
	*smaller = empty_values + *smaller
	for i := 0 ; i < len(*larger); i++ {
		if (*larger)[i] != (*smaller)[i] {
			counter += 1
		}
	}
	return counter
}

func main() {
	fmt.Println(hammingDistance(3,5))
}