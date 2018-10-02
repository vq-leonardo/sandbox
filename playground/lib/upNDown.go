package main

import (
	"strings"
)

func Arrange(s string) string {
	sl := strings.Split(s, " ")

	for i, _ := range sl {
		if i+1 < len(sl) {
			if (i%2 == 0 && len(sl[i]) > len(sl[i+1])) || (i%2 != 0 && len(sl[i]) < len(sl[i+1])) {
				temp := sl[i]
				sl[i] = sl[i+1]
				sl[i+1] = temp
			}
		}

		if i%2 == 0 {
			sl[i] = strings.ToLower(sl[i])
		} else {
			sl[i] = strings.ToUpper(sl[i])
		}
	}

	return strings.Join(sl, " ")
}

func mainSort() {
	str := Arrange("who hit retaining The That a we taken")
	println(str)
}
