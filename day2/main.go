package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pwd struct {
	min  int
	max  int
	char string
	pass string
}

func checkPart1(arr []pwd) int {

	var counter = 0

	for _, v := range arr {
		c := strings.Count(v.pass, v.char)
		if c <= v.max && c >= v.min {
			counter = counter + 1
		}
	}
	return counter
}

func checkPart2(arr []pwd) int {

	var counter = 0

	for _, v := range arr {
		// c := strings.Count(v.pass, v.char)
		if (string(v.pass[v.max-1]) == v.char &&
			string(v.pass[v.min-1]) != v.char) ||
			(string(v.pass[v.max-1]) != v.char &&
				string(v.pass[v.min-1]) == v.char) {

			counter = counter + 1
		}
	}
	return counter
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer file.Close()

	var arr []pwd

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		var pass pwd
		p := strings.Split(s, " ")
		first := strings.Split(p[0], "-")
		pass.min, _ = strconv.Atoi(first[0])
		pass.max, _ = strconv.Atoi(first[1])
		pass.char = strings.Split(p[1], ":")[0]
		pass.pass = p[2] //password
		arr = append(arr, pass)
	}

	fmt.Println(checkPart1(arr))
	fmt.Println(checkPart2(arr))

}
