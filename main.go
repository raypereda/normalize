package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// tr '[:upper:]' '[:lower:]' |
// sed -e 's/^/__label__/g' |
// sed -e "s/'/ ' /g"
// 	-e 's/"//g'
// 	-e 's/\./ \. /g'
// 	-e 's/<br \/>/ /g'
// 	-e 's/,/ , /g'
// 	-e 's/(/ ( /g'
// 	-e 's/)/ ) /g'
// 	-e 's/\!/ \! /g'
// 	-e 's/\?/ \? /g'
// 	-e 's/\;/ /g'
// 	-e 's/\:/ /g'
// tr -s " " | myshuf

func main() {
	space := regexp.MustCompile(`\s+`)

	// `ABC'''"""...,,,()!?;:    END`
	// normalized: `__label__ABC ' ' ' . . . , , , ( ) ! ? END`

	file := os.Stdin
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ToUpper(line)
		line = "__label__" + line
		line = strings.ReplaceAll(line, "'", " ' ")
		line = strings.ReplaceAll(line, `"`, "")
		line = strings.ReplaceAll(line, ".", " . ")
		line = strings.ReplaceAll(line, "<br />", " ")
		line = strings.ReplaceAll(line, ",", " , ")
		line = strings.ReplaceAll(line, "(", " ( ")
		line = strings.ReplaceAll(line, ")", " ) ")
		line = strings.ReplaceAll(line, "!", " ! ")
		line = strings.ReplaceAll(line, "?", " ? ")
		line = strings.ReplaceAll(line, ";", " ")
		line = strings.ReplaceAll(line, ":", " ")
		line = space.ReplaceAllString(line, " ")
		fmt.Println(line)
	}
}
