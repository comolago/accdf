package domain

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

/*func TestNewTestCase(t *testing.T) {
	fmt.Printf("\nTestNewTestCase\n\n")
	fmt.Print("Filling a Testcase Struct with data and converting it to XML ...\n\n")
	tc := new(TestCase)
	tc.Benchmark = "abenchmark"
	tc.Title = "Testcase Title"
	tc.Name = "example-testcase"
	tc.Id = "0000-3333-4445-2221"
	tc.Description = "An example testcase"
	tc.Detail.Minimum_health_percentage = "66%"
	tc.Detail.Responsible = "Telecom"
	tc.Detail.Severity = "critical"
	tc.Inject.AddClass("An example TAG class")
	tc.AddTest("first-test", "First Test")
	var step TestStep
	step.Dst_IP = "192.168.5.6"
	step.Id = "1"
	step.Port = "389"
	step.Protocol = "tcp"
	tc.Tests[0].AddTestStep(step)
	var parameter Parameter
	parameter.Name = "DST_IP"
	parameter.Engine = "XPATH"
	parameter.Source = "topology.xml"
	parameter.Query = "/Topology/Server[${TARGET_HOST}]/Peers[@port=389,@port=636]/@ip"
	tc.Tests[0].AddParameter(parameter)
	var dependency Dependency
	dependency.Name = "testcase2"
	dependency.TestSuite = "testsuite3"
	dependency.Title = "TestCase Title"
	dependency.Type = "testcase"
	tc.AddDependency(dependency)
	testToXML(t, tc)
	testToJson(t, tc)
}*/

func TestLoadFromFile(t *testing.T) {
	fmt.Printf("\nTestLoadFromFile\n")
	BenchmarkFile, err := filepath.Abs("testcases_test.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	file, err := os.Open(BenchmarkFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	tc := new(TestCase)
	err = tc.FromXML(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	testToXML(t, tc)
	testToJson(t, tc)
}

func testToXML(t *testing.T, tc *TestCase) {
	fmt.Printf("\ntestToXML()\n\n")
	data, _ := tc.ToXML()
	fmt.Printf("\n%s\n", data)
}

func testToJson(t *testing.T, tc *TestCase) {
	fmt.Printf("\ntestToJson()\n\n")
	data, _ := tc.ToJson()
	fmt.Printf("\n%s\n", data)
}
