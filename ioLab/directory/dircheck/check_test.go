package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Fixture struct {
	TestSummaries []testSummary `json:"test_cases"`
}
type testSummary struct {
	Title     string     `json:"title"`
	TestCases []testCase `json:"case"`
}
type testCase struct {
	Result bool   `json:"result"`
	Path   string `json:"path"`
}

var fixtures *Fixture

func TestMain(m *testing.M) {
	prepareTestData()
	os.Exit(m.Run())
}

func prepareTestData() {
	b, err := ioutil.ReadFile("testdata/fixture.json")
	if err != nil {
		log.Fatal(err)
	}
	f := new(Fixture)
	if err := json.Unmarshal(b, f); err != nil {
		log.Fatal(err)
	}
	//log.Printf("fixture: %#v", f)
	fixtures = f
}

func TestIsExists(t *testing.T) {
	for _, summary := range fixtures.TestSummaries {
		t.Logf("Running test case `%s`", summary.Title)
		for _, testCase := range summary.TestCases {
			assert.Equal(t, testCase.Result, IsExists(testCase.Path))
		}
	}
}
