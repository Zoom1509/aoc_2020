package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// --- Day 7: Handy Haversacks ---
// You land at the regional airport in time for your next flight. In fact, it looks like you'll even have time to grab some food: all flights are currently delayed due to issues in luggage processing.

// Due to recent aviation regulations, many rules (your puzzle input) are being enforced about bags and their contents; bags must be color-coded and must contain specific quantities of other color-coded bags. Apparently, nobody responsible for these regulations considered how long they would take to enforce!

// For example, consider the following rules:

// light red bags contain 1 bright white bag, 2 muted yellow bags.
// dark orange bags contain 3 bright white bags, 4 muted yellow bags.
// bright white bags contain 1 shiny gold bag.
// muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
// shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
// dark olive bags contain 3 faded blue bags, 4 dotted black bags.
// vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
// faded blue bags contain no other bags.
// dotted black bags contain no other bags.
// These rules specify the required contents for 9 bag types. In this example, every faded blue bag is empty, every vibrant plum bag contains 11 bags (5 faded blue and 6 dotted black), and so on.

// You have a shiny gold bag. If you wanted to carry it in at least one other bag, how many different bag colors would be valid for the outermost bag? (In other words: how many colors can, eventually, contain at least one shiny gold bag?)

// In the above rules, the following options would be available to you:

// A bright white bag, which can hold your shiny gold bag directly.
// A muted yellow bag, which can hold your shiny gold bag directly, plus some other bags.
// A dark orange bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
// A light red bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
// So, in this example, the number of bag colors that can eventually contain at least one shiny gold bag is 4.

// How many bag colors can eventually contain at least one shiny gold bag? (The list of rules is quite long; make sure you get all of it.)

func main() {
	part1("input.txt")
	// part2()
}

func part1(inp string) {
	file, err := os.Open(inp)
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make(map[string][]string)

	for scanner.Scan() {
		s := scanner.Text()
		tmp := strings.Split(s, " bags contain")

		if len(tmp) == 0 {
			continue
		}
		bag := tmp[0]
		canhold := tmp[1]

		if canhold == "no other bags." || bag == "shiny gold bag" {
		} else {
			re := regexp.MustCompile(`[0-9] ([a-z]* [a-z]*) bag`)
			matches := re.FindAllStringSubmatch(canhold, -1)

			for _, s := range matches {
				m[bag] = append(m[bag], s[1])
			}

		}
	}

	var canHoldGold func(string) bool
	// should loop over map, recursively check subentry
	canHoldGold = func(s string) bool {
		// fmt.Println("Check ", s)
		if v, ok := m[s]; ok {
			for _, sub := range v {
				// fmt.Println("\t check ", sub)
				if sub == "shiny gold" {
					fmt.Println("check found in ", s)
					return true
				}
				return canHoldGold(sub)
			}
		}
		return false
	}

	c := 0
	for k, _ := range m {
		if canHoldGold(k) {
			c++
		}
	}
	fmt.Printf("%v", c)

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
