use consistent_hash::consistent_hash::find_next_3_node;
// use kv_service::pb::key_value_service_client::KeyValueServiceClient;
use consistent_hash::RequestType;
use consistent_hash::gossip::ServerInfo;
use consistent_hash::pb::{
    key_value_service_client::KeyValueServiceClient,
    service_discovery_client::ServiceDiscoveryClient,
};
use std::env;
use tonic::Request;

// #[tonic::async_trait]
// impl

// async fn get_responsible_server(proxy_address: &str, key: &str) -> Result<String, Box<dyn std::error::Error>> {
//     let mut client = ServiceDiscoveryClient::connect(format!("http://{}", proxy_address)).await?;

//     let request = Request::new(GetServerRequest {
//         key: key.to_string(),
//     });

//     let response = client.get_server(request).await?;
//     let server_info = response.into_inner().server.ok_or("No server found")?;

//     Ok(server_info.address)
// }

async fn run_quorum_based_request(
    server_info: &Vec<ServerInfo>,
    key: &String,
    value: Option<String>,
) -> Result<(), Box<dyn std::error::Error>> {
    let mut success_count = 0;
    let nodes = find_next_3_node(server_info, key);
    for node in nodes {
        let mut client = KeyValueServiceClient::connect(format!("http://{}", node)).await?;
        let request = match value {
            Some(v) => client.create(key.clone(), v).await,
            None => client.read(key.clone()).await,
        };
        match request {
            Ok(response) => {
                success_count += 1;
                println!("Response from {}: {:?}", node, response);
            }
            Err(e) => {
                eprintln!("Error from {}: {:?}", node, e);
            }
        }
    }
    if success_count < 2 {
        eprintln!("Quorum not reached. Only {} out of 3 nodes responded successfully.", success_count);
        return Err("Quorum not reached".into());
    }
    println!("Quorum reached with {} successful responses.", success_count);
    return Ok(());
}

async fn query_server_info(
    proxy_address: &String,
) -> Result<Vec<ServerInfo>, Box<dyn std::error::Error>> {
    let mut client = KeyValueServiceClient::connect(format!("http://{}", proxy_address)).await?;
    let server_info = client.get_server_info().await?;
    Ok(server_info)
}

async fn request(
    request_type: &RequestType,
    proxy_address: &String,
    key: &String,
    value: Option<String>,
) -> Result<(), Box<dyn std::error::Error>> {
    let server_info = query_server_info(proxy_address).await?;
    run_quorum_based_request(&server_info, key, value).await?;
    Ok(())
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args: Vec<String> = env::args().collect();
    if args.len() < 4 {
        eprintln!("Usage: client <proxy_address> <operation> <key> [value]");
        eprintln!("Operations: create(c), read(r), update(u), delete(d)");
        eprintln!("Example: client 127.0.0.1:50051 create mykey myvalue");
        return Ok(());
    }

    let proxy_address = &args[1];
    let operation = &args[2];
    let key = &args[3];
    let value = args.get(4).cloned().unwrap_or_default();

    let request_type = match operation.as_str() {
        "create" | "c" => RequestType::Create,
        "read" | "r" => RequestType::Read,
        "update" | "u" => RequestType::Update,
        "delete" | "d" => RequestType::Delete,
        _ => {
            eprintln!("Invalid operation: {}", operation);
            return Ok(());
        }
    };

    request(&request_type, proxy_address, key, Some(value)).await?;

    Ok(())
}
