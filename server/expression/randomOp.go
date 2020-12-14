package expression

import (
	"math/rand"
	"strconv"
	"time"
)

//	the expression send to the front
type Ex struct {
	Expression string `json:"Expression"`
	Ans        string `json:"Answer"`
}

//	level表示难度即几个运算数，max表示运算数中的最大值,nums表示出多少道题目
//	前端传回的数据
type LevelChoose struct {
	Level 	int	`json:"level"`
	Max		int	`json:"max"`
	Nums	int	`json:"nums"`
}

//首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用

func StartExamine(ch LevelChoose) []byte  {
	allExpressions:=make([]Ex,0)
	ExpressionGenerate(ch.Level,ch.Max,ch.Nums,&allExpressions)
	return GenerateWholeJsonExpression(0,"",allExpressions)
}

//	生成满足要求的运算式子,最多四则运算
//	level表示难度即几个运算数，max表示运算数中的最大值,nums表示出多少道题目
func ExpressionGenerate(level int, max int, nums int, allExpressions *[]Ex){
	//:=make([]Ex,nums)
	//var allExpressions	[]Ex
	for ; nums > 0; nums-- {
		//Operands:=make([]int,0)
		//Operators:=make([]string,0)
		//var temp Ex
		var Operands,Operators,ans = Expression(level, max)
		//var Operands,Operators,ans = Expression(level, max, &Operands, &Operators)
		temp:=GetExStruct(Operands,Operators,ans)
		*allExpressions=append(*allExpressions,temp)
	}
}


//返回Ex结构体
func GetExStruct(operands []int, operators []string, ans int) Ex {
	var exp	Ex
	exp.Expression=ExpToString(operands,operators)
	exp.Ans=strconv.Itoa(ans)
	return exp
	//return &exp
}

//返回运算式的string
func ExpToString(operands []int, operators []string) string {
	exp:=strconv.Itoa(operands[0])
	for i := 0; i < len(operators); i++ {
		exp=exp+operators[i]+strconv.Itoa(operands[i+1])
	}
	return exp
}

//	生成满足要求的运算式子,最多四则运算
//	level表示难度，max表示运算数中的最大值
func Expression(level int, max int) ([]int,[]string,int){
	for {
		Operands:=make([]int,0)
		Operators:=make([]string,0)
		//operands,operators:=GetExpression(level,max, Operands,Operators)
		GetExpression(level,max, &Operands,&Operators)
		ans,succeed:=CheckAns(Operands,Operators,200)
		if succeed{
			return Operands,Operators,ans
		}
		continue
	}
}


//	level表示难度，max表示运算数中的最大值
//	增加了slice的值,需要传指针
func GetExpression(level int, max int, Operands *[]int, Operators *[]string){
	//Operands:=GetOperands(level,max)
	GetOperands(level,max,Operands)
	//Operators:=GetOperators(level-1)
	GetOperators(level-1,Operators)
	//return Operands,Operators
}


//	获得一个操作数，大小在[1,max]
//	返回int
func Operand(max int) int{
	rand.New(rand.NewSource(time.Now().UnixNano()))
	//rand.Seed(time.Now().Unix())
	return rand.Intn(max) + 1
}

func GetOperands(times int, max int, operands *[]int){
	for ; times > 0; times-- {
		*operands=append(*operands,Operand(max))
	}
}


//	获得一个运算符
//	返回一个string
func Operator() string{
	operators := "+-*/"
	//rand.Seed(time.Now().Unix())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return string(operators[rand.Intn(4)])
}

func GetOperators(times int, operands *[]string){
	//var operands []string
	for ; times > 0; times-- {
		*operands=append(*operands,Operator())
	}
}
