package processor

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
	keywordsFound, processedString := ProcessKeywords(string(bytes), map[string]string{"यासाठी": "for", "जर": "if", "शेवट": "end", "छाप": "print", "आत": "in"})
	if keywordsFound != 7 {
		t.Error("Keyword found is", keywordsFound, ", instead of 7")
	}
	log.Println(processedString)
}
