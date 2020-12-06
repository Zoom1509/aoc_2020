package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part2() {
	// file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer file.Close()

	// var arr []Passport
	var blocks [][][]string

	scanner := bufio.NewScanner(file)

	var arr [][]string

	for scanner.Scan() {
		s := scanner.Text()

		if len(s) != 0 {

			arr = append(arr, strings.Split(s, ""))
		}

		if len(s) == 0 {
			blocks = append(blocks, arr)
			arr = make([][]string, 0)
		}
	}
	blocks = append(blocks, arr)

	sum := 0
	for _, block := range blocks {
		// for _, b := range block {
		sum = sum + perUniqe(block)
		// fmt.Printf("%+v", block)
		// // sum = sum + unique(strings.Split(block, ""))
		// fmt.Println("------------------")
		// fmt.Println(sum)
		// fmt.Println("------------------")

		// }
	}
	fmt.Println(sum)

	// fmt.Println(blocks)

	// fmt.Println(checkPart2(arr))
	// fmt.Println(checkPart2(arr))

}
func part1() {
	// file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer file.Close()

	// var arr []Passport
	var blocks []string

	scanner := bufio.NewScanner(file)

	tmp := ""

	for scanner.Scan() {
		s := scanner.Text()

		if len(s) != 0 {
			tmp = tmp + " " + s
		}

		if len(s) == 0 {
			blocks = append(blocks, tmp)
			tmp = ""
		}
	}
	blocks = append(blocks, tmp)

	sum := 0
	for _, block := range blocks {
		fmt.Println(block)
		fmt.Println("//")
		sum = sum + unique(strings.Split(block, ""))
		fmt.Println("------------------")
		fmt.Println(sum)
		fmt.Println("------------------")

	}
	fmt.Println(sum)
	fmt.Println(len(blocks))

	// fmt.Println(blocks)

	// fmt.Println(checkPart2(arr))
	// fmt.Println(checkPart2(arr))

}
func perUniqe(a [][]string) int {
	// keys := make(map[string]bool)
	// list := []string{}
	c := 0
	if len(a) == 1 {
		c = unique(a[0])
	} else {
		for _, first_entry := range a[0] {
			tc := 0
			for _, sec_arr := range a[1:] {
				for _, sec_entry := range sec_arr {
					fmt.Println(first_entry, " / ", sec_entry)
					if first_entry == sec_entry {
						tc++
						continue
					}
				}
			}
			if tc == len(a)-1 {
				c++
			}
		}
	}
	return c
}

func unique(intSlice []string) int {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value && entry != "" && entry != " " {
			keys[entry] = true
			fmt.Print(entry)
			list = append(list, entry)
		}
	}
	fmt.Printf("// %+v \n//%+v", keys, list)
	fmt.Println("Got: ", intSlice, "gave:", len(list))
	if len(list) != len(keys) {
		fmt.Println("################################################")
	}
	return len(list)
}
