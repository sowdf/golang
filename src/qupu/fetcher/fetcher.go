package fetcher

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) (contents []byte, err error) {
	//创建客户端
	client := &http.Client{}
	//创建请求
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	//设置请求头
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.28 Safari/537.36")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	request.Header.Add("Cookie", "__cfduid=def3c13d276a10848b8f1258e083d36e21548585278; Hm_lvt_dca7dc99d8ac55393ef7fbc057d85ffb=1548585281,1548586001; bdshare_firstime=1549453858855; PHPSESSID=uitopl4vcj2fi34ek6pci6ivr3; Hm_lvt_40108d7e4cc326e04eecdd70da888247=1553696084,1553696118,1553696841; Murl=-m0XB7fGVrb8WLe%3DkTv8V3XF%2F7v2NypDU0yRAYh0P%2Ft1WLa--tqjr93M%2F7ELNRS4B; damon_token_2019=160105452; Hm_lpvt_40108d7e4cc326e04eecdd70da888247=1553736667")
	request.Header.Add("Host", "www.qupu123.com")

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
