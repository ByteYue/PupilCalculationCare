package expression

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func GenerateWholeJsonExpression(allExpression []Ex) {
	js := make(map[string]interface{})
	jsons := make([]interface{},0)
	for flag:=0;flag<len(allExpression);flag++ {
		temp:=make([]interface{},0)
		temp=append(temp,"expression"+strconv.Itoa(flag))
		//_,p:=ToJson(Expression())
		temp=append(temp,allExpression[flag])
		jsons= append(jsons, temp)
	}
	js["expressions"]=jsons
	//println(jsons)
	bt, errs := json.Marshal(js)
	if errs!=nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println(string(bt))
	fmt.Print(len(jsons))
}
