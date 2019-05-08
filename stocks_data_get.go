package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func httpGet(id string) {
	url := "http://hq.sinajs.cn/list=" + id
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	//fmt.Println(string(body))
	ss := strings.Split(string(body), "\"")
	items := strings.Split(ss[1], ",")

	yesterday, _ := strconv.ParseFloat(items[2], 32)
	now, _ := strconv.ParseFloat(items[3], 32)
	top, _ := strconv.ParseFloat(items[4], 32)
	bottom, _ := strconv.ParseFloat(items[5], 32)
	rate := (now - yesterday) / yesterday * 100.0

	fmt.Printf("now: %.2f, rate: %.2f, top: %.2f, bottom: %.2f\n", now, rate, top, bottom)
}

func main() {
	httpGet("sh601318")
}
