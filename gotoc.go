package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Toc struct {
	header []int
	title  []string
}

func writeHeader() string {
	header := "**Table of Contents**"
	return header
}

func removeSpecialCharacters(str string) string {
	r, _ := regexp.Compile("([a-zA-Z]+)")
	cleanedStr := r.FindAllString(str, -1)
	return strings.Join(cleanedStr, " ")
}

func main() {
	var startsWith bool
	var line string
	var doctoc = new(Toc)

	file, err := os.Open("example.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		startsWith = strings.HasPrefix(line, "#")
		if startsWith == true {
			fmt.Println("TRUE")
			headerCount := strings.Count(line, "#")
			fmt.Println(strings.Count(line, "#"))
			content := line[headerCount:]
			fmt.Println(content)
			cleanedContent := removeSpecialCharacters(content)
			fmt.Println(cleanedContent)
			doctoc.header = append(doctoc.header, headerCount)
			doctoc.title = append(doctoc.title, cleanedContent)
		}
	}

	fmt.Println(doctoc)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
