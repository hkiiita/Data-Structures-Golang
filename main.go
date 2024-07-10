package main

import (
	"fmt"

	"github.com/hkiiita/data-structures-golang/stack"
)

func main() {
	s := stack.NewStack()

	value, err := s.Peek()
	if err != nil {
		fmt.Printf("Error Peeking : %+v\n", err)
	} else {
		println(value)
	}

	value, err = s.Pop()
	if err != nil {
		fmt.Printf("Error popping value : %+v\n", err)
	} else {
		println(value)
	}

	for _, v := range []int{12, 34, 56, 78, 43423, 567, 879} {
		s.Push(v)

	}

	value, err = s.Peek()
	if err != nil {
		fmt.Printf("Error Peeking : %+v\n", err)
	} else {
		println(value)
	}

	value, err = s.Pop()
	if err != nil {
		fmt.Printf("Error popping value : %+v\n", err)
	} else {
		println(value)
	}

	value, err = s.Peek()
	if err != nil {
		fmt.Printf("Error Peeking : %+v\n", err)
	} else {
		println(value)
	}

}
