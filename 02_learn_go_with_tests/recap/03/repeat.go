// package iteration

// func Repeat(c string, n int) string {

// 	var repeated string

// 	for i := 0; i < n; i++ {
// 		repeated += c
// 	}

// 	return repeated
// }

package iteration

import "strings"

func Repeat(c string, n int) string {

	var repeated strings.Builder

	for i := 0; i < n; i++ {
		repeated.WriteString(c)
	}

	return repeated.String()
}
