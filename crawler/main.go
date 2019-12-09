package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp , err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()



	if resp.StatusCode == http.StatusOK {

		all , err := ioutil.ReadAll(resp.Body)
		reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())

		if err != nil {
			panic(err)
		}
		printCityList(all)
		//fmt.Printf("%s",all)
	}
}
func printCityList(content []byte) {
	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

	all := compile.FindAllSubmatch(content, -1)
	for _ , m := range all {
		fmt.Printf("url is %s && city name is %s \n" , m[1],m[2] )
	}
}
