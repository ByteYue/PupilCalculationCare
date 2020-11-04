package expression

import "github.com/yue/PupilCalculation/stack"


//	获得运算符的优先级
func priority(op string) int{
	if op == "+" || op == "-" {
		return	0
	}else{
		return 	1
	}
}

func CheckAns(operands []int,operators []string, max int) (int,bool) {
	var	operandsStack	stack.Stack
	var operatorsStack	stack.Stack
	//var temp int
	for i:=0; i < 2*len(operators)+1 ; {
		if i%2 == 0 {//操作数栈
			operandsStack.Push(operands[i/2])
			i++
		}else {//操作符栈
			//当前操作符栈为空,或者current op优先级比栈顶大
			if operatorsStack.Size() == 0|| priority(operatorsStack.Top().(string))<priority(operators[i/2]){
				operatorsStack.Push(operators[i/2])
				i++
				continue
			}else {//current op优先级比栈顶小
				op:=operatorsStack.Pop().(string)
				num1:=operandsStack.Pop().(int)
				num2:=operandsStack.Pop().(int)
				if op == "/" && (num2 < num1||num2%num1>0) {// a/b when a<b or a%b!=0
					return 0,false
				}
				operandsStack.Push(getAnswer(num2,num1,op))
			}
		}
	}
	for ;operatorsStack.Size()>0; {
		op:=operatorsStack.Pop().(string)
		num1:=operandsStack.Pop().(int)
		num2:=operandsStack.Pop().(int)
		if op == "/" && (num2 < num1||num2%num1>0) {
			return 0,false
		}
		operandsStack.Push(getAnswer(num2,num1,op))
	}
	ans:=operandsStack.Pop().(int)
	if ans>=0&&ans<=max {
		return ans,true
	}else {
		return 0,false
	}
}


//计算答案
//前面保证了不会传入op为空和op2为0
func getAnswer(op1 int, op2 int, op string) int {
	if op == "+" {
		return op1+op2
	}
	if op == "-" {
		return op1-op2
	}
	if op == "*" {
		return op1*op2
	}
	return op1/op2
}
