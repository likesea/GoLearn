package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type div struct {
	XMLName xml.Name `xml:"profile"`
	// First letter must be capital. Cannot use `content`
	Content string `xml:",innerxml"`
}

func main() {
	d := div{}
	xmlContent, _ := ioutil.ReadFile("ConfigProfile.xml")
	err := xml.Unmarshal(xmlContent, &d)
	if err != nil {
		panic(err)
	}
	fmt.Println("XMLName:", d.XMLName)
	fmt.Println("Content:", d.Content)
}
