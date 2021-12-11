package main

import (
	"dahua/readexecl"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var linenum string
var sessionid string

var updata_url string = "http://192.168.1.108/RPC2"

func main()  {
	args := os.Args

	if args == nil || len(args) < 2 {
		fmt.Println("COMMAND 行号  sessionid")
		return
	}
	linenum = args[1]
	sessionid = args[2]

	readexecl.Readexecl(linenum)
	res1 := POST_data(updata_url,sessionid)
	a := strings.Contains(res1,"Invalid session in request data")
	res2 := POST_zhuce(updata_url,sessionid)
	b := strings.Contains(res2,"Invalid session in request data")
	res3 := POST_NTP(updata_url,sessionid)
	c := strings.Contains(res3,"Invalid session in request data")
	//fmt.Println(res3)
	if a == false && b == false && c == false {
		fmt.Println(readexecl.Readdata.Addr,readexecl.Readdata.Zoneid)
		fmt.Println("数据更新成功~！")
	}else {
		fmt.Println("数据更新失败，sessionid 错误~！")
	}
}

func POST_data(url string,sessionid string) string {
	client := &http.Client{}
	data := fmt.Sprintf("{\"method\":\"system.multicall\",\"params\":[{\"method\":\"configManager.setConfig\",\"params\":{\"name\":\"ChannelTitle\",\"table\":[{\"Name\":\"固\"}],\"options\":[]},\"id\":530,\"session\":\"%s\"},{\"method\":\"configManager.setConfig\",\"params\":{\"name\":\"VideoWidget\",\"table\":[{\"ChannelTitle\":{\"BackColor\":[0,0,0,128],\"EncodeBlend\":true,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":true,\"Rect\":[148,7511,1773,7928],\"TextAlign\":2},\"Covers\":[{\"BackColor\":[0,0,0,128],\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[0,0,1024,1024]},{\"BackColor\":[0,0,0,128],\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[1024,1024,2048,2048]},{\"BackColor\":[0,0,0,128],\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[2048,2048,3072,3072]},{\"BackColor\":[0,0,0,128],\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[3072,3072,4096,4096]}],\"CustomTitle\":[{\"BackColor\":[0,0,0,128],\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[5321,7450,7931,7868],\"Text\":\"\",\"TextAlign\":2},{\"BackColor\":[0,0,0,128],\"EncodeBlend\":true,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":true,\"Rect\":[5319,7445,7929,7862],\"Text\":\"河北|秦皇岛|北戴河|%s|%s\",\"TextAlign\":2},{\"BackColor\":[0,0,0,128],\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[5321,7450,7931,7868],\"Text\":\"\",\"TextAlign\":2},{\"BackColor\":[0,0,0,128],\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[5321,7450,7931,7868],\"Text\":\"\",\"TextAlign\":2}],\"FontBorder\":true,\"FontSize\":0,\"FontSizeExtra1\":0,\"FontSizeExtra2\":0,\"FontSizeExtra3\":0,\"FontSizeScale\":1,\"FontSizeSnapshot\":0,\"OSDMobileState\":{\"BackColor\":[0,1,1,128],\"EncodeBlend\":true,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":true,\"Rect\":[6295,920,7929,1351]},\"PTZCoordinates\":{\"BackColor\":[0,1,1,128],\"DisplayTime\":0,\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[541,3320,541,3320]},\"PTZDirection\":{\"BackColor\":[0,1,1,128],\"DisplayTime\":5,\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[541,3320,541,3320]},\"PTZOSDMenu\":{\"BackColor\":[0,1,1,128],\"DisplayTime\":0,\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[541,3320,541,3320]},\"PTZOSDMenuViaApp\":{\"BackColor\":[0,1,1,128],\"DisplayTime\":60,\"EncodeBlend\":true,\"EncodeBlendExtra1\":true,\"EncodeBlendExtra2\":true,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":true,\"Rect\":[281,1477,2897,1542]},\"PTZPreset\":{\"BackColor\":[0,1,1,128],\"DisplayTime\":5,\"EncodeBlend\":false,\"FrontColor\":[255,255,255,3],\"PreviewBlend\":false,\"Rect\":[475,1442,475,1442]},\"PTZZoom\":{\"BackColor\":[0,1,1,128],\"DisplayTime\":0,\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[541,3320,541,3320]},\"PictureTitle\":{\"BackColor\":[0,0,0,128],\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"Name\":\"osd.bmp\",\"PreviewBlend\":false,\"Rect\":[0,0,128,128]},\"PtzPattern\":{\"BackColor\":[0,1,1,128],\"DisplayTime\":0,\"EncodeBlend\":true,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[541,3320,541,3320]},\"PtzRS485Detect\":{\"BackColor\":[0,1,1,128],\"DisplayTime\":0,\"EncodeBlend\":true,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[541,3320,541,3320]},\"Temperature\":{\"BackColor\":[0,1,1,128],\"DisplayTime\":0,\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[541,3320,541,3320],\"TemperatureUnit\":\"Centigrade\"},\"TimeTitle\":{\"BackColor\":[0,0,0,128],\"EncodeBlend\":true,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":true,\"Rect\":[5319,352,7929,769],\"ShowWeek\":false,\"WeekPosition\":\"Right\"},\"UserDefinedTitle\":[{\"BackColor\":[0,0,0,128],\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[148,352,1773,769],\"Text\":\"\"},{\"BackColor\":[0,0,0,128],\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[148,352,1773,769],\"Text\":\"\"},{\"BackColor\":[0,0,0,128],\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[148,352,1773,769],\"Text\":\"\"},{\"BackColor\":[0,0,0,128],\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":false,\"Rect\":[148,352,1773,769],\"Text\":\"\"}],\"VoltageStatus\":{\"BackColor\":[0,1,1,128],\"DisplayTime\":0,\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"PreviewBlend\":true,\"Rect\":[541,3320,541,3320]},\"WideHeightRatio\":0}],\"options\":[]},\"id\":531,\"session\":\"%s\"},{\"method\":\"configManager.setConfig\",\"params\":{\"name\":\"FaceFlowStat\",\"table\":[{\"BackColor\":[0,0,0,128],\"EncodeBlend\":false,\"FrontColor\":[255,255,255,0],\"Rect\":[148,263,1773,680],\"ShowFaceDetection\":true,\"ShowFaceRecognition\":false,\"TextAlign\":0}],\"options\":[]},\"id\":532,\"session\":\"%s\"}],\"id\":533,\"session\":\"%s\"}",sessionid,readexecl.Readdata.Zoneid,readexecl.Readdata.Addr,sessionid,sessionid,sessionid)
	reqbody := strings.NewReader(data)
	req, _ := http.NewRequest("POST", url, reqbody)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	req.Header.Set("username","admin")
	req.Header.Set("secure","")
	req.Header.Set("DWebClientSessionID",sessionid)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return ""
	}

	if resp.StatusCode != 200 {
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


func POST_zhuce(url string,sessionid string) string {
	client := &http.Client{}
	data := fmt.Sprintf("{\"method\":\"configManager.setConfig\",\"params\":{\"name\":\"DVRIP\",\"table\":{\"MCASTAddress\":\"224.1.0.0\",\"MCASTEnable\":true,\"MCASTPort\":36666,\"MaxConnections\":10,\"MonMode\":\"TCP\",\"RegisterServer\":{\"DeviceID\":\"%s\",\"Enable\":true,\"Servers\":[{\"Address\":\"121.22.104.243\",\"Port\":%s}]},\"SSLPort\":37776,\"StreamPolicy\":\"None\",\"TCPPort\":37777,\"UDPPort\":37778},\"options\":[]},\"id\":415,\"session\":\"%s\"}",readexecl.Readdata.Subid,readexecl.Readdata.Port,sessionid)
	reqbody := strings.NewReader(data)
	req, _ := http.NewRequest("POST", url, reqbody)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	req.Header.Set("username","admin")
	req.Header.Set("secure","")
	req.Header.Set("DWebClientSessionID",sessionid)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return ""
	}

	if resp.StatusCode != 200 {
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

func POST_NTP(url string,sessionid string) string {
	client := &http.Client{}
	data := fmt.Sprintf("{\"method\":\"system.multicall\",\"params\":[{\"method\":\"configManager.setConfig\",\"params\":{\"name\":\"Locales\",\"table\":{\"DSTEnable\":false,\"DSTEnd\":{\"Day\":2,\"Hour\":0,\"Minute\":0,\"Month\":1,\"Week\":0,\"Year\":2021},\"DSTStart\":{\"Day\":1,\"Hour\":0,\"Minute\":0,\"Month\":1,\"Week\":0,\"Year\":2021},\"DateSeparator\":\"\",\"TimeFormat\":\"yyyy-MM-dd HH:mm:ss\"},\"options\":[]},\"id\":374,\"session\":\"2ea1a7c1e7e4b5b6cd27e444d8e748f6\"},{\"method\":\"configManager.setConfig\",\"params\":{\"name\":\"NTP\",\"table\":{\"Address\":\"time1.aliyun.com\",\"Enable\":true,\"Port\":123,\"TimeZone\":13,\"TimeZoneDesc\":\"Beijing,Chongqing,HongKong,Urumqi\",\"UpdatePeriod\":10},\"options\":[]},\"id\":375,\"session\":\"2ea1a7c1e7e4b5b6cd27e444d8e748f6\"}],\"id\":376,\"session\":\"%s\"}",sessionid)
	reqbody := strings.NewReader(data)
	req, _ := http.NewRequest("POST", url, reqbody)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	req.Header.Set("username","admin")
	req.Header.Set("secure","")
	req.Header.Set("DWebClientSessionID",sessionid)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return ""
	}

	if resp.StatusCode != 200 {
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