# Package scope
We can share variables  and functions across multiples go files as long as they use the same package 

We just need to make sure that we run them all in 1 command eg:
```sh
go run main.go greetings.go
```

'''Note!''' The variables should not be declared inside the function other otherwise you will get an undefined error