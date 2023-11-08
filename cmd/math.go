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
	eval     uint16
}

func (p polynomial) String() string {
	lhs := p.lhs
	switch p.lhs.(type) {
	case *polynomial:
		lhs = "(" + p.lhs.(*polynomial).String() + ")"
	}

	rhs := p.rhs
	switch p.rhs.(type) {
	case *polynomial:
		rhs = "(" + p.rhs.(*polynomial).String() + ")"
	}

	return fmt.Sprintf("%v %v %v", lhs, p.operator, rhs)
}

func generatePrimitive(operator string, count int) (results []*polynomial) {
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
		var eval uint16
		switch f {
		case 0:
			flag = "+"
			eval = a + b
		case 1:
			flag = "-"
			if a < b {
				b, a = a, b
			}
			eval = a - b
		case 2:
			flag = "×"

			if a > 10 && b > 10 {
				if b > a {
					a %= 10
				} else {
					b %= 10
				}
			}

			eval = a * b

		case 3:
			flag = "÷"
			b %= 10
			if b == 0 {
				b++
			}
			a *= b
			eval = a / b
		}

		if a <= 1 || b <= 1 {
			goto RETRY
		}

		results = append(results, &polynomial{operator: flag, lhs: a, rhs: b, eval: eval})
	}
	return
}

func generateExpr(parent []*polynomial, nestedLevel int) {
	if nestedLevel == 0 {
		return
	}

	polys := generatePrimitive("", len(parent))

	for i := 0; i < len(parent); i++ {
		parentEval := parent[i].eval
		var expr *polynomial
		fmt.Println(parentEval, polys[i].eval)
		if parentEval > polys[i].eval {
			expr = &polynomial{operator: "+", lhs: polys[i], rhs: (parentEval - polys[i].eval), eval: parentEval}
		} else if parentEval < polys[i].eval {
			expr = &polynomial{operator: "-", lhs: polys[i], rhs: (polys[i].eval - parentEval), eval: parentEval}
		}

		switch _rand() % 2 {
		case 0:
			parent[i].lhs = expr
		case 1:
			parent[i].rhs = expr
		}
	}

	generateExpr(polys, nestedLevel-1)
	return
}

func _rand() uint16 {
	bts := make([]byte, 2)
	io.ReadFull(rand.Reader, bts)
	return binary.LittleEndian.Uint16(bts)
}
