package expression

import (
	"math/rand"
	"strconv"
	"time"
)

//Ex the expression send to the front
type Ex struct {
	ID         int    `json:"id"`
	Expression string `json:"Expression"`
	Ans        string `json:"Answer"`
	Judge      string `json:"Judge"`
	Point      int    `json:"Point"`
}

//LevelChoose level表示难度即几个运算数，max表示运算数中的最大值,nums表示出多少道题目
//	TODO 加上   symbol:符号(1:+ 2:- 3:* 4:/)	anssize:答案范围
//	前端传回的数据
type LevelChoose struct {
	Symbol  []int8 `json:"symbols"`
	Level   int    `json:"level"`
	Max     int    `json:"max"`
	Nums    int    `json:"nums"`
	Anssize int    `json:"anssize"`
}

//StartExamine 首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用
//StartExamine
func StartExamine(ch LevelChoose, DIY bool) ExJSON {
	allExpressions := make([]Ex, 0)
	var opers string
	if DIY == false {
		opers = "+-*/"
	} else {
		for _, elem := range ch.Symbol {
			opers += symbolToOp(elem)
		}
	}
	ExpressionGenerate(ch.Level, ch.Max, ch.Nums, &allExpressions, &opers)
	return GenerateJSONToFront(allExpressions)
}

func symbolToOp(op int8) string {
	if op == 1 {
		return "+"
	} else if op == 2 {
		return "-"
	} else if op == 3 {
		return "*"
	} else {
		return "/"
	}
}

//ExpressionGenerate 生成满足要求的运算式子,最多四则运算
//ExpressionGenerate level表示难度即几个运算数，max表示运算数中的最大值,nums表示出多少道题目
func ExpressionGenerate(level int, max int, nums int, allExpressions *[]Ex, ops *string) {
	//:=make([]Ex,nums)
	//var allExpressions	[]Ex
	for index := 1; index <= nums; index++ {
		//Operands:=make([]int,0)
		//Operators:=make([]string,0)
		//var temp Ex
		var Operands, Operators, ans = Expression(level, max, ops)
		//var Operands,Operators,ans = Expression(level, max, &Operands, &Operators)
		temp := GetExStruct(Operands, Operators, ans, index)
		*allExpressions = append(*allExpressions, temp)
	}
}

//GetExStruct 返回Ex结构体
func GetExStruct(operands []int, operators []string, ans int, ID int) Ex {
	var exp Ex
	exp.ID = ID
	exp.Expression = ExpToString(operands, operators)
	exp.Ans = strconv.Itoa(ans)
	exp.Point = 0
	return exp
	//return &exp
}

//ExpToString 返回运算式的string
func ExpToString(operands []int, operators []string) string {
	exp := strconv.Itoa(operands[0])
	for i := 0; i < len(operators); i++ {
		exp = exp + operators[i] + strconv.Itoa(operands[i+1])
	}
	return exp
}

//Expression 生成满足要求的运算式子,最多四则运算
//Expression	level表示难度，max表示运算数中的最大值
func Expression(level int, max int, ops *string) ([]int, []string, int) {
	for {
		Operands := make([]int, 0)
		Operators := make([]string, 0)
		//operands,operators:=GetExpression(level,max, Operands,Operators)
		GetExpression(level, max, &Operands, &Operators, ops)
		ans, succeed := CheckAns(Operands, Operators, 200)
		if succeed {
			return Operands, Operators, ans
		}
		continue
	}
}

//GetExpression level表示难度，max表示运算数中的最大值
//GetExpression	增加了slice的值,需要传指针
//GetExpression TODO:增加判断是否DIY->是否调用operatorspecial
func GetExpression(level int, max int, Operands *[]int, Operators *[]string, ops *string) {
	//Operands:=GetOperands(level,max)
	GetOperands(level, max, Operands)
	//Operators:=GetOperators(level-1)
	GetOperators(level-1, Operators, ops)
	//return Operands,Operators
}

//Operand 获得一个操作数，大小在[1,max]
//GetOperands返回int
func Operand(max int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	//rand.Seed(time.Now().Unix())
	return rand.Intn(max) + 1
}

//GetOperands 生成运算数
func GetOperands(times int, max int, operands *[]int) {
	for ; times > 0; times-- {
		*operands = append(*operands, Operand(max))
	}
}

//OperatorSpecial 自定义题目难度时使用
//	由前端发回的json中的symbol数组确认str的值
func OperatorSpecial(str *string) string {
	//rand.Seed(time.Now().Unix())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return string((*str)[rand.Intn(4)])
}

//GetOperators 生成运算符切片
func GetOperators(times int, operands *[]string, ops *string) {
	//var operands []string
	for ; times > 0; times-- {
		*operands = append(*operands, OperatorSpecial(ops))
	}
}

//Operator 获得一个运算符 废弃
//	返回一个string
func Operator(operators *string) string {
	//operators := "+-*/"
	//rand.Seed(time.Now().Unix())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return string((*operators)[rand.Intn(4)])
}
