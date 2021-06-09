package main

import "fmt"

func HelloWorld(nome string) string {
	return "Hello, " + nome
}

func main() {
	fmt.Println(HelloWorld("Elton"))
}
