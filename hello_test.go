package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"./mypackage"
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
	bytes, readError := ioutil.ReadFile("numbers.txt")

	if readError != nil {
		panic(readError)
	}

	fmt.Println(bytes)
	fmt.Println(len(bytes))
	fmt.Println(string(bytes))

	// Output:
	// [32 95 10 124 32 124 10 124 95 124 10 10 10]
	// 13
	//	 _
	//	| |
	//	|_|

}

func ExampleFunctions() {
	fmt.Println(mypackage.Mysplit("oh no"))
	fmt.Println(mypackage.Myjoin(mypackage.Mysplit("oh no")))
	// Output:
	// [oh no]
	// oh|no
}
