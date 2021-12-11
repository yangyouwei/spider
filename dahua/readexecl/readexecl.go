package readexecl

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type Datastr struct {
	Addr string
	Zoneid string
	Port string
	Subid string
}

var Readdata Datastr

func Readexecl(n string) {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	a := f.GetSheetName(1)
	Readdata.Addr = f.GetCellValue(a,"B"+n)+f.GetCellValue(a,"C"+n)+f.GetCellValue(a,"D"+n)+f.GetCellValue(a,"E"+n)+f.GetCellValue(a,"F"+n)
	Readdata.Zoneid = f.GetCellValue(a,"G"+n)+f.GetCellValue(a,"H"+n)
	Readdata.Port = f.GetCellValue(a,"N"+n)
	Readdata.Subid = f.GetCellValue(a,"J"+n)
}

