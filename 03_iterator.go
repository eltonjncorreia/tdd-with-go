package main

var countRepeating = 5

func Repeat(n int, name string) string {
	var r string

	if n != 0 {
		countRepeating = n
	}

	for i := 0; i < countRepeating; i++ {
		r += name
	}
	return r
}
