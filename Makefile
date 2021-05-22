IMAGE := go-jwt-auth
gen:
	@protoc --gofast_out=plugins=grpc:pb proto/*.proto 
run-dev:
	@reflex -r "\.go$$" -s -- sh -c "go run ./"
build-local:
	@docker build --target=dev \
		-t $(IMAGE):dev . 

	
