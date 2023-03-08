package main

import (
    "fmt"
    "errors"
)

type Superdivider struct {
	answer int
	err error
}

func (s *Superdivider) divide(x int) {
	if s.err != nil {
		return
	} else if x == 0 {
		s.err = errors.New("math: Sorry, you cannot div by 0.")
		return
	}
	s.answer = s.answer / x
}

func divide(a, b, c int){
	z := Superdivider{a, nil}
	z.divide(b)
	z.divide(c)

	if z.err != nil {
		fmt.Println(z.err)
		return
	}
	fmt.Println(z.answer)
}

func main() {
	divide(100, 2, 2)
	divide(100,0, 20)
	divide(100,5,0)
}