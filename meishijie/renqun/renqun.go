package renqun

import (
	"fmt"
	"meishijie/utilib"
	"net/url"
)

//id
//婴儿 597
//哺乳期 596
//幼儿 595
//产妇 594
//孕妇 593

//url
//https://newapi.meishi.cc/Search/recipe_7_1_0?source=android&format=json&fc=msjandroid&lat=39.900087&lon=119.508859&cityCode=148&token=

//body
//cid=597&keyword=%E5%A9%B4%E5%84%BF&order=-hot&page=1&per_page=10&source=android&format=json&lat=39.900087&lon=119.508859&cityCode=148
var id []string = []string{"597","596","595","594","593"}


func Getdata(m map[string]utilib.Caipin)  {
	u:="https://newapi.meishi.cc/Search/recipe_7_1_0?source=android&format=json&fc=msjandroid&lat=39.904973&lon=116.39852&cityCode=110&token="

	for _,x:=range id{
		ret := ctxrenqun(fmt.Sprint(x))
		bstring := fmt.Sprintf("cid=%s&keyword=",x)
		b := url.QueryEscape(ret)
		bodystring := bstring+b

		for i:=0;i<10;i++ {
			e := fmt.Sprintf("&order=-hot&page=%v&per_page=100&source=android&format=json&lat=39.904973&lon=116.39852&cityCode=110",i+1)
			//fmt.Println("-------")
			//fmt.Println(bodystring+e)
			//fmt.Println(u)
			jsondata := utilib.GetJson(u,bodystring+e)
			l := utilib.ParseJsonList(jsondata,"data.items")

			//fmt.Println(bodystring+e)
			if len(l) == 0{
				fmt.Println("finish")
				continue
			}

			for _,y:=range l{
				if fmt.Sprint(y.Get("type")) != "1" {
					continue
				}

				if v, ok := m[fmt.Sprint(y.Get("recipe.id"))]; ok {
					v.Renqun = append(v.Renqun,ret)
					m[fmt.Sprint(y.Get("recipe.id"))]=v
					//fmt.Println(m[fmt.Sprint(y.Get("recipe.id"))],"重复")
					continue
				}else {
					var caipin utilib.Caipin
					caipin.Name = fmt.Sprint(y.Get("recipe.title"))
					caipin.Img = fmt.Sprint(y.Get("recipe.img"))
					//ret := ctxrenqun(fmt.Sprint(x))
					caipin.Renqun = append(caipin.Renqun,ret)
					m[fmt.Sprint(y.Get("recipe.id"))] = caipin
					//fmt.Println(m[fmt.Sprint(y.Get("recipe.id"))],"新增")
					//fmt.Println(y.Get("recipe.id"),ret)
				}
			}
		}
	}
}


//婴儿 597
//哺乳期 596
//幼儿 595
//产妇 594
//孕妇 593
func ctxrenqun(id string) string {
	switch id {
	case fmt.Sprint(597):
		return "婴儿"
	case fmt.Sprint(596):
		return "哺乳期"
	case fmt.Sprint(595):
		return "幼儿"
	case fmt.Sprint(594):
		return "产妇"
	case fmt.Sprint(593):
		return "孕妇"
	default:
		return ""
	}
}
