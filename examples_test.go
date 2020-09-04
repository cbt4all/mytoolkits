package mytoolkits

import (
	"fmt"

	"github.com/cbt4all/mytoolkits"
)

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

	// Output:
	// [aaaaaa bbbbbb cccccc dddddd]

}

func ExampleSaveStringToFile() {
	// Create a string of data
	str := "aaaaaa bbbbbb cccccc"

	// Save the string str to the file myfile.txt as a string
	mytoolkits.SaveStringToFile("./myfile.txt", str)
}

func ExampleSaveSliceStrToFile() {
	// Create a slice of string
	strSlice := []string{
		"aaaaaa",
		"bbbbbb",
		"cccccc",
		"aaaaaa",
		"dddddd",
		"cccccc",
	}

	// Save the slice of string  to the file myfile.txt as a string
	mytoolkits.SaveSliceStrToFile("./myfile.txt", strSlice)
}

func ExampleReadTextFileLoadToMem() {
	// assume that the file myfile.txt in in the current directory
	// and the content of the file is aaaaaa bbbbbb cccccc
	// Read the file and returns a string and a slice of string
	str, sliceStr := mytoolkits.ReadTextFileLoadToMem("./myfile.txt")
	fmt.Println(str)
	fmt.Println(sliceStr)

	//Output will be:
	//aaaaaa bbbbbb cccccc
	//[aaaaaa bbbbbb cccccc]
}

func ExampleFindPatternString() {
	// Assume that we have a text like below (note this is a string not a slice of string)
	// We ou need to find a string with has specific format or pattern. The patterns starts with specific string
	// and ends with another string. For example, here we are trying to find:
	// <response status="success"><result>DATA1</result></response>
	// in the given text. Our string starts with "<response status="success"><result>"
	// and ends with "</result></response>"
	strdata := `garbage string like blah blah not important
		<response status="success"><result>DATA1</result></response>
		garbage string like blah blah not important`

	// the function FindPatternString gets three parameters: FindPatternString(s, start, end string). In this example:
	// s will be strdata
	// start will be <response status="success"><result>
	// end will be </result></response>
	var strSlice []string
	var idx int
	l := len(strdata)
	for i := 0; i < l; i++ {
		iresult, sresult := mytoolkits.FindPatternString(strdata[idx:], `<response status="success"><result>`, `</result></response>`)
		if iresult[0] != -1 && iresult[1] != -1 {
			idx = idx + iresult[1]
			strSlice = append(strSlice, sresult)
		} else {
			break
		}
	}

	fmt.Println(strSlice)
	// Output will be: [<response status="success"><result>DATA1</result></response>]
}

func ExampleFindPatternSliceString() {
	// To understand this example, you need to check the example of the function FindPatternString
	// Assume that we have a text like below (note this is a string not a slice of string)
	// We ou need to find all strings with has specific format or pattern. The patterns starts with specific string
	// and ends with another string. For example, here we are trying to find:
	//  <response status="success"><result>DATA1</result></response>
	//	<response status="success"><result>DATA2</result></response>
	//	<response status="success"><result>DATA3</result></response>
	//	<response status="success"><result>DATA4</result></response>
	// in the given text. Our strings all start with "<response status="success"><result>"
	// and end with "</result></response>"
	strdata := `garbage string like blah blah not important text
		<response status="success"><result>DATA1</result></response>
		garbage string like blah blah not important text
		<response status="success"><result>DATA2</result></response>
		<response status="success"><result>DATA3</result></response>
		garbage string like blah blah not important text
		garbage string like blah blah not important text
		<response status="success"><result>DATA4</result></response>
		garbage string like blah blah not important text`

	// the function FindPatternSliceString gets three parameters: FindPatternString(s, start, end string). In this example:
	// s will be strdata
	// start will be <response status="success"><result>
	// end will be </result></response>
	strSlice := mytoolkits.FindPatternSliceString(strdata, `<response status="success"><result>`, `</result></response>`)

	for _, item := range strSlice {
		fmt.Println(item)
	}

	// Output will be:
	//<response status="success"><result>DATA1</result></response>
	//<response status="success"><result>DATA2</result></response>
	//<response status="success"><result>DATA3</result></response>
	//<response status="success"><result>DATA4</result></response>
}

func ExampleFindBetweenString() {
	// Assume that we have a text like below (note this is a string not a slice of string)
	strdata := `aaaaaaaaa bbbbbbb ccccccc index1 dddddd eeeeee ffffff ggggg index2 
	hhhhhh iiiiii jjjjjj `

	iresult, str := mytoolkits.FindBetweenString(strdata, " index1 ", " index2 ")

	if iresult[0] != -1 && iresult[1] != -1 {
		fmt.Println("index1 is in: ", iresult[0])
		fmt.Println("index2 is in: ", iresult[1])
		fmt.Println("The string is: ", str)
	} else {
		fmt.Println("Nothing is found")
	}
}
