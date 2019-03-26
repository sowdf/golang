package parser

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func createFile(contents []byte, filename string) {
	file, e := os.Create(filename)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprintln(writer, string(contents))
}

func TestParseCityList(t *testing.T) {
	contents, e := ioutil.ReadFile("citylist_test_data.html")
	//contents, e := fetcher.Fetch("http://www.zhenai.com/zhenghun");
	if e != nil {
		panic(e)
	}
	//fmt.Printf("%s\n",contents);
	result := ParseCityList(contents)

	const cityListSize = 470

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("Url 第%d个Url与%sUrl不相符，获得url是：%s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Requests) != cityListSize {
		t.Errorf("城市个数应该有 %d,但是请求个数只有 %d", cityListSize, len(result.Requests))
	}

	//创建文件
	//createFile(contents,"citylist_test_data.html")
}
