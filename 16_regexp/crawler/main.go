package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"regexp"
)

func main() {
	url := "https://tv.sohu.com/hotmovie/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	sHTML, _ := ioutil.ReadAll(resp.Body)
	// sHTML, _ := ioutil.ReadFile("data.txt")

	// data-pb-txid="toppage_film_play" pb-click="">灵欲专车</a>
	reg := regexp.MustCompile(`title="(.*)">(.*)</a>`)
	// reg := regexp.MustCompile("title=\"(.+)\">(.+)</a>")
	result := reg.FindAllStringSubmatch(string(sHTML), -1)

	for _, item := range result {
		fmt.Println(item[1])
	}
}
