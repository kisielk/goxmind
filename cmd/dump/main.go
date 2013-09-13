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
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	content := gomind.XMapContent{}
	dec := xml.NewDecoder(f)
	err = dec.Decode(&content)
	if err != nil {
		log.Fatal(err)
	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "\t")
	err = enc.Encode(content)
	if err != nil {
		log.Fatal(err)
	}
}
