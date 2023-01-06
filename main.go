package main

import (
	"fmt"
)

func solve(s string, t string) string {
	// write code here
	i := len(s) - 1
	j := len(t) - 1
	max := j
	if i > j {
		max = i
	}
	var low int
	var high int
	result := make([]byte, max+2)
	pos := max + 1
	for i >= 0 && j >= 0 {
		sum := int(s[i]-'0') + int(t[j]-'0') + high
		low = sum % 10
		high = sum / 10
		result[pos] = byte(low + '0')
		pos--
		i--
		j--
	}
	for i >= 0 {
		sum := int(s[i]-'0') + high
		low = sum % 10
		high = sum / 10
		result[pos] = byte(low + '0')
		pos--
		i--
	}
	for j >= 0 {
		sum := int(t[j]-'0') + high
		low = sum % 10
		high = sum / 10
		result[pos] = byte(low + '0')
		pos--
		j--
	}
	if high > 0 {
		result[pos] = byte(high + '0')
		return string(result)
	} else {
		return string(result[1:])
	}
}

func main() {
	s := "0"
	t := "0"
	r := solve(s, t)
	fmt.Printf("%s", r)

}
