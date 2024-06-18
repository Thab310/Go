# Notes

1. It is good to seperate your domain code from your outside world code(side-effects)

2. Writing a test is just like writing a function, with a few rules

* It needs to be in a file whose name ends with `_test.go`

* The test function must start with the word `Test`

* The test function takes one argument only `t *testing.T`

To test run:
```bash
go test
```
