package mytoolkits_test

import (
	"fmt"

	"github.com/cbt4all/mytoolkits"
)

func Example_removeDuplicatesFromSliceString() {

	// Using RemoveDuplicatesFromSliceString...
	// Create a slice of string
	strSlice := []string{
		"aaaaaa",
		"bbbbbb",
		"cccccc",
		"aaaaaa",
		"dddddd",
		"cccccc",
	}

	// RemoveDuplicatesFromSliceString gets a slice of string and remove all duplicates in it
	tmpSlice := mytoolkits.RemoveDuplicatesFromSliceString(strSlice)
	fmt.Println(tmpSlice)

	// Output:
	// [aaaaaa bbbbbb cccccc dddddd]
}

func Example_saveStringToFile() {

	// Using SaveStringToFile...
	// Create a string of data
	str := "aaaaaa bbbbbb cccccc"

	// SaveStringToFile saves the string str to the file myfile.txt as a string
	mytoolkits.SaveStringToFile("./myfile.txt", str)

}
