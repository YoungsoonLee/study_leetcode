package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	res := bytes.Buffer{}
	p := numRows*2 - 2 // !!!!
	fmt.Println(p)

	// 처음 선 처리
	for i := 0; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	//중간 선 처리 !!!!!
	for r := 1; r <= numRows-2; r++ {
		res.WriteByte(s[r])
		for k := p; k-r < len(s); k += p {
			fmt.Println(p, k, r)
			res.WriteByte(s[k-r])
			if k+r < len(s) {
				res.WriteByte(s[k+r])
			}
		}
	}

	// 마지막 처리
	for i := numRows - 1; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	fmt.Println(res.String())
	return res.String()
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	convert(s, numRows)
}
