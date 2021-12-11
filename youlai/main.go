package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
	"youlai/utillib"
)
//execl def
//c docter
//d 指标正常值
//e 指标概述
//f 指标偏高
//g 指标偏低
//h 指标阳性
//i 相关检查

type Docter struct {
	Name string
	Img string
	Title string
	Hospital string
	Department string
}

type Zbzcz struct {
	name string
	content string
}

type Content struct {
	Name string
	Subcontent Subcontent
}

type Subcontent struct {
	Content []string
	Image string
}

type Ydtype struct {
	Name string
	contents []Content
}

type FullData struct {
	Docter Docter
	zbzcz Zbzcz
	YD1 Ydtype
	YD2 Ydtype
	YD3 Ydtype
	YD4 yd4
}

type yd4 struct {
	name string
	data string
}

type Fulldata_string struct {
	name string
	content string
}

type Fulldatas_strings struct {
	Docter string
	Datas []Fulldata_string
}

func main() {
	urls := execlread("./youlai.xlsx")
	for n,i := range urls{
		fmt.Println(n+1,i)
		htmlroot := utillib.GetRootHtml(i)
		if htmlroot == ""{
			continue
		}

		root, _ := htmlquery.Parse(strings.NewReader(htmlroot))
		var fd FullData
		fd.GET_YD(root)
		fd.getdocter(root)
		var zbzcz Zbzcz
		zbzcz.GETzhbzcz(root)
		fd.zbzcz = zbzcz
		a := fd.TurnToString()
		a.Writetoexecl(n+1)
	}
}

func (this *Zbzcz)GETzhbzcz(nodes *html.Node)  {
	ifnull := htmlquery.Find(nodes,"//*[@id=\"zbzcz\"]")
	if ifnull == nil {
		this.name = "指标正常值"
		return
	}
	zbzcz := htmlquery.Find(nodes,"/html/body/section[3]/div[2]/dl")
	//fmt.Println(htmlquery.InnerText(zbzcz[0]))
	this.name = "指标正常值"
	if zbzcz == nil {
		zbzcz = htmlquery.Find(nodes,"/html/body/section/div[2]/dl")
		this.content = utillib.Removespaceline(utillib.Removespace(htmlquery.InnerText(zbzcz[0])))
	}else {
		this.content = utillib.Removespaceline(utillib.Removespace(htmlquery.InnerText(zbzcz[0])))
	}
}

func (this *Fulldatas_strings)Writetoexecl(n int)  {
	//  execl列写入
	//c docter
	//d 指标正常值
	//e 指标概述
	//f 指标偏高
	//g 指标偏低
	//h 指标阳性
	//i 相关检

	f, err := excelize.OpenFile("youlai.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	a :=f.GetSheetName(1)

	//医生信息
	f.SetCellValue(a, "c"+fmt.Sprint(n), this.Docter)

	for _,i :=range this.Datas{
		switch i.name {
		case "指标正常值":
			f.SetCellValue(a, "d"+fmt.Sprint(n),i.content)
		case "指标概述":
			f.SetCellValue(a, "e"+fmt.Sprint(n),i.content)
		case "指标偏高":
			f.SetCellValue(a, "f"+fmt.Sprint(n),i.content)
		case "指标偏低":
			f.SetCellValue(a, "g"+fmt.Sprint(n),i.content)
		case "指标阳性":
			f.SetCellValue(a, "h"+fmt.Sprint(n),i.content)
		case "相关检查":
			f.SetCellValue(a, "i"+fmt.Sprint(n),i.content)
		}
	}

	if err := f.SaveAs("youlai.xlsx"); err != nil {
		fmt.Println(err)
	}
}

//读取excel url
func execlread(s string) []string {
	f, err := excelize.OpenFile(s)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	a :=f.GetSheetName(1)
	var urls []string
	rows := f.GetRows(a)
	for _, row := range rows {
		urls = append(urls,row[0])
	}
	return urls
}

///html/body/section[3]/div[1]/div[2]
func (this *Docter)GET_VALUE(node *html.Node)  {
	//fmt.Println(htmlquery.InnerText(node))    //html/body/section/div[1]/div[2]
	docter_node := htmlquery.Find(node, "/html/body/section[3]/div[1]/div[2]")

	if docter_node == nil {
		docter_node = htmlquery.Find(node, "/html/body/section/div[1]/div[2]")
		img_node :=htmlquery.Find(docter_node[0],"//img")
		for _,i := range img_node{
			this.Img = i.Attr[2].Val
		}
		name_node := htmlquery.Find(docter_node[0],"//strong")
		this.Name = htmlquery.InnerText(name_node[0])
		span_node := htmlquery.Find(docter_node[0],"//span")
		this.Title = htmlquery.InnerText(span_node[0])
		this.Hospital = htmlquery.InnerText(span_node[1])
		this.Department = htmlquery.InnerText(span_node[2])
		return
	}
	img_node :=htmlquery.Find(docter_node[0],"//img")
	for _,i := range img_node{
		this.Img = i.Attr[2].Val
	}
	name_node := htmlquery.Find(docter_node[0],"//strong")
	this.Name = htmlquery.InnerText(name_node[0])
	span_node := htmlquery.Find(docter_node[0],"//span")
	this.Title = htmlquery.InnerText(span_node[0])
	this.Hospital = htmlquery.InnerText(span_node[1])
	this.Department = htmlquery.InnerText(span_node[2])
}

func (this *FullData)getdocter(root *html.Node)  {
	this.Docter.GET_VALUE(root)
}

func (this *FullData)GET_YD(root *html.Node)  {
	yd := htmlquery.Find(root,"//*[@id=\"yyContent\"]")
	ydnode := htmlquery.Find(yd[0],"./div")
	for _,i :=range ydnode{
		a := i.Attr
		switch a[0].Val {
		case "yd1":
			this.YD1 = GetContet(i)
		case "yd2":
			this.YD2 = GetContet(i)
		case "yd3":
			this.YD3 = GetContet(i)
		case "yd4":
			this.YD4=getyd4(i)
		}
	}
}

func getyd4(nodes *html.Node) yd4 {
	var a yd4
	name_node := htmlquery.Find(nodes,"//h3")
	a.name = htmlquery.InnerText(name_node[0])

	yd4nodes := htmlquery.Find(nodes,"//h4")
	for _,i := range yd4nodes{
		a.data =  a.data+htmlquery.InnerText(i)+"\n"
	}

	//fmt.Println(a)
	return a
}

func GetContet(nodes *html.Node) Ydtype {
	var yd Ydtype

	var cs []Content
	nodeslice := htmlquery.Find(nodes,"./*")
	for _,i:= range nodeslice{

		var c Content
		if i.Data == "h3" {
			namestring := htmlquery.InnerText(i)
			namestring = utillib.StringFmt(namestring)
			yd.Name = namestring
		}

		if i.Data == "h4" {
			namestring := htmlquery.InnerText(i)
			namestring = utillib.StringFmt(namestring)
			c.Name = "<h2>"+namestring+"</h2>"+"\n"
		}else {
			var subcontet Subcontent
			if i.Data == "div" {
				contentstring :=htmlquery.InnerText(i)
				contentstring = utillib.StringFmt(contentstring)+"\n"
				subcontet.Content = append(subcontet.Content,contentstring)
				image_node := htmlquery.Find(i, "//img")
				if image_node != nil {
					img := image_node[0].Attr
					subcontet.Image = img[0].Val+"\n"
				}
			}
			c.Subcontent = subcontet
		}

		cs = append(cs,c)
	}
	yd.contents = cs
	return yd
}

func (this *FullData)TurnToString() Fulldatas_strings {
	var yd Fulldatas_strings

	docter := this.Docter.Name+"\n"+this.Docter.Title+"\n"+this.Docter.Hospital+"\n"+this.Docter.Department+"\n"+this.Docter.Img
	//医生信息
	yd.Docter =docter

	var zczb Fulldata_string
	zczb.name = this.zbzcz.name
	zczb.content = this.zbzcz.content
	yd.Datas = append(yd.Datas,zczb)

	var yd1 Fulldata_string
	for _,i :=range this.YD1.contents{
		yd1.content = yd1.content + i.ContentToString()
	}
	yd1.name = this.YD1.Name
	yd.Datas = append(yd.Datas,yd1)

	var yd2 Fulldata_string
	for _,i :=range this.YD2.contents{
		yd2.content = yd2.content + i.ContentToString()
	}
	yd2.name =this.YD2.Name
	yd.Datas = append(yd.Datas,yd2)

	var yd3 Fulldata_string
	for _,i :=range this.YD3.contents{
		yd3.content = yd3.content + i.ContentToString()
	}
	yd3.name = this.YD3.Name
	yd.Datas = append(yd.Datas,yd3)

	var yd4d Fulldata_string
	yd4d.name = this.YD4.name
	yd4d.content = this.YD4.data
	yd.Datas = append(yd.Datas,yd4d)

	return yd
}

func (this *Content)ContentToString()string  {
	var c string
	c = this.Name
	for _,i := range this.Subcontent.Content {
		c = c+ i
	}
	c = c + this.Subcontent.Image
	return c
}