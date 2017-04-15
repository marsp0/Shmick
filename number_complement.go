package main

import "fmt"
import "strconv"

func findComplement(num int) int {
    var x = strconv.FormatInt(int64(num),2)
    var result string
    for i := 0 ; i < len(x) ; i++ {
    	if x[i] == '1'{
    		result += "0"
    	} else {
    		result += "1"
    	}
    }
    var i, _ = strconv.ParseInt(result, 2, 32)
    var p = int(i)
    return p
}

func main() {
	fmt.Println(findComplement(5))
}