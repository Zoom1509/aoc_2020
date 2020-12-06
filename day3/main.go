package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const TREE = "#"
const FREE = "."
const FOUND = "X"
const EMPTY = "O"

func checkPart1(arr [][]string) int {

	var counter = 0

	var y = 0

	for x := 0; x < len(arr); x++ {
		// fmt.Println(x, y, string(arr[x][y]))
		if string(arr[x][y]) == TREE {
			counter++
			arr[x][y] = FOUND
		} else {
			arr[x][y] = EMPTY
		}

		if y+3 < len(arr[x]) {
			y = y + 3
		} else {
			y = y - len(arr[x]) + 3
		}
		// for y := 0; y < len(arr[0]); y+3 {
		// 	if string(arr[x][y]) == TREE {
		// 		counter++
		// 	}
		// }
		fmt.Println(arr[x])
	}
	return counter
}

func checkSlope(arr [][]string, right int, down int) int {

	var counter = 0

	var y = 0
	fmt.Println("- - - - - - - - - - - - - r: ", right, " / d:", down)

	for x := 0; x < len(arr); x += down {
		if string(arr[x][y]) == TREE {
			counter++
			// arr[x][y] = FOUND
		} else {
			// arr[x][y] = EMPTY
		}

		if y+right < len(arr[x]) {
			y = y + right
		} else {
			y = y - len(arr[x]) + right
		}
		fmt.Println(arr[x])
	}
	return counter
}

func checkPart2(arr [][]string) int {

	// Right 1, down 1.
	// Right 3, down 1. (This is the slope you already checked.)
	// Right 5, down 1.
	// Right 7, down 1.
	// Right 1, down 2.
	return checkSlope(arr, 1, 1) *
		checkSlope(arr, 3, 1) *
		checkSlope(arr, 5, 1) *
		checkSlope(arr, 7, 1) *
		checkSlope(arr, 1, 2)
}

// func checkPart2(arr []string) int {

// }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer file.Close()

	var arr [][]string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		// var y []string
		y := strings.Split(s, "")
		arr = append(arr, y)
	}

	fmt.Println(checkPart2(arr))
	// fmt.Println(checkPart2(arr))

}
