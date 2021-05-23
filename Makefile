IMAGE := go-jwt-auth
gen:
	@protoc --gofast_out=plugins=grpc:pb proto/*.proto
run-dev:
	@reflex -r "\.go$$" -s -- sh -c "go run ./"
build-local:
	@docker build --target=dev \
		-t $(IMAGE):dev . 
gen-test:
	@protoc -I . --grpc-gateway_out ./gen/go \
    	--grpc-gateway_opt logtostderr=true \
    	--grpc-gateway_opt paths=source_relative \
    	--grpc-gateway_opt generate_unbound_methods=true \
    	proto/user_service.proto
