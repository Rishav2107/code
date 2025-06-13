package main

import (
	"fmt"
	"strings"
)

func main() {
	//we have to think of string in a logical way and physical way
	//strings in go is all unicode
	//unicode is a technique to represent all international characters

	//Types related to strings
	//byte: a synonym for uint8
	//rune: a synonym for int32 for characters
	//string: an immutable sequence of characters
	// 		physically: a sequence of bytes(UTF-8 encoding)
	//		logically a sequence of (unicode) runes

	s := "€lite"
	fmt.Printf("%8T %[1]v %d\n", s, len(s)) //len of string is length of byte string
	fmt.Printf("%8T %[1]v\n", []rune(s))
	b := []byte(s)
	fmt.Printf("%8T %[1]v %d\n", b, len(b))

	//string structure
	s1 := "hello, world"
	hello := s1[:5]
	//world := s1[7:]

	//The internal string rep is a pointer and a length
	//		 _______________
	//eg:s1 |data(pointer)	|
	//      |len 12         |
	//      |_______________|
	//		    _______________
	//eg:hello |data(pointer)  |
	//         |len 5          |
	//         |_______________|
	//		 	_______________
	//eg:world |data(pointer)  |
	//         |len 5          |
	//         |_______________|

	//s1[5] = 'a' //not allowed as string is immutable
	s1 = "es" + s1 //copied, no changes in the original string. s1 now points to a new string address
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("hello: %v\n", hello)

	s1 = strings.ToUpper(s1) //copied, no changes in the original string. s1 now points to a new string address
	fmt.Printf("s1: %v\n", s1)
}

/*
  string €lite 7
 []int32 [8364 108 105 116 101]
 []uint8 [226 130 172 108 105 116 101] 7
s1: eshello, world
hello: hello
s1: ESHELLO, WORLD
*/
