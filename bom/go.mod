module pmt/bom

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.38.0
	pmt/product v0.0.0-00010101000000-000000000000
)

replace pmt/product => ../product
