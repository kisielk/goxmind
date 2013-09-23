package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"github.com/kisielk/goxmind"
	"log"
	"os"
)

type encoder interface {
	Encode(v interface{}) error
}

func main() {
	format := flag.String("f", "xml", "format to output, one of: xml, json")
	flag.Parse()
	mind, err := goxmind.Open(flag.Arg(0))

	var enc encoder
	switch *format {
	case "xml":
		xmlEnc := xml.NewEncoder(os.Stdout)
		xmlEnc.Indent("", "\t")
		enc = xmlEnc
	case "json":
		jsonEnc := json.NewEncoder(os.Stdout)
		enc = jsonEnc
	default:
		log.Fatalln("invalid format:", *format)
	}

	content, err := mind.Content()
	if err != nil {
		log.Fatal(err)
	}
	err = enc.Encode(content)
	if err != nil {
		log.Fatal(err)
	}
}
