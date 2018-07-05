extern crate futures;
extern crate futures_cpupool;

extern crate grpc;
extern crate protobuf;
extern crate reverse;

use std::env;
use std::thread;

use reverse::reverse::*;
use reverse::reverse_grpc::*;

#[macro_use]
extern crate log;
extern crate env_logger;

struct ReverseServiceImpl;

impl Reverse for ReverseServiceImpl {
    fn reverse(
        &self,
        _m: grpc::RequestOptions,
        req: ReverseRequest,
    ) -> grpc::SingleResponse<ReverseResponse> {
        info!("received - {}", req.text);
        let mut r = ReverseResponse::new();
        r.text = reverse_text(req.text);
        info!("responding with - {}", r.text);
        grpc::SingleResponse::completed(r)
    }
}

fn reverse_text(text: String) -> String {
    text.chars().rev().collect::<String>()
}

fn main() {
    env_logger::init();

    #[allow(unused_assignments)]
    let mut port: u16 = 80;

    match env::var("PORT") {
        Ok(val) => port = val.parse::<u16>().unwrap(),
        Err(_e) => port = 80,
    }

    info!("starting service on port {}", port);

    let mut server = grpc::ServerBuilder::new_plain();
    server.http.set_port(port);
    server.add_service(ReverseServer::new_service_def(ReverseServiceImpl));
    server.http.set_cpu_pool_threads(4);
    let _server = server.build().expect("server");

    loop {
        thread::park();
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_reverse() {
        let reversed_text = reverse_text(String::from("test"));
        assert_eq!(reversed_text, String::from("tset"));
    }
}
