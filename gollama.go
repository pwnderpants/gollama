package gollama

import "fmt"

type GollamaTest struct {
	Message string
	Name    string
}

func NewLlama(n string) *GollamaTest {
	return &GollamaTest{
		Message: "Hello",
		Name:    n,
	}
}

func (g *GollamaTest) Hello() {
	fmt.Printf("%s %s!", g.Message, g.Name)
}
