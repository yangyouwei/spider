package utillib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func Removespaceline(s string) string {
	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
	return re.ReplaceAllString(s, "")
}
func Removespace(s string) string {
	return strings.ReplaceAll(s," ","")
}

func RemoveEnter(s string) string {
	return strings.ReplaceAll(s,"\n","")
}

func StringFmt(s string)string  {
	s = Removespaceline(s)
	return RemoveEnter(s)
}

//获取网页html
func GetRootHtml(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return ""
	}

	if resp.StatusCode != 200 {
		return ""
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return ""
	}
	return  string(body)
}
