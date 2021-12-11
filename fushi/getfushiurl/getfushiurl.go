package getfushiurl

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
)

var lines [][]string

type caipin struct {
	id string
	name string
	smallpic string
	bigpic string
	video string
	keyword string
	synianling string
	comment string
	shicai string
	buzhou []string
	fitaget string
}

var Caipin []caipin

func GetYueling()  {
	s := Readtemp()
	for _,i:=range s{
		u := fmt.Sprintf("https://papi.mama.cn/wap/auxiliary_food/get_recipe_detail?recipe_id=%s&uid=101062945&type=recipe",i[0])
		body := getjson(u)
		fmt.Println(u,"-------------")
		handeljson(body)
	}
	writetotemp(Caipin)
}

func handeljson(s string)  {
	var a caipin
	idres :=gjson.Get(s,"data.recipe_list.id")
	nameres :=gjson.Get(s,"data.recipe_list.recipe_name")
	keywordres := gjson.Get(s,"data.recipe_list.recipe_keyword")
	smallpicres := gjson.Get(s,"data.recipe_list.recipe_share_img")
	bigpicres :=gjson.Get(s,"data.recipe_list.recipe_img")
	videores := gjson.Get(s,"data.recipe_list.recipe_video")
	commentres := gjson.Get(s,"data.recipe_list.recipe_effect")
	shicaires:= gjson.Get(s,"data.recipe_list.recipe_ingredients")
	buzhoures:= gjson.Get(s,"data.recipe_list.recipe_steps")
	fitageres := gjson.Get(s,"data.recipe_list.fit_age")
	a.id = fmt.Sprint(idres)
	a.name = fmt.Sprint(nameres)
	a.bigpic = fmt.Sprint(bigpicres)
	a.smallpic = fmt.Sprint(smallpicres)
	for _,i:= range keywordres.Array(){
		a.keyword = a.keyword+"\n"+fmt.Sprint(i)
	}
	a.video = fmt.Sprint(videores)
	a.comment = fmt.Sprint(commentres)
	a.fitaget = fmt.Sprint(fitageres)
	a.shicai = fmt.Sprint(shicaires)
	for _,i:= range buzhoures.Array(){
		if fmt.Sprint(i)== "" {
			continue
		}
		t1 := strings.Split(fmt.Sprint(i),"src=\"")
		if len(t1)==1 {
			c := strings.Split(fmt.Sprint(i),"</p>")
			c1 := strings.ReplaceAll(c[0],"<p>","")
			c00:= strings.ReplaceAll(c1,"<br />","")
			strongtag := strings.ReplaceAll(c00,"<strong>","")
			strongtag1 := strings.ReplaceAll(strongtag,"</strong>","")
			divtag := strings.ReplaceAll(strongtag1,"<div>","")
			divtag1 := strings.ReplaceAll(divtag,"</div>","")
			c2 := strings.ReplaceAll(divtag1,"&nbsp;","")
			a.buzhou = append(a.buzhou,c2)
			continue
		}
		t := strings.Split(t1[1],"\"")
		fmt.Println(t[0],"t3")

		c := strings.Split(fmt.Sprint(i),"</p>")
		if len(strings.Split(c[0],"img"))>1 {
			c0:= strings.Split(fmt.Sprint(i),"<img")
			c1 := strings.ReplaceAll(c0[0],"<p>","")
			c00:= strings.ReplaceAll(c1,"<br />","")
			strongtag := strings.ReplaceAll(c00,"<strong>","")
			strongtag1 := strings.ReplaceAll(strongtag,"</strong>","")
			divtag := strings.ReplaceAll(strongtag1,"<div>","")
			divtag1 := strings.ReplaceAll(divtag,"</div>","")
			c2 := strings.ReplaceAll(divtag1,"&nbsp;","")
			bz := c2+"\n"+fmt.Sprint(t[0])
			a.buzhou = append(a.buzhou,bz)
			continue
		}
		c1 := strings.ReplaceAll(c[0],"<p>","")
		c00:= strings.ReplaceAll(c1,"<br />","")
		strongtag := strings.ReplaceAll(c00,"<strong>","")
		strongtag1 := strings.ReplaceAll(strongtag,"</strong>","")
		divtag := strings.ReplaceAll(strongtag1,"<div>","")
		divtag1 := strings.ReplaceAll(divtag,"</div>","")
		c2 := strings.ReplaceAll(divtag1,"&nbsp;","")
		bz := c2+"\n"+fmt.Sprint(t[0])
		a.buzhou = append(a.buzhou,bz)
	}
	Caipin = append(Caipin,a)
}

func getyueling(s string)  {
	res := gjson.Get(s,"data.link_list.3.info.open")
	for _,i:=range res.Array(){
		fmt.Println(i.Get("title"))
		fmt.Println(i.Get("url"))
		//writetotemp(fmt.Sprint(i.Get("title")),fmt.Sprint(i.Get("url")),n)
	}
}

func getchangjian(s string,n int)  {
	res := gjson.Get(s,"data.level_info")
	titleres := gjson.Get(s,"data.title")
	for _,i:=range res.Array(){
		fmt.Println(titleres,i.Get("id"),i.Get("name"))
		//writetotemp(fmt.Sprint(titleres),fmt.Sprint(i.Get("name")),fmt.Sprint(i.Get("id")),n)
	}
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

func writetotemp(s []caipin)  {
	f, err := excelize.OpenFile("temp.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	a := f.GetSheetName(4)
	for n,i:=range s{
		f.SetCellValue(a,"D"+fmt.Sprint(n+1),i.id)
		f.SetCellValue(a,"E"+fmt.Sprint(n+1),i.smallpic)
		f.SetCellValue(a,"F"+fmt.Sprint(n+1),i.bigpic)
		f.SetCellValue(a,"G"+fmt.Sprint(n+1),i.name)
		f.SetCellValue(a,"H"+fmt.Sprint(n+1),i.video)
		f.SetCellValue(a,"I"+fmt.Sprint(n+1),i.keyword)
		f.SetCellValue(a,"J"+fmt.Sprint(n+1),i.fitaget)
		f.SetCellValue(a,"K"+fmt.Sprint(n+1),i.comment)
		f.SetCellValue(a,"L"+fmt.Sprint(n+1),i.shicai)
		f.SetSheetRow(a,"M"+fmt.Sprint(n+1),&i.buzhou)
		//i.buzhou
		fmt.Println(i.buzhou)
	}
	if err := f.SaveAs("temp.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func Readtemp() [][]string {
	f, err := excelize.OpenFile("temp.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	a := f.GetSheetName(6)

	s := f.GetRows(a)
	return s
}