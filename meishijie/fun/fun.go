package fun

import (
	"fmt"
	"meishijie/utilib"
)

//50页  200个
//url https://newapi.meishi.cc/Search/recipe_7_1_0
//post  cid=2&keyword=%E5%AE%B6%E5%B8%B8%E8%8F%9C&order=-hot&page=1&per_page=200&source=android&format=json&lat=39.904973&lon=116.39852&cityCode=110
//type =1
//lable 不是空的话  为图一

var fulldata [][]string

func Done(m map[string]utilib.Caipin)  {
	for i:=0;i<51;i++ {
		Getdata(i+1,m)
	}
}

func GetListcheshi()  {
	u:="https://newapi.meishi.cc/Search/recipe_7_1_0"
	bodystring := "cid=2&keyword=%E5%AE%B6%E5%B8%B8%E8%8F%9C&order=-hot&page=1&per_page=200&source=android&format=json&lat=39.904973&lon=116.39852&cityCode=110"
	jsondata := utilib.GetJson(u,bodystring)
	l := utilib.ParseJsonList(jsondata,"data.items")
	if len(l) == 0{
		fmt.Println("finish")
		return
	}

	for _,i:=range l{
		if fmt.Sprint(i.Get("type")) != "1" {
			continue
		}
		fmt.Println(i.Get("recipe.id"),i.Get("recipe.title"),i.Get("recipe.img"))
	}

}

func Getdata(n int,m map[string]utilib.Caipin)  {
	u:="https://newapi.meishi.cc/Search/recipe_7_1_0"
	a:= `cid=2&keyword=%E5%AE%B6%E5%B8%B8%E8%8F%9C&order=-hot&`
	b := fmt.Sprintf("page=%v&per_page=200&source=android&format=json&lat=39.904973&lon=116.39852&cityCode=110",n)
	jsondata := utilib.GetJson(u,a+b)
	l := utilib.ParseJsonList(jsondata,"data.items")


	if len(l) == 0{
		fmt.Println(n,"家常菜 items none")
		return
	}

	for _,i:=range l{
		if fmt.Sprint(i.Get("type")) != "1" {
			continue
		}

		var caipin utilib.Caipin
		if _, ok := m[fmt.Sprint(i.Get("recipe.id"))]; ok {
			continue
		}
		caipin.Name = fmt.Sprint(i.Get("recipe.title"))
		caipin.Img = fmt.Sprint(i.Get("recipe.img"))
		caipin.Jiachangcai = "家常菜"
		m[fmt.Sprint(i.Get("recipe.id"))] =caipin

	}
	fmt.Println("page : ",n)
}
