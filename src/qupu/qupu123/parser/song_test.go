package parser

import (
	"fmt"
	"qupu/common"
	"qupu/fetcher"
	"testing"
)

const filename = "song_test_data.html"

func TestParseSong(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.qupu123.com/minge/sanzi/p316606.html")

	if err != nil {
		panic(err)
	}

	common.CreateFile(filename, contents)

	parseResult := ParseSong(contents, "http://www.qupu123.com/tongsu/sanzi/p336145.html")

	fmt.Printf("%v\n", parseResult.Item)

}
