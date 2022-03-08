# Go - Json -> Struct -> New Json

Get json data from json file and create a new json file with the same data.

Packages used:

["encoding/json"](https://pkg.go.dev/encoding/json), ["io/ioutil"](https://pkg.go.dev/io/ioutil)

```go
cd golang\json

docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
go version
```
