package change

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strings"
)

func Done()  {
	a := readtable()
	Changedata(a)
}

func Changedata(s [][]string)  {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	a := f.GetSheetName(1)
	for n,i:=range s{
		line := i[12:]
		res := handline(line)
		f.SetSheetRow(a,"M"+fmt.Sprint(n+1),&res)
	}
	if err := f.SaveAs("data.xlsx"); err != nil {
		fmt.Println(err)
	}
}


func readtable() [][]string {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	a := f.GetSheetName(1)

	s := f.GetRows(a)
	return s
}

func handline(s []string) []string {
	var a []string
	for _,i:=range s{
		cline := strings.Split(i,"\n")
		for _,x:=range cline{
			if strings.HasPrefix(x,"http") {
				continue
			}
			a =append(a,x)
		}
	}
	return a
}