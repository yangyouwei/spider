package downloadpic

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"net/http"
	"os"
	"strings"
)

func Getpic()  {
	//for n,i := range ReadUrl("renqun1.xlsx"){
	//	if i[3] != "" {
	//		fmt.Println(n)
	//		httpget(i[3],"D:\\meishijie\\renqun\\cover",n)
	//	}else {
	//		continue
	//	}
	//}

	for n,i := range ReadUrl("renqun1.xlsx"){
		var us []string
		for _,x:=range i{
			if strings.HasPrefix(x,"成品图") {
				a := strings.Split(x,"" +"\n")
				for _,y:=range a{
					if	strings.HasPrefix(y,"http"){
						us = append(us,y)
					}else {
						continue
					}
				}
				for z,m:=range us{
					fmt.Println(n,m)
					httpget(m,"D:\\meishijie\\renqun\\chengpin",n,z)
				}
				continue
			}
		}
	}
}

func httpget(u string,p string,n int,picn int) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 创建一个文件用于保存
	out, err := os.Create(fmt.Sprintf(p+"\\%v_%v.jpg",n+2,picn+1))
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

func ReadUrl(name string) [][]string  {
	f, err := excelize.OpenFile(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	a := f.GetSheetName(1)

	s := f.GetRows(a)
	return s
}
