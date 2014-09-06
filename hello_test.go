package main

import (
	"fmt"
	"io/ioutil"
	"./mypackage"
	"strings"
)

// The Bla1 example tests if the output matches
func ExampleBla1() {
	fmt.Println("Bla")

	// Output:
	//  Bla
}

func ExampleMap() {
	m := make(map[string]int)
	m[" _ " +
	  "| |" +
	  "|_|" ] = 0

	m["   " +
	  "  |" +
	  "  |" ] = 1

	m[" _ " +
	  " _|" +
	  "|_ " ] = 2

	m[" _ " +
	  " _]" +
	  " _|" ] = 3

	m["   " +
	  "|_|" +
	  "  |" ] = 4

	m[" _ " +
	  "|_ " +
	  " _|"] = 5

	m[" _ " +
	  "|_ " +
	  "|_|" ] = 6

	m[" _ " +
	  "  |" +
	  "  |" ] = 7

	m[" _ " +
	  "|_|" +
	  "|_|" ] = 8

	m[" _ " +
	  "|_|" +
	  " _|" ] = 9
}

func ExampleReadFile() {
	// http://golang.org/pkg/io/ioutil/#ReadFile
	bytes, readError := ioutil.ReadFile("numbers_original.txt")

	if readError != nil {
		panic(readError)
	}

	lines:= strings.Split(string(bytes), "\n")
	key:= 0

	var number_lines [3]string

	for _, val:= range lines {
		if key == 3 {
			GetNumber(number_lines)
			key = 0
		} else {
			number_lines[key] = val
			key++
		}
	}

//	fmt.Println(bytes)
//	fmt.Println(len(bytes))
//	fmt.Println(string(bytes))

	// Output:
	// [32 95 10 124 32 124 10 124 95 124 10 10 10]
	// 13
	//	 _
	//	| |
	//	|_|

}

func GetNumber (values [3]string) {
	fmt.Println(values)
}

func ExampleFunctions() {
	fmt.Println(mypackage.Mysplit("oh no"))
	fmt.Println(mypackage.Myjoin(mypackage.Mysplit("oh no")))
	// Output:
	// [oh no]
	// oh|no
}
