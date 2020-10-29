package expression

import (
	"encoding/json"
	"fmt"
)


type ExGroup struct {
	ID  int `json:"id"`
	Exp Ex  `json:"exp"`//注意要大写才能被外界访问
}

func Fileprocess()  {
	
}

//将运算式生成json数据
func GenerateWholeJsonExpression(allExpression []Ex) []byte {
	js := make(map[string]interface{})
	jsons := make([]interface{},0)
	for flag:=0;flag<len(allExpression);flag++ {
		temp:=make([]interface{},0)
		var elem ExGroup
		elem.ID=flag
		elem.Exp =allExpression[flag]
		temp=append(temp,elem)
		//temp=append(temp,"expression"+strconv.Itoa(flag))
		//_,p:=ToJson(Expression())
		//temp=append(temp,allExpression[flag])
		jsons= append(jsons, temp)
	}
	js["expressions"]=jsons
	//println(jsons)
	bt, errs := json.Marshal(js)
	if errs!=nil {
		fmt.Println("json marshal error:", errs)
	}
	return bt
	//fmt.Println(string(bt))
	//fmt.Print(len(jsons))
}
