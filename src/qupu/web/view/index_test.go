package view

import (
	"model"
	"os"
	model2 "qupu/web/model"
	"testing"
)

func TestIndex(t *testing.T) {

	resultView := CreateSearchResultView("index.html")

	file, err := os.Create("index_test.html")

	expected := model.Spectrum{
		Id:          "336145",
		Type:        "qupu123",
		Title:       "​30女儿情（双谱）",
		From:        "新编中国声乐作品第15集  ",
		FromUrl:     "http://www.qupu123.com/tongsu/sanzi/p336145.html",
		Composition: "许镜清",
		Uploader:    "瓦莲金娜AA  ",
		Views:       3561,
		UploadDate:  "2018-05-21",
		ImageUrls:   []string{"http://www.qupu123.com/Public/Uploads/2018/05/21/5b02162abcb9b.jpg", "http://www.qupu123.com/Public/Uploads/2018/05/21/5b0216332b8d1.jpg"},
		ClickNum:    0,
		Lyrics:      "杨洁",
	}

	if err != nil {
		panic(err)
	}

	data := model2.SearchResult{}

	for i := 0; i < 10; i++ {
		data.Songs = append(data.Songs, expected)
	}

	err = resultView.Render(file, data)

	if err != nil {
		panic(err)
	}

}
