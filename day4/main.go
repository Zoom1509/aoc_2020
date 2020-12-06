package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID)
type Passport struct {
	byr byr
	iyr iyr
	eyr eyr
	hgt hgt
	hcl hcl
	ecl ecl
	pid pid
	cid cid
}
type PField interface {
	valid() bool
}

type byr struct {
	v string
}
type iyr struct{ v string }
type eyr struct{ v string }
type hgt struct{ v string }
type hcl struct{ v string }
type ecl struct{ v string }
type pid struct{ v string }
type cid struct{ v string }

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.

func (f iyr) valid() bool {
	if len(f.v) == 0 {
		return false
	}

	num, err := strconv.Atoi(f.v)
	if err != nil {
		return false
	}

	if num >= 2010 && num <= 2020 {
		return true
	}
	return false
}

func (f eyr) valid() bool {
	if len(f.v) == 0 {
		return false
	}

	num, err := strconv.Atoi(f.v)
	if err != nil {
		return false
	}

	if num >= 2020 && num <= 2030 {
		return true
	}
	return false
}

func (f hgt) valid() bool {

	if len(f.v) == 0 {
		return false
	}

	if strings.Contains(f.v, "cm") {
		cm := strings.Split(f.v, "cm")
		num, err := strconv.Atoi(cm[0])
		if err != nil {
			return false
		}
		if num >= 150 && num <= 193 {
			return true
		}
	} else if strings.Contains(f.v, "in") {
		cm := strings.Split(f.v, "in")
		num, err := strconv.Atoi(cm[0])
		if err != nil {
			return false
		}
		if num >= 59 && num <= 76 {
			return true
		}
	}
	return false
}
func (f hcl) valid() bool {
	if len(f.v) != 7 {
		return false
	}

	matched, err := regexp.Match("#[[:xdigit:]]{6}", []byte(f.v))
	if err != nil {
		fmt.Println(err)
	}

	return matched

}
func (f ecl) valid() bool {
	allowed := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, a := range allowed {
		if a == f.v {
			return true
		}
	}
	return false
}
func (f pid) valid() bool {
	if len(f.v) != 9 {
		return false
	}
	matched, _ := regexp.Match("[[:digit:]]", []byte(f.v))
	return matched

}
func (f cid) valid() bool { return true }

func (f byr) valid() bool {
	if len(f.v) == 0 {
		return false
	}

	num, err := strconv.Atoi(f.v)
	if err != nil {
		return false
	}

	if num >= 1920 && num <= 2002 {
		return true
	}
	return false
}

func (p Passport) Valid() bool {

	fmt.Println("p.byr :", p.byr.v, "is", p.byr.valid())
	fmt.Println("p.iyr :", p.iyr.v, "is", p.iyr.valid())
	fmt.Println("p.eyr :", p.eyr.v, "is", p.eyr.valid())
	fmt.Println("p.hgt :", p.hgt.v, "is", p.hgt.valid())
	fmt.Println("p.hcl :", p.hcl.v, "is", p.hcl.valid())
	fmt.Println("p.ecl :", p.ecl.v, "is", p.ecl.valid())
	fmt.Println("p.pid :", p.pid.v, "is", p.pid.valid())
	fmt.Println("p.cid :", p.cid.v, "is", p.cid.valid())

	fmt.Println("-----------------------")
	return p.byr.valid() &&
		p.iyr.valid() &&
		p.eyr.valid() &&
		p.hgt.valid() &&
		p.hcl.valid() &&
		p.ecl.valid() &&
		p.pid.valid() &&
		p.cid.valid()
}

func (p Passport) New(s string) Passport {

	parts := strings.Fields(s)

	for _, part := range parts {

		entry := strings.Split(part, ":")

		switch entry[0] {
		case "byr":
			p.byr.v = entry[1]
		case "iyr":
			p.iyr.v = entry[1]
		case "eyr":
			p.eyr.v = entry[1]
		case "hgt":
			p.hgt.v = entry[1]
		case "hcl":
			p.hcl.v = entry[1]
		case "ecl":
			p.ecl.v = entry[1]
		case "pid":
			p.pid.v = entry[1]
		case "cid":
			p.cid.v = entry[1]
		}
		// fmt.Printf("%s %s %+v\n", entry[0], entry[1], p)

	}
	// fmt.Println("----------------------")
	return p

}

func main() {
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

	counter := 0
	for _, block := range blocks {
		var p Passport
		p = p.New(block)
		if p.Valid() {
			counter++
		}
		// fmt.Println(block)
		// fmt.Println(p)
	}
	fmt.Println(counter)

	// fmt.Println(blocks)

	// fmt.Println(checkPart2(arr))
	// fmt.Println(checkPart2(arr))

}
