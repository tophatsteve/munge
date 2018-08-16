DEPENDENCIES := \
	github.com/golang/protobuf/proto \
	golang.org/x/net/context \
	google.golang.org/grpc \
	github.com/julienschmidt/httprouter

.PHONY: build clean godeps all

godeps: 
	go get $(DEPENDENCIES)

build: godeps
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./frontend/frontend ./frontend/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./reverse/cmd/reverse ./reverse/cmd/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./capitalise/cmd/capitalise ./capitalise/cmd/main.go

# clean:
# 	rm ./frontend/frontend
