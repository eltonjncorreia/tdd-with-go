package main

import "fmt"

const GreetingInEnglish = "Hello"

func HelloWorld(name string) string {
	if name == "" {
		return GreetingInEnglish + ", world"
	}
	return GreetingInEnglish + ", " + name
}

func main() {
	fmt.Println(HelloWorld("Elton"))
}
