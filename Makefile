.DEFAULT_GOAL := help


.PHONY: help
help:
	@echo "Usage: "
	@echo "make	generate-login:		...	generate login proto"
	@echo "make	clean-login:		...	clean login proto generated files"
	@echo ""


LOGIN_PROTO=./model/login.proto


# Generates protobufs and gRPC client and server for the geak
.PHONY:
generate-login:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ${LOGIN_PROTO}

.PHONY:
clean-login:
	rm -rf ./model/*.pb.go

