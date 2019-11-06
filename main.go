package main

import (
	"fmt"

	"go.uber.org/dig"
)

var container = dig.New()

func init() {
	container.Provide(func() int {
		return 0
	})
	container.Provide(NewItemA)
	container.Provide(NewItemB)
	container.Provide(NewItemC)
	container.Provide(NewItemD)
	container.Provide(NewItemE)
}

func main() {
	var ie *ItemE
	if err := container.Invoke(func(e *ItemE) {
		fmt.Printf("%v\n", e)
		ie = e
	}); err != nil {
		fmt.Println(err)
	}

	ie.A.Name = "a"
	fmt.Println(ie)
}

func NewItemA() ItemA {
	return ItemA{Name: ""}
}

type ItemA struct {
	Name string
}

func NewItemB(A ItemA, B int) ItemB {
	return ItemB{A: A, B: B}
}

type ItemB struct {
	A ItemA
	B int
}

func NewItemC(A ItemA, B ItemB) ItemC {
	return ItemC{A: A, B: B}
}

type ItemC struct {
	A ItemA
	B ItemB
}

func NewItemD(A ItemA, B ItemB, C ItemC) ItemD {
	return ItemD{A: A, B: B, C: C}
}

type ItemD struct {
	A ItemA
	B ItemB
	C ItemC
}

func NewItemE(A ItemA, B ItemB, C ItemC, D ItemD) *ItemE {
	return &ItemE{A: A, B: B, C: C, D: D}
}

type ItemE struct {
	A ItemA
	B ItemB
	C ItemC
	D ItemD
}
