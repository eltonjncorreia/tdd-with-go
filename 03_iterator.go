package main

const countRepeating = 5

func Repeat(name string) string {
	var r string

	for i := 0; i < countRepeating; i++ {
		r += name
	}
	return r
}
