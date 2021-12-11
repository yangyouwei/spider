package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

//医生
type doctor struct {
	//姓名
	name string
	//头像
	img string
	//职位
	title string
	//医院
	hospital string
	//科室
	department string
}
type subcontent struct {
	h4 string
	content string
}
type subcontent1 struct {
	h4 string
	p  []string
}
//id = ycjg
type unomalresult struct {
	h4 string
	div []string
}
//认识检查
type msg1 struct {
	img string
	div []subcontent
}
//检查前
type msg2 struct {
	img string
	div []subcontent
}
//检查中
type msg3 struct {
	img string
	div []subcontent
}
//检查后
type msg4 struct {
	img string
	div []subcontent
}
//检查报告解读
type msg5 struct {
	img string
	jdms subcontent
	jcsj subcontent
	cjzd subcontent
	zcjg subcontent1
	ycjg unomalresult
}
//相关检查项
type msg6 struct {
	img string
	a []string
}
//参考文献
type msg7 struct {
	h5 string
	div string
}
//全部信息
type FullInfo struct {
	//医生信息
	doctor
	//认识检查
	msg1
	//检查前
	msg2
	//检查中
	msg3
	//检查后
	msg4
	//检查报告解读
	msg5
	//相关检查项
	msg6
	//参考文献
	msg7
}

func main()  {
	urls := execlread("./有来医生测试.xlsx")
	for n,i := range urls {
		fmt.Println(i)
		rootnode := getnode(download(i))

		var fullinfo FullInfo
		fullinfo.getfull(rootnode)
		writetofile(fullinfo,n+1)
	}
	//	rootnode := getnode(download(urls[1]))
	//
	//	var fullinfo FullInfo
	//	fullinfo.getfull(rootnode)
	//	fmt.Println(fullinfo.doctor)
	//	writetofile(fullinfo,1+1)

}
//获取execl 一个url的全部要爬取的内容
func (this *FullInfo)getfull(rootnode *html.Node)  {
	var doctor doctor
	doctor.getdockerinfo(rootnode)

	var msg1 msg1
	msg1.getmsg1(rootnode)
	//fmt.Println(msg1)

	var msg2 msg2
	msg2.getmsg2(rootnode)
	//fmt.Println(msg2)

	var msg3 msg3
	msg3.getmsg3(rootnode)
	//fmt.Println(msg3)

	var msg4 msg4
	msg4.getmsg4(rootnode)
	//fmt.Println(msg4)

	var msg5 msg5
	msg5.getmsg5(rootnode)
	//fmt.Println(msg5)

	var msg6 msg6
	msg6.getmsg6(rootnode)
	//fmt.Println(msg6)

	var msg7 msg7
	msg7.getmsg7(rootnode)
	//fmt.Println(msg7)
	this.msg1 =msg1
	this.msg2 =msg2
	this.msg3 = msg3
	this.msg4 = msg4
	this.msg5 = msg5
	this.msg6 = msg6
	this.msg7 = msg7
	this.doctor = doctor
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
//get body
func download(url string)string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
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
//获取*html.node   root
func getnode(html string) *html.Node {
	root, _ := htmlquery.Parse(strings.NewReader(html))
	return root
}
//获取医生信息
func (this *doctor)getdockerinfo(root *html.Node)  {
	tr := htmlquery.Find(root, "/html/body/section[3]/div[1]/div[2]/a/dl")
	if tr !=nil {
		for _,i := range tr {
			head_image_html := htmlquery.Find(i, "./dt/img")
			head_image := head_image_html[0].Attr
			this.img = head_image[2].Val

			item := htmlquery.Find(i, "./dd/div/strong")
			name := htmlquery.InnerText(item[0])
			this.name = name

			item1 := htmlquery.Find(i, "./dd/div/span")
			title := htmlquery.InnerText(item1[0])
			hospital := htmlquery.InnerText(item1[1])
			department := htmlquery.InnerText(item1[2])

			this.title = title
			this.hospital = hospital
			this.department = department
		}
	}else {
		tr = htmlquery.Find(root, "/html/body/section[3]/div[2]/div[2]/a/dl")
		for _,i := range tr {
			head_image_html := htmlquery.Find(i, "./dt/img")
			fmt.Println(head_image_html)
			head_image := head_image_html[0].Attr
			this.img = head_image[2].Val

			item := htmlquery.Find(i, "./dd/div/strong")
			name := htmlquery.InnerText(item[0])
			this.name = name

			item1 := htmlquery.Find(i, "./dd/div/span")
			title := htmlquery.InnerText(item1[0])
			hospital := htmlquery.InnerText(item1[1])
			department := htmlquery.InnerText(item1[2])

			this.title = title
			this.hospital = hospital
			this.department = department
		}
	}

}
func (this *doctor)turnstring() string {
	s := this.name+"\n"+this.img+"\n"+this.title+"\n"+this.hospital+"\n"+this.department
	return removespaceline(s)
}
//获取msg1信息
func (this *msg1)getmsg1(root *html.Node)  {
	tr := htmlquery.Find(root, "//*[@id=\"yd1\"]")
	for _,i := range tr{
		h4_title := htmlquery.Find(i, "./h4")
		for _,t :=range h4_title {
			var divinfo subcontent
			divinfo.h4 = htmlquery.InnerText(t)
			this.div = append(this.div,divinfo)
		}

		contentinfo := htmlquery.Find(i, "./div")
		for n,d := range contentinfo{
			this.div[n].content = strings.ReplaceAll(htmlquery.InnerText(d)," ","")
		}

		images := htmlquery.FindOne(i, "//img")
		this.img = htmlquery.SelectAttr(images, "src")
	}
}
// "<h2>"+i.h4+"</h2>"
func (this *msg1)turnstring()string  {
	var s string
	for _,i := range this.div{
		s = s + "<h2>"+i.h4+"</h2>" +"\n"+i.content+"\n"
	}
	s = s + this.img
	return removespaceline(s)
}

//获取msg2信息
func (this *msg2)getmsg2(root *html.Node)  {
	tr := htmlquery.Find(root, "//*[@id=\"yd2\"]")
	for _,i := range tr{
		h4_title := htmlquery.Find(i, "./h4")
		for _,t :=range h4_title {
			var divinfo subcontent
			divinfo.h4 = htmlquery.InnerText(t)
			this.div = append(this.div,divinfo)
		}

		contentinfo := htmlquery.Find(i, "./div")
		for n,d := range contentinfo{
			this.div[n].content = strings.ReplaceAll(htmlquery.InnerText(d)," ","")
		}

		images := htmlquery.FindOne(i, "//img")
		this.img = htmlquery.SelectAttr(images, "src")
	}
}
func (this *msg2)turnstring() string {
	var s string
	for _,i := range this.div{
		s = s+"<h2>"+i.h4+"</h2>"+"\n"+i.content+"\n"
	}
	s = s+ this.img
	return removespaceline(s)
}
//获取msg3信息
func (this *msg3)getmsg3(root *html.Node)  {
	tr := htmlquery.Find(root, "//*[@id=\"yd3\"]")
	for _,i := range tr{
		h4_title := htmlquery.Find(i, "./h4")
		for _,t :=range h4_title {
			var divinfo subcontent
			divinfo.h4 = htmlquery.InnerText(t)
			this.div = append(this.div,divinfo)
		}

		contentinfo := htmlquery.Find(i, "./div")
		for n,d := range contentinfo{
			this.div[n].content = strings.ReplaceAll(htmlquery.InnerText(d)," ","")
		}

		images := htmlquery.FindOne(i, "//img")
		this.img = htmlquery.SelectAttr(images, "src")
	}
}
func (this *msg3)turnstring()string  {
	var s string
	for _,i:= range this.div{
		s = s +"<h2>"+i.h4+"</h2>"+"\n"+i.content+"\n"
	}
	s = s + this.img
	return removespaceline(s)
}
//获取msg4信息
func (this *msg4)getmsg4(root *html.Node)  {
	tr := htmlquery.Find(root, "//*[@id=\"yd4\"]")
	for _,i := range tr{
		h4_title := htmlquery.Find(i, "./h4")
		for _,t :=range h4_title {
			var divinfo subcontent
			divinfo.h4 = htmlquery.InnerText(t)
			this.div = append(this.div,divinfo)
		}

		contentinfo := htmlquery.Find(i, "./div")
		for n,d := range contentinfo{
			this.div[n].content = strings.ReplaceAll(htmlquery.InnerText(d)," ","")
		}

		images := htmlquery.FindOne(i, "//img")
		this.img = htmlquery.SelectAttr(images, "src")
	}
}
func (this *msg4)turnstring()string  {
	var s string
	for _,i:=range this.div{
		s = s+"<h2>"+i.h4+"</h2>"+"\n"+i.content+"\n"
	}
	s = s+this.img
	return removespaceline(s)
}
//获取msg5信息
func (this *msg5)getmsg5(root *html.Node)  {
	tr := htmlquery.Find(root, "//*[@id=\"yd5\"]")
	for _,i := range tr{
		images := htmlquery.FindOne(i, "//img")
		this.img = htmlquery.SelectAttr(images, "src")

		//简短描述
		//*[@id="jdms"]
		//*[@id="yd5"]/div[1]
		jdms := htmlquery.Find(i, "//*[@id=\"jdms\"]")
		if jdms != nil {
			this.jdms.h4="<h2>"+htmlquery.InnerText(jdms[0])+"</h2>"
			jdmsdiv := htmlquery.Find(i, "//*[@id=\"yd5\"]/div[1]")
			this.jdms.content = htmlquery.InnerText(jdmsdiv[0])
		}

		//检查所见
		////*[@id="jcsj"]
		//*[@id="yd5"]/div[2]

		jcsj := htmlquery.Find(i, "//*[@id=\"jcsj\"]")
		if jcsj != nil {
			this.jcsj.h4 = "<h2>"+htmlquery.InnerText(jcsj[0])+"</h2>"
			jscjdiv := htmlquery.Find(i, "//*[@id=\"yd5\"]/div[2]")
			this.jcsj.content = strings.ReplaceAll(htmlquery.InnerText(jscjdiv[0])," ","")
		}


		//常见诊断
		//*[@id="cjzd"]
		//*[@id="yd5"]/div[3]
		cjzd := htmlquery.Find(i, "//*[@id=\"cjzd\"]")
		this.cjzd.h4 = "<h2>"+htmlquery.InnerText(cjzd[0])+"</h2>"
		cjzddiv := htmlquery.Find(i, "//*[@id=\"yd5\"]/div[3]")
		this.cjzd.content = strings.ReplaceAll(htmlquery.InnerText(cjzddiv[0])," ","")

		//表示正常的结果
		//*[@id="zcjg"]
		//*[@id="yd5"]/div[4]
		zcjg:= htmlquery.Find(i, "//*[@id=\"zcjg\"]")
		if zcjg != nil {
			this.zcjg.h4 = "<h2>"+htmlquery.InnerText(zcjg[0])+"</h2>"
			zcjgdiv := htmlquery.Find(i, "//*[@id=\"yd5\"]/div[4]/p")
			for _,i:= range zcjgdiv{
				this.zcjg.p = append(this.zcjg.p,strings.ReplaceAll(htmlquery.InnerText(i)," ",""))
			}
		}

		//表示异常的结果
		//*[@id="ycjg"]
		//*[@id="yd5"]/div[5]
		ycjg := htmlquery.Find(i, "//*[@id=\"ycjg\"]")
		this.ycjg.h4 = "<h2>"+htmlquery.InnerText(ycjg[0])+"</h2>"
		ycjgdiv := htmlquery.Find(i, "//*[@id=\"yd5\"]/div")
		for n,i :=range ycjgdiv{
			if n >= 4 {
				this.ycjg.div = append(this.ycjg.div,strings.ReplaceAll(htmlquery.InnerText(i)," ",""))
			}
		}
	}

}
func (this *msg5)turnstring()string  {
	s :=""
	if this.jdms.h4 == "" {
		s =this.jcsj.h4+"\n"+this.jcsj.content+"\n"+this.cjzd.h4+"\n"+this.cjzd.content+"\n"+this.zcjg.h4+"\n"
	}else{
		s = this.jdms.h4+"\n"+this.jdms.content+"\n"+this.jcsj.h4+"\n"+this.jcsj.content+"\n"+this.cjzd.h4+"\n"+this.cjzd.content+"\n"+this.zcjg.h4+"\n"
	}

	if this.jcsj.h4 == "" {
		s = this.jdms.h4+"\n"+this.jdms.content+"\n"+this.cjzd.h4+"\n"+this.cjzd.content+"\n"+this.zcjg.h4+"\n"
	}else {
		s =this.jdms.h4+"\n"+this.jdms.content+"\n"+this.jcsj.h4+"\n"+this.jcsj.content+"\n"+this.cjzd.h4+"\n"+this.cjzd.content+"\n"+this.zcjg.h4+"\n"

	}

	if this.zcjg.h4 != "" {
		for _,i:=range this.zcjg.p{
			s = s+i+"\n"
		}
	}

	for _,i:=range this.zcjg.p{
		s = s+i+"\n"
	}

	s = s + "\n"+this.ycjg.h4+"\n"
	for _,i:= range this.ycjg.div{
		s =s + i +"\n"
	}
	s = s + this.img
	return removespaceline(s)
}
//获取msg6信息
func (this *msg6)getmsg6(root *html.Node)  {
	tr := htmlquery.Find(root, "//*[@id=\"yd6\"]")
	for _,i := range tr{
		images := htmlquery.FindOne(i, "//img")
		this.img = htmlquery.SelectAttr(images, "src")

		a := htmlquery.Find(i,"./div/a")
		for _,n := range a{
			s := strings.ReplaceAll(htmlquery.InnerText(n)," ","")
			this.a = append(this.a,s)
		}
	}
}
func (this *msg6)turnstring()string  {
	var s string
	for _,i:=range this.a{
		s = s +i+"\n"
	}
	s =s+this.img
	return removespaceline(s)
}
//获取msg7信息
func (this *msg7)getmsg7(root *html.Node)  {
	tr := htmlquery.Find(root, "//*[@id=\"yyEnd\"]")
	for _,i := range tr{
		h5 := htmlquery.Find(i,"./h5")
		this.h5 =  htmlquery.InnerText(h5[0])

		div := htmlquery.Find(i,"./div")
		this.div = strings.ReplaceAll(htmlquery.InnerText(div[0])," ","")
	}
}
func (this *msg7)turnstring() string {
	s := this.h5+"\n"+this.div
	return removespaceline(s)
}
//写入表格
func writetofile(s FullInfo,n int)  {
	f, err := excelize.OpenFile("有来医生测试.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	a :=f.GetSheetName(1)
	//医生信息
	f.SetCellValue(a, "c"+fmt.Sprint(n), s.doctor.turnstring())
	//认识检查
	f.SetCellValue(a, "D"+fmt.Sprint(n), s.msg1.turnstring())
	// 检查前
	f.SetCellValue(a, "E"+fmt.Sprint(n), s.msg2.turnstring())
	//检查中
	f.SetCellValue(a, "F"+fmt.Sprint(n), s.msg3.turnstring())
	//检查后
	f.SetCellValue(a, "G"+fmt.Sprint(n), s.msg4.turnstring())
	//检查报告解读
	f.SetCellValue(a, "H"+fmt.Sprint(n), s.msg5.turnstring())
	//相关检查项
	f.SetCellValue(a, "I"+fmt.Sprint(n), s.msg6.turnstring())
	//参考资料
	f.SetCellValue(a, "J"+fmt.Sprint(n), s.msg7.turnstring())
	// 根据指定路径保存文件
	if err := f.SaveAs("有来医生测试.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func removespaceline(s string) string {
	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
	return re.ReplaceAllString(s, "")
}