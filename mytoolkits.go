package mytoolkits

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// MyLogger ...

// RemoveDuplicatesFromSliceString ....
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

//------------------------------------------------------------------------------

// SaveStringToFile ...
func SaveStringToFile(filepath, values string) {

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

// ReadTextFileLoadToMem ...
func ReadTextFileLoadToMem(filename string) (string, []string) {

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

		var idx int
		l := len(output)
		for i := 0; i < l; i++ {
			iresult, sresult := FindPatternString(output[idx:], `<response status="success"><result>`, `</result></response>`)
			idx = idx + iresult[1]
			fmt.Println(sresult)
		}
	*/
}
