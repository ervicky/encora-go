# Encora GO

## 1. Parse the text strings into a tree of node structs.
#### For parser-service run the following command
```bash
go build -o ./bin/parser-service -i ./cmd/parser-service/ && ./bin/parser-service
```
##### For test case run the following command
```bash
go test -c -o ./bin/parser-service-test github.com/encora-go/cmd/parser-service
go tool test2json -t ./bin/parser-service-test -test.v
```


## 2. Adjust the code, so that the double() function can only be called 5 times concurrently.
#### For counter-service run the following command
```bash
go build -o ./bin/counter-service -i ./cmd/counter-service/ && ./bin/counter-service
```
##### For test case run the following command
```bash
go test -c -o ./bin/counter-service-test github.com/encora-go/cmd/counter-service
go tool test2json -t ./bin/counter-service-test -test.v
```


## 3. Find the right package for following:
This can be found in the [package-selection](/cmd/package-selection/ReadMe.md).