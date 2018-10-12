package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Calculator ...
// type Calculator interface {
// 	Excute()
// 	Plus()
// 	Minus()
// 	Divide()
// 	Multiple()
// }

type calculator struct {
	s1 string
	s2 string
}

func main() {
	cal := &calculator{
		s1: "3000",
		s2: "200000",
	}

	result := cal.Excute("plus")
	fmt.Println(result)
}

func (c *calculator) Excute(operator string) string {
	var result []string
	lenS1 := len(c.s1)
	lenS2 := len(c.s2)
	lenLoop := 0
	if lenS1 > lenS2 {
		lenLoop = lenS1
	} else {
		lenLoop = lenS2
	}

	operator = strings.ToLower(operator)
	switch operator {
	case "plus":
		result = plus(c, lenLoop)
		break
	}

	return strings.Join(result, "")
}

func plus(c *calculator, lenLoop int) (result []string) {
	remember := false

	arrS1 := strings.Split(c.s1, "")
	arrS2 := strings.Split(c.s2, "")
	lenArr1 := len(arrS1)
	lenArr2 := len(arrS2)

	for i := 0; i < lenLoop; i++ {
		itemS1, itemS2 := 0, 0
		if i < lenArr2 {
			itemS1, _ = strconv.Atoi(arrS2[lenArr2-i-1])
		}
		if i < lenArr1 {
			itemS2, _ = strconv.Atoi(arrS1[lenArr1-i-1])
		}

		item := itemS1 + itemS2
		if remember {
			item++
			remember = false
		}
		if item > 9 {
			item -= 10
			remember = true
		}
		strItem := []string{strconv.Itoa(item)}
		result = append(strItem, result...)
	}
	return
}

// func minus(c *calculator, lenLoop int) (result []string) {
// 	remember := false
// 	arrS1 := strings.Split(c.s1, "")
// 	arrS2 := strings.Split(c.s2, "")
// 	lenArr1 := len(arrS1)
// 	lenArr2 := len(arrS2)

// 	for i := 0; i < lenLoop; i++ {
// 		itemS1, itemS2 := 0, 0
// 		if i < lenArr2 {
// 			itemS1, _ = strconv.Atoi(arrS2[lenArr2-i-1])
// 		}
// 		if i < lenArr1 {
// 			itemS2, _ = strconv.Atoi(arrS1[lenArr1-i-1])
// 		}

// 		item := itemS1 - itemS2
// 	}

// 	return
// }
