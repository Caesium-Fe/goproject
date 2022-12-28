package learn

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Cnm() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("err", err)
	}
	closer := resp.Body
	bytes, err := ioutil.ReadAll(closer)
	fmt.Println(string(bytes))
}
