package mytoolkits

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// ExampleRemoveDuplicatesFromSliceString gets a slice of string and remove duplicate
// string from it
func ExampleRemoveDuplicatesFromSliceString(sliceStr []string) []string {

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

//------------------------------------------------------------------------------

// ExampleSaveStringToFile gets filepath (path+filename) and values that should be stored
func ExampleSaveStringToFile(filepath, values string) {

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

//------------------------------------------------------------------------------

// ExampleReadTextFileLoadToMem ...
func ExampleReadTextFileLoadToMem(filename string) (string, []string) {

	var slcResult []string
	var sResult string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		tmpstr := scanner.Text()
		slcResult = append(slcResult, tmpstr)
		sResult = sResult + tmpstr
	}
	return sResult, slcResult
}

//------------------------------------------------------------------------------

// ExampleFindPatternString ....
func ExampleFindPatternString(s, start, end string) ([]int, string) {

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

//------------------------------------------------------------------------------

// ExampleFindPatternSliceString ....
func ExampleFindPatternSliceString(s, start, end string) []string {

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

//------------------------------------------------------------------------------

// ExampleSaveSliceStrToFile ...
func ExampleSaveSliceStrToFile(fname string, slc []string) {
	var strTmp string
	for _, item := range slc {
		strTmp = strTmp + item + "\n"
	}
	SaveStringToFile(fname, strTmp)
}

//------------------------------------------------------------------------------

// ExampleFindBetweenString ...
func ExampleFindBetweenString(s, start, end string) ([]int, string) {

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
