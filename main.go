package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"bytes"
	"log"
	"os"
	"regexp"
)

func main() {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	newFileNames := getFormatedFileName()

	sourceNames := make([]string, 0)
	regexFileName := regexp.MustCompile(`^\d+\.[a-zA-Z]+$`)
	for _, entry := range entries {
		fileName := entry.Name()
		matched := regexFileName.MatchString(fileName)
		if matched {
			sourceNames = append(sourceNames, fileName)
		}
	}

	// sort targetFiles by name
	sort.Slice(sourceNames, func(i, j int) bool {
		return getNumberedFileName(sourceNames[i]) < getNumberedFileName(sourceNames[j])
	})

	for i, sourceFile := range sourceNames {
		sourceName := strings.Split(sourceFile, ".")
		oldSourceName := sourceName[0]
		ext := sourceName[1]

		if i >= len(newFileNames) {
			log.Println("The given file names are not enough for the source file names. Ignoring the rest of the files.")
			break
		}

		newName := fmt.Sprintf("%s-%s.%s", oldSourceName, newFileNames[i], ext)
		err = os.Rename(fmt.Sprintf("./%s", sourceFile), fmt.Sprintf("./%s", newName))
		if err != nil {
			log.Println("Error renaming file", err)
		}
		log.Println("Renamed", sourceFile, "to", newName)
	}

}

func getNumberedFileName(name string) int {
	num := strings.Split(name, ".")[0]
	res, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal("getNumberedFileName() - error converting string to int", err)
	}
	return res
}

func getFormatedFileName() []string {
	contents, err := os.ReadFile("./name.txt")
	if err != nil {
		log.Fatal("Error reading file name.txt", err)
	}

	// Split content into lines and trim spaces
	lines := bytes.Split(bytes.TrimSpace(contents), []byte("\n"))

	// Convert [][]byte to []string and format each line
	stringLines := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmedLine := bytes.TrimSpace(line)
		if len(trimmedLine) > 0 {
			stringLines = append(stringLines, toDashCase(trimNonLetterPrefix(string(trimmedLine))))
		}
	}

	return stringLines
}

var vietnameseCharMap = map[rune]string{
	'à': "a", 'á': "a", 'ạ': "a", 'ả': "a", 'ã': "a",
	'â': "a", 'ầ': "a", 'ấ': "a", 'ậ': "a", 'ẩ': "a", 'ẫ': "a",
	'ă': "a", 'ằ': "a", 'ắ': "a", 'ặ': "a", 'ẳ': "a", 'ẵ': "a",
	'è': "e", 'é': "e", 'ẹ': "e", 'ẻ': "e", 'ẽ': "e",
	'ê': "e", 'ề': "e", 'ế': "e", 'ệ': "e", 'ể': "e", 'ễ': "e",
	'ì': "i", 'í': "i", 'ị': "i", 'ỉ': "i", 'ĩ': "i",
	'ò': "o", 'ó': "o", 'ọ': "o", 'ỏ': "o", 'õ': "o",
	'ô': "o", 'ồ': "o", 'ố': "o", 'ộ': "o", 'ổ': "o", 'ỗ': "o",
	'ơ': "o", 'ờ': "o", 'ớ': "o", 'ợ': "o", 'ở': "o", 'ỡ': "o",
	'ù': "u", 'ú': "u", 'ụ': "u", 'ủ': "u", 'ũ': "u",
	'ư': "u", 'ừ': "u", 'ứ': "u", 'ự': "u", 'ử': "u", 'ữ': "u",
	'ỳ': "y", 'ý': "y", 'ỵ': "y", 'ỷ': "y", 'ỹ': "y",
	'đ': "d",
}

func removeDiacritics(s string) string {
	var result strings.Builder
	for _, r := range s {
		lowerR := unicode.ToLower(r)
		if replacement, ok := vietnameseCharMap[lowerR]; ok {
			if unicode.IsUpper(r) {
				result.WriteString(strings.ToUpper(replacement))
			} else {
				result.WriteString(replacement)
			}
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func toDashCase(s string) string {
	s = removeDiacritics(s)
	s = strings.ToLower(s)

	// Replace spaces and other non-alphanumeric characters with dashes
	var result strings.Builder
	var lastWasDash bool
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(r)
			lastWasDash = false
		} else if !lastWasDash {
			result.WriteRune('-')
			lastWasDash = true
		}
	}

	return strings.Trim(result.String(), "-")
}

func trimNonLetterPrefix(s string) string {
	for i, r := range s {
		if unicode.IsLetter(r) {
			return s[i:]
		}
	}
	return "" // Return empty string if no letter is found
}
