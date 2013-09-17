package goxmind

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
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
	StructureClass string  `xml:"structure-class,attr"`
	Timestamp      int     `xml:"timestamp,attr"`
	Title          string  `xml:"title"`
	Children       []Topic `xml:"children>topics>topic"`
	Href           string  `xml:"http://www.w3.org/1999/xlink href,attr"`
}

type XMind struct {
	f *zip.ReadCloser
}

func (m *XMind) Close() error {
	return m.f.Close()
}

func (m *XMind) Content() (*XMapContent, error) {
	for _, f := range m.f.File {
		if f.Name == "content.xml" {
			content, err := f.Open()
			if err != nil {
				return nil, err
			}
			dec := xml.NewDecoder(content)

			result := new(XMapContent)
			err = dec.Decode(result)
			return result, err
		}
	}
	return nil, fmt.Errorf("no content found")
}

func Open(name string) (*XMind, error) {
	r, err := zip.OpenReader(name)
	return &XMind{r}, err
}
