package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//get()
	//getWithParam()
	post()
}

func get() {
	response, err := http.Get("https://lab.yiidii.cn/api/free/wb/hot")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}

func getWithParam() {
	// 接口
	uri, err := url.ParseRequestURI("https://lab.yiidii.cn/api/free/wb/hot")
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
		return
	}
	// 参数
	data := url.Values{}
	data.Set("topN", "3")

	// 设置参数
	uri.RawQuery = data.Encode()

	fmt.Println(uri.String())

	response, err := http.Get(uri.String())
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}

func post() {
	url := "https://lab.yiidii.cn/api/wm"
	contentType := "application/json"
	data := `{
		"bizCode": "douYinRmWaterMark",
		"links": [
			"https://v.douyin.com/NwscyGy/"
		]
	}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

func customClient() {
	//client := &http.Client{
	//	CheckRedirect: redirectPolicyFunc,
	//}
	//resp, err := client.Get("http://example.com")
	//req, err := http.NewRequest("GET", "http://example.com", nil)
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	//resp, err := client.Do(req)
}

/*
要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个Transport：

Client和Transport类型都可以安全的被多个goroutine同时使用。出于效率考虑，应该一次建立、尽量重用。

*/
func customTransport() {
	//tr := &http.Transport{
	//	TLSClientConfig:    &tls.Config{RootCAs: pool},
	//	DisableCompression: true,
	//}
	//client := &http.Client{Transport: tr}
	//resp, err := client.Get("https://example.com")
}
