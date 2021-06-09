package main

func Repeat(name string) string {
	var r string

	for i := 0; i < 5; i++ {
		r += name
	}
	return r
}
