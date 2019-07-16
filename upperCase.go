package main

import (
	"fmt"
	"unicode"
)

const xiaomiId = "9hf0P6MBHM7OY6UOWm"
const myId = "KaRinA23906"
const expected = "9hf0_p6_m_b_h_m7_o_y6_u_o_wm"

func main() {
	fmt.Println(IdToLowercase(xiaomiId))
	fmt.Println(IdToLowercase(myId))
}

func IdToLowercase(id string) string {
	id2 := ""
	for _, c := range id {
		if unicode.IsUpper(c) {
			id2 += "_"
			id2 += string(unicode.ToLower(c))
		} else {
			id2+=string(c)
		}
	}
	return id2
}
