package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func parseSeat(s string) int {

	codes := strings.Split(s, "")

	var l = 0
	var r = 127
	var x = 8
	var y = 0
	var m = 0
	var n = 0

	fmt.Println("--------------", s)
	for _, code := range codes {
		switch code {
		case "F":
			r = (r + l) >> 1
			m = l
		case "B":
			l = (r + l) >> 1
			m = r
		case "R":
			y = (x + y) >> 1
			n = x
		case "L":
			x = (x + y) >> 1
			n = y
		}
		fmt.Println(code, r, l, x, y, m, n)
	}
	fmt.Println("--------------")

	return (m * 8) + n
}

func main() {
	// file, err := os.Open("test.txt")

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer file.Close()

	var seats []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		seats = append(seats, s)
	}

	m := make(map[int]string)
	for _, seat := range seats {
		m[parseSeat(seat)] = seat

	}

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}

	for i := 0; i < 922; i++ {
		_, f := m[i]
		if !f {
			_, up := m[i+1]
			_, down := m[i-1]
			if up && down {
				fmt.Println(i, " looks good", m[i], m[i-1], m[i+1])
			}
		}
	}
}
