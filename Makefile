BINARY_NAME=envtoschema

build:
	go build -o ${BINARY_NAME} cmd/envtoschema/main.go

test:
	go test ./...

run:
	go build -o ${BINARY_NAME} cmd/envtoschema/main.go
	./${BINARY_NAME} $(ARGS)

clean:
	go clean
	rm ${BINARY_NAME}