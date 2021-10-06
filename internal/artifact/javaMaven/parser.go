package javaMaven

import (
	"io/ioutil"
	"encoding/xml"
)

type properties struct {
	XMLName         xml.Name    `xml:"properties"`
	ArtifactPath    string      `xml:"app.release.artifact"`
}

type Project struct {
	XMLName        xml.Name      `xml:"project"`
	Name           string        `xml:"name"`
	Version        string        `xml:"version"`
	Description    string        `xml:"description"`
	Properties     properties    `xml:"properties"`
}

func ParseFile(filePath string) (Project, error) {
	var Properties Project

	xmlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Properties, err
	}

	err = xml.Unmarshal(xmlFile, &Properties)

	return Properties, err
}