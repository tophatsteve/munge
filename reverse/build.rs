extern crate protoc_rust_grpc;

// generate grpc stubs
fn main() {
    protoc_rust_grpc::run(protoc_rust_grpc::Args {
        out_dir: "src",
        includes: &[],
        input: &["reverse.proto"],
        rust_protobuf: true,
    }).expect("protoc-rust-grpc");
}
