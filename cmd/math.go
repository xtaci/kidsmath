package cmd

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
)

type polynomial struct {
	operator string
	lhs      interface{}
	rhs      interface{}
}

func (p polynomial) String() string {
	lhs := p.lhs
	switch p.lhs.(type) {
	case polynomial:
		lhs = "(" + p.lhs.(polynomial).String() + ")"
	}

	rhs := p.rhs
	switch p.rhs.(type) {
	case polynomial:
		rhs = "(" + p.rhs.(polynomial).String() + ")"
	}

	return fmt.Sprintf("%v %v %v", lhs, p.operator, rhs)
}

// generate random simple algebra quiz
func generate(operator string, count int) (results []polynomial) {
	for i := 0; i < count; i++ {
	RETRY:
		var a, b, f uint16
		for a <= 1 || b <= 1 {
			a = _rand() % 100
			b = _rand() % 100
		}

		f = _rand() % 4

		if operator != "" {
			switch operator {
			case "+":
				f = 0
			case "-":
				f = 1
			case "*":
				f = 2
			case "/":
				f = 3
			}
		}

		var flag string
		switch f {
		case 0:
			flag = "+"
		case 1:
			flag = "-"
			if a < b {
				b, a = a, b
			}
		case 2:
			flag = "×"

			if a > 10 && b > 10 {
				if b > a {
					a %= 10
				} else {
					b %= 10
				}
			}

		case 3:
			flag = "÷"
			b %= 10
			a *= b
		}

		if a <= 1 || b <= 1 {
			goto RETRY
		}

		results = append(results, polynomial{operator: flag, lhs: a, rhs: b})
	}
	return
}

func polyGenerate(parent []polynomial, nestedlevel int) {
	polys_a := generate("+", len(parent)/2)
	polys_b := generate("*", len(parent)-len(polys_a))
	polys := append(polys_a, polys_b...)

	for i := 0; i < len(parent); i++ {
		right_or_left := _rand() % 2
		switch right_or_left {
		case 0:
			parent[i].lhs = polys[i]
		case 1:
			parent[i].rhs = polys[i]
		}
	}
	return
}

func _rand() uint16 {
	bts := make([]byte, 2)
	io.ReadFull(rand.Reader, bts)
	return binary.LittleEndian.Uint16(bts)
}
