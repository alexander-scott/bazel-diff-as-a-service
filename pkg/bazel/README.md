# Working with the proto file

## Source proto

Source is from https://github.com/bazelbuild/bazel/blob/master/src/main/protobuf/build.proto

Need to add to the proto:

```proto
option go_package = "./bazel";
```

Note: Ensure the bazel version of the workspace matches the content of the proto file!

## Deps

```bash
$ sudo apt install -y protobuf-compiler
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Generate go file from proto file

```bash
$ cd pkg/bazel
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative build.proto
```
