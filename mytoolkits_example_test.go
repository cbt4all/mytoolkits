package mytoolkits

import (
	"fmt"

	"github.com/cbt4all/mytoolkits"
)

// A simple example of how to use RemoveDuplicatesFromSliceString
func ExampleRemoveDuplicatesFromSliceString() {

	// Create a slice of string
	strSlice := []string{
		"aaaaaa",
		"bbbbbb",
		"cccccc",
		"aaaaaa",
		"dddddd",
		"cccccc",
	}

	//Remove all duplicates in the slice strSlice and save the result in tmpSlice
	tmpSlice := mytoolkits.RemoveDuplicatesFromSliceString(strSlice)
	fmt.Println(tmpSlice)

	// Output will be:
	//[aaaaaa bbbbbb cccccc dddddd]
}

// A simple example of how to use SaveStringToFile
func Example_saveStringToFile() {

	// Create a string of data
	str := "aaaaaa bbbbbb cccccc"

	// Save the string str to the file myfile.txt as a string
	mytoolkits.SaveStringToFile("./myfile.txt", str)
}

func Example_readTextFileLoadToMem() {

	// assume that the file myfile.txt in in the current directory
	// and the content of the file is aaaaaa bbbbbb cccccc
	// Read the file and returns a string and a slice of string
	str, slicestr := mytoolkits.ReadTextFileLoadToMem("./myfile.txt", str)

	//Output will be:
	//aaaaaa bbbbbb cccccc
	//[aaaaaa bbbbbb cccccc]
}
