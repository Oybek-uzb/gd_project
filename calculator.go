package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "2+2, 5-5, 8*6, 8/4, 89*65, 24/5"
	exps := strings.Split(str, ", ")
	var result string

	res := make(chan string)
	for _, exp := range exps {
		go calculateStr(exp, res)
	}

	for _ = range exps {
		result = fmt.Sprint(<-res, " ", result)
	}

	fmt.Println("Results: ", result)
}

func calculateStr(expression string, ch chan string) {
	indexOfOperator := strings.IndexAny(expression, "+-*/")
	operator := string(expression[indexOfOperator])
	args := strings.Split(expression, operator)

	a, _ := strconv.ParseFloat(args[0], 32)
	b, _ := strconv.ParseFloat(args[1], 32)

	switch operator {
	case "+":
		ch <- fmt.Sprint(expression, "=", a+b)
	case "-":
		ch <- fmt.Sprint(expression, "=", a-b)
	case "*":
		ch <- fmt.Sprint(expression, "=", a*b)
	case "/":
		ch <- fmt.Sprint(expression, "=", a/b)
	}
}
