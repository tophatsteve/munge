DEPENDENCIES := \
	github.com/golang/protobuf/proto \
	golang.org/x/net/context \
	google.golang.org/grpc \
	github.com/julienschmidt/httprouter

.PHONY: build clean godeps containerize

godeps: 
	go get $(DEPENDENCIES)

build: godeps
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./frontend/frontend ./frontend/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./reverse/cmd/reverse ./reverse/cmd/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./capitalise/cmd/capitalise ./capitalise/cmd/main.go

containerize: build
	docker build -t tophatsteve/frontend:latest ./frontend
	docker build -t tophatsteve/reverse:latest ./reverse/cmd
	docker build -t tophatsteve/capitalise:latest ./capitalise/cmd

clean:
	rm -f ./frontend/frontend
	rm -f ./reverse/cmd/reverse
	rm -f ./capitalise/cmd/capitalise
