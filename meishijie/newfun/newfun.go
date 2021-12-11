package newfun

import (
	"fmt"
	"meishijie/fun"
	"meishijie/renqun"
	"meishijie/utilib"
)

var fulldata map[string]utilib.Caipin
var renqunstr [][]string
var jiachang [][]string

func Done()  {
	fulldata = make(map[string]utilib.Caipin)
	Jiachangcai(fulldata)
	renqun.Getdata(fulldata)

	//for k,i:=range fulldata{
	//	fmt.Println(k,i.Name,i.Renqun)
	//}

	writeTOfile(fulldata)
	//utilib.WriteToExecl(renqunstr,"renqun.xlsx")
	//utilib.WriteToExecl(jiachang,"jiachang.xlsx")

}

func Jiachangcai(m map[string]utilib.Caipin)  {
	for i:=0;i<52;i++ {
		fun.Getdata(i+1,m)
	}
}

func writeTOfile(m map[string]utilib.Caipin)  {
	var n = 0
	for k,v:=range m{
		if len(v.Renqun) == 0 {
			//fmt.Println("renqun")
			var a []string
			a =append(a,k)
			a = append(a,v.Name)
			a = append(a,v.Img)
			jiachang = append(jiachang,a)
		}else {
			if v.Jiachangcai != "" {
				//fmt.Println("jiachang")
				var a []string
				a =append(a,k)
				a = append(a,v.Name)
				a = append(a,v.Img)
				jiachang = append(jiachang,a)
			}
				var a []string
				a =append(a,k)
				a = append(a,v.Name)
				a = append(a,v.Img)
				var b string
				for _,i:=range v.Renqun{
					b = b + i+"\n"
				}
			a=append(a,b)
			fmt.Println("===")
			fmt.Println(b)
			fmt.Println("====")
			renqunstr = append(renqunstr,a)
		}
		n++
	}
}


