package domain

type Platform struct {
	Id      string `xml:"id,attr"      json:"id,attr"`
	Label   string `xml:"label,attr"   json:"label,attr"`
	Version string `xml:"version,attr" json:"version,attr"`
}

type Require struct {
	Type string `xml:"type,attr"               json:"type,att"`
	Id   string `xml:"id,attr"                 json:"id,attr"`
}

type Fingerprint struct {
	Type string `xml:"type,attr"               json:"type,attr"`
	Hash string `xml:"hash,attr"               json:"hash,attr"`
}

type Benchmark struct {
	Id           string        `xml:"id,attr"                 json:"id,attr"`
	Name         string        `xml:"name,attr"               json:"name,attr"`
	Platforms    []Platform    `xml:"Platforms>Platform"      json:"Platforms>Platform"`
	Requires     []Require     `xml:"Requires>Require"        json:"Requires>Require"`
	Fingerprints []Fingerprint `xml:"Fingerprints>Fingerprint json:"Fingerprints>Fingerprint"`
	Privileges   string        `xml:"Privileges"              json:"Privileges"`
	Description  string        `xml:"Description"             json:"Description"`
	Rationale    string        `xml:"Rationale"               json:"Rationale"`
}

func (b *Benchmark) AddPlatform(id string, label string, version string) {
	var p Platform
	p.Id = id
	p.Label = label
	p.Version = version
	b.Platforms = append(b.Platforms, p)
}

func (b *Benchmark) GetId() string {
	return b.Id
}

func (b *Benchmark) SetId(i string) {
	b.Id = i
}

func (b *Benchmark) GetName() string {
	return b.Name
}

func (b *Benchmark) SetName(n string) {
	b.Name = n
}
