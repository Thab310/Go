/*
package main

import "fmt"

func main() {
	fmt.Println("Here we go again :)")
}
*/

// It is a good practice to seperate your domain code from the outside-world(side-effets)

////////////////////////////////////////////////////////
/*
package main

import "fmt"

func Hello() string {
	return "Here we go again :)"
}

func main() {
	fmt.Println(Hello())
}
*/
////////////////////////////////////////////////////////

//It's worth thinking about creating constants to capture the meaning of values and sometimes to aid performance.
/*
package main

const EnglishHelloPrefix = "Hello"

func Hello(n string) string {
	return EnglishHelloPrefix + " " + n
}
*/
////////////////////////////////////////////////////////

package main

const (
	Spanish = "Spanish"
	Venda   = "Venda"

	EnglishHelloPrefix   = "Hello"
	SpanishHelloPrefix   = "Hola"
	TshiVendaHelloPrefix = "Ndaa"
)

func Hello(n, l string) string {
	if n == "" {
		n = "world"
	}

	return GreetingPrefix(l) + " " + n
}

func GreetingPrefix(l string) (p string) {
	switch l {
	case Spanish:
		p = SpanishHelloPrefix
	case Venda:
		p = TshiVendaHelloPrefix
	default:
		p = EnglishHelloPrefix
	}

	return
}
