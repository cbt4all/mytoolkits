package mytoolkits

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// RemoveDuplicatesFromSliceString gets a slice of string and remove duplicate
// string from it
func RemoveDuplicatesFromSliceString(sliceStr []string) []string {

	m := make(map[string]bool)

	for _, item := range sliceStr {
		if _, DoesExist := m[item]; !DoesExist {
			m[item] = true
		}
	}

	var result []string

	for item := range m {
		result = append(result, item)
	}
	return result
}

// SaveStringToFile gets filename (path+filename) and values that should be stored
func SaveStringToFile(filename, values string) {

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed open the file:", err)
	}
	defer f.Close()

	if _, err := f.WriteString(values); err != nil {
		log.Println(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

// ReadTextFileLoadToMem gets filepath (path+filename), reads it and
// returns it as a string as well as a slice of string
func ReadTextFileLoadToMem(filename string) (string, []string) {

	// Slice of string version of the file
	var slcResult []string

	// String version of the file
	var sResult string

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a bufio.Scanner to read the file
	scanner := bufio.NewScanner(file)

	// Criteria of having a new line in the file being read
	// is to have new-line charecters (\n,\r or both)
	scanner.Split(bufio.ScanLines)

	// Read the file line by line
	for scanner.Scan() {
		tmpstr := scanner.Text()
		slcResult = append(slcResult, tmpstr) // Creating slice of the string
		sResult = sResult + tmpstr            // Creating a string of the file
	}
	return sResult, slcResult
}

// FindPatternString ....
func FindPatternString(s, start, end string) ([]int, string) {

	var iresult []int
	var sresult string

	sidx := strings.Index(s, start)
	eidx := strings.Index(s, end) + len(end)

	iresult = append(iresult, sidx)
	iresult = append(iresult, eidx)

	if sidx != -1 && eidx != -1 {
		sresult = string(s[sidx:eidx])
	}

	return iresult, sresult

	/*
		An Example of how to use in main function

		var strSlice []string
		var idx int
		l := len(output)
		for i := 0; i < l; i++ {
			iresult, sresult := mytoolkits.FindPatternString(output[idx:], `<response status="success"><result>`, `</result></response>`)
			if iresult[0] != -1 && iresult[1] != -1 {
				idx = idx + iresult[1]
				strSlice = append(strSlice, sresult)
			} else {
				break
			}
		}
	*/
}

// FindPatternSliceString ....
func FindPatternSliceString(s, start, end string) []string {

	var SliceResutl []string
	var idx int
	l := len(s)

	for i := 0; i < l; i++ {
		iresult, sresult := FindPatternString(s[idx:], start, end)
		if iresult[0] != -1 && iresult[1] != -1 {
			idx = idx + iresult[1]
			SliceResutl = append(SliceResutl, sresult)
		} else {
			break
		}
	}
	return SliceResutl
}

// SaveSliceStrToFile ...
func SaveSliceStrToFile(fname string, slc []string) {
	var strTmp string
	for _, item := range slc {
		strTmp = strTmp + item + "\n"
	}
	SaveStringToFile(fname, strTmp)
}

// FindBetweenString ...
func FindBetweenString(s, start, end string) ([]int, string) {

	var iresult []int
	var sresult string

	sidx := strings.Index(s, start)
	eidx := strings.Index(s, end)
	if end == "\n\r" {
		eidx = len(s)
	}

	iresult = append(iresult, sidx)
	iresult = append(iresult, eidx)

	if sidx != -1 && eidx != -1 {
		sresult = string(s[sidx+len(start) : eidx])
	}

	return iresult, sresult

	/*
		An Example of how to use in main function

		iresult, tmp := FindBetweenString(str, " from ", " to ")
			if iresult[0] != -1 && iresult[1] != -1 {
				fmt.Println(tmp)
			} else {
				fmt.Println("ERROR")
			}
	*/
}
