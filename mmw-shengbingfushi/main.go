package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

type fenlei struct {
	name string
	jibing string
}

var fulldata []fenlei


func main()  {
	jibing := []string{"感冒","咳嗽","便秘","腹泻","积食","免疫力","开胃","清火去燥"}
	for _,i := range jibing{
		getall(i)
	}
	//fmt.Println(fulldata)
	wirtetoexecl()
}

func getall(s string)  {
	switch s {
	case "感冒":
		url := fmt.Sprintf("https://papi.mama.cn/wap/auxiliary_food/get_recipe_pagination?id=%s&page=1&pageSize=200&type=auxiliary_food","39")
		jstring := getjson(url)
		getname(jstring,s)
	case "咳嗽":
		url := fmt.Sprintf("https://papi.mama.cn/wap/auxiliary_food/get_recipe_pagination?id=%s&page=1&pageSize=200&type=auxiliary_food","40")
		jstring :=getjson(url)
		getname(jstring,s)
	case "便秘":
		url := fmt.Sprintf("https://papi.mama.cn/wap/auxiliary_food/get_recipe_pagination?id=%s&page=1&pageSize=200&type=auxiliary_food","41")
		jstring :=getjson(url)
		getname(jstring,s)
	case "腹泻":
		url := fmt.Sprintf("https://papi.mama.cn/wap/auxiliary_food/get_recipe_pagination?id=%s&page=1&pageSize=200&type=auxiliary_food","42")
		jstring :=getjson(url)
		getname(jstring,s)
	case "积食":
		url := fmt.Sprintf("https://papi.mama.cn/wap/auxiliary_food/get_recipe_pagination?id=%s&page=1&pageSize=200&type=auxiliary_food","43")
		jstring :=getjson(url)
		getname(jstring,s)
	case "免疫力":
		url := fmt.Sprintf("https://papi.mama.cn/wap/auxiliary_food/get_recipe_pagination?id=%s&page=1&pageSize=200&type=auxiliary_food","44")
		jstring :=getjson(url)
		getname(jstring,s)
	case "开胃":
		url := fmt.Sprintf("https://papi.mama.cn/wap/auxiliary_food/get_recipe_pagination?id=%s&page=1&pageSize=200&type=auxiliary_food","45")
		jstring :=getjson(url)
		getname(jstring,s)
	case "清火去燥":
		url := fmt.Sprintf("https://papi.mama.cn/wap/auxiliary_food/get_recipe_pagination?id=%s&page=1&pageSize=200&type=auxiliary_food","46")
		jstring :=getjson(url)
		getname(jstring,s)
	default:
		return
	}
}

//获取接口json数据，字符串类型
func getjson(u string) string {
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}

func getname (jstring string,jibing string)  {
	for _,i:= range gjson.Get(jstring,"data.recipe_info").Array() {
		var a fenlei
		a.jibing = jibing
		a.name = fmt.Sprint(i.Get("recipe_name"))
		fulldata = append(fulldata,a)
	}
}

func wirtetoexecl()  {
	f := excelize.NewFile()

	f.NewSheet("Sheet1")

	for n,i:= range fulldata{
		f.SetCellValue("Sheet1","A"+fmt.Sprint(n+1),i.jibing)
		f.SetCellValue("Sheet1","B"+fmt.Sprint(n+1),i.name)
	}

	err := f.SaveAs("data.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("finished.")
}

