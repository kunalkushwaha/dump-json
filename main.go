package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"flag"
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
	flag.Usage = usage
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			break
		}
		err = json.Indent(&jsonText, line, "", "\t")
		if err != nil {
			// In case of json file, just print text as it is, excluding the blank lines
			if len(line) == 0 {
				continue
			}
			fmt.Println(string(line))
		} else {
			jsonText.WriteTo(os.Stdout)
		}
	}
}
