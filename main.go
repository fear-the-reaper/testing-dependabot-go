package main

import "github.com/testing-dependabot-go/something"

func main() {
	hWorld := something.HelloWorld{
		Msg: "I saw you dancing in a crowded room",
	}
	hWorld.Greet()
}
