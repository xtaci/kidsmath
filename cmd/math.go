package cmd

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
)

type expression struct {
	operator string
	lhs      interface{}
	rhs      interface{}
	eval     uint16
}

func (p *expression) String() string {
	lhs := p.lhs
	switch p.lhs.(type) {
	case *expression:
		lhs = "(" + p.lhs.(*expression).String() + ")"
	}

	rhs := p.rhs
	switch p.rhs.(type) {
	case *expression:
		rhs = "(" + p.rhs.(*expression).String() + ")"
	}

	return fmt.Sprintf("%v %v %v", lhs, p.operator, rhs)
}

func generatePrimitive(operator string, count int, n int, m int) (results []*expression) {
	for i := 0; i < count; i++ {
	RETRY:
		var a, b, f uint16
		for a <= 1 || b <= 1 {
			a = _rand() % uint16(n)
			b = _rand() % uint16(m)
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

			eval = a * b

		case 3:
			flag = "÷"
			if a < b {
				a, b = b, a
			}

			a = b * (a / b)
			eval = a / b
		}

		if a <= 1 || b <= 1 {
			goto RETRY
		}

		results = append(results, &expression{operator: flag, lhs: a, rhs: b, eval: eval})
	}
	return
}

func generateExpr(parent []*expression, nestedLevel int, n int, m int) {
	if nestedLevel == 0 {
		return
	}

	exprs := generatePrimitive("", len(parent), n, m)

	for i := 0; i < len(parent); i++ {

		r := _rand() % 2
		parentEval := parent[i].lhs.(uint16)
		if r == 1 {
			parentEval = parent[i].rhs.(uint16)
		}

		expr := exprs[i]
		if parentEval > exprs[i].eval {
			expr = &expression{operator: "+", lhs: exprs[i], rhs: (parentEval - exprs[i].eval), eval: parentEval}
		} else if parentEval < exprs[i].eval {
			expr = &expression{operator: "-", lhs: exprs[i], rhs: (exprs[i].eval - parentEval), eval: parentEval}
		}

		switch r {
		case 0:
			parent[i].lhs = expr
		case 1:
			parent[i].rhs = expr
		}

	}

	generateExpr(exprs, nestedLevel-1, n, m)
	return
}

func _rand() uint16 {
	bts := make([]byte, 2)
	io.ReadFull(rand.Reader, bts)
	return binary.LittleEndian.Uint16(bts)
}
