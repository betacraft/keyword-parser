package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"parser/processor"
)

var keywords = map[string]string{"यासाठी": "for",
	"जर":   "if",
	"शेवट": "end",
	"छाप":  "print",
	"आत":   "in"}

func main() {
	args := os.Args[1:]
	//TODO  handle argument length
	bytes, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println("Error reading input file ", err.Error())
		return
	}
	_, processedString := processor.ProcessKeywords(string(bytes), keywords)
	defer os.Remove(args[0] + ".tmp")
	err = ioutil.WriteFile(args[0]+".tmp", []byte(processedString), os.ModePerm)
	if err != nil {
		fmt.Println("Error writing tmp file ", err.Error())
		return
	}

	fmt.Println("Going run translated file - \n ", processedString)

	out, err := exec.Command("ruby", args[0]+".tmp").CombinedOutput()
	if err != nil {
		fmt.Println("Error running ruby script  ", err.Error(), string(out))
		return
	}
	fmt.Println("Ruby script output - \n %s", out)

}
