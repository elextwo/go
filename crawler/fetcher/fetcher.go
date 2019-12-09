package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

func fetch (url string) ([]byte ,error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil , err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil , fmt.Errorf("wrong status code : %d" ,resp.StatusCode )
	}

}
func determineEncoding( r io.Reader) encoding.Encoding {
	bytes, e := bufio.NewReader(r).Peek(1024)
	if e != nil {
		panic(e)
	}
	charset.Dete
}