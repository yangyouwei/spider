package yuezi

import (
	"fmt"
	"fushi/download"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/antchfx/htmlquery"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type candian struct {
	name string
	keyword []string
	miaoshu string
	shicai string
	buzhou []string
	bigpci string
	smalpic string
	video string
}

var cs []candian

func Done()  {
	//for n,i:=range urls{
	//	switch n {
	//	case 0:
	//		handeljson(getjson(i),"主食")
	//	case 1:
	//		handeljson(getjson(i),"主菜")
	//	case 2:
	//		handeljson(getjson(i),"炖汤")
	//	case 3:
	//		handeljson(getjson(i),"饮品")
	//	}
	//}

	//ss := readtable()
	//for n,i:=range ss{
	//	urlstr := fmt.Sprintf("https://papi.mama.cn/wap/auxiliary_food/get_recipe_detail?recipe_id=%s&type=moon_meal",i[1])
	//	handeljson(getjson(urlstr),n)
	//}
	//
	//wirtetofile()

	download.DoneDownload()

	//ss := readtable()
	//for n,i:=range ss{
	//	var lists []string
	//	for _,y:=range i{
	//		if	y == "" {
	//			continue
	//		}
	//		a := getstring(y)
	//		//fmt.Println(y)
	//		lists = append(lists,a)
	//	}
	//	fmt.Println(n,lists)
	//	//wirtetofile(lists,n)
	//	for x,z:=range lists{
	//		if z == "" {
	//			continue
	//		}
	//		download.Downloadfile(z,n,x)
	//	}
	//}
}

func getstring(s string) string {
	nodep,err := htmlquery.Parse(strings.NewReader(s))
	if err !=nil {
		fmt.Println(err)
	}
	imgnode :=htmlquery.FindOne(nodep,"//img")
	if imgnode == nil{
		return ""
	}
	img := htmlquery.SelectAttr(imgnode, "src")
	//fmt.Println(htmlquery.InnerText(nodep))
	return img
}

func RemoveEnter(s string) string {
	return strings.ReplaceAll(s,"\n","")
}

func Removespaceline(s string) string {
	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
	return re.ReplaceAllString(s, "")
}

func handeljson(s string,n int)  {
	//a := gjson.Get(s,"data.recipe_info")
	//for _,i:=range a.Array(){
	//	fmt.Println(fenlei,i.Get("recipe_id"))
	//}

	var a candian
	//fmt.Println(gjson.Get(s,"data.recipe_list.video_type"),n)

	a.name = fmt.Sprint(gjson.Get(s,"data.recipe_list.recipe_name"))
	k := gjson.Get(s,"data.recipe_list.recipe_keyword")
	for _,i := range k.Array() {
		a.keyword = append(a.keyword,fmt.Sprint(i))
	}
	a.smalpic = fmt.Sprint(gjson.Get(s,"data.recipe_list.recipe_share_img"))
	a.bigpci = fmt.Sprint(gjson.Get(s,"data.recipe_list.recipe_img"))
	a.miaoshu = fmt.Sprint(gjson.Get(s,"data.recipe_list.recipe_effect"))
	a.shicai = fmt.Sprint(gjson.Get(s,"data.recipe_list.recipe_ingredients"))
	for _,x:= range gjson.Get(s,"data.recipe_list.recipe_steps").Array(){
		a.buzhou = append(a.buzhou,fmt.Sprint(x))
	}
	if fmt.Sprint(gjson.Get(s,"data.recipe_list.video_type"))  == "3"{
		a.video = fmt.Sprint(gjson.Get(s,"data.recipe_list.recipe_video"))
	}
	fmt.Println(a)
	cs=append(cs,a)
}

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

func readtable() [][]string {
	f, err := excelize.OpenFile("yuezi.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	a := f.GetSheetName(3)

	s := f.GetRows(a)
	return s
}

func wirtetofile(s []string,n int)  {
	f, err := excelize.OpenFile("yuezi.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	a := f.GetSheetName(1)
	f.SetSheetRow(a,"J"+fmt.Sprint(n+1),&s)

	if err := f.SaveAs("yuezi.xlsx"); err != nil {
		fmt.Println(err)
	}
}