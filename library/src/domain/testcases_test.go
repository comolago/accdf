package domain

import (
	"fmt"
	"os"
	"path/filepath"
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

func testcases_test() {
	fmt.Printf("\ntestcases_test\n")
	BenchmarkFile, err := filepath.Abs("../../test/xml/testcases/000.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputFile, err := os.Open(BenchmarkFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer inputFile.Close()
	tc := new(TestCase)
	err = tc.FromXML(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data, _ := tc.ToXML()
	outputFile, err := os.Create("../../tmp/xml/testcases/000.xml")
	check(err)
	defer outputFile.Close()
	n, err := outputFile.Write(data)
	check(err)
	fmt.Printf("wrote %d bytes\n", n)
	outputFile.Sync()
}
