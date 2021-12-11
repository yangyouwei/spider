package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
)

type Medicines struct {
	Id           string `json:"id"`         //药品id
	Name         string `json:"name"`       //名称
	Pic          string `json:"pic"`        //图片
	Adaptation   string `json:"adaptation"` //适用症
	Component    string `json:"component"`  //成分
	Catalog_id   string `json::catalog_id`  //疾病类型
	Catalog_name string //疾病类型
	Months       string `json:"months"`  //月份
	MedType      string `json:"type"`    //类型 1 otc
	Content      string `json:"content"` //说明书id
}

type Disease struct {
	Id   string `json:"id"`
	Name string `json:"catalog_name"`
	Num  string
}

var rows []Medicines

var jibing string = `{
    "code": 0,
    "msg": "success",
    "data": {
        "list": [
            {
                "id": "12",
                "catalog_name": "感冒发烧",
                "icon": "https://pt-images1.cdnmama.com/admin/medicine/4a8f4c5da7603cd36c0e53e2b11d5004_w96X96.png"
            },
            {
                "id": "15",
                "catalog_name": "腹泻便秘",
                "icon": "https://pt-images4.cdnmama.com/admin/medicine/0a4faf1c4095aa2b3962e30bb8cca75d_w96X96.png"
            },
            {
                "id": "14",
                "catalog_name": "止咳化痰",
                "icon": "https://pt-images2.cdnmama.com/admin/medicine/2e2f62d62e708e451284f589219742e6_w96X96.png"
            },
            {
                "id": "13",
                "catalog_name": "清热解毒",
                "icon": "https://pt-images1.cdnmama.com/admin/medicine/14de5493ed6fca66936180f1fac514a5_w96X96.png"
            },
            {
                "id": "1",
                "catalog_name": "健脾消食",
                "icon": "https://pt-images1.cdnmama.com/admin/medicine/15d7030d814b098dabe6348742b1fd58_w64X64.png"
            },
            {
                "id": "18",
                "catalog_name": "五官口鼻",
                "icon": "https://pt-images2.cdnmama.com/admin/medicine/558b61f0f1533196078555134c319866_w96X96.png"
            },
            {
                "id": "16",
                "catalog_name": "胃痛腹痛",
                "icon": "https://pt-images4.cdnmama.com/admin/medicine/b08531968eec2a8be3e641947af294e0_w96X96.png"
            },
            {
                "id": "17",
                "catalog_name": "皮肤用药",
                "icon": "https://pt-images4.cdnmama.com/admin/medicine/7bd127bc1f4d0f721a426e652ab5c59b_w96X96.png"
            },
            {
                "id": "19",
                "catalog_name": "跌打外伤",
                "icon": "https://pt-images1.cdnmama.com/admin/medicine/54c710038f8aa81d4726068f24ee95ba_w96X96.png"
            },
            {
                "id": "20",
                "catalog_name": "营养补充",
                "icon": "https://pt-images.cdnmama.com/admin/medicine/46bfa85a3fd84f13259488e3de28cceb_w96X96.png"
            },
            {
                "id": "22",
                "catalog_name": "小儿驱虫",
                "icon": "https://pt-images4.cdnmama.com/admin/medicine/3bac03cfd2043d8548c4dc92a27ddcd8_w64X64.png"
            }
        ]
    }
}`

var yaopin [][13]string

func main() {
	//获取疾病列表
	//a := GetJsonString()
	//fmt.Println(a)

	//疾病列表写入execl
	//datasjson := gjson.Get(jibing,`data.list`)
	//for n,i := range datasjson.Array(){
	//	var a Disease
	//	a.Id = fmt.Sprint(i.Get("id"))
	//	a.Name = fmt.Sprint(i.Get("catalog_name"))
	//	WriteToExecl(n+1,a)
	//}

	//读取疾病列表获取药列表写入
	//s := ReadExecl()
	//parsejsonn(s)

	////获取说明书
	//s := ReadExecl()
	//for n, i := range s {
	//	for nf, field := range i {
	//		if nf == 0 {
	//			jsonstr := GetJsonString(field)
	//			res := gjson.Get(jsonstr, "data.detail")
	//			ress := res.Array()
	//			fmt.Println(ress[0].Get("content"))
	//			writecontent(fmt.Sprint(n+1), fmt.Sprint(ress[0].Get("content")))
	//		}
	//	}
	//}

	//【通用名称】
	//【商品名称】
	//【用药禁忌】
	//【用法用量】
	//【药品性状】
	//【包装规格】
	//【不良反应】
	//【贮藏条件】
	//【有效期】
	//【注意事项】
	//【相互作用】
	//【批准文号】
	//【生产厂商】


	s := ReadExecl()
	slipstring(s)
	//fmt.Println(yaopin)
	yaopinwrite(yaopin)
}

func slipstring(s [][]string)  {
	for _,n := range s{
		var a [13]string
		aline := strings.Split(n[0],"<br/>")
		//fmt.Println(aline)
		for _,i :=range aline{
			switch  {
			case strings.HasPrefix(i,"【通用名称】"):
				//fmt.Println(strings.Split(i,"】")[1])
				a[0] = strings.Split(i,"】")[1]
			case strings.HasPrefix(i,"【商品名称】"):
				//fmt.Println(st1ings.Split(i,"】")[1])
				a[1] = strings.Split(i,"】")[1]
			case strings.HasPrefix(i,"【用药禁忌】"):
				//fmt.Println(strings.Split(i,"】")[1])
				a[2] = strings.Split(i,"】")[1]
			case strings.HasPrefix(i,"【用法用量】"):
				//fmt.Println(strings.Split(i,"】")[1])
				a[3] = strings.Split(i,"】")[1]
			case strings.HasPrefix(i,"【药品性状】"):
				//fmt.Println(strings.Split(i,"】")[1])
				a[4] = strings.Split(i,"】")[1]
			case strings.HasPrefix(i,"【包装规格】"):
				//fmt.Println(strings.Split(i,"】")[1])
				a[5] = strings.Split(i,"】")[1]
			case strings.HasPrefix(i,"【不良反应】"):
				//fmt.Println(strings.Split(i,"】")[1])
				a[6] = strings.Split(i,"】")[1]
			case strings.HasPrefix(i,"【贮藏条件】"):
				//fmt.Println(strings.Split(i,"】")[1])
				a[7] = strings.Split(i,"】")[1]
			case strings.HasPrefix(i,"【有效期】"):
				//fmt.Println(strings.Split(i,"】")[1])
				a[8] = strings.Split(i,"】")[1]
			case strings.HasPrefix(i,"【注意事项】"):
				//fmt.Println(strings.Split(i,"】")[1])
				a[9] = strings.Split(i,"】")[1]
			case strings.HasPrefix(i,"【相互作用】"):
				//fmt.Println(strings.Split(i,"】")[1])
				a[10] = strings.Split(i,"】")[1]
			case strings.HasPrefix(i,"【批准文号】"):
				//fmt.Println(strings.Split(i,"】")[1])
				a[11] = strings.Split(i,"】")[1]
			case strings.HasPrefix(i,"【生产厂商】"):
				//fmt.Println(strings.Split(i,"】")[1])
				a[12] = strings.Split(i,"】")[1]
			}
		}
		yaopin = append(yaopin,a)
	}
}

func yaopinwrite(s [][13]string) {
	f, err := excelize.OpenFile("常用药查询.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	a := f.GetSheetName(1)
	for n, i := range s {
		for x, y := range i {
			switch x {
			case 0:
				f.SetCellValue(a, "I"+fmt.Sprint(n+1), y)
			case 1:
				f.SetCellValue(a, "J"+fmt.Sprint(n+1), y)
			case 2:
				f.SetCellValue(a, "K"+fmt.Sprint(n+1), y)
			case 3:
				f.SetCellValue(a, "L"+fmt.Sprint(n+1), y)
			case 4:
				f.SetCellValue(a, "M"+fmt.Sprint(n+1), y)
			case 5:
				f.SetCellValue(a, "N"+fmt.Sprint(n+1), y)
			case 6:
				f.SetCellValue(a, "O"+fmt.Sprint(n+1), y)
			case 7:
				f.SetCellValue(a, "P"+fmt.Sprint(n+1), y)
			case 8:
				f.SetCellValue(a, "Q"+fmt.Sprint(n+1), y)
			case 9:
				f.SetCellValue(a, "R"+fmt.Sprint(n+1), y)
			case 10:
				f.SetCellValue(a, "S"+fmt.Sprint(n+1), y)
			case 11:
				f.SetCellValue(a, "T"+fmt.Sprint(n+1), y)
			case 12:
				f.SetCellValue(a, "U"+fmt.Sprint(n+1), y)
			}

		}

	}
	if err := f.SaveAs("常用药查询.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func writecontent(n string, s string) {
	f, err := excelize.OpenFile("常用药查询.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	a := f.GetSheetName(1)

	f.SetCellValue(a, "H"+fmt.Sprint(n), s)

	if err := f.SaveAs("常用药查询.xlsx"); err != nil {
		fmt.Println(err)
	}
}

//func parsejsonn(s [][]string) {
//	for _, i := range s {
//		var name string
//		var nameid string
//		var num string
//		for n,d :=range i{
//			switch n {
//			case 0:
//				name =d
//			case 1:
//				nameid = d
//			case 2:
//				num = d
//			}
//		}
//		jsonstring := GetJsonString(name,num,nameid)
//		writetofile(jsonstring,name)
//	}
//}

func writetofile(s string, c string) {
	one := gjson.Get(s, `data.list`)
	for _, i := range one.Array() {
		var a Medicines
		a.Name = fmt.Sprint(i.Get("name"))
		a.Catalog_name = c
		a.Pic = fmt.Sprint(i.Get("pic"))
		a.Id = fmt.Sprint(i.Get("id"))
		a.Adaptation = fmt.Sprint(i.Get("adaptation"))
		a.Component = fmt.Sprint(i.Get("component"))
		a.MedType = fmt.Sprint(i.Get("type"))
		a.Months = fmt.Sprint(i.Get("months"))
		rows = append(rows, a)
	}
	AppendWrite(rows)
}

func AppendWrite(m []Medicines) {
	f, err := excelize.OpenFile("常用药查询.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	a := f.GetSheetName(1)
	for n, i := range m {
		fmt.Println("qq", i)
		f.SetCellValue(a, "A"+fmt.Sprint(n+1), i.Id)
		f.SetCellValue(a, "B"+fmt.Sprint(n+1), i.Catalog_name)
		f.SetCellValue(a, "C"+fmt.Sprint(n+1), i.Name)
		f.SetCellValue(a, "D"+fmt.Sprint(n+1), i.MedType)
		f.SetCellValue(a, "E"+fmt.Sprint(n+1), i.Months)
		f.SetCellValue(a, "F"+fmt.Sprint(n+1), i.Component)
		f.SetCellValue(a, "G"+fmt.Sprint(n+1), i.Adaptation)
		f.SetCellValue(a, "I"+fmt.Sprint(n+1), i.Pic)
	}

	if err := f.SaveAs("常用药查询.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func GetJsonString(id string) string {
	//获取疾病类型列表
	//resp, err := http.Get("https://papi.mama.cn/wap/medicine/getCatalog")

	//获取疾病类型下的药品列表
	//url := fmt.Sprintf("https://papi.mama.cn/wap/medicine/getListByCatalog?page=1&page_size=%s&catalog_id=%s", num, catalog_id)
	//resp, err := http.Get(url)

	//获取药品信息
	url := fmt.Sprint("https://papi.mama.cn/wap/medicine/getMedicineById?medicine_id=" + id)
	//fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return string(body)
}

func WriteToExecl(n int, s Disease) {
	f, err := excelize.OpenFile("wangyunlist.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	a := f.GetSheetName(1)
	f.SetCellValue(a, "a"+fmt.Sprint(n), s.Name)
	f.SetCellValue(a, "b"+fmt.Sprint(n), s.Id)
	f.SetCellValue(a, "c"+fmt.Sprint(n), s.Num)

	if err := f.SaveAs("wangyunlist.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func ReadExecl() [][]string {
	f, err := excelize.OpenFile("常用药查询.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	a := f.GetSheetName(4)

	s := f.GetRows(a)
	return s
}
