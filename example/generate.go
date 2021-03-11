package example

//go:generate protoc --go-ascii_out=. --go-ascii_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative -I . example.proto
