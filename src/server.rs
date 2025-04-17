use chrono::Utc;
use consistent_hash::consistent_hash::ConsistentHash;
use consistent_hash::pb::{
    Key, KeyValue, RegisterRequest, Response, ResponseType, ServerInfo,
    key_value_service_server::{KeyValueService, KeyValueServiceServer},
    service_discovery_client::ServiceDiscoveryClient,
};
use consistent_hash::protocol::{KvOp, KvRequest};
use serde_json;
use std::collections::HashMap;
use std::env;
use std::net::SocketAddr;
use std::sync::Arc;
use tokio::sync::RwLock;
use tonic::{Request, Response as TonicResponse, Status, transport::Server};
use uuid::Uuid;

#[derive(Debug)]
struct ServerState {
    kv_store: RwLock<HashMap<String, String>>,
    server_id: String,
    address: String,
}

#[derive(Debug)]
struct KeyValueServiceImpl {
    state: Arc<ServerState>,
}

#[tonic::async_trait]
impl KeyValueService for KeyValueServiceImpl {
    async fn create(&self, request: Request<KeyValue>) -> Result<TonicResponse<Response>, Status> {
        let kv = request.into_inner();
        let mut store = self.state.kv_store.write().await;

        let already_exists = store.contains_key(&kv.key);
        store.insert(kv.key.clone(), kv.value.clone());

        println!("创建键值对: {} = {}", kv.key, kv.value);

        Ok(TonicResponse::new(Response {
            kv: Some(kv),
            response_type: if already_exists {
                ResponseType::KeyNotExist as i32
            } else {
                ResponseType::Success as i32
            },
        }))
    }

    async fn read(&self, request: Request<Key>) -> Result<TonicResponse<KeyValue>, Status> {
        let key = request.into_inner().key;
        let store = self.state.kv_store.read().await;

        match store.get(&key) {
            Some(value) => {
                println!("读取键值对: {} = {}", key, value);

                Ok(TonicResponse::new(KeyValue {
                    key,
                    value: value.clone(),
                }))
            }
            None => Err(Status::not_found(format!("Key not found: {}", key))),
        }
    }

    async fn update(&self, request: Request<KeyValue>) -> Result<TonicResponse<Response>, Status> {
        let kv = request.into_inner();
        let mut store = self.state.kv_store.write().await;

        let existed = store.contains_key(&kv.key);
        if existed {
            store.insert(kv.key.clone(), kv.value.clone());
            println!("更新键值对: {} = {}", kv.key, kv.value);

            Ok(TonicResponse::new(Response {
                kv: Some(kv),
                response_type: ResponseType::Success as i32,
            }))
        } else {
            println!("更新失败，键不存在: {}", kv.key);

            Ok(TonicResponse::new(Response {
                kv: Some(kv),
                response_type: ResponseType::KeyNotExist as i32,
            }))
        }
    }

    async fn delete(&self, request: Request<Key>) -> Result<TonicResponse<Response>, Status> {
        let key = request.into_inner().key;
        let mut store = self.state.kv_store.write().await;

        let removed = store.remove(&key);

        if removed.is_some() {
            println!("删除键值对: {}", key);

            Ok(TonicResponse::new(Response {
                kv: Some(KeyValue {
                    key: key.clone(),
                    value: "".to_string(),
                }),
                response_type: ResponseType::Success as i32,
            }))
        } else {
            println!("删除失败，键不存在: {}", key);

            Ok(TonicResponse::new(Response {
                kv: Some(KeyValue {
                    key: key.clone(),
                    value: "".to_string(),
                }),
                response_type: ResponseType::KeyNotExist as i32,
            }))
        }
    }
}

async fn register_with_proxy(
    proxy_addr: &str,
    server_id: &str,
    server_addr: &str,
) -> Result<(), Box<dyn std::error::Error>> {
    let mut client = ServiceDiscoveryClient::connect(format!("http://{}", proxy_addr)).await?;

    let request = tonic::Request::new(RegisterRequest {
        server: Some(ServerInfo {
            server_id: server_id.to_string(),
            address: server_addr.to_string(),
            weight: 1,
        }),
    });

    let response = client.register(request).await?;

    println!("注册到代理: {:?}", response.into_inner());

    Ok(())
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args: Vec<String> = env::args().collect();
    if args.len() < 3 {
        eprintln!("Usage: server <server_address> <proxy_address>");
        eprintln!("Example: server 127.0.0.1:50052 127.0.0.1:50051");
        return Ok(());
    }

    let server_addr = &args[1];
    let proxy_addr = &args[2];

    let server_id = Uuid::new_v4().to_string();

    let state = Arc::new(ServerState {
        kv_store: RwLock::new(HashMap::new()),
        server_id: server_id.clone(),
        address: server_addr.to_string(),
    });

    println!("服务器启动于: {}, ID: {}", server_addr, server_id);

    register_with_proxy(proxy_addr, &server_id, server_addr).await?;

    let service = KeyValueServiceImpl { state };
    let addr: SocketAddr = server_addr.parse()?;

    Server::builder()
        .add_service(KeyValueServiceServer::new(service))
        .serve(addr)
        .await?;

    Ok(())
}
