use consistent_hash::consistent_hash::ConsistentHash;
use consistent_hash::pb::{
    DeregisterRequest, DeregisterResponse, GetAllServersRequest, GetServerRequest,
    GetServerResponse, RegisterRequest, RegisterResponse, ServerInfo, ServerList,
    key_value_service_client::KeyValueServiceClient,
    service_discovery_client::ServiceDiscoveryClient,
    service_discovery_server::{ServiceDiscovery, ServiceDiscoveryServer},
};
use std::collections::HashMap;
use std::sync::Arc;
use tokio::sync::RwLock;
use tonic::{Request, Response, Status, transport::Server};
use uuid::Uuid;

#[derive(Debug)]
struct ProxyState {
    hash_ring: ConsistentHash,
    servers: RwLock<HashMap<String, ServerInfo>>,
}

// 代理服务实现
#[derive(Debug)]
struct ProxyService {
    state: Arc<ProxyState>,
}

#[tonic::async_trait]
impl ServiceDiscovery for ProxyService {
    async fn register(
        &self,
        request: Request<RegisterRequest>,
    ) -> Result<Response<RegisterResponse>, Status> {
        let server_info = request.into_inner().server;

        // if server_info.server_id.is_empty() || server_info.address.is_empty() {
        //     return Ok(Response::new(RegisterResponse {
        //         success: false,
        //         message: "Server ID and address are required".to_string(),
        //     }));
        // }

        // // 添加服务器到哈希环
        // self.state.hash_ring.add_node(&server_info.address).await;

        // // 将服务器信息保存到内存中
        // {
        //     let mut servers = self.state.servers.write().await;
        //     servers.insert(server_info.server_id.clone(), server_info);
        // }

        // println!("Server registered: {}", &server_info.address);

        Ok(Response::new(RegisterResponse {
            success: true,
            message: "Server registered successfully".to_string(),
        }))
    }

    async fn deregister(
        &self,
        request: Request<DeregisterRequest>,
    ) -> Result<Response<DeregisterResponse>, Status> {
        let server_id = request.into_inner().server_id;

        let address = {
            let servers = self.state.servers.read().await;
            match servers.get(&server_id) {
                Some(server) => server.address.clone(),
                None => {
                    return Ok(Response::new(DeregisterResponse {
                        success: false,
                        message: "Server not found".to_string(),
                    }));
                }
            }
        };

        // 从哈希环中移除服务器
        self.state.hash_ring.remove_node(&address).await;

        // 从内存中删除服务器信息
        {
            let mut servers = self.state.servers.write().await;
            servers.remove(&server_id);
        }

        println!("Server deregistered: {}", &address);

        Ok(Response::new(DeregisterResponse {
            success: true,
            message: "Server deregistered successfully".to_string(),
        }))
    }

    async fn get_server(
        &self,
        request: Request<GetServerRequest>,
    ) -> Result<Response<GetServerResponse>, Status> {
        let key = request.into_inner().key;

        // 使用一致性哈希获取负责此键的服务器
        let server_address = match self.state.hash_ring.get_node(&key).await {
            Some(address) => address,
            None => {
                return Err(Status::not_found("No servers available"));
            }
        };

        // 查找对应的服务器信息
        let servers = self.state.servers.read().await;
        for server in servers.values() {
            if server.address == server_address {
                return Ok(Response::new(GetServerResponse {
                    server: Some(server.clone()),
                }));
            }
        }

        Err(Status::internal(
            "Server found in hash ring but not in registry",
        ))
    }

    async fn get_all_servers(
        &self,
        _request: Request<GetAllServersRequest>,
    ) -> Result<Response<ServerList>, Status> {
        let servers = self.state.servers.read().await;

        let server_list = ServerList {
            servers: servers.values().cloned().collect(),
        };

        Ok(Response::new(server_list))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "0.0.0.0:50051".parse()?;
    println!("Proxy server listening on {}", addr);

    let state = Arc::new(ProxyState {
        hash_ring: ConsistentHash::new(10), // 每个服务器有10个虚拟节点
        servers: RwLock::new(HashMap::new()),
    });

    let proxy_service = ProxyService {
        state: state.clone(),
    };

    let service = ServiceDiscoveryServer::new(proxy_service);

    Server::builder().add_service(service).serve(addr).await?;

    Ok(())
}
