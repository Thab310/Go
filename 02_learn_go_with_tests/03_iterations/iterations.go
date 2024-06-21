package main

const repeatedCount = 5

func Repeat(n string) string {
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += n
	}
	return repeated
}
