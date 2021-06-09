package main

import "fmt"

const HelloInEnglish = "Hello"

func HelloWorld(nome string) string {
	if nome == "" {
		return HelloInEnglish + ", world"
	}
	return HelloInEnglish + ", " + nome
}

func main() {
	fmt.Println(HelloWorld("Elton"))
}
