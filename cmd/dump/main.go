package main

import (
	"encoding/xml"
	"flag"
	"gomind"
	"log"
	"os"
)

func main() {
	flag.Parse()
	mind, err := gomind.Open(flag.Arg(0))

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "\t")
	content, err := mind.Content()
	if err != nil {
		log.Fatal(err)
	}
	err = enc.Encode(content)
	if err != nil {
		log.Fatal(err)
	}
}
