package assignment

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAddUint32(t *testing.T) {
	/*
		Sum uint32 numbers, return uint32 sum value and boolean overflow flag
		cases need to pass:
			math.MaxUint32, 1 => 0, true
			1, 1 => 2, false
			42, 2701 => 2743, false
			42, math.MaxUint32 => 41, true
			4294967290, 5 => 4294967295, false
			4294967290, 6 => 0, true
			4294967290, 10 => 4, true


	*/

	type input struct {
		number1 uint32
		number2 uint32
	}

	type expect struct {
		expectSum      uint32
		expectOverflow bool
	}

	cases := []struct {
		caseId int
		input
		expect
	}{
		{
			1,
			input{number1: math.MaxUint32, number2: 1},
			expect{0, true},
		}, {
			2,
			input{number1: 1, number2: 1},
			expect{2, false},
		}, {
			3,
			input{number1: 42, number2: 2701},
			expect{2743, false},
		}, {
			4,
			input{number1: 42, number2: math.MaxUint32},
			expect{41, true},
		}, {
			5,
			input{number1: 4294967290, number2: 5},
			expect{4294967295, false},
		}, {
			6,
			input{number1: 4294967290, number2: 6},
			expect{0, true},
		}, {
			7,
			input{number1: 4294967290, number2: 10},
			expect{4, true},
		}, {
			8,
			input{number1: 4294967295, number2: 0},
			expect{4294967295, false},
		}, {
			9,
			input{number1: 0, number2: 4294967295},
			expect{4294967295, false},
		}, {
			10,
			input{number1: 1, number2: math.MaxUint32},
			expect{0, true},
		}, {
			11,
			input{number1: 2147483647, number2: 2147483647},
			expect{4294967294, false},
		}, {
			12,
			input{number1: 2147483647, number2: 2147483648},
			expect{4294967295, false},
		}, {
			13,
			input{number1: 2147483648, number2: 2147483647},
			expect{4294967295, false},
		}, {
			14,
			input{number1: 2147483648, number2: 2147483648},
			expect{0, true},
		},
	}

	for _, c := range cases {
		sum, overflow := AddUint32(c.number1, c.number2)
		assert.Equal(t, c.expectSum, sum, fmt.Sprintf("%s%d", "caseId:", c.caseId))
		assert.Equal(t, c.expectOverflow, overflow, fmt.Sprintf("%s%d", "caseId:", c.caseId))
	}

}

func TestCeilNumber(t *testing.T) {
	/*
		Ceil the number within 0.25
		cases need to pass:
			42.42 => 42.50
			42 => 42
			42.01 => 42.25
			42.24 => 42.25
			42.25 => 42.25
			42.26 => 42.50
			42.55 => 42.75
			42.75 => 42.75
			42.76 => 43
			42.99 => 43
			43.13 => 43.25
	*/

	cases := []struct {
		caseId int
		input  float64
		expect float64
	}{
		{
			1,
			42.42,
			42.50,
		},
		{
			2,
			42,
			42,
		},
		{
			3,
			42.01,
			42.25,
		},
		{
			4,
			42.24,
			42.25,
		},
		{
			5,
			42.25,
			42.25,
		},
		{
			6,
			42.26,
			42.50,
		},
		{
			7,
			42.55,
			42.75,
		},
		{

			8,
			42.75,
			42.75,
		},
		{
			9,
			42.76,
			43,
		},
		{
			10,
			42.99,
			43,
		},
		{
			11,
			43.13,
			43.25,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s %d", "caseId :", c.caseId), func(t *testing.T) {
			point := CeilNumber(c.input)
			assert.Equal(t, c.expect, point)
		})
	}

}

func TestAlphabetSoup(t *testing.T) {
	/*
		String with the letters in alphabetical order.
		cases need to pass:
		 	"hello" => "ehllo"
			""    => ""
			"h"   => "h"
			"ab"  => "ab"
			"ba"  => "ab"
			"bac" => "abc"
			"cba" => "abc"
	*/

	cases := []struct {
		caseId int
		input  string
		expect string
	}{
		{
			1,
			"hello",
			"ehllo",
		}, {
			2,
			"",
			"",
		}, {
			3,
			"h",
			"h",
		}, {4,
			"ab",
			"ab",
		}, {
			5,
			"ba",
			"ab",
		}, {
			6,
			"bac",
			"abc",
		}, {
			7,
			"cba",
			"abc",
		},
	}

	for _, c := range cases {
		result := AlphabetSoup(c.input)
		assert.Equal(t, c.expect, result, fmt.Sprintf("%s%d", "caseId: ", c.caseId))
	}

}

func TestStringMask(t *testing.T) {
	/*
		Replace after n(uint) character of string with '*' character.
		cases need to pass:
			"!mysecret*", 2 => "!m********"
			"", n(any positive number) => "*"
			"a", 1 => "*"
			"string", 0 => "******"
			"string", 3 => "str***"
			"string", 5 => "strin*"
			"string", 6 => "******"
			"string", 7(bigger than len of "string") => "******"
			"s*r*n*", 3 => "s*r***"
	*/

	cases := []struct {
		caseId       int
		inputString  string
		inputNumber  uint
		expectString string
	}{
		{
			1,
			"!mysecret*",
			2,
			"!m********",
		},
		{

			2,
			"",
			5,
			"*",
		},
		{
			3,
			"a",
			1,
			"*",
		},
		{4,
			"string",
			0,
			"******",
		},
		{
			5,
			"string",
			3,
			"str***",
		},
		{
			6,
			"string",
			5,
			"strin*",
		},
		{
			7,
			"string",
			6,
			"******",
		},
		{
			9,
			"string",
			7,
			"******",
		},
		{
			10,
			"s*r*n*",
			3,
			"s*r***",
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s %d", "caseId :", c.caseId), func(t *testing.T) {
			result := StringMask(c.inputString, c.inputNumber)

			assert.Equal(t, c.expectString, result)
		})
	}

}

//func TestWordSplit(t *testing.T) {
//	words := "apple,bat,cat,goodbye,hello,yellow,why"
//	/*
//		Your goal is to determine if the first element in the array can be split into two words,
//		where both words exist in the dictionary(words variable) that is provided in the second element of array.
//
//		cases need to pass:
//			[2]string{"hellocat",words} => hello,cat
//			[2]string{"catbat",words} => cat,bat
//			[2]string{"yellowapple",words} => yellow,apple
//			[2]string{"",words} => not possible
//			[2]string{"notcat",words} => not possible
//			[2]string{"bootcamprocks!",words} => not possible
//	*/
//
//	cases := []struct {
//		caseId       int
//		inputString  [2]string
//		expectString string
//	}{
//		{
//			1,
//			[2]string{"hellocat", words},
//			"hello,cat",
//		},
//		{
//
//			2,
//			[2]string{"catbat", words},
//			"cat,bat",
//		},
//		{
//			3,
//			[2]string{"yellowapple", words},
//			"yellow,apple",
//		},
//		{4,
//			[2]string{"", words},
//			"not possible",
//		},
//		{
//			5,
//			[2]string{"notcat", words},
//			"not possible",
//		},
//		{
//			6,
//			[2]string{"bootcamprocks!", words},
//			"not possible",
//		},
//	}
//
//	for _, c := range cases {
//		t.Run(fmt.Sprintf("%s %d", "caseId :", c.caseId), func(t *testing.T) {
//			result := WordSplit(c.inputString)
//
//			assert.Equal(t, c.expectString, result)
//		})
//	}
//
//}

//func TestVariadicSet(t *testing.T) {
//	/*
//		FINAL BOSS ALERT :)
//		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )
//
//		Convert inputs to set(no duplicate element)
//		cases need to pass:
//			4,2,5,4,2,4 => []interface{4,2,5}
//			"bootcamp","rocks!","really","rocks!" => []interface{"bootcamp","rocks!","really"}
//			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
//	*/
//	type input []interface {
//	}
//	type expect []interface {
//	}
//
//	cases := []struct {
//		caseId int
//		input  input
//		expect expect
//	}{
//		{
//			1,
//			append(input{}, 4, 2, 5, 4, 2, 4),
//			append(expect{}, 4, 2, 5),
//		},
//		{
//
//			1,
//			append(input{}, "bootcamp", "rocks!", "really", "rocks!"),
//			append(expect{}, "bootcamp", "rocks!", "really"),
//		},
//		{
//			1,
//			append(input{}, 1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"),
//			append(expect{}, 1, uint32(1), "first", 2, uint32(2), "second"),
//		},
//	}
//
//	for _, c := range cases {
//		t.Run(fmt.Sprintf("%s %d", "caseId :", c.caseId), func(t *testing.T) {
//			set := VariadicSet(c.input)
//			assert.Equal(t, []interface{}{4, 2, 5}, set)
//		})
//	}
//
//}
