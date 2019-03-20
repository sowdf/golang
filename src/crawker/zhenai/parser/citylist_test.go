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
	expectedCitys := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("Url 第%d个Url与%sUrl不相符，获得url是：%s", i, url, result.Requests[i].Url)
		}
	}

	for i, city := range expectedCitys {
		if result.Items[i] != city {
			t.Errorf("Url 第%d个城市与%s城市不相符，获得城市是：%s", i, city, result.Items[i])
		}
	}
	if len(result.Requests) != cityListSize {
		t.Errorf("城市个数应该有 %d,但是请求个数只有 %d", cityListSize, len(result.Requests))
	}

	if len(result.Items) != cityListSize {
		t.Errorf("城市个数应该有 %d,但是请求个数只有 %d", cityListSize, len(result.Items))
	}
	//创建文件
	//createFile(contents,"citylist_test_data.html")
}
