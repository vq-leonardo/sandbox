package main

import (
	"fmt"
	"strconv"
	"strings"
)

type calculator struct {
	s1 string
	s2 string
}

func main() {
	cal := &calculator{
		s1: "30032432423423000",
		s2: "200043243243232",
	}

	result := cal.Excute("minus")
	fmt.Println(result)
}

func (c *calculator) Excute(operator string) string {
	var result string
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
	case "minus":
		result = minus(c, lenLoop)
		break
	}

	return result
}

func plus(c *calculator, lenLoop int) string {
	remember := false
	var result []string

	arrS1 := strings.Split(c.s1, "")
	arrS2 := strings.Split(c.s2, "")
	lenArr1 := len(arrS1)
	lenArr2 := len(arrS2)

	for i := 0; i < lenLoop; i++ {
		itemS1, itemS2 := 0, 0
		if i < lenArr1 {
			itemS1, _ = strconv.Atoi(arrS1[lenArr1-i-1])
		}
		if i < lenArr2 {
			itemS2, _ = strconv.Atoi(arrS2[lenArr2-i-1])
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
	return strings.Join(result, "")
}

func minus(c *calculator, lenLoop int) (strR string) {
	remember := false
	negative := false
	var result []string
	numS1, _ := strconv.Atoi(c.s1)
	numS2, _ := strconv.Atoi(c.s2)
	if numS1 < numS2 {
		negative = true
		c.s1, c.s2 = c.s2, c.s1
	}
	arrS1 := strings.Split(c.s1, "")
	arrS2 := strings.Split(c.s2, "")
	lenArr1 := len(arrS1)
	lenArr2 := len(arrS2)

	for i := 0; i < lenLoop; i++ {
		itemS1, itemS2 := 0, 0
		if i < lenArr1 {
			itemS1, _ = strconv.Atoi(arrS1[lenArr1-i-1])
		}
		if i < lenArr2 {
			itemS2, _ = strconv.Atoi(arrS2[lenArr2-i-1])
		}

		if remember {
			itemS2++
			remember = false
		}
		item := itemS1 - itemS2
		if itemS1 < itemS2 {
			item = itemS1 + 10 - itemS2
			remember = true
		}

		strItem := []string{strconv.Itoa(item)}
		result = append(strItem, result...)
	}

	strR = strings.Join(result, "")
	if negative {
		strR = "-" + strR
	}

	return
}

func multiple(c *calculator) (strR string) {

	return
}
