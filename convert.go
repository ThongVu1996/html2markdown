package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: convert <input.html> <output.md>")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	data, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	md, err := htmltomarkdown.ConvertString(string(data))
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(output, []byte(md), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✔ Converted:", input, "→", output)
}

