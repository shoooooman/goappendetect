# goappendetect
Detect assignments in functions as follows:
```go
func f(s []int) {
    s = append(s, 1) // NG
}

func g(s *[]int) {
    *s = append(*s, 1) // OK
}
```
It assumes the cases where the slice should have been passed by reference but are passed by value.

# Install
```sh
go get github.com/shoooooman/goappendetect/cmd/goappendetect
```

# Usage
```sh
go vet -vettool=($which goappendetect) .
```
