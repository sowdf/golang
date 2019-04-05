package fetcher

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) (contents []byte, err error) {
	<-rateLimiter
	//创建客户端
	client := &http.Client{}
	//创建请求
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	//设置请求头
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.28 Safari/537.36")
	request.Header.Add("Cookie", "timestamp=02030357048468189705; _ga=GA1.2.805527182.1553010937; ciu_key=ac97d6f4-be4a-4bdd-ba8a-ec69cbfa3751$112.5.201.212; ticket=afe2cd4f-f7c6-424a-bf3e-94dd17c58a10; inviteriid=1064280; JSESSIONID=bjsqrcaausktdzd3pu9mxlk6")
	request.Header.Add("Origin", "https://m.tititoy2688.com")
	request.Header.Add("Host", "m.tititoy2688.com")
	//请求

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	//关闭读取
	defer response.Body.Close()

	//读取内容到buffer 里面
	bodyReader := bufio.NewReader(response.Body)
	//获取编码
	e := determineEncoding(bodyReader)
	//转换
	newReader := transform.NewReader(bodyReader, e.NewEncoder())

	return ioutil.ReadAll(newReader)

}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, e := r.Peek(1024)
	if e != nil {
		return unicode.UTF8
	}

	e2, _, _ := charset.DetermineEncoding(bytes, "")

	return e2
}
