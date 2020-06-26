
##############################################################################################
RemoveDuplicatesFromSliceString

The function removeDuplicatesFromSliceString gets a slice of string, removes the duplicate data from it
and return a slice of strings without any duplicate data

//---------------------------------
Simple example of a main.go file
//----------------------------------

func main() {

	sliceStringTmp := []string{
		"111111",
		"222222",
		"333333",
		"444444",
		"111111",
		"555555",
		"666666",
		"333333",
		"333333",
		"555555",
	}
	sliceStrIPs := removeDuplicatesFromSliceString(sliceStringTmp)
	fmt.Println(sliceStrIPs)
}

##############################################################################################
SaveStringToFile

The function saveStringToFile get two strings: a filepath and a value. Then save the value/data to the file which should be mentioned in the filepath

//---------------------------------
Simple example of a main.go file
//---------------------------------

func main() {
	saveStringToFile("./folder1/file1.txt", "value or data")
}

//---------------------------------

Version 0.001:
- The path of the file should already be there. In this version we don't create any folder.

##############################################################################################
ReadTextFileLoadToMem

The function readTextFileLoadToMem takes a string (which is a file name - can be included the path and the file name).
Then read the file as a string file and return a slice of string

//---------------------------------
Simple example of a main.go file
//---------------------------------

func main() {
	sliceStr := readTextFileLoadToMem("IPs.txt")
	for _, item := range sliceStr {
		fmt.Println(item)
	}
}

//---------------------------------

Version 0.001
- the function gets a string (filename) and returns slice of string

##############################################################################################

