package main

import (
	"github.com/yue/PupilCalculation/expression"
)

type Exp struct {
	Times	string	`json:"Nums"`
	Expree	expression.Ex	`json:"Exp"`
}

func main() {
	//expression.GenerateWholeJsonExpression(expression.StartExamine(2,10,10))
	expression.StartExamine(4,10,10)
}
