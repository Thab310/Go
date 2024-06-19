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
3. Sometimes it is useful to group tests around a "thing" and then have subtests describing different scenarios. A benefit of this approach is you can set up shared code that can be used in the other tests.

4. Discipline (flow)

* Write a test

* Make the compiler pass

* Run the test, see that it fails and check the error message is meaningful

* Write enough code to make the test pass

* Refactor