# go_test
Just a simple test in golang.

### Cross-Compile Linux Architecture
```bash
$ GOOS=linux GOARCH=amd64 go build test.go
```
### Cross-Compile Mipsle Architecture (Creator Ci40)
```bash
$ GOOS=linux GOARCH=mipsle go build test.go
```
### Run
```bash
* $ ./test
or
* $ go run test.go
```