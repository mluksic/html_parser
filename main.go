package main

import (
	"bytes"
	"fmt"
	link "github.com/mluksic/html_parser/pkg"
	"log"
	"os"
)

func main() {
	file, _ := os.ReadFile("examples/exp2/exp2.html")
	r := bytes.NewReader(file)

	links, err := link.Parse(r)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(links)
}
