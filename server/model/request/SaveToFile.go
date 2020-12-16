package request

import (
	"bufio"
	"fmt"
	"os"
)

func SaveToFile(expression []string, name string) {
	path := name
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("err=%v", err)
		return
	}
	defer file.Close()
	n := 0
	length := len(expression)
	for str := ""; n < length; n++ {
		str = expression[n]
		str += "\n"
		writer := bufio.NewWriter(file)
		count, w_err := writer.WriteString(str)
		writer.Flush()
		if w_err != nil || count == 0 {
			return
		}
	}
}
