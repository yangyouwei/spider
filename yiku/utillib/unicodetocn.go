package utillib

import (
	"sync"
)


func RangeString(stringchan chan string,s string,wg *sync.WaitGroup)  {
	for _,i := range s {
		stringchan <- string(i)
	}
	close(stringchan)
	wg.Done()
}

func TurnTocn(c chan string,s *string,wg *sync.WaitGroup)  {
	var fullstring string
	var tmpstr string
	for  {
		i, isClose := <-c
		if !isClose {
			*s = fullstring
			//fmt.Println("*s",&s)
			break
		}
		if i == "\\" {
			tmpstr = "\\"
			u, isClose := <-c
			if !isClose {
				//fmt.Println(fullstring)
				break
			}
			if u == "u" {
				var substring string = "u"
				for i:=0;i<=3;i++{
					us, isClose := <-c
					if !isClose {
						break
					}
					substring = substring + us
				}
				tmpstr = tmpstr+substring
				fullstring = fullstring + UnicodeToCN(tmpstr)
				continue
				//fmt.Println(UnicodeToCN(tmpstr))
			}
			fullstring = fullstring + u
		}
		fullstring = fullstring + i
	}
	wg.Done()
}

func FullTurnToCn(s string,cn *string){
	stringchan := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go TurnTocn(stringchan,cn,&wg)
	go RangeString(stringchan,s,&wg)
	wg.Wait()
}
