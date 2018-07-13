all: clean build

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./frontend/frontend ./frontend/main.go
	docker run --rm -it -v  $(pwd)/reverse/:/home/rust/src ekidd/rust-musl-builder cargo build --release

clean:
	rm ./frontend/frontend
	rm -rf ./reverse/target/x86_64-unknown-linux-musl
	