package gomind

import (
	"encoding/xml"
)

type XMapContent struct {
	XMLName xml.Name `xml:"xmap-content"`
	Sheets  []Sheet  `xml:"sheet"`
}

type Sheet struct {
	ID        string `xml:"id,attr"`
	Title     string `xml:"title"`
	Timestamp int    `xml:"timestamp,attr"`
	Topic     Topic  `xml:"topic"`
}

type Topic struct {
	ID             string  `xml:"id,attr"`
	StructureClass string  `xml:"structure-class"`
	Timestamp      int     `xml:"timestamp,attr"`
	Title          string  `xml:"title"`
	Children       []Topic `xml:"children>topics>topic"`
	Href           string  `xml:"http://www.w3.org/1999/xlink href,attr"`
}
