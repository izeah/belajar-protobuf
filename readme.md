Install Protoc

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

Compile Protobuf

`protoc --proto_path=. ./*.proto --go_out=.`

Execute

`go run main.go`
