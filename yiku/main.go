package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"yiku/getlist"
)

//var api string ="https://api.meditool.cn/apicheckquick/inspectionmsg"

type Sublist struct {
	InspectionID string
	TypeName string
	TypeFID string
	IsLower string
	DataType string
	IsFinish string
}

type datastruct struct {
	InspectionID string
	TypeName string
	TypeFID string
	IsLower string
	DataType string
	IsFinish string
	List []Sublist
}

type fullcheckdata struct {
	ItemName string
	ItemMsg string
	Img []string
}

type onecheckdata struct {
	TypeName string
	CheckName string
	InspectionID string
	Fulldatas []fullcheckdata
}


func main()  {
	//写入到execl 创建list表
	//c := make(chan []string)
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//go WriteListToExecl(c,&wg)
	//go HandleData(c,&wg)
	//wg.Wait()


	//全文unicode转换中文
	//var a string
	//utillib.FullTurnToCn(getlist.Datas,&a)
	//fmt.Println("aaa",a)

	GetListData()
}

func GetListData()  {
//	POST /apicheckquick/inspectionmsg? HTTP/1.1
//	User-Agent: zhenlipai/4.9.5(Linux; Android 4.4.4)
//	Content-Length: 85
//	Content-Type: application/x-www-form-urlencoded
//  Host: api.meditool.cn
//  Connection: Keep-Alive

	//datatype=0&checktypeid=3950&usertoken=698569fd391e0603fc11d5270f6001a7&userid=1203561
	//api_url := "https://api.meditool.cn/apicheckquick/inspectionmsg"
	var url string = "https://api.meditool.cn/apicheckquick/inspectionmsg"
	b :=[]datastruct{}
	err :=json.Unmarshal([]byte(getlist.CheckList),&b)
	if err != nil {
		fmt.Println(err)
	}

	//a := getcheckdata(url,b[0].List[0].InspectionID)
	//fmt.Println(a)
	num := 0
	for _,i :=range b{
		for _,x := range i.List {
			checkdata := onecheckdata{}
			checkdata.TypeName = i.TypeName
			checkdata.InspectionID = x.InspectionID
			checkdata.CheckName = x.TypeName
			checkdata.Fulldatas = getcheckdata(url,x.InspectionID)
			fmt.Println(num+1,checkdata.CheckName)
			checkdata.WriteToExecl(num+1)
			num = num+1
		}
	}
}

func (this *onecheckdata)WriteToExecl(n int)  {
	f, err := excelize.OpenFile("yiku_list.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	a :=f.GetSheetName(1)
	f.SetCellValue(a, "a"+fmt.Sprint(n), this.TypeName)
	f.SetCellValue(a, "b"+fmt.Sprint(n), this.CheckName)
	for _,i := range this.Fulldatas{
		switch i.ItemName {
		case "样本要求":
			imgs := joinsting(i.Img)
			f.SetCellValue(a, "c"+fmt.Sprint(n), i.ItemMsg+"\n"+imgs)
		case "测定方法":
			imgs := joinsting(i.Img)
			f.SetCellValue(a, "d"+fmt.Sprint(n), i.ItemMsg+"\n"+imgs)
		case "参考值":
			imgs := joinsting(i.Img)
			f.SetCellValue(a, "e"+fmt.Sprint(n), i.ItemMsg+"\n"+imgs)
		case "临床意义":
			imgs := joinsting(i.Img)
			f.SetCellValue(a, "f"+fmt.Sprint(n), i.ItemMsg+"\n"+imgs)
		case "相关疾病":
			imgs := joinsting(i.Img)
			f.SetCellValue(a, "g"+fmt.Sprint(n), i.ItemMsg+"\n"+imgs)
		case "检查目的":
			imgs := joinsting(i.Img)
			f.SetCellValue(a, "h"+fmt.Sprint(n), i.ItemMsg+"\n"+imgs)
		case "检查时机":
			imgs := joinsting(i.Img)
			f.SetCellValue(a, "i"+fmt.Sprint(n), i.ItemMsg+"\n"+imgs)
		case "检验过程":
			imgs := joinsting(i.Img)
			f.SetCellValue(a, "j"+fmt.Sprint(n), i.ItemMsg+"\n"+imgs)
		case "检验原理":
			imgs := joinsting(i.Img)
			f.SetCellValue(a, "k"+fmt.Sprint(n), i.ItemMsg+"\n"+imgs)
		case "注意事项":
			imgs := joinsting(i.Img)
			f.SetCellValue(a, "l"+fmt.Sprint(n), i.ItemMsg+"\n"+imgs)
		}
		if err := f.SaveAs("yiku_list.xlsx"); err != nil {
			fmt.Println(err)
		}
	}
}

func joinsting(strslice []string)string  {
	s := ""
	for _,i:=range strslice{
		s = s + i+"\n"
	}
	return s
}

func getcheckdata(url string,checkid string) []fullcheckdata {
	bodystring := PostRequest(url,checkid)
	datasjson := gjson.Get(bodystring,`data`)
	fulldatas := []fullcheckdata{}
	for _,i := range datasjson.Array() {
		onedata := fullcheckdata{}
		for n,d := range i.Array() {
			if n == 0 {
				onedata.ItemMsg = fmt.Sprint(d.Get("ItemMsg"))
				onedata.ItemName = fmt.Sprint(d.Get("ItemName"))
			}else {
				onedata.Img = append(onedata.Img,fmt.Sprint(d.Get("ItemMsg")))
			}
		}
		fulldatas = append(fulldatas,onedata)
	}
	return fulldatas
}

func PostRequest(url string,checkid string) string {
	//var bodystring string
	client := &http.Client{}
	requestbody := fmt.Sprintf("datatype=0&checktypeid=%s&usertoken=698569fd391e0603fc11d5270f6001a7&userid=1203561",checkid)
	//fmt.Println(requestbody)
	req, err := http.NewRequest("POST", url, strings.NewReader(requestbody))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent","zhenlipai/4.9.5(Linux; Android 4.4.4)")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}


//write to typename and check list
func HandleData(c chan []string,wg *sync.WaitGroup)  {
	b :=[]datastruct{}
	err :=json.Unmarshal([]byte(getlist.CheckList),&b)
	if err != nil {
		fmt.Println(err)
	}
	for _,i := range b{
		 t := i.TypeName
		for _,i:= range i.List {
			var s []string
			s = append(s,t)
			s = append(s,i.TypeName)
			s = append(s,i.InspectionID)
			c <- s
		}
	}
	close(c)
	wg.Done()
}

func WriteListToExecl(c chan []string,wg *sync.WaitGroup)  {
	n := 1
	for  {
		i, isClose := <- c
		if !isClose {
			break
		}
		WriteToExecl(n,i)
		fmt.Println(i,n)
		n = n+1
	}
	wg.Done()
}

func WriteToExecl(n int,s []string) {
	f, err := excelize.OpenFile("yiku_list.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	a :=f.GetSheetName(1)
	f.SetCellValue(a, "a"+fmt.Sprint(n), s[0])
	f.SetCellValue(a, "b"+fmt.Sprint(n), s[1])
	f.SetCellValue(a, "c"+fmt.Sprint(n), s[2])

	if err := f.SaveAs("yiku_list.xlsx"); err != nil {
		fmt.Println(err)
	}
}







