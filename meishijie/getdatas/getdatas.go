package getdatas

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tidwall/gjson"
	"meishijie/utilib"
)

//id=1657873&time=0&from_home_like=0&type=1&source=android&format=json&lat=39.900087&lon=119.508859&cityCode=148
//https://newapi.meishi.cc/recipe/detail?source=android&format=json&fc=msjandroid&lat=39.900087&lon=119.508859&cityCode=148&token=



//data
//	id
//	title  名称

//	story_content  介绍

//tag
// rate  难易程度
//	technology 技巧
//	taste  味道
//	cooking_time   烹饪时间

//recipe_type=1 非视频
//	cover_img.big  封面图片

//recipe_type=3 视频
// 图片
//video.img
//视频地址   vendor_video

//原料
//主料辅料
//main_ingredient     secondary_ingredient
//步骤
//img   cook_steps[].pic_urls[].big
//文字   cook_steps[].content

//烹饪技巧
//advanced_skills


//营养
//https://newapi.meishi.cc/Recipe/YingYang?source=android&format=json&fc=msjandroid&lat=39.900087&lon=119.508859&cityCode=148&token=
//news_id=1657873&type=YYList&source=android&format=json&lat=39.900087&lon=119.508859&cityCode=148
//data.lable[]

type caipindata struct {
	id string
	name string
	cover_img string //封面
	video string  //视频
	tagrate string //tag 烹饪难度
	tagtechnology string //tag 烹饪技巧
	tagtaste string //tag 味道
	tagcooking_time string //烹饪时间
	story_content string  //内容介绍
	amount string
	main_ingredient []string //主要原料
	secondary_ingredient []string //辅料
	advanced_skills string //烹饪技巧
	yingyang []string  //营养标签
	cook_steps []string //步骤 文字+图片url
}
var fulldatas []caipindata

func Done()  {
	//a := readtable("renqun.xlsx")
	//for _,i:=range a{
	//	getjson(i[0])
	//}
	//writedata(fulldatas,"renqun1.xlsx")
}

func getjson(id string)  {
	u :=`https://newapi.meishi.cc/recipe/detail`
	bodystring := fmt.Sprintf("id=%s&time=0&from_home_like=0&type=1&source=android&format=json&lat=39.900087&lon=119.508859&cityCode=148",id)
	jsondata := utilib.GetJson(u,bodystring)
	a := handeljson(jsondata)
	a.yingyang = yingyang(id)
	//fmt.Println(a.cook_steps)
	fulldatas = append(fulldatas,a)
	//writedata(aa,"jiachang1.xlsx")
	fmt.Println(a)
}

func handeljson(s string) caipindata  {
	var a caipindata
	a.id = fmt.Sprint(gjson.Get(s,"data.id"))
	a.name =fmt.Sprint(gjson.Get(s,"data.title"))
	fmt.Println("http get id: ",a.id)
	a.cover_img = fmt.Sprint(gjson.Get(s,"data.cover_img.big"))

	a.video = fmt.Sprint(gjson.Get(s,"data.video.vendor_video"))
	//标签
	a.tagrate = fmt.Sprint(gjson.Get(s,"data.rate"))
	a.tagtechnology =  fmt.Sprint(gjson.Get(s,"data.technology"))
	a.tagtaste = fmt.Sprint(gjson.Get(s,"data.taste"))
	a.tagcooking_time = fmt.Sprint(gjson.Get(s,"data.cooking_time"))
	//营养标签
	//a.yingyang

	//介绍
	a.story_content = fmt.Sprint(gjson.Get(s,"data.story_content"))

	a.amount = fmt.Sprint(gjson.Get(s,"data.amount"))
	//原料
	//a.main_ingredient
	zhuyao := gjson.Get(s,"data.main_ingredient").Array()
	var zhuyaoslice []string
	for _,i:=range zhuyao{
		n := fmt.Sprint(i.Get("title"))
		w := fmt.Sprint(i.Get("amount"))
		zhuyaoslice = append(zhuyaoslice,n+w)
	}
	a.main_ingredient = zhuyaoslice
	//a.secondary_ingredient
	fuzhu :=gjson.Get(s,"data.secondary_ingredient").Array()
	var fuzhuslice []string
	for _,i:=range fuzhu{
		n := fmt.Sprint(i.Get("title"))
		w := fmt.Sprint(i.Get("amount"))
		fuzhuslice = append(fuzhuslice,n+w)
	}
	a.secondary_ingredient = fuzhuslice
	//步骤
	//a.cook_steps
	buzhou := gjson.Get(s,"data.cook_steps").Array()
	var buzhouslice []string
	for n,i:=range buzhou{
		if n == len(buzhou)-1 {
			c := "成品图"
			urls := i.Get("pic_urls").Array()
			var url string
			for _,x:=range urls{
				if url == "" {
					url = fmt.Sprint(x.Get("big"))
				}else {
					url = url+"\n"+fmt.Sprint(x.Get("big"))
				}
			}
			buzhouslice = append(buzhouslice,c+"\n"+url)
		}else {
			c := fmt.Sprint(i.Get("content"))
			buzhouslice = append(buzhouslice,c)
		}
	}
	a.cook_steps = buzhouslice

	//烹饪技巧
	//a.advanced_skills
	jiqiao := gjson.Get(s,"data.advanced_skills").Array()
	var jiqiaostring string
	for _,i:=range jiqiao{
		if jiqiaostring == "" {
			jiqiaostring = fmt.Sprint(i.Get("content"))
		}else {
			jiqiaostring = jiqiaostring + "\n"+ fmt.Sprint(i.Get("content"))
		}
	}
	a.advanced_skills = jiqiaostring
	return a
}

func yingyang(id string) []string {
	var a []string
	u := "https://newapi.meishi.cc/Recipe/YingYang"
	bodystring :=fmt.Sprintf("news_id=%s&type=YYList&source=android&format=json&lat=39.900087&lon=119.508859&cityCode=148",id)
	jsondata := utilib.GetJson(u,bodystring)
	y := gjson.Get(jsondata,"data.label").Array()
	if len(y) == 0 {
		return a
	}else {
		for _,i :=range y{
			a = append(a,fmt.Sprint(i))
		}
	}
	return a
}

func writedata(s []caipindata,name string)  {
	f, err := excelize.OpenFile(name)
	if err != nil {
		fmt.Println(err)
	}
	a := f.GetSheetName(1)

	for n,i:= range s{
		strings := turntoline(i)
		f.SetSheetRow(a,"B"+fmt.Sprint(n+1),&strings)
		fmt.Println(i.name)
	}

	if err := f.SaveAs(name); err != nil {
		fmt.Println(err)
	}
}

//id string
//name string
//cover_img string //封面
//toppic string  //首页大图
//tagrate string //tag 烹饪难度
//tagtechnology string //tag 烹饪技巧
//tagtaste string //tag 味道
//tagcooking_time string //烹饪时间
//story_content string  //内容介绍
//main_ingredient []string //主要原料
//secondary_ingredient []string //辅料
//advanced_skills string //烹饪技巧
//yingyang []string  //营养标签
//cook_steps []string //步

func turntoline (c caipindata) []string {
	var a []string
	a = append(a,c.id)
	a = append(a,c.name)
	a = append(a,c.cover_img)
	a = append(a,slicetolin(c.yingyang))
	a = append(a,c.tagrate)
	a = append(a,c.tagtechnology)
	a = append(a,c.tagtaste)
	a = append(a,c.tagcooking_time)
	a = append(a,c.story_content)
	a = append(a,c.amount)
	a = append(a,slicetolin(c.main_ingredient))
	a = append(a,slicetolin(c.secondary_ingredient))
	a = append(a,c.advanced_skills)
	for _,i:=range c.cook_steps{
		a = append(a,i)
	}
	return a
}

func slicetolin(s []string)string  {
	var a string
	for _,i:=range s{
		if a == "" {
			a = i
		}else {
			a = a+"\n"+i
		}
	}
	return a
}

func readtable(name string) [][]string {
	f, err := excelize.OpenFile(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	a := f.GetSheetName(1)

	s := f.GetRows(a)
	return s
}