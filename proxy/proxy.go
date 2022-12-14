package proxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
)

const (
	gStaticUa = "Mozilla/5.0 (Linux; U; Android 4.4.4; zh-cn; M351 Build/KTU84P) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30"
)

var (
	gQgNetKey = "7AC773EC"
)

type IpGetRespT struct {
	Code   int    `json:"Code"`
	TaskID string `json:"TaskID"`
	Num    int    `json:"Num"`
	Data   []struct {
		IP       string `json:"IP"`
		Port     string `json:"port"`
		Deadline string `json:"deadline"`
		Host     string `json:"host"`
	} `json:"Data"`
}

func HttpProxy(addr, reqUrl string) {
	urli := url.URL{}
	urlproxy, _ := urli.Parse("http://" + addr)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
		Timeout: 5 * time.Second,
	}
	rqt, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		logs.Error("请求失败:%v", err)
		return
	}

	rqt.Header.Add("User-Agent", gStaticUa)
	//处理返回结果
	response, _ := client.Do(rqt)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	logs.Debug("http:", string(body))
	return

}

type IpInfoT struct {
	Ip        string `json:"ip"`
	Country   string `json:"country"`
	Area      string `json:"area"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Isp       string `json:"isp"`
	Timestamp int    `json:"timestamp"`
}

//本机IP
func Httplocal() (*IpInfoT, error) {
	client := &http.Client{}
	rqt, err := http.NewRequest("GET", "http://myip.top", nil)
	if err != nil {
		logs.Debug("请求失败:%v", err)
		return nil, err
	}

	rqt.Header.Add("User-Agent", "Lingjiang")
	//处理返回结果
	response, _ := client.Do(rqt)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {

		return nil, err
	}
	res := &IpInfoT{}
	json.Unmarshal(body, res)
	logs.Debug("本机:", string(body))
	return res, nil

}

func addLocalIpToWhite() {

	local, err := Httplocal()
	if err != nil {
		logs.Error("get local ip addr failed.")
	}

	url := fmt.Sprintf("https://proxy.qg.net/whitelist/add?Key=%s&IP=%s", gQgNetKey, local.Ip)
	body, err := httplib.Get(url).Bytes()
	if err != nil {
		logs.Error("添加白名单失败:%v", err)
		return
	}
	res := &IpGetRespT{}
	json.Unmarshal(body, res)
	if res.Code != 0 {
		logs.Error("添加白名单失败:%v", string(body))
	} else {
		logs.Debug("添加本机IP：[%v]到白名单", local.Ip)
	}
}

func RunProxyGet() {

	addLocalIpToWhite()

	reqUrl := fmt.Sprintf("https://proxy.qg.net/allocate?Key=%s&Num=%d", gQgNetKey, 10)
	resp, err := httplib.Get(reqUrl).
		Bytes()
	if err != nil {
		logs.Error("请求失败:%v", err)
		return
	}

	logs.Debug("收到IP:%v", string(resp))

	response := IpGetRespT{}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logs.Error("结果解析错误:%v", err)
	}

	for _, ip := range response.Data {
		HttpProxy(ip.Host, "https://myip.top")
	}
}

func RunProxyTask() {
	UseDps("o9hc0gc0l7enc936s65e", "x7pe9y578p7d50wcwpnr6k81whnnyn05")

}
