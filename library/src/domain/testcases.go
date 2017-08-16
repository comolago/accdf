package domain

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

type Details struct {
	Severity                  string `xml:"severity,attr"            json:"severity,attr"`
	Minimum_health_percentage string `xml:"minimum_health,attr" json:"minimum_health,attr"`
	Responsible               string `xml:"responsible,attr"               json:"responsible,attr"`
	//_                   string `xml:"#text"               json:"#text"`
}

type Test struct {
	Name       string      `xml:"name,attr" json:"name,attr"`
	Title      string      `xml:"title,attr" json:"title,attr"`
	TestSteps  []TestStep  `xml:"TestSteps>TestStep"  json:"TestSteps>TestStep"`
	Parameters []Parameter `xml:"Parameters>Parameter"  json:"Parameters>Parameter"`
}

func (t *Test) AddTestStep(step TestStep) {
	t.TestSteps = append(t.TestSteps, step)
}

func (t *Test) AddParameter(parameter Parameter) {
	t.Parameters = append(t.Parameters, parameter)
}

type TestStep struct {
	Id       string `xml:"id,attr"          json:"id,attr"`
	Dst_IP   string `xml:"dst_ip,attr"          json:"dst_ip,attr"`
	Protocol string `xml:"protocol,attr"    json:"protocol,attr"`
	Port     string `xml:"port,attr"        json:"port,attr"`
	//	Description string `xml:"description,attr" json:"description,attr"`
}

type Parameter struct {
	Name   string `xml:"name,attr"          json:"name,attr"`
	Engine string `xml:"engine,attr"        json:"engine,attr"`
	Source string `xml:"source,attr"        json:"source,attr"`
	Query  string `xml:"query,attr"         json:"query,attr"`
}

type Inject struct {
	Classes []Class `xml:"Classes"             json:"Classes"`
}

func (i *Inject) AddClass(class string) {
	var cls Class
	cls.Class = class
	i.Classes = append(i.Classes, cls)
}

type Class struct {
	Class string `xml:"Class"               json:"Class"`
}

type TestCase struct {
	Id           string       `xml:"id,attr"             json:"id,attr"`
	Name         string       `xml:"name,attr"           json:"name,attr"`
	Title        string       `xml:"title,attr"          json:"title,attr"`
	Benchmark    string       `xml:"benchmark,attr"      json:"benchmark,attr"`
	Description  string       `xml:"Description"         json:"Description"`
	Detail       Details      `xml:"Details"             json:"Details"`
	Tests        []Test       `xml:"Tests>Test"          json:"Tests>Test"`
	Inject       Inject       `xml:"Inject"              json:"Inject"`
	Dependencies []Dependency `xml:"Dependencies>Dependency"  json:"Dependencies>Dependency"`
}

type Dependency struct {
	Name      string `xml:"name,attr"          json:"name,attr"`
	Type      string `xml:"type,attr"          json:"type,attr"`
	TestSuite string `xml:"testsuite,attr"     json:"testsuite,attr"`
	Title     string `xml:"title,attr"         json:"title,attr"`
}

func (t *TestCase) AddDependency(dep Dependency) {
	t.Dependencies = append(t.Dependencies, dep)
}

func (t *TestCase) AddTest(name string, title string) {
	var ts Test
	ts.Name = name
	ts.Title = title
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
