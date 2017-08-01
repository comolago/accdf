package domain

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

type Details struct {
	Severity            string `xml:"severity,attr"            json:"severity,attr"`
	Tolerancepercentage string `xml:"tolerancepercentage,attr" json:"tolerancepercentage,attr"`
	Owner               string `xml:"owner,attr"               json:"owner,attr"`
	//_                   string `xml:"#text"               json:"#text"`
}

type Test struct {
	Name      string     `xml:"name,attr" json:"name,attr"`
	Label     string     `xml:"label,attr" json:"label,attr"`
	TestSteps []TestStep `xml:"TestSteps>TestStep"  json:"TestSteps>TestStep"`
}

type TestStep struct {
	Id          string `xml:"id,attr"          json:"id,attr"`
	IP          string `xml:"ip,attr"          json:"ip,attr"`
	Protocol    string `xml:"protocol,attr"    json:"protocol,attr"`
	Port        string `xml:"port,attr"        json:"port,attr"`
	Description string `xml:"description,attr" json:"description,attr"`
}

type TestCase struct {
	Id          string  `xml:"id,attr"             json:"id,attr"`
	Name        string  `xml:"name,attr"           json:"name,attr"`
	Benchmark   string  `xml:"benchmark,attr"      json:"benchmark,attr"`
	Description string  `xml:"Description"         json:"Description"`
	Detail      Details `xml:"Details"           json:"Details"`
	Tests       []Test  `xml:"Tests>Test"          json:"Tests>Test"`
}

func (t *TestCase) AddTest(name string, label string) {
	var ts Test
	ts.Name = name
	ts.Label = label
	t.Tests = append(t.Tests, ts)
}

func (t *TestCase) GetId() string {
	return t.Id
}

func (t *TestCase) SetId(i string) {
	t.Id = i
}

func (t *TestCase) GetName() string {
	return t.Name
}

func (t *TestCase) SetName(n string) {
	t.Name = n
}

func (t *TestCase) FromXML(reader io.Reader) error {
	if err := xml.NewDecoder(reader).Decode(&t); err != nil {
		return err
	}
	return nil
}

func (t *TestCase) FromJson(reader io.Reader) error {
	if err := json.NewDecoder(reader).Decode(&t); err != nil {
		return err
	}
	return nil
}

func (t *TestCase) ToXML() ([]byte, error) {
	data, err := xml.Marshal(&t)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t *TestCase) ToJson() ([]byte, error) {
	data, err := json.Marshal(&t)
	if err != nil {
		return nil, err
	}
	return data, nil
}
