package assignment

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {

	sum := x + y
	if (sum > x) == (y > 0) {
		return sum, false
	}
	return sum, true
}

func CeilNumber(f float64) float64 {
	f = math.Ceil(f*4) / 4
	return f
}

func AlphabetSoup(s string) string {

	var intArray []int

	for _, runeValue := range s {
		intArray = append(intArray, int(runeValue))
	}

	sort.Ints(intArray)

	var runeArray []rune
	for _, r := range intArray {
		runeArray = append(runeArray, rune(r))
	}

	newS := string(runeArray)
	return newS
}

func StringMask(s string, n uint) string {

	l := len([]rune(s))
	if l == 0 {
		return "*"
	} else if l <= int(n) {
		n = 0
	}

	var runeArray []rune
	for i, runeValue := range s {

		if i < int(n) {
			runeArray = append(runeArray, runeValue)
		} else {
			runeArray = append(runeArray, 42)
		}
	}

	return string(runeArray)
}

func WordSplit(arr [2]string) string {

	var searchString string
	var foundString string

	str := arr[0]
	s := arr[1]

	searchString = ""
	foundString = ""
	for _, runeValue := range str {

		searchString += string([]rune{runeValue})
		if strings.Contains(s, searchString) {
			foundString = searchString
		}

	}
	fmt.Println(foundString)
	return ""
}

func VariadicSet(i ...interface{}) []interface{} {
	return nil
}
