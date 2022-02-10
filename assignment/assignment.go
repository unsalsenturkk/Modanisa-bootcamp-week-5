package assignment

import (
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

	var r []rune

	for _, runeValue := range s {
		r = append(r, runeValue)
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
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

	words := strings.Split(arr[1], ",")
	str := arr[0]

	var searchString string
	var foundStrings []string
	for _, runeValue := range str {

		searchString += string([]rune{runeValue})
	searchStringFounded:
		for _, w := range words {
			if strings.EqualFold(searchString, w) {
				foundStrings = append(foundStrings, searchString)
				searchString = ""
				break searchStringFounded
			}
		}
	}
	if len(foundStrings) != 2 {
		return "not possible"
	}

	return strings.Join(foundStrings, ",")
}

func VariadicSet(i ...interface{}) []interface{} {
	keys := make(map[interface{}]struct{})
	list := []interface{}{}

	for _, v := range i {
		if _, value := keys[v]; !value {
			keys[v] = struct{}{}
			list = append(list, v)
		}
	}
	return list
}
