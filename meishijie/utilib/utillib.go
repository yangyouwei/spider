package utilib

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tidwall/gjson"
	"net/http"
	"io/ioutil"
	"strings"
)



func GetJson(url string,bodyc string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, strings.NewReader(bodyc))
	// 自定义Header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded",)
	req.Header.Set("User-Agent", "model:android;Version:meishij7.3.4;udid:867886020672907;channel:huawei;imsi:null;statusbarHeight:25",)
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

func ParseJsonObject(j string,keystring string) string {
	return  fmt.Sprint(gjson.Get(j,keystring))
}

func ParseJsonList(j string,keystring string) []gjson.Result {
	return gjson.Get(j,keystring).Array()
}

func WriteToExecl(s [][]string,name string)  {

	f, err := excelize.OpenFile(name)
	if err != nil {
		fmt.Println(err)
	}
	a := f.GetSheetName(1)

	for n,i:= range s {
		if len(i) >3 {
			for y,x:=range i{
				switch y {
				case 0:
					f.SetCellValue(a,"A"+fmt.Sprint(n+1),x)
				case 1:
					f.SetCellValue(a,"B"+fmt.Sprint(n+1),x)
				case 2:
					f.SetCellValue(a,"C"+fmt.Sprint(n+1),x)
				case 3:
					//fmt.Println("----")
					//fmt.Println(x)
					//fmt.Println("----------")
					f.SetCellValue(a,"D"+fmt.Sprint(n+1),x)
				}
			}
		}
		f.SetSheetRow(a,"A"+fmt.Sprint(n+1),&i)
	}

	if err := f.SaveAs(name); err != nil {
		fmt.Println(err)
	}
}

func WriteMapToExecl(m Caipin,id string,name string,n int)  {

	f, err := excelize.OpenFile(name)
	if err != nil {
		fmt.Println(err)
	}
	a := f.GetSheetName(1)
	var line []string
	line =append(line,id)
	line =append(line,m.Name)
	line =append(line,m.Img)

	if len(m.Renqun) != 0 {
		for _,i:= range m.Renqun{
			a := a+i+"\n"
			line =append(line,a)
		}
	}

	f.SetSheetRow(a,"A"+fmt.Sprint(n+1),&line)

	if err := f.SaveAs(name); err != nil {
		fmt.Println(err)
	}
}