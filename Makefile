package = github.com/deepthawtz/percrpc

.PHONY: install release test

protoc:
	protoc -I ./percentage/ ./percentage/percentage.proto --go_out=plugins=grpc:percentage
