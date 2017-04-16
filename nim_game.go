package main

import "fmt"

func canWinNim(n int) bool {
	if n <= 3 {
		return true
	}
    return n%4 != 0
}

func main() {
	fmt.Println(canWinNim(6))
}
