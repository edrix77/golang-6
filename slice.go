package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string
	three := 0
	x := make([]int, 3)
	fmt.Println(len(x))
	prin(&input)

	for input != "X" {
		var i, err = strconv.Atoi(input)

		if err != nil {
			fmt.Println("Unable to parse your input as a number. Error: ", err)
		} else if three < len(x) {
			x[three] = i
		} else {
			x = append(x, i)
			sort.Ints(x)
		}
		three = three + 1
		prin(&input)
	}

	fmt.Println(x)

}

func prin(input *string) {
	fmt.Printf("Enter a number: ")
	fmt.Scan(input)
}
