module github.com/pandajc/go-demo/grpc-svr

go 1.16

require (
	google.golang.org/grpc v1.43.0
	github.com/pandajc/go-demo/common v0.0.0-incompatible
	github.com/pandajc/go-demo/protos v0.0.0-incompatible
)
replace (
	github.com/pandajc/go-demo/common  => ../common
	github.com/pandajc/go-demo/protos  => ../protos
)
