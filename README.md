![Diagrams](/images/2560px-Go_Logo_Blue.svg.png)

Material Used:
1. Videos
- [Net Ninja](https://www.youtube.com/watch?v=etSN4X_fCnM&list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM)

- [Go: The Complete Developer's Guide (Golang)](https://www.udemy.com/course/go-the-complete-developers-guide/)

- [Learn go for beginners crash course](https://www.udemy.com/course-dashboard-redirect/?course_id=4001218)

2. Docs
- [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests)
- [Effective Go](https://go.dev/doc/effective_go)
- [A tour of Go](https://go.dev/tour/welcome/1)
- [Gophercises](https://gophercises.com/)

>[!TIP]
There is no shame in using a linter :)

Install the linter:
```bash
brew install golangci-lint
```
run:
```bash
go mod init <module>
golangci-lint run
```

>[!TIP]
In Go, the capitalization of the first letter of a function (or any identifier) determines its visibility and accessibility outside of its defining package. This is a key aspect of Go's visibility rules and is part of the language's convention for organizing code. Here's the explanation:

- `Capitalized Functions`: Exported, accessible from other packages. aka `Public functions`
- `Lowercase Functions`: Unexported, private to the package. aka `Private functions`




