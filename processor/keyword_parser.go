package processor

import (
	"bytes"
	"strconv"
	"strings"
)

// Intention is to pass this a whole file as string, and a map of translated keywords and original english keywords
// Function is supposed to find out occurences of translated keywords, and later on replace them with original
// english keywords. This should also be tunable for laungages. This one is using ruby as a target.

func ProcessKeywords(input string, keywords map[string]string, language string) (int, string) {
	var quotes int
	var singleQuote int
	var blockCommentStart int
	var blockCommentEnd int
	var keywordsFound int
	var processedStringBuffer bytes.Buffer
	var LineCommentString string
	switch language {
	case "ruby":
		LineCommentString = "#" //for ruby
	default:
		LineCommentString = "//"
	}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if strings.HasPrefix(strings.Trim(line, " "), LineCommentString) {
			processedStringBuffer.WriteString(line + "\n")
			continue
		}
		fields := strings.Fields(line)
		for index, field := range fields {
			quotes += strings.Count(field, "\"")
			singleQuote += strings.Count(field, "'")
			blockCommentStart += strings.Count(field, "/*")
			blockCommentEnd += strings.Count(field, "*/")
			if quotes%2 == 0 && singleQuote%2 == 0 && blockCommentStart == blockCommentEnd {
				//token is Ok to check for reserved word, as it is not comments or inside a string
				if keyword := keywords[field]; keyword != "" {
					if index == (len(fields) - 1) {
						processedStringBuffer.WriteString(keyword)
					} else {
						processedStringBuffer.WriteString(keyword + " ")
					}
					keywordsFound += 1
					continue
				} else {
					found := false
					for k, v := range keywords {
						if _, err := strconv.Atoi(v); err == nil {
							if strings.Contains(field, k) {
								field = strings.Replace(field, k, v, -1)
								found = true
							}
						}
					}
					if found {
						processedStringBuffer.WriteString(field + " ")
						continue
					}
				}
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
