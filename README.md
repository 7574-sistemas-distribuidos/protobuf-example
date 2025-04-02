# Protobuf Example - Golang

## Install protocol buffers interpreter for Golang

Run the following command to install the Go protocol buffers plugin:

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

## Run the protocol buffers compiler

Now run the compiler, specifying the source directory (where your application’s source code lives – the current directory is used if you don’t provide a value), the destination directory (where you want the generated code to go; often the same as $SRC_DIR), and the path to your .proto. In this case, you would invoke:

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/person.proto
protoc -I=. --go_out=. ./person.proto

## References

https://protobuf.dev/