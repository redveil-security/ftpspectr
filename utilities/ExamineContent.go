package utilities

import (
	//"bytes"
	"regexp"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Data struct {
	Pattern []string `yaml:"patterns"`
}

// Read YAML config file with regex patterns
func ExtractPatterns(filepath string) []string {
	var pattern Data
	fHandle, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("[!] The following error was encountered when attempting to read the pattern file. This likely occurred because either no config file was passed or the config file passed doesn't exist: %s\n", err.Error())
	}
	err = yaml.Unmarshal(fHandle, &pattern)
	return pattern.Pattern
}

func ExamineContents(filecontent string, patternFile ...string) (bool, string) {
	// Check if patternFile is passed as arg
	// if not, load default patterns
	// otherwise, load patterns in from config file
	patterns := []string{}
	
	if len(patternFile[0]) == 0 { // First index of patternFile (string array) should be config file being passed
		// Load default regex patterns
		fmt.Println("[-] No config file passed, using default patterns..")
		patterns = []string {
			`\b\d{3}-\d{2}-\d{4}\b`, // SSN regex
			`\b[\w.-]+:[^\s:@]{1,100}\b`, // username:password regex	
		}
	} else {
		// Read regex from config file
		patterns = ExtractPatterns(patternFile[0])
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
