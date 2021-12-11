package handeldata

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strings"
)

func Done()  {
	t1 :=readtable1()
	t3:=readtable3()
	for n,i := range t3{
		for _,y:=range t1{
			for _,m:= range strings.Split(i[3]," ") {
				if m == y[2] {
					if i[1]==y[1] {
						continue
					}
					if i[1] != "" {
						fenlei := i[1]+"\n"+y[1]
						wirtetofile(fenlei,fmt.Sprint(n+1))
					}
					fenlei := y[1]
					wirtetofile(fenlei,fmt.Sprint(n+1))
				}
			}
		}
	}
}



func wirtetofile(s string,n string)  {
	f, err := excelize.OpenFile("temp.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	a := f.GetSheetName(4)
	f.SetCellValue(a,"B"+n,s)
	fmt.Println("wirte line num",n)
	if err := f.SaveAs("temp.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func readtable1() [][]string {
	f, err := excelize.OpenFile("temp.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	a := f.GetSheetName(1)

	s := f.GetRows(a)
	return s
}

func readtable2() [][]string {
	f, err := excelize.OpenFile("temp.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	a := f.GetSheetName(2)

	s := f.GetRows(a)
	return s
}

func readtable3() [][]string {
	f, err := excelize.OpenFile("temp.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	a := f.GetSheetName(4)

	s := f.GetRows(a)
	return s
}