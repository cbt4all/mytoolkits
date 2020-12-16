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

// SaveStringToFile gets fname (path+filename) and values that should be stored
// in the file
func SaveStringToFile(fname, values string) {

	f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
func ReadTextFileLoadToMem(fname string) (string, []string) {

	// Slice of string version of the file
	var slcResult []string

	// String version of the file
	var sResult string

	// Open the file
	file, err := os.Open(fname)
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

// FindPatternString find a specific patterns in a text. To do this, it gets a string 's'
// and 'start', 'end' as two other string as prapeters and search in the entier 's' to find a substring
// and begins with 'start' and finishes with 'end'. Note 'start' ≠ 'end'
// If nothing is found {-1,-1} is returned
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

	// iresult is the index of 'start' and 'end'
	// sresult is the string found between 'start' and 'end'
	return iresult, sresult
}

// FindPatternSliceString uses the function FindPatternString to find all strings with a specific paterns in
// a string 's'. The patern should begin with 'start' and finish with 'end' while 'start' and 'end' are the function
// parameter. Note 'start' ≠ 'end'
// If nothing is found, nil is returned
func FindPatternSliceString(s, start, end string) []string {

	var sliceResult []string
	var idx int
	l := len(s)

	for i := 0; i < l; i++ {
		iresult, sresult := FindPatternString(s[idx:], start, end)
		if iresult[0] != -1 && iresult[1] != -1 {
			idx = idx + iresult[1]
			sliceResult = append(sliceResult, sresult)
		} else {
			break
		}
	}
	return sliceResult
}

// SaveSliceStrToFile gets fname (path+filename) and value as slice of string that should be stored
// in the file
func SaveSliceStrToFile(fname string, value []string) {
	var strTmp string
	for _, item := range value {
		strTmp = strTmp + item + "\n"
	}
	SaveStringToFile(fname, strTmp)
}

// FindBetweenString gets three strings 's', 'start' and 'end' as parameters then search 's' to
// find a string begins with 'start' and after that continue searching to find a string begins
// with 'end' and returns the index of what is found
// If nothing is found {-1,-1} is returned
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

	// iresult is the index of 'start' and 'end'
	// sresult is the string found between 'start' and 'end'
	return iresult, sresult
}
