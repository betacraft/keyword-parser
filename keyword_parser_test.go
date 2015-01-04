package parser

import (
	"io/ioutil"
	"testing"
)

func TestKeywordParser(t *testing.T) {
	bytes, err := ioutil.ReadFile("test.rb")
	if err != nil {
		t.Error("Error reading file test.rb ", err.Error())
	}
	keywordsFound := ParseKeywords(string(bytes), map[string]string{"यासाठी": "for"})
	if keywordsFound != 1 {
		t.Error("Keyword found is", keywordsFound, ", instead of 1")
	}
}
