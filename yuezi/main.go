package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/antchfx/htmlquery"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//这个爬虫是，先获取大的分类json数据，取得大的分类，调理项目。。。。
//再获取每个调理项目下有很多蔬菜类型。每个蔬菜类型下有很多该蔬菜做的菜。都是json数据。
//每个菜品有个一个json数据。包含了菜名，id，步骤 食材
//步骤里的内容是html文本。需要解析出文本和图片

var (
	endpiont string = "https://papi.mama.cn/wap/subject/subject_detail?subject_id=142&video_type=1&rel=1&mode=2&bb_birthday=2021-01-19&page=1&perpage=20&isMiniApp=0&device_id=&source="

	detail string = "https://papi.mama.cn/wap/auxiliary_food/get_recipe_detail?recipe_id=1781&uid=101062945&type=pregnancy_recipe"
)

//一个类型的完整数据
type fulldata struct {
	month     string
	month_url string
	fenlei    []caipinlist
}

//菜品分类的数据结构
type caipinlist struct {
	shucaifenlei string
	fenlei_id    string
	caipins      []caipin
}

//菜品的数据结构
type caipin struct {
	id                 string
	recipe_img         string
	recipe_name        string
	recipe_ingredients string //食材
	steps              []step
}

//烹饪步骤的数据结构
type step struct {
	content string
	images  string
}

//全部抓取下来的数据
var full []fulldata

func main() {
	//获取月份 url
	firstlist_json := getjson(endpiont)
	getylsp_list(firstlist_json)

	//获取月份蔬菜分类
	getsecondlist()

	//获取菜品详细信息
	getdetail()

	//写入execl，并下载图片
	writetoexecl()
}

func getsecondlist() {
	for n, i := range full {
		num := strings.Split(strings.Split(i.month_url, "id=")[1], "&")[0]
		u := fmt.Sprintf("https://papi.mama.cn/wap/auxiliary_food/get_recipe_info?id=%s&page=1&pageSize=50&type=pregnancy_recipe", num)
		secondlist := getjson(u)
		var cslice []caipinlist
		for _, x := range gjson.Get(secondlist, "data.level_info").Array() {
			var c caipinlist
			c.fenlei_id = fmt.Sprint(x.Get("id"))
			c.shucaifenlei = fmt.Sprint(x.Get("name"))
			cslice = append(cslice, c)
		}
		for y, i := range cslice {
			var cps []caipin
			id := i.fenlei_id
			tmstr := fmt.Sprintf("data.recipe_info.%s", id)
			for _, m := range gjson.Get(secondlist, tmstr).Array() {
				var c caipin
				c.id = fmt.Sprint(m.Get("recipe_id"))
				c.recipe_name = fmt.Sprint(m.Get("recipe_name"))
				c.recipe_img = fmt.Sprint(m.Get("recipe_img"))
				cps = append(cps, c)
			}
			cslice[y].caipins = cps
		}
		full[n].fenlei = cslice
	}
}

func getdetail() {
	for n, i := range full {
		for x, y := range i.fenlei {
			for a, b := range y.caipins {
				u := fmt.Sprintf("https://papi.mama.cn/wap/auxiliary_food/get_recipe_detail?recipe_id=%s&uid=101062945&type=pregnancy_recipe", b.id)
				//fmt.Println(u)
				j := getjson(u)
				full[n].fenlei[x].caipins[a].recipe_ingredients = fmt.Sprint(gjson.Get(j, "data.recipe_list.recipe_ingredients"))
				full[n].fenlei[x].caipins[a].steps = getstep(j)
				//fmt.Println(full[n].fenlei[x].caipins[a].steps)
			}
		}
	}
}

//解析json数据获取步骤信息。
func getstep(s string) []step {
	var ss []step
	for _, i := range gjson.Get(s, "data.recipe_list.recipe_steps").Array() {
		var s step
		s.content, s.images = getcontentAndImage(fmt.Sprint(i))
		ss = append(ss, s)
	}
	return ss
}

func getcontentAndImage(s string) (c string, i string) {
	i = getimg(s)
	c = getcontent(s)
	return c, i
}

//获取html文本中的文字，去掉标签和格式。
func getcontent(s string) string {
	var content string
	nodep, err := htmlquery.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	}
	content = htmlquery.InnerText(nodep)
	return content
}

//解析html页面中的图片连接，步骤里的图片
func getimg(s string) string {
	nodep, err := htmlquery.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	}
	imgnode := htmlquery.FindOne(nodep, "//img")
	if imgnode == nil {
		return ""
	}
	img := htmlquery.SelectAttr(imgnode, "src")
	//fmt.Println(htmlquery.InnerText(nodep))
	return img
}

//将full 这个全局变量写入execl
//有数据是写入的，图片直接下载重命名按序号+2 潇哥那边定的+2
//步骤图片是行号+步骤number
func writetoexecl() {
	f, err := excelize.OpenFile("yuezi.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	a := f.GetSheetName(2)
	var n = 0
	for _, i := range full {
		for _, y := range i.fenlei {
			for _, m := range y.caipins {
				fmt.Println(m)
				f.SetCellValue(a, "A"+fmt.Sprint(n+1), i.month)
				f.SetCellValue(a, "B"+fmt.Sprint(n+1), y.shucaifenlei)
				f.SetCellValue(a, "C"+fmt.Sprint(n+1), m.recipe_name)
				f.SetCellValue(a, "D"+fmt.Sprint(n+1), m.recipe_img)
				img_path := "C:\\Users\\yyw\\go\\src\\yuezi\\tiaoli\\coverimg\\"
				Downloadfile(m.recipe_img, img_path+fmt.Sprint(n+2)+".jpg")
				f.SetCellValue(a, "E"+fmt.Sprint(n+1), m.recipe_ingredients)
				var ss []string
				for z, b := range m.steps {
					ss = append(ss, b.content)
					stepimg := "C:\\Users\\yyw\\go\\src\\yuezi\\tiaoli\\stepimg\\"
					Downloadfile(b.images, stepimg+fmt.Sprint(n+2)+"-"+fmt.Sprint(z+1)+".jpg")
				}
				f.SetSheetRow(a, "F"+fmt.Sprint(n+1), &ss)
				n = n + 1
				fmt.Println(n)
			}
		}
	}

	if err := f.SaveAs("yuezi.xlsx"); err != nil {
		fmt.Println(err)
	}
}

//下载文件存到指定位置
func Downloadfile(s string, n string) error {
	resp, err := http.Get(s)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 创建一个文件用于保存
	out, err := os.Create(n)
	if err != nil {
		return err
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	//fmt.Printf("%v_%v.jpg \n",n+2,x+1)
	fmt.Println("download ", n, "finished.")
	return nil
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

func getylsp_list(s string) {
	for _, i := range gjson.Get(s, "data.link_list.2.info.open").Array() {
		var f fulldata
		f.month = fmt.Sprint(i.Get("title"))
		f.month_url = fmt.Sprint(i.Get("url"))
		full = append(full, f)
	}
}
