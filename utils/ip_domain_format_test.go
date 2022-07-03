package utils

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	fmt.Println("https://baidu.com")
}

func TestURLFormatAssign(t *testing.T) {
	valid,url := URLFormatAssign("http://google.com") 
	fmt.Println(valid,url)

	valid,url = URLFormatAssign("https://taobao.com") 
	fmt.Println(valid,url)

	valid,url = URLFormatAssign("http://google.com:1080") 
	fmt.Println(valid,url)

	valid,url = URLFormatAssign("https://taobao.com:1080") 
	fmt.Println(valid,url)

	valid,url = URLFormatAssign("https://192.168.0.1") 
	fmt.Println(valid,url)

	valid,url = URLFormatAssign("https://192.168.2.1:1080") 
	fmt.Println(valid,url)

	valid,url = URLFormatAssign("http://192.168.0.1") 
	fmt.Println(valid,url)


	valid,url = URLFormatAssign("http://192.168.0.1:1080") 
	fmt.Println(valid,url)


	valid,url = URLFormatAssign("google.com") 
	fmt.Println(valid,url)

	valid,url = URLFormatAssign("google.com:8080") 
	fmt.Println(valid,url)

	valid,url = URLFormatAssign("192.168.0.1") 
	fmt.Println(valid,url)

	valid,url = URLFormatAssign("192.168.0.1:1080") 
	fmt.Println(valid,url)

	valid,url = URLFormatAssign("jd.com")
	fmt.Println(valid,url)

	// if !valid {
    //     t.Errorf("got %v, wanted %v",valid,true)
    // }
}
