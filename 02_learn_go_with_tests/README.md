# LEARN GO WITH TESTS
Go is a good language for learning Test Driven Development `TDD` because it is a simple language to learn and testing is built-in.

>[!NOTE]
Every test has a cost, soo use them wisely.

run the following command to check test coverage
```bash
go test -cover
```

Go can let you write [variadic functions](https://gobyexample.com/variadic-functions) that can take a variable number of arguments.

```go
package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0

    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
```