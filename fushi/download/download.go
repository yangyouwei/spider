package download

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"net/http"
	"os"
	"regexp"
)

func DoneDownload()  {
	for n,i:= range ReadExecl(){
		//var content []string
		//var urls []string
		//for _,y:= range i{
		//	s := strings.Split(Removespaceline(y),"\n")
		//	if len(s) == 1 {
		//		content = append(content,s[0])
		//		urls = append(urls,"")
		//	}else {
		//		content = append(content,s[0])
		//		urls =append(urls,s[1])
		//	}
		//}
		//for x,i:=range urls{
		//	if i=="" {
		//		continue
		//	}
		//	Downloadfile(i,n,x)
		//}
		fmt.Println(i[7])
		if	i[7] == ""{
			continue
		}
		Downloadfile(i[7],n)
	}
}

func Removespaceline(s string) string {
	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
	return re.ReplaceAllString(s, "")
}

func Downloadfile(s string,n int) error {
	resp, err := http.Get(s)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 创建一个文件用于保存
	out, err := os.Create(fmt.Sprintf("D:\\yuezi\\video\\%v.mp4",n+2))
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
	return nil
}

func WriteToExecl(name []string,num string)  {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	a := f.GetSheetName(4)
	f.SetCellValue(a,"K"+num,name)
	//f.SetCellValue(a,"M"+num,"1")

	if err := f.SaveAs("data.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func wirtetemp(n string)  {
	f, err := excelize.OpenFile("temp.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	a := f.GetSheetName(1)
	f.SetCellValue(a,"F"+n,"1")

	if err := f.SaveAs("temp.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func ReadExecl() [][]string {
	f, err := excelize.OpenFile("yuezi.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	a := f.GetSheetName(1)

	s := f.GetRows(a)
	return s
}
