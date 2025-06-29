package utilities

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ParseInputFile(filepath string) []string {
	var inputs []string
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}
	for _, d := range strings.Split(string(data), "\n") {
		inputs = append(inputs, d)
	}
	return inputs

}
