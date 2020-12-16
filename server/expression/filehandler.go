package expression

// ExGroup 对应exp和id
type ExGroup struct {
	Exp Ex  `json:"exp"` //注意要大写才能被外界访问
	ID  int `json:"id"`
}

//ExJSON 对应一维数组名
type ExJSON struct {
	Expressions []Ex `json:"expressions"`
}

//GenerateJSONToFront 生成运算式的json给前端
func GenerateJSONToFront(allExpression []Ex) ExJSON {
	var transdata ExJSON
	transdata.Expressions = allExpression
	return transdata
}

//GenerateWholeJSONExpression 将运算式生成json数据
//times:用户第几次使用,name:用户名 用户一同生成文件名
//abondoned now, 生成的JSON太丑
func GenerateWholeJSONExpression(times int, name string, allExpression []Ex) map[string]interface{} {
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
	//bt, errs := json.Marshal(js)
	//if errs != nil {
	//	fmt.Println("json marshal error:", errs)
	//}
	return js
}
