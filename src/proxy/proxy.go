package main

import (
	"crawker/fetcher"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

var ipRe = regexp.MustCompile(`<td>([\d]+\.[\d]+\.[\d]+\.[\d]+)</td>
      <td>([\d]+)</td>`)

func readMaze(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]string, row)
	for i := range maze {
		maze[i] = make([]string, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%s", &maze[i][j])
		}
	}

	return maze
}

func main() {

	//
	//112.85.128.71:9999
	proxyFetch("222.94.149.206", "3128", "131111333114562223431")

	return
	contents, e := fetcher.Fetch("https://www.xicidaili.com/nn/")
	matchs := ipRe.FindAllSubmatch(contents, -1)
	//sfmt.Printf("%s\n",matchs)
	if e != nil {
		panic(e)
	}
	for i, item := range matchs {
		ip := string(item[1])
		port := string(item[2])
		func(i int) {
			fmt.Printf("ip:%s:%s\n", ip, port)

		}(i)
	}
	//fmt.Printf("%s",contents)

}

func proxyFetch(ip string, port string, i string) {
	req_url := "https://m.tititoy2688.com//user/register?userName=czhui11148" + i + "&password=a123456&captcha=&inviteCode=346017"
	fmt.Printf("url: %s", req_url)

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://" + ip + ":" + port)
	}
	transport := &http.Transport{Proxy: proxy}

	c := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", req_url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3724.8 Safari/537.36")
	if err != nil {
		fmt.Println(err)
	}

	res, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", body)
}
