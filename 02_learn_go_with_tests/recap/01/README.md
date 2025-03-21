# 01
In Go (and several other programming languages), a `function` or `variable` that starts with a capital letter indicates that it is `"exported"` or `"public"` - `meaning it can be accessed from other packages.` This is a key part of Go's visibility rules:

`(capitalized)` - Can be accessed from other packages that import your package
`myFunction()`(lowercase) - Can only be used within the same package

## Writing tests
Writing a test is just like writing a function, with a few rules

* It needs to be in a file with a name like xxx_test.go

* The test function must start with the word Test

* The test function takes one argument only t *testing.T

* To use the *testing.T type, you need to import "testing", like we did with fmt in the other file

t.Helper() is needed to tell the test suite that this method is a helper. By doing this, when it fails, the line number reported will be in our function call rather than inside our test helper.
