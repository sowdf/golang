package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCategory(t *testing.T) {

	//contents, err := fetcher.Fetch("http://www.qupu123.com/qiyue/jita/")
	contents, err := ioutil.ReadFile("category_test_data.html")

	if err != nil {
		panic(err)
	}

	//common.CreateFile("category_test_data.html",contents)

	parseResult := ParseCategory(contents)

	for _, item := range parseResult.Requests {
		fmt.Printf("%v"+
			"\n", item)
	}

}
