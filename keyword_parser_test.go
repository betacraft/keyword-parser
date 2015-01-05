package parser

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestProcessKeywords(t *testing.T) {
	bytes, err := ioutil.ReadFile("test.rb")
	if err != nil {
		t.Error("Error reading file test.rb ", err.Error())
	}
	keywordsFound, processedString := ProcessKeywords(string(bytes), map[string]string{"यासाठी": "for"})
	if keywordsFound != 1 {
		t.Error("Keyword found is", keywordsFound, ", instead of 1")
	}
	log.Println(processedString)
}
