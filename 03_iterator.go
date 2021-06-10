package main

var countRepeating = 5

func Repeat(name string, n ...int) string {
	var r string

	if n != nil {
		countRepeating = n[0]
	}

	for i := 0; i < countRepeating; i++ {
		r += name
	}
	return r
}
