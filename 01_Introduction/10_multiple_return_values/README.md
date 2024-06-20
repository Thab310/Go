# Key Differences between return and non-return type functions
## Return Type Functions:

Perform operations and return a value.
Must specify the type of the return value.
Example: 
```go
func add(a int, b int) int {
     return a + b 
}
```

## Non-Return Functions:

Perform operations without returning a value.
Do not specify a return type.
Example: 
```go
func greet(name string) {
     fmt.Println("Hello, " + name + "!") 
}
```
In summary:

Use return type functions when you need to get a result back from the function.
Use non-return functions when you only need to perform an operations.

>[TIP]
* Naked Return Example
```GO
package main

import "fmt"

// Function with named return values
func add(a, b int) (sum int) {
    sum = a + b
    return // Naked return
}

func main() {
    result := add(3, 4)
    fmt.Println("Sum:", result) // Output: Sum: 7
}
```
* Explicit Return Example
```GO
package main

import "fmt"

// Function with named return values but using explicit return
func add(a, b int) (sum int) {
    sum = a + b
    return sum // Explicit return
}

func main() {
    result := add(3, 4)
    fmt.Println("Sum:", result) // Output: Sum: 7
}
```