package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"parser/processor"
	"strings"
)

var keywords = map[string]string{}

func readKeywords(name string) error {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		keywords[fields[0]] = fields[1]
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Arguments - parser source_file language keywords_file")
		return
	}
	bytes, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println("Error reading input file ", err.Error())
		return
	}
	if err = readKeywords(args[2]); err != nil {
		fmt.Println("Error reading keywords ", err)
		return
	}
	_, processedString := processor.ProcessKeywords(string(bytes), keywords, args[1])
	defer os.Remove(args[0] + ".tmp")
	err = ioutil.WriteFile(args[0]+".tmp", []byte(processedString), os.ModePerm)
	if err != nil {
		fmt.Println("Error writing tmp file ", err.Error())
		return
	}

	fmt.Println("Going to run translated file - \n ", processedString)

	out, err := exec.Command("ruby", args[0]+".tmp").CombinedOutput()
	if err != nil {
		fmt.Println("Error running ruby script  ", err.Error(), string(out))
		return
	}
	fmt.Println("Ruby script output - \n ", string(out))

}
