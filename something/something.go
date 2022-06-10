package something

import "fmt"

type HelloWorld struct {
	Msg string
}

func (helloWorld *HelloWorld) Greet() {
	fmt.Println(helloWorld.Msg)
}
