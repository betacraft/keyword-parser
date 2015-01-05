package parser

import (
	"bytes"
	"log"
	"strings"
)

// Intention is to pass this a whole file as string, and a map of translated keywords and original english keywords
// Function is supposed to find out occurences of translated keywords, and later on replace them with original
// english keywords. This should also be tunable for laungages. This one is using ruby as a target.

func ProcessKeywords(input string, keywords map[string]string) (int, string) {
	var quotes int
	var singleQuote int
	var blockCommentStart int
	var blockCommentEnd int
	var keywordsFound int
	var processedStringBuffer bytes.Buffer

	LineCommentString := "#" //for ruby
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if strings.HasPrefix(strings.Trim(line, " "), LineCommentString) {
			log.Println("skip line because commented ", line)
			processedStringBuffer.WriteString(line + "\n")
			continue
		}
		log.Println("chekcing line ", line)
		fields := strings.Fields(line)
		for index, field := range fields {
			quotes += strings.Count(field, "\"")
			singleQuote += strings.Count(field, "'")
			blockCommentStart += strings.Count(field, "/*")
			blockCommentEnd += strings.Count(field, "*/")
			log.Println("checking word ", field)
			if quotes%2 == 0 && singleQuote%2 == 0 && blockCommentStart == blockCommentEnd {
				log.Println("word without quotes, and comments found")
				//token is Ok to check for reserved word, as it is not comments or inside a string
				if keyword := keywords[field]; keyword != "" {
					log.Println("Keyword found ", field, " for ", keyword)
					if index == (len(fields) - 1) {
						processedStringBuffer.WriteString(keyword)
					} else {
						processedStringBuffer.WriteString(keyword + " ")
					}
					keywordsFound += 1
					continue
				}
			} else {

				log.Println("this word in in quotes, or comments")
			}
			if index == (len(fields) - 1) {
				processedStringBuffer.WriteString(field)
			} else {
				processedStringBuffer.WriteString(field + " ")
			}
		}
		processedStringBuffer.WriteString("\n")
	}
	return keywordsFound, processedStringBuffer.String()
}
