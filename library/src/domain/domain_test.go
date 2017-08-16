package domain

import (
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func TestDomain(t *testing.T) {
	testcases_test()
	benchmarks_test()
}
