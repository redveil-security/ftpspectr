package utilities

import (
	//"bytes"
	"regexp"
	//"fmt"
)

func ExamineContents(filecontent string) (bool, string) {
	// Add more regex patterns here
	patterns := []string {
		`\b\d{3}-\d{2}-\d{4}\b`, // SSN regex
		`\b[\w.-]+:[^\s:@]{1,100}\b`, // username:password regex	
	}
	// Compile each pattern
	var regexList []*regexp.Regexp
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		regexList = append(regexList, re)
	}
	
	for _, re := range regexList {
		matches := re.FindAllString(filecontent, -1)
		for _, match := range matches {
			return true, match
		}
	}
	// no match
	return false, ""	
}
