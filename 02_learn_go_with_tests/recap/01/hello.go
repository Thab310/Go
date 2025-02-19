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

package main

func Hello(n string) string {
	return "Here we go again :( " + n
}
