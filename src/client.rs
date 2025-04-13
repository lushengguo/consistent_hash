use tokio::net::TcpStream;
use tokio::io::{AsyncReadExt, AsyncWriteExt};
use consistent_hash::protocol::Protocol;
use serde_json;

#[tokio::main]
async fn main() {
    let mut args = std::env::args();
    args.next(); // Skip the executable name

    let key = args.next().expect("Key is required");
    let value = args.next().unwrap_or_default();

    let servers = vec!["127.0.0.1:8080", "127.0.0.1:8081", "127.0.0.1:8082"];

    for server in &servers {
        if let Ok(mut stream) = TcpStream::connect(server).await {
            let protocol = Protocol {
                key: key.clone(),
                value: value.clone(),
                error: None,
            };

            let request = serde_json::to_vec(&protocol).unwrap();
            stream.write_all(&request).await.unwrap();

            let mut buffer = vec![0; 1024];
            let n = stream.read(&mut buffer).await.unwrap();

            let response: Protocol = serde_json::from_slice(&buffer[..n]).unwrap();

            if let Some(error) = response.error {
                eprintln!("Server {}: Error: {}", server, error);
            } else {
                println!("Server {}: Key: {}, Value: {}", server, response.key, response.value);
                return;
            }
        }
    }

    eprintln!("Failed to connect to any server.");
}