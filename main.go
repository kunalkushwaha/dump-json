package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Simple tool to print formatted json string from stdin")
	fmt.Println("")
	fmt.Println("  $ cmd | dump-json")
	fmt.Println("  $ dump-json < <unformatted-json-file>")
	fmt.Println("")
}
func main() {

	var jsonText bytes.Buffer
	var jsonFile string
	flag.Usage = usage
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			break
		}
		jsonFile = jsonFile + string(line)
	}
	err := json.Indent(&jsonText, []byte(jsonFile), "", "\t")
	if err != nil {
		fmt.Println("BAD Json", jsonFile)
	} else {
		jsonText.WriteTo(os.Stdout)
	}

}
