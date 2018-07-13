DEPENDENCIES := \
	github.com/golang/protobuf/proto \
	golang.org/x/net/context \
	google.golang.org/grpc \
	github.com/julienschmidt/httprouter

.PHONY: build clean godeps

build:
	godeps
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./frontend/frontend ./frontend/main.go
	# docker run --rm -it -v  $(pwd)/reverse/:/home/rust/src ekidd/rust-musl-builder cargo build --release

# clean:
# 	rm ./frontend/frontend
# 	rm -rf ./reverse/target/x86_64-unknown-linux-musl
	
godeps: 
	go get $(DEPENDENCIES)