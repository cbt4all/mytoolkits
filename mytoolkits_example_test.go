package mytoolkits

import (
	"fmt"

	"github.com/cbt4all/mytoolkits"
)

// ExampleRemoveDuplicatesFromSliceString gets a slice of string and removes all duplicates
// string from it
func ExampleRemoveDuplicatesFromSliceString() {

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

	// RemoveDuplicatesFromSliceString gets a slice of string and removes all duplicates in it
	tmpSlice := mytoolkits.RemoveDuplicatesFromSliceString(strSlice)
	fmt.Println(tmpSlice)

	// Output:	[aaaaaa bbbbbb cccccc dddddd]
}
