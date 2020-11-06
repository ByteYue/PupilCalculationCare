package expression

import (
	"encoding/json"
	"fmt"
)

type ExGroup struct {
	ID  int `json:"id"`
	Exp Ex  `json:"exp"` //注意要大写才能被外界访问
}

//func Fileprocess() bool {
//
//}

//func saveFile(data []byte,) bool {
//	var name string//准备用用户名和第几次使用生成文件名
//}

//将运算式生成json数据
//times:用户第几次使用,name:用户名 用户一同生成文件名
func GenerateWholeJsonExpression(times int, name string, allExpression []Ex) []byte {
	js := make(map[string]interface{})
	jsons := make([]interface{}, 0)
	for flag := 0; flag < len(allExpression); flag++ {
		temp := make([]interface{}, 0)
		var elem ExGroup
		elem.ID = flag
		elem.Exp = allExpression[flag]
		temp = append(temp, elem)
		//temp=append(temp,"expression"+strconv.Itoa(flag))
		//_,p:=ToJson(Expression())
		//temp=append(temp,allExpression[flag])
		jsons = append(jsons, temp)
	}
	js["expressions"] = jsons
	//println(jsons)
	bt, errs := json.Marshal(js)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	//err=
	return bt
	//fmt.Println(string(bt))
	//fmt.Print(len(jsons))
}
