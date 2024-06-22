package main

func Repeat(n string, repeatedCount int) string {
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += n
	}
	return repeated
}
