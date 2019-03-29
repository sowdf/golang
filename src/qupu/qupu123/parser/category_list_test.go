package parser

import (
	"qupu/common"
	"qupu/fetcher"
	"testing"
)

func TestParseCategoryList(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.qupu123.com/")

	if err != nil {
		panic(err)
	}

	const num = 42

	common.CreateFile("category_list_test_data.html", contents)

	parseResult := ParseCategoryList(contents, "")

	if len(parseResult.Requests) != num {
		t.Errorf("获取到的个数是：%d，实际的个数是%d", len(parseResult.Requests), num)
	}
}
