package getlist

var usertoken string = "datatype=0&checktypeid=3950&usertoken=698569fd391e0603fc11d5270f6001a7&userid=1203561"

//func RquestPOST(s string) string {
//	client := &http.Client{}
//
//	req, err := http.NewRequest("POST", s, strings.NewReader(usertoken))
//	if err != nil {
//		// handle error
//	}
//
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	req.Header.Set("User-Agent","zhenlipai/4.9.5(Linux; Android 4.4.4)")
//
//	resp, err := client.Do(req)
//
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		// handle error
//	}
//	return string(body)
//}


var CheckList string = `[
  {
    "InspectionID": "3917",
    "TypeName": "血常规",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "3950",
        "TypeName": "血常规",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3953",
        "TypeName": "红细胞计数",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3954",
        "TypeName": "红细胞比积（PCV）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3955",
        "TypeName": "红细胞分布宽度（RDW）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3956",
        "TypeName": "平均红细胞体积（MCV）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3957",
        "TypeName": "平均红细胞血红蛋白含量（MCH）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3958",
        "TypeName": "平均红细胞血红蛋白浓度（MCHC）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3959",
        "TypeName": "血红蛋白浓度（HGB）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3960",
        "TypeName": "白细胞计数（WBC）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3961",
        "TypeName": "中性粒细胞计数（NEUT）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3962",
        "TypeName": "嗜酸性粒细胞数",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3963",
        "TypeName": "血小板压积（PCT）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3964",
        "TypeName": "淋巴细胞计数（LY）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3965",
        "TypeName": "平均血小板体积（MPV）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3966",
        "TypeName": "血小板计数（PLT）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3967",
        "TypeName": "嗜碱性粒细胞计数（B）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3968",
        "TypeName": "嗜酸性粒细胞计数（E）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3969",
        "TypeName": "中性粒细胞比例（NEUT%）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3970",
        "TypeName": "淋巴细胞比值（LY%）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3971",
        "TypeName": "单核细胞计数（MONO）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4164",
        "TypeName": "网织红细胞计数（RC）",
        "TypeFID": "3917",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3918",
    "TypeName": "尿常规",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "3972",
        "TypeName": "尿常规",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3973",
        "TypeName": "尿中红细胞（RBC，BLO）",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3974",
        "TypeName": "尿酸碱度（pH）",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3975",
        "TypeName": "尿中白细胞（WBC，LEU）",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3976",
        "TypeName": "尿蛋白定量（PRO）",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3977",
        "TypeName": "尿胆红素（BIL）",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3978",
        "TypeName": "尿结晶",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3979",
        "TypeName": "尿比重（SG）",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3980",
        "TypeName": "尿沉渣管型",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3981",
        "TypeName": "浊度",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3982",
        "TypeName": "尿液颜色（UCO）",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3983",
        "TypeName": "尿胆原（URO）",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3984",
        "TypeName": "尿糖（GLU）",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3986",
        "TypeName": "尿酮体（KET）",
        "TypeFID": "3918",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3919",
    "TypeName": "便常规",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "3987",
        "TypeName": "便常规",
        "TypeFID": "3919",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3988",
        "TypeName": "粪便颜色",
        "TypeFID": "3919",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3989",
        "TypeName": "粪便白细胞",
        "TypeFID": "3919",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3990",
        "TypeName": "粪便红细胞",
        "TypeFID": "3919",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3991",
        "TypeName": "粪便隐血试验（OBT）",
        "TypeFID": "3919",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3992",
        "TypeName": "幽门螺杆菌",
        "TypeFID": "3919",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3993",
        "TypeName": "粪寄生虫卵",
        "TypeFID": "3919",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4166",
        "TypeName": "粪便性状",
        "TypeFID": "3919",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3924",
    "TypeName": "C-反应蛋白",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "3994",
        "TypeName": "C-反应蛋白",
        "TypeFID": "3924",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4180",
        "TypeName": "C反应蛋白检验（CRP）",
        "TypeFID": "3924",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3925",
    "TypeName": "血栓三项",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4019",
        "TypeName": "血栓三项",
        "TypeFID": "3925",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4020",
        "TypeName": "血清纤维蛋白降解产物测定",
        "TypeFID": "3925",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4021",
        "TypeName": "抗凝血酶III活性和抗凝血酶III抗原",
        "TypeFID": "3925",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4167",
        "TypeName": "凝血酶时间",
        "TypeFID": "3925",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3926",
    "TypeName": "乙肝五项",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4022",
        "TypeName": "乙肝两对半",
        "TypeFID": "3926",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4023",
        "TypeName": "乙型肝炎病毒表面抗体（HBsAb）",
        "TypeFID": "3926",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4024",
        "TypeName": "乙型肝炎病毒核心抗体（HBcAb）",
        "TypeFID": "3926",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4025",
        "TypeName": "乙型肝炎病毒e抗原（HBeAg）",
        "TypeFID": "3926",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4026",
        "TypeName": "乙型肝炎e抗体",
        "TypeFID": "3926",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4168",
        "TypeName": "乙型肝炎病毒表面抗原（HBsAg）",
        "TypeFID": "3926",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3927",
    "TypeName": "凝血常规",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4028",
        "TypeName": "凝血常规",
        "TypeFID": "3927",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4029",
        "TypeName": "凝血酶时间",
        "TypeFID": "3927",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4030",
        "TypeName": "凝血酶原时间（PT）",
        "TypeFID": "3927",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4031",
        "TypeName": "纤维蛋白原（Fg,FIB）",
        "TypeFID": "3927",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4032",
        "TypeName": "血浆D-二聚体测定",
        "TypeFID": "3927",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4181",
        "TypeName": "活化部分凝血活酶时间（APTT）",
        "TypeFID": "3927",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4189",
        "TypeName": "血浆凝血酶原片段1+2检测",
        "TypeFID": "3927",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3928",
    "TypeName": "甲功五项",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4033",
        "TypeName": "甲状腺功能检查",
        "TypeFID": "3928",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4034",
        "TypeName": "三碘甲状腺原氨酸（T<sub>3</sub>）",
        "TypeFID": "3928",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4035",
        "TypeName": "甲状腺素（T<sub>4</sub>）",
        "TypeFID": "3928",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4036",
        "TypeName": "促甲状腺激素（TSH）",
        "TypeFID": "3928",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4203",
        "TypeName": "抗甲状腺球蛋白抗体",
        "TypeFID": "3928",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4204",
        "TypeName": "抗甲状腺过氧化酶抗体",
        "TypeFID": "3928",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3929",
    "TypeName": "电解质",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4038",
        "TypeName": "电解质检查",
        "TypeFID": "3929",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4039",
        "TypeName": "血清镁",
        "TypeFID": "3929",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4040",
        "TypeName": "血清钙",
        "TypeFID": "3929",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4041",
        "TypeName": "血清钠",
        "TypeFID": "3929",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4042",
        "TypeName": "血清钾",
        "TypeFID": "3929",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4043",
        "TypeName": "血清无机磷",
        "TypeFID": "3929",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4173",
        "TypeName": "血清氯化物",
        "TypeFID": "3929",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3930",
    "TypeName": "血气分析",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4044",
        "TypeName": "血气分析",
        "TypeFID": "3930",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4045",
        "TypeName": "氧分压",
        "TypeFID": "3930",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4046",
        "TypeName": "二氧化碳分压（PCO<sub>2</sub>）",
        "TypeFID": "3930",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4047",
        "TypeName": "酸碱度值测定（PH）",
        "TypeFID": "3930",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4048",
        "TypeName": "剩余碱（BE,BD）",
        "TypeFID": "3930",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4049",
        "TypeName": "实际碳酸氢根(AB)",
        "TypeFID": "3930",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4050",
        "TypeName": "二氧化碳总量",
        "TypeFID": "3930",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4206",
        "TypeName": "血氧饱和度",
        "TypeFID": "3930",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4207",
        "TypeName": "血清铁",
        "TypeFID": "3930",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4208",
        "TypeName": "标准碳酸氢根（SB）",
        "TypeFID": "3930",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4209",
        "TypeName": "阴离子间隙",
        "TypeFID": "3930",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4210",
        "TypeName": "血氧含量",
        "TypeFID": "3930",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4211",
        "TypeName": "血清总铁结合力",
        "TypeFID": "3930",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3931",
    "TypeName": "贫血原因检测",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4052",
        "TypeName": "贫血检测",
        "TypeFID": "3931",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4053",
        "TypeName": "血清铁蛋白",
        "TypeFID": "3931",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4054",
        "TypeName": "维生素B<sub>12</sub>",
        "TypeFID": "3931",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4055",
        "TypeName": "叶酸",
        "TypeFID": "3931",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4056",
        "TypeName": "血清不饱和铁结合力",
        "TypeFID": "3931",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4057",
        "TypeName": "红细胞渗透脆性试验",
        "TypeFID": "3931",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4058",
        "TypeName": "蔗糖溶血试验",
        "TypeFID": "3931",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4059",
        "TypeName": "尿含铁血黄素定性试验",
        "TypeFID": "3931",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "5650",
        "TypeName": "血清总铁结合力",
        "TypeFID": "3931",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3932",
    "TypeName": "免疫球蛋白检测",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4060",
        "TypeName": "免疫球蛋白",
        "TypeFID": "3932",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4061",
        "TypeName": "补体C<sub>3</sub>测定",
        "TypeFID": "3932",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4062",
        "TypeName": "血清免疫球蛋白M(IgM)",
        "TypeFID": "3932",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4063",
        "TypeName": "血清免疫球蛋白G(IgG)",
        "TypeFID": "3932",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4064",
        "TypeName": "血清免疫球蛋白A(IgA)",
        "TypeFID": "3932",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4065",
        "TypeName": "血清免疫球蛋白E(IgE)",
        "TypeFID": "3932",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4066",
        "TypeName": "血清β<sub>2</sub>微球蛋白（β<sub>2</sub>-MG）",
        "TypeFID": "3932",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4544",
        "TypeName": "血清免疫球蛋白D(IgD)",
        "TypeFID": "3932",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4545",
        "TypeName": "血清补体C<sub>4</sub>含量测定",
        "TypeFID": "3932",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3933",
    "TypeName": "病毒检查",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4067",
        "TypeName": "乙型肝炎病毒表面抗原（HBsAg）",
        "TypeFID": "3933",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4068",
        "TypeName": "EB病毒抗体（EBV-Ab）",
        "TypeFID": "3933",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4069",
        "TypeName": "梅毒螺旋体抗体",
        "TypeFID": "3933",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4070",
        "TypeName": "抗艾滋病（AIDS）抗体",
        "TypeFID": "3933",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4071",
        "TypeName": "抗巨细胞病毒抗体（CMV-AB）",
        "TypeFID": "3933",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4072",
        "TypeName": "抗巨细胞病毒IgM抗体",
        "TypeFID": "3933",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4073",
        "TypeName": "乙型肝炎病毒核心抗体（HBcAb）",
        "TypeFID": "3933",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4074",
        "TypeName": "丙型肝炎抗体检测（抗-HCV-Ig）",
        "TypeFID": "3933",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4075",
        "TypeName": "乙型肝炎病毒e抗原（HBeAg）",
        "TypeFID": "3933",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4076",
        "TypeName": "乙型肝炎e抗体",
        "TypeFID": "3933",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4077",
        "TypeName": "乙型肝炎病毒表面抗体（HBsAb）",
        "TypeFID": "3933",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "14300",
        "TypeName": "病毒检查",
        "TypeFID": "3933",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3934",
    "TypeName": "24h尿蛋白测定",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4078",
        "TypeName": "尿肌酐",
        "TypeFID": "3934",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4079",
        "TypeName": "尿酸",
        "TypeFID": "3934",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4080",
        "TypeName": "尿蛋白定量",
        "TypeFID": "3934",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6837",
        "TypeName": "24h尿蛋白",
        "TypeFID": "3934",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3935",
    "TypeName": "自身抗体",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4081",
        "TypeName": "抗核抗体（ANA）",
        "TypeFID": "3935",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4082",
        "TypeName": "抗核抗体（ANA）或抗核因子",
        "TypeFID": "3935",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4083",
        "TypeName": "抗双链DNA抗体",
        "TypeFID": "3935",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4084",
        "TypeName": "抗着丝点抗体",
        "TypeFID": "3935",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4085",
        "TypeName": "抗中性粒细胞胞浆",
        "TypeFID": "3935",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4086",
        "TypeName": "抗PM-1抗体",
        "TypeFID": "3935",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4087",
        "TypeName": "抗Sa抗体",
        "TypeFID": "3935",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4088",
        "TypeName": "抗Scl-70抗体",
        "TypeFID": "3935",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4089",
        "TypeName": "抗SS-A和SS-B抗体",
        "TypeFID": "3935",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4090",
        "TypeName": "抗U1RNP抗体",
        "TypeFID": "3935",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "14302",
        "TypeName": "自身抗体检测",
        "TypeFID": "3935",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3936",
    "TypeName": "血型",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4091",
        "TypeName": "ABO血型鉴定",
        "TypeFID": "3936",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4092",
        "TypeName": "Rh血型鉴定",
        "TypeFID": "3936",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "14305",
        "TypeName": "血型测定",
        "TypeFID": "3936",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3937",
    "TypeName": "肿瘤标记物检验",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4093",
        "TypeName": "癌抗原72-4（CA72-4）",
        "TypeFID": "3937",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4094",
        "TypeName": "癌抗原15-3（CA15-3）",
        "TypeFID": "3937",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4095",
        "TypeName": "神经元特异性烯醇化酶（NSE）",
        "TypeFID": "3937",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4096",
        "TypeName": "血清铁蛋白（SF）",
        "TypeFID": "3937",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4097",
        "TypeName": "癌胚抗原（CEA）",
        "TypeFID": "3937",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4098",
        "TypeName": "人绒毛膜促性腺激素（HCG）",
        "TypeFID": "3937",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4099",
        "TypeName": "癌抗原125（CA125）",
        "TypeFID": "3937",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4101",
        "TypeName": "癌抗原19-9（CA19-9）",
        "TypeFID": "3937",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4102",
        "TypeName": "鳞状上皮细胞癌抗原（SCC）",
        "TypeFID": "3937",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4103",
        "TypeName": "细胞角蛋白19片段抗原（CYFRA<sub>21-1</sub>）",
        "TypeFID": "3937",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6984",
        "TypeName": "血清甲胎蛋白(AFP)",
        "TypeFID": "3937",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "14308",
        "TypeName": "肿瘤标记物",
        "TypeFID": "3937",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3938",
    "TypeName": "血清四项",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4104",
        "TypeName": "乙型肝炎病毒表面抗原（HBsAg）",
        "TypeFID": "3938",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4105",
        "TypeName": "梅毒螺旋体抗体",
        "TypeFID": "3938",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4106",
        "TypeName": "抗艾滋病（AIDS）抗体",
        "TypeFID": "3938",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4107",
        "TypeName": "丙型肝炎抗体检测（抗-HCV-Ig）",
        "TypeFID": "3938",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "14309",
        "TypeName": "血清四项",
        "TypeFID": "3938",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3939",
    "TypeName": "普通生化检验",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4110",
        "TypeName": "血清钙",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4111",
        "TypeName": "血清钠",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4112",
        "TypeName": "血清钾",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4113",
        "TypeName": "血清镁",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4114",
        "TypeName": "血清无机磷",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4115",
        "TypeName": "血液葡萄糖",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4116",
        "TypeName": "淀粉酶",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4119",
        "TypeName": "血清肌酸激酶同工酶（CKI）",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4120",
        "TypeName": "尿谷草转移酶",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4121",
        "TypeName": "单胺氧化酶",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4122",
        "TypeName": "鳞状上皮细胞癌抗原（SCC）",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4124",
        "TypeName": "低密度脂蛋白胆固醇（LDL-C）",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4125",
        "TypeName": "血肌酐（Cr，Crea）",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4126",
        "TypeName": "脂蛋白a（Lp-a）",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4128",
        "TypeName": "载脂蛋白B(ApoB)",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4130",
        "TypeName": "脂肪酶",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4131",
        "TypeName": "血清碱性磷酸酶（ALP）",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4133",
        "TypeName": "血清尿酸（UA）",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4135",
        "TypeName": "血清唾液酸测定",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4136",
        "TypeName": "血清总蛋白（TP）",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4137",
        "TypeName": "血清白蛋白（ALB,A）",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4138",
        "TypeName": "糖化血清蛋白",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4140",
        "TypeName": "载脂蛋白AI",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "14310",
        "TypeName": "普通生化检验",
        "TypeFID": "3939",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3940",
    "TypeName": "脑脊液常规",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4142",
        "TypeName": "脑脊液常规检查（CSF）",
        "TypeFID": "3940",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4143",
        "TypeName": "脑脊液颜色",
        "TypeFID": "3940",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4144",
        "TypeName": "脑脊液细胞计数（CST）",
        "TypeFID": "3940",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4145",
        "TypeName": "脑脊液蛋白定性",
        "TypeFID": "3940",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4146",
        "TypeName": "脑脊液细胞分类计数",
        "TypeFID": "3940",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "14102",
        "TypeName": "脑脊液透明度",
        "TypeFID": "3940",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3941",
    "TypeName": "脑脊液生化检查",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4147",
        "TypeName": "脑脊液蛋白定量",
        "TypeFID": "3941",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4148",
        "TypeName": "脑脊液葡萄糖",
        "TypeFID": "3941",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4149",
        "TypeName": "脑脊液氯化物",
        "TypeFID": "3941",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3942",
    "TypeName": "HLA",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4150",
        "TypeName": "白细胞抗原系统（HLA）",
        "TypeFID": "3942",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4151",
        "TypeName": "人白细胞抗原（HLA-B27）",
        "TypeFID": "3942",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4152",
        "TypeName": "HLA定型",
        "TypeFID": "3942",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "14299",
        "TypeName": "白细胞抗原（HLA）",
        "TypeFID": "3942",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3943",
    "TypeName": "纤溶系统",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4153",
        "TypeName": "纤溶系统检查",
        "TypeFID": "3943",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4154",
        "TypeName": "血清纤维蛋白降解产物测定",
        "TypeFID": "3943",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4155",
        "TypeName": "α<sub>2</sub>-纤溶酶抑制抗原",
        "TypeFID": "3943",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "14301",
        "TypeName": "纤维蛋白溶解时间",
        "TypeFID": "3943",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3944",
    "TypeName": "尿培养",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4156",
        "TypeName": "真菌培养检查",
        "TypeFID": "3944",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4157",
        "TypeName": "尿液细菌培养",
        "TypeFID": "3944",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4158",
        "TypeName": "中段尿细菌培养计数",
        "TypeFID": "3944",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "14303",
        "TypeName": "尿培养",
        "TypeFID": "3944",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3945",
    "TypeName": "性腺五项",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4159",
        "TypeName": "性腺五项",
        "TypeFID": "3945",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4160",
        "TypeName": "孕酮（P）",
        "TypeFID": "3945",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4161",
        "TypeName": "睾酮（T）",
        "TypeFID": "3945",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4162",
        "TypeName": "黄体生成素（LH）",
        "TypeFID": "3945",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4163",
        "TypeName": "雌二醇（E<sub>2</sub>）",
        "TypeFID": "3945",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "14304",
        "TypeName": "促卵泡激素（FSH）",
        "TypeFID": "3945",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3946",
    "TypeName": "肝功能",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "3995",
        "TypeName": "肝功能",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3996",
        "TypeName": "总胆红素（TBIL,STB）",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3997",
        "TypeName": "总胆汁酸（TBA）",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3998",
        "TypeName": "白蛋白(A，Alb)",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "3999",
        "TypeName": "血清总蛋白（TP）",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4000",
        "TypeName": "丙氨酸氨基转移酶；谷丙转氨酶（ALT）",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4001",
        "TypeName": "乳酸脱氢酶（LDH，LD）",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4002",
        "TypeName": "淀粉酶（Amy,AMS）",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4003",
        "TypeName": "脂肪酶（LPS,LIP）",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4004",
        "TypeName": "血清碱性磷酸酶（ALP）",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4005",
        "TypeName": "血清肌酸激酶同工酶",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4006",
        "TypeName": "肌酸激酶(CPK)",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4007",
        "TypeName": "尿谷草转移酶(UGOT)",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4008",
        "TypeName": "血清γ-谷氨酰基转移酶(γ-GT)",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4009",
        "TypeName": "血清不饱和铁结合力",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4174",
        "TypeName": "直接胆红素(CB)",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6877",
        "TypeName": "脑磷脂胆固醇絮状试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6878",
        "TypeName": "血清谷草转氨酶与血清谷-丙转氨酶比值",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6879",
        "TypeName": "间接胆红素(IBIL)",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6881",
        "TypeName": "利多卡因代谢试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6883",
        "TypeName": "非结合胆红素（SIB,IBIL）",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6884",
        "TypeName": "维生素K反应试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6885",
        "TypeName": "吲哚氰绿试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6886",
        "TypeName": "肝腹水的检查",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6887",
        "TypeName": "血氨",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6888",
        "TypeName": "氨基比林呼气试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6889",
        "TypeName": "双向琼脂扩散试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6891",
        "TypeName": "蛋白电泳",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6892",
        "TypeName": "反映肝脏间质变化的试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6894",
        "TypeName": "肝促凝血活酶试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6895",
        "TypeName": "抗戊型肝炎病毒抗体",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6896",
        "TypeName": "动脉血酮体比测定",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6897",
        "TypeName": "反映肝细胞损伤的试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6898",
        "TypeName": "血清丙氨酸氨基转移同工酶",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6899",
        "TypeName": "阿-曼氏试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6900",
        "TypeName": "反映肝脏排泄功能的试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6901",
        "TypeName": "乙型肝炎抗原抗体检测",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6902",
        "TypeName": "肝纤维化指标检查",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6903",
        "TypeName": "色氨酸耐量试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6904",
        "TypeName": "丁型肝炎抗体",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6905",
        "TypeName": "胰高血糖素负荷试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6906",
        "TypeName": "1分钟胆红素",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6908",
        "TypeName": "δ-胆红素",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6909",
        "TypeName": "反映肝脏贮备功能的试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6910",
        "TypeName": "硫酸锌浊度试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6912",
        "TypeName": "乳酸脱氢酶同工酶",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6913",
        "TypeName": "血清天冬氨酶氨基转移酶同工酶",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6914",
        "TypeName": "絮状沉淀试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6915",
        "TypeName": "血清胆汁酸测定（TBA）",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6916",
        "TypeName": "氨基酸清除率测定",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6917",
        "TypeName": "直接胆红素与总胆红素比值",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6918",
        "TypeName": "直接胆红素/间接胆红素比值",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6919",
        "TypeName": "定量肝功能试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6920",
        "TypeName": "胆固醇酯与总胆固醇比值",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6921",
        "TypeName": "对氨马尿酸合成试验",
        "TypeFID": "3946",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3947",
    "TypeName": "肾功能",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4010",
        "TypeName": "肾功能检查",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4011",
        "TypeName": "血清尿酸（UA）",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4012",
        "TypeName": "尿尿素",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4175",
        "TypeName": "血肌酐（Cr，Crea）",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4191",
        "TypeName": "血尿素氮（BUN）",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4192",
        "TypeName": "尿蛋白",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4193",
        "TypeName": "选择性蛋白尿指数",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4194",
        "TypeName": "β<sub>2</sub>微球蛋白清除试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4195",
        "TypeName": "血内生肌酐清除率",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4196",
        "TypeName": "尿素氮/肌酐比值",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4197",
        "TypeName": "酚红排泄试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6922",
        "TypeName": "尿微量转铁蛋白",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6923",
        "TypeName": "尿CO<sub>2</sub>测定",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6924",
        "TypeName": "尿道抬举试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6925",
        "TypeName": "滤过钠排泄分数ＦＥＮa",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6926",
        "TypeName": "血清尿素",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6928",
        "TypeName": "肾小球滤过分数(GFF)",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6929",
        "TypeName": "可乐定试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6930",
        "TypeName": "肾小球滤过率（GFR）",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6932",
        "TypeName": "开博通试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6933",
        "TypeName": "二氧化碳结合力(CO<sub>2</sub>CP)",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6934",
        "TypeName": "放射性核素肾扫描",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6935",
        "TypeName": "高钠试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6936",
        "TypeName": "肾血流图",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6937",
        "TypeName": "左肾静脉“胡桃夹”综合征检查",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6938",
        "TypeName": "间碘苄胍γ-闪烁照相",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6939",
        "TypeName": "导尿术",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6940",
        "TypeName": "禁水试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6941",
        "TypeName": "肾癌检查",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6942",
        "TypeName": "尿β2微球蛋白清除率",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6943",
        "TypeName": "肌丙素试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6944",
        "TypeName": "螺内酯试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6945",
        "TypeName": "尿流率测定",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6946",
        "TypeName": "经皮肾镜检查",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6947",
        "TypeName": "氯化铵负荷试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6948",
        "TypeName": "血清阳离子测定",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6949",
        "TypeName": "尿道功能检查",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6950",
        "TypeName": "肾脏超音波",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6951",
        "TypeName": "卧立位试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6952",
        "TypeName": "尿液N-乙酰-β-D-氨基葡萄糖苷酶（UNAG）",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6953",
        "TypeName": "角尺度试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6954",
        "TypeName": "尿TH糖蛋白（THP）",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6955",
        "TypeName": "肾血流量（RPF）",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6956",
        "TypeName": "选择蛋白指数（SPI）",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6957",
        "TypeName": "尿浓缩与稀释功能试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6958",
        "TypeName": "卡托普利肾图试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6959",
        "TypeName": "尿白蛋白清除率",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6960",
        "TypeName": "肾图检查",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6961",
        "TypeName": "无离子水（自由水）清除率",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6962",
        "TypeName": "偏振光显微镜检查",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6963",
        "TypeName": "肾盂内压测定",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6964",
        "TypeName": "盐水输注试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6965",
        "TypeName": "分侧肾静脉肾素活性测定",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6966",
        "TypeName": "静脉注射肾盂摄影检查",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6967",
        "TypeName": "尿稀释试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6968",
        "TypeName": "棉签试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6969",
        "TypeName": "卡托普利试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6970",
        "TypeName": "血清阴离子测定",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6971",
        "TypeName": "中剂量地塞米松抑制试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6972",
        "TypeName": "血清肌酐",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6973",
        "TypeName": "浓缩试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6974",
        "TypeName": "肾小球功能试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6975",
        "TypeName": "肾上腺素试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6976",
        "TypeName": "尿道冲洗试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6977",
        "TypeName": "阴道尿道瘘显象",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6978",
        "TypeName": "补液试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6979",
        "TypeName": "HCO<sub>3</sub><sup>-</sup>重吸收试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6980",
        "TypeName": "肾盂镜",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6981",
        "TypeName": "急性肾功能衰竭诊断试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6982",
        "TypeName": "速尿激发试验",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "6983",
        "TypeName": "尿丙氨酸氨基肽酶（UAAP）",
        "TypeFID": "3947",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3948",
    "TypeName": "血糖",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4013",
        "TypeName": "血糖检查",
        "TypeFID": "3948",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4014",
        "TypeName": "血液葡萄糖",
        "TypeFID": "3948",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4184",
        "TypeName": "糖化血清蛋白",
        "TypeFID": "3948",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  },
  {
    "InspectionID": "3949",
    "TypeName": "血脂",
    "TypeFID": "0",
    "IsLower": "0",
    "DataType": "0",
    "IsFinish": "0",
    "list": [
      {
        "InspectionID": "4015",
        "TypeName": "血脂检查",
        "TypeFID": "3949",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4016",
        "TypeName": "血浆总胆固醇（TC）",
        "TypeFID": "3949",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4017",
        "TypeName": "高密度脂蛋白胆固醇（HDL-C）",
        "TypeFID": "3949",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4018",
        "TypeName": "低密度脂蛋白胆固醇（LDL-C）",
        "TypeFID": "3949",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      },
      {
        "InspectionID": "4183",
        "TypeName": "甘油三酯",
        "TypeFID": "3949",
        "IsLower": "1",
        "DataType": "0",
        "IsFinish": "1"
      }
    ]
  }
]`

var Datas string = `{
	"status": 0,
	"msg": "success",
	"data": [
		[{
			"InfoID": "13762",
			"ItemID": "23",
			"ItemMsg": "\u8840\u5e38\u89c4",
			"ItemName": "\u68c0\u9a8c\u9879\u76ee"
		}],
		[{
			"InfoID": "13820",
			"ItemID": "24",
			"ItemMsg": "\u9759\u8109\u8840\/\u672b\u68a2\u8840",
			"ItemName": "\u6837\u672c\u8981\u6c42"
		}],
		[{
			"InfoID": "13821",
			"ItemID": "25",
			"ItemMsg": "\u5168\u81ea\u52a8\u8840\u6db2\u7ec6\u80de\u4eea\u8ba1\u6570\u6cd5\uff1b\u76ee\u89c6\u8ba1\u6570\u6cd5\uff1b\u6c30\u5316\u9ad8\u94c1\u8840\u7ea2\u86cb\u767d\u6cd5\uff1b\u6e29\u6c0f\u6cd5\uff1b\u6bdb\u7ec6\u8840\u7ba1\u6cd5\uff1bMiller\u7aa5\u76d8\u6cd5\uff1b\u714c\u7126\u6cb9\u84dd\u67d3\u8272\u767e\u5206\u8ba1\u6570\u6cd5\u3002",
			"ItemName": "\u6d4b\u5b9a\u65b9\u6cd5"
		}],
		[{
			"InfoID": "13822",
			"ItemID": "26",
			"ItemMsg": "\u8840\u5e38\u89c4\u5404\u68c0\u9a8c\u9879\u76ee\u53c2\u8003\u503c\u8303\u56f4\u5982\u56fe",
			"ItemName": "\u53c2\u8003\u503c"
		}, {
			"InfoID": "117289",
			"ItemID": "26",
			"ItemMsg": "http:\/\/meditool.cn\/uploadfiles\/inspectionpic\/3950_K4IESU.jpg"
		}, {
			"InfoID": "117290",
			"ItemID": "26",
			"ItemMsg": "http:\/\/meditool.cn\/uploadfiles\/inspectionpic\/3950_H2XJ4P.jpg"
		}, {
			"InfoID": "117291",
			"ItemID": "26",
			"ItemMsg": "http:\/\/meditool.cn\/uploadfiles\/inspectionpic\/3950_MFHGW9.jpg"
		}, {
			"InfoID": "117292",
			"ItemID": "26",
			"ItemMsg": "http:\/\/meditool.cn\/uploadfiles\/inspectionpic\/3950_7YQDH0.jpg"
		}],
		[{
			"InfoID": "13827",
			"ItemID": "28",
			"ItemMsg": "\u5bf9\u5e94\u68c0\u67e5\u9879\u76ee\u53ca\u5176\u4e34\u5e8a\u610f\u4e49\u3001\u4e34\u5e8a\u75be\u75c5\u5bf9\u5e94\u5982\u56fe\u6240\u793a\u3002\u8840\u5e38\u89c4\u68c0\u67e5\u5728\u5168\u8eab\u4f53\u68c0\u4e2d\u662f\u57fa\u672c\u7684\u4f53\u68c0\u9879\u76ee\uff0c\u5b83\u7684\u610f\u4e49\u5728\u4e8e\u53ef\u4ee5\u53d1\u73b0\u8bb8\u591a\u5168\u8eab\u6027\u75be\u75c5\u7684\u65e9\u671f\u8ff9\u8c61\uff0c\u8bca\u65ad\u662f\u5426\u8d2b\u8840\uff0c\u662f\u5426\u6709\u8840\u6db2\u7cfb\u7edf\u75be\u75c5\uff0c\u53cd\u5e94\u9aa8\u9ad3\u7684\u9020\u8840\u529f\u80fd\u7b49\u3002",
			"ItemName": "\u4e34\u5e8a\u610f\u4e49"
		}, {
			"InfoID": "13857",
			"ItemID": "28",
			"ItemMsg": "http:\/\/meditool.cn\/uploadfiles\/inspectionpic\/3950_U0J52R.jpg"
		}, {
			"InfoID": "13858",
			"ItemID": "28",
			"ItemMsg": "http:\/\/meditool.cn\/uploadfiles\/inspectionpic\/3950_UP3BV6.jpg"
		}, {
			"InfoID": "13859",
			"ItemID": "28",
			"ItemMsg": "http:\/\/meditool.cn\/uploadfiles\/inspectionpic\/3950_X46C4J.jpg"
		}, {
			"InfoID": "13860",
			"ItemID": "28",
			"ItemMsg": "http:\/\/meditool.cn\/uploadfiles\/inspectionpic\/3950_7F93FV.jpg"
		}],
		[{
			"InfoID": "13861",
			"ItemID": "29",
			"ItemMsg": "\u4e0b\u80a2\u9759\u8109\u66f2\u5f20\uff1b\u8111\u8840\u6813\u5f62\u6210\uff1bIgA\u80be\u75c5\uff1b\u6025\u6027\u80be\u5c0f\u7403\u80be\u708e\uff1b\u7535\u51b0\u7bb1\u98df\u7269\u4e2d\u6bd2\uff1b\u6025\u6027\u8f93\u5375\u7ba1\u708e\uff1b\u4e73\u7a81\u708e\uff1b\u809d\u708e\uff1b\u5c40\u7076\u8282\u6bb5\u6027\u80be\u5c0f\u7403\u786c\u5316\uff1b\u6bdb\u7ec6\u80de\u767d\u8840\u75c5\uff0c\u598a\u5a20\u5408\u5e76\u767d\u8840\u75c5\uff1b\u524d\u5ead\u5927\u817a\u764c\uff1b\u8d56\u7279\u7efc\u5408\u5f81\u6027\u5de9\u819c\u708e\uff1b\u6025\u6027\u6d77\u7ef5\u7aa6\u6813\u585e\u6027\u9759\u8109\u708e\uff1b\u7736\u9aa8\u819c\u4e0b\u8113\u80bf\uff1b\u4e73\u7cdc\u6cfb\uff1b\u80ba\u708e\uff1b\u767e\u8349\u67af\u4e2d\u6bd2\uff1b\u5e93\u6b23\u7efc\u5408\u5f81\uff1b\u5e72\u71e5\u7efc\u5408\u5f81\uff1b\u80a5\u80d6\u75c7\uff1b\u7532\u72b6\u817a\u529f\u80fd\u51cf\u9000\u75c7\uff1b\u963b\u585e\u6027\u7761\u7720\u547c\u5438\u6682\u505c\u7efc\u5408\u5f81\uff1b\u514b\u7f57\u6069\u75c5\uff1b\u8fc7\u654f\u6027\u7d2b\u765c\uff1b\u8111\u51fa\u8840\uff1b\u80c6\u56ca\u708e\uff1b\u652f\u6c14\u7ba1\u708e\uff1b\u80be\u764c\uff1b\u89c6\u795e\u7ecf\u810a\u9ad3\u708e\uff1b\u7ec6\u83cc\u6027\u89d2\u819c\u708e\uff1b\u80ba\u708e\uff1b\u767e\u65e5\u54b3\uff1b\u513f\u7ae5\u9aa8\u9ad3\u589e\u751f\u5f02\u5e38\u7efc\u5408\u5f81\uff1b\u80ba\u52a8\u8109\u9ad8\u538b\uff1b\u766b\u75eb\uff1b\u6d41\u611f\uff1b\u54ee\u5598\uff1b\u7279\u53d1\u6027\u810a\u67f1\u4fa7\u51f8\uff1b\u6d88\u5316\u4e0d\u826f\uff1b\u9ec4\u75b8\uff1b\u5ddd\u5d0e\u75c5\uff1b\u6cd5\u5e03\u91cc\u75c5\uff1b\u7cfb\u7edf\u6027\u7ea2\u6591\u72fc\u75ae\uff1b\u6025\u6027\u80c3\u80a0\u708e\uff1b\u793e\u533a\u83b7\u5f97\u6027\u80ba\u708e\uff1b\u611f\u67d3\u6027\u4f11\u514b\uff1b\u5fc3\u529b\u8870\u7aed\uff1b\u8113\u6bd2\u75c7\uff1b\u6025\u6027\u9611\u5c3e\u708e\uff1b\u8d2b\u8840\uff1b\u4e0a\u6d88\u5316\u9053\u51fa\u8840\uff1b\u8179\u75db\uff1b\u9ad8\u8840\u538b\uff1b\u8c35\u5984\uff1b\u70df\u96fe\u75c5\uff1b\u809d\u8870\u7aed\uff1b\u6025\u8179\u75c7\uff1b\u9885\u5185\u52a8\u8109\u7624\uff1b\u4e09\u53c9\u795e\u7ecf\u75db\uff1b\u9aa8\u9ad3\u708e\uff1b\u767b\u9769\u70ed\uff1b\u9ec4\u70ed\u75c5\uff1b\u94dc\u7eff\u5047\u5355\u80de\u83cc\u611f\u67d3\uff1b\u975e\u6dcb\u7403\u83cc\u6027\u5c3f\u9053\u708e\uff1b\u6025\u6027\u4e0a\u547c\u5438\u9053\u611f\u67d3\uff1b\u72c2\u72ac\u75c5",
			"ItemName": "\u76f8\u5173\u75be\u75c5"
		}],
		[{
			"InfoID": "13863",
			"ItemID": "32",
			"ItemMsg": "\u9759\u8109\u91c7\u8840\u591a\u91c7\u7528\u4f4d\u4e8e\u4f53\u8868\u7684\u6d45\u9759\u8109\uff0c\u901a\u5e38\u91c7\u7528\u8098\u90e8\u9759\u8109\u3001\u624b\u80cc\u9759\u8109\u3001\u5185\u8e1d\u9759\u8109\u6216\u80a1\u9759\u8109\u3002\u8098\u524d\u9759\u8109\u662f\u7edd\u5927\u591a\u6570\u4eba(\u5a74\u5e7c\u513f\u9664\u5916)\u7684\u9996\u9009\u91c7\u8840\u90e8\u4f4d\u3002\u6b64\u5904\u4e00\u822c\u8840\u7ba1\u8f83\u660e\u663e,\u75bc\u75db\u611f\u8f83\u8f7b,\u64cd\u4f5c\u65b9\u4fbf\u6613\u884c\u3002\u5c0f\u513f\u53ef\u91c7\u9888\u5916\u9759\u8109\u8840\u6db2\u3002\u5176\u6b21\u9009\u62e9\u80a1\u9759\u8109\u3002\r\n\u65b9\u6cd5\uff1a\r\n(1) \u5907\u9f50\u7528\u7269\uff0c\u6807\u672c\u5bb9\u5668\u4e0a\u8d34\u597d\u6807\u7b7e\uff0c\u6838\u5bf9\u65e0\u8bef\u540e\u5411\u60a3\u8005\u89e3\u91ca\u4ee5\u53d6\u5f97\u5408\u4f5c\u3002\u9732\u51fa\u60a3\u8005\u624b\u81c2\uff0c\u9009\u62e9\u9759\u8109\uff0c\u4e8e\u9759\u8109\u7a7f\u523a\u90e8\u4f4d\u4e0a\u65b9\u7ea64-6cm\u5904\u624e\u7d27\u6b62\u8840\u5e26\uff0c\u5e76\u5631\u60a3\u8005\u63e1\u7d27\u62f3\u5934\uff0c\u4f7f\u9759\u8109\u5145\u76c8\u663e\u9732\u3002\r\n(2) \u5e38\u89c4\u6d88\u6bd2\u76ae\u80a4\uff0c\u5f85\u5e72\u3002\r\n(3) \u5728\u7a7f\u523a\u90e8\u4f4d\u4e0b\u65b9\uff0c\u4ee5\u5de6\u624b\u62c7\u6307\u62c9\u7d27\u76ae\u80a4\u5e76\u56fa\u5b9a\u9759\u8109\uff0c\u53f3\u624b\u6301\u6ce8\u5c04\u5668\uff0c\u9488\u5934\u659c\u9762\u5411\u4e0a\u4e0e\u76ae\u80a4\u621015\u5ea6-30\u5ea6\uff0c\u5728\u9759\u8109\u4e0a\u6216\u65c1\u4fa7\u523a\u5165\u76ae\u4e0b\uff0c\u518d\u6cbf\u9759\u8109\u8d70\u5411\u6f5c\u884c\u523a\u5165\u9759\u8109\uff0c\u89c1\u56de\u8840\u540e\u5c06\u9488\u5934\u7565\u653e\u5e73\uff0c\u7a0d\u524d\u884c\u56fa\u5b9a\u4e0d\u52a8\uff0c\u62bd\u8840\u81f3\u9700\u8981\u91cf\u65f6\uff0c\u653e\u677e\u6b62\u8840\u5e26\uff0c\u5631\u60a3\u8005\u677e\u62f3\uff0c\u5e72\u68c9\u7b7e\u6309\u538b\u7a7f\u523a\u70b9\uff0c\u8fc5\u901f\u62d4\u51fa\u9488\u5934\uff0c\u5e76\u5c06\u60a3\u8005\u524d\u81c2\u5c48\u66f2\u538b\u8feb\u7247\u523b\u3002\r\n(4) \u5378\u4e0b\u9488\u5934\uff0c\u5c06\u8840\u6db2\u6cbf\u7ba1\u58c1\u7f13\u7f13\u6ce8\u5165\u5bb9\u5668\u5185\uff0c\u5207\u52ff\u5c06\u6ce1\u6cab\u6ce8\u5165\uff0c\u4ee5\u514d\u6eb6\u8840\u3002\u5bb9\u5668\u5185\u653e\u6709\u73bb\u7483\u73e0\u65f6\u5e94\u8fc5\u901f\u6447\u52a8\uff0c\u4ee5\u9664\u53bb\u7ea4\u7ef4\u86cb\u767d\u539f\uff1b\u5982\u7cfb\u6297\u51dd\u8bd5\u7ba1\uff0c\u5e94\u5728\u53cc\u624b\u5185\u65cb\u8f6c\u6413\u52a8\uff0c\u4ee5\u9632\u51dd\u56fa\uff1b\u5982\u7cfb\u5e72\u71e5\u8bd5\u7ba1\uff0c\u4e0d\u5e94\u6447\u52a8\uff1b\u5982\u7cfb\u6db2\u4f53\u57f9\u517b\u57fa\uff0c\u5e94\u4f7f\u8840\u6db2\u4e0e\u57f9\u517b\u6db2\u6df7\u5300\uff0c\u5e76\u5728\u8840\u6db2\u6ce8\u5165\u57f9\u517b\u74f6\u524d\u540e\uff0c\u7528\u706b\u7130\u6d88\u6bd2\u74f6\u53e3\uff0c\u6ce8\u610f\u52ff\u4f7f\u74f6\u585e\u63a5\u89e6\u8840\u6db2\u3002\r\n\u62bd\u8840\u91cf\u7684\u591a\u5c11\u662f\u6839\u636e\u5316\u9a8c\u5185\u5bb9\u7684\u4e0d\u540c\u53ca\u9879\u76ee\u7684\u591a\u5c11\u6765\u51b3\u5b9a\u7684\uff0c\u4e00\u822c\u57285ml\u5de6\u53f3\u3002",
			"ItemName": "\u68c0\u9a8c\u8fc7\u7a0b"
		}],
		[{
			"InfoID": "13865",
			"ItemID": "34",
			"ItemMsg": "\u68c0\u67e5\u524d\uff1a\r\n(1) \u62bd\u8840\u524d\u4e00\u5929\u4e0d\u5403\u8fc7\u4e8e\u6cb9\u817b\u3001\u9ad8\u86cb\u767d\u98df\u7269\uff0c\u907f\u514d\u5927\u91cf\u996e\u9152\u3002\u8840\u6db2\u4e2d\u7684\u9152\u7cbe\u6210\u5206\u4f1a\u76f4\u63a5\u5f71\u54cd\u68c0\u9a8c\u7ed3\u679c\u3002\r\n(2) \u4f53\u68c0\u524d\u4e00\u5929\u7684\u665a\u516b\u65f6\u4ee5\u540e\uff0c\u5e94\u5f00\u59cb\u7981\u98df12\u5c0f\u65f6\uff0c\u4ee5\u514d\u5f71\u54cd\u68c0\u6d4b\u7ed3\u679c\u3002\r\n(3) \u62bd\u8840\u65f6\u5e94\u653e\u677e\u5fc3\u60c5\uff0c\u907f\u514d\u56e0\u6050\u60e7\u9020\u6210\u8840\u7ba1\u7684\u6536\u7f29\uff0c\u589e\u52a0\u91c7\u8840\u7684\u56f0\u96be\u3002\r\n\r\n\u68c0\u67e5\u540e\uff1a\u3000\u3000\r\n(1) \u62bd\u8840\u540e\uff0c\u9700\u5728\u9488\u5b54\u5904\u8fdb\u884c\u5c40\u90e8\u6309\u538b3-5\u5206\u949f\uff0c\u8fdb\u884c\u6b62\u8840\u3002\u6ce8\u610f\uff1a\u4e0d\u8981\u63c9\uff0c\u4ee5\u514d\u9020\u6210\u76ae\u4e0b\u8840\u80bf\u3002\r\n(2) \u6309\u538b\u65f6\u95f4\u5e94\u5145\u5206\u3002\u5404\u4eba\u7684\u51dd\u8840\u65f6\u95f4\u6709\u5dee\u5f02\uff0c\u6709\u7684\u4eba\u9700\u8981\u7a0d\u957f\u7684\u65f6\u95f4\u65b9\u53ef\u51dd\u8840\u3002\u6240\u4ee5\u5f53\u76ae\u80a4\u8868\u5c42\u770b\u4f3c\u672a\u51fa\u8840\u5c31\u9a6c\u4e0a\u505c\u6b62\u538b\u8feb\uff0c\u53ef\u80fd\u4f1a\u56e0\u672a\u5b8c\u5168\u6b62\u8840\uff0c\u800c\u4f7f\u8840\u6db2\u6e17\u81f3\u76ae\u4e0b\u9020\u6210\u9752\u6de4\u3002\u56e0\u6b64\u6309\u538b\u65f6\u95f4\u957f\u4e9b\uff0c\u624d\u80fd\u5b8c\u5168\u6b62\u8840\u3002\u5982\u6709\u51fa\u8840\u503e\u5411\uff0c\u66f4\u5e94\u5ef6\u957f\u6309\u538b\u65f6\u95f4\u3002\r\n(3) \u62bd\u8840\u540e\u51fa\u73b0\u6655\u9488\u75c7\u72b6\u5982\uff1a\u5934\u6655\u3001\u773c\u82b1\u3001\u4e4f\u529b\u7b49\u5e94\u7acb\u5373\u5e73\u5367\u3001\u996e\u5c11\u91cf\u7cd6\u6c34\uff0c\u5f85\u75c7\u72b6\u7f13\u89e3\u540e\u518d\u8fdb\u884c\u4f53\u68c0\u3002\r\n(4) \u82e5\u5c40\u90e8\u51fa\u73b0\u6de4\u8840\uff0c24\u5c0f\u65f6\u540e\u7528\u6e29\u70ed\u6bdb\u5dfe\u6e7f\u6577\uff0c\u53ef\u4fc3\u8fdb\u5438\u6536\u3002",
			"ItemName": "\u6ce8\u610f\u4e8b\u9879"
		}]
	],
	"enitemmsg": null,
	"iscollect": 0,
	"shareurl": "http:\/\/meditool.cn\/Preview\/checkquick?msgid=3950&datatype=0",
	"sharemsg": {
		"shareimg": "http:\/\/meditool.cn\/shareimg\/jianyandefault.png",
		"sharetitle": null,
		"sharedesc": "\u533b\u5e93\u533b\u5b66\u68c0\u9a8c\uff0c\u62e5\u6709\u6700\u5168\u6700\u4e13\u4e1a\u7684\u68c0\u9a8c\u77e5\u8bc6\uff0c\u4e13\u4e3a\u533b\u52a1\u5de5\u4f5c\u8005\u6253\u9020\u3002"
	}
}`