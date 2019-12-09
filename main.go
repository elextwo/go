package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://www.weibo.com")
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {

		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n" ,bytes)
}
func printCityList(content []byte) {
	regexp.Compile(``)
}