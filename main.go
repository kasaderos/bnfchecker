package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kasaderos/parsegen"
)

var filename = flag.String("f", "", "bnf file")

func main() {
	flag.Parse()
	fmt.Println(*filename)
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	bnfBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	parser, err := parsegen.Generate(bnfBytes)
	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(os.Stdin)
	input := make([]byte, 0, 1024)
	fmt.Println("enter input:")
	for {
		c, err := in.ReadByte()
		if err == io.EOF {
			break
		}
		input = append(input, c)
	}
	pd, err := parser.Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	pd.Print()
}
