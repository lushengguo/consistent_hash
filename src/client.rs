use consistent_hash::pb::{
    key_value_service_client::KeyValueServiceClient,
    service_discovery_client::ServiceDiscoveryClient,
    KeyValue, Key, GetServerRequest,
};
use std::env;
use tonic::Request;

async fn get_responsible_server(proxy_addr: &str, key: &str) -> Result<String, Box<dyn std::error::Error>> {
    let mut client = ServiceDiscoveryClient::connect(format!("http://{}", proxy_addr)).await?;
    
    let request = Request::new(GetServerRequest {
        key: key.to_string(),
    });
    
    let response = client.get_server(request).await?;
    let server_info = response.into_inner().server.ok_or("No server found")?;
    
    Ok(server_info.address)
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args: Vec<String> = env::args().collect();
    if args.len() < 4 {
        eprintln!("Usage: client <proxy_address> <operation> <key> [value]");
        eprintln!("Operations: create, read, update, delete");
        eprintln!("Example: client 127.0.0.1:50051 create mykey myvalue");
        return Ok(());
    }
    
    let proxy_addr = &args[1];
    let operation = &args[2];
    let key = &args[3];
    let value = args.get(4).cloned().unwrap_or_default();
    
    println!("使用代理 {} 执行 {} 操作，键: {}, 值: {}", proxy_addr, operation, key, value);
    
    // 通过代理获取负责该键的服务器地址
    let server_addr = match get_responsible_server(proxy_addr, key).await {
        Ok(addr) => addr,
        Err(e) => {
            eprintln!("获取服务器地址失败: {:?}", e);
            return Err(e);
        }
    };
    
    println!("键 {} 由服务器 {} 负责", key, server_addr);
    
    // 连接到负责该键的服务器
    let mut client = KeyValueServiceClient::connect(format!("http://{}", server_addr)).await?;
    
    match operation.to_lowercase().as_str() {
        "create" => {
            let request = Request::new(KeyValue {
                key: key.to_string(),
                value: value.to_string(),
            });
            
            let response = client.create(request).await?;
            println!("创建响应: {:?}", response);
        }
        "read" => {
            let request = Request::new(Key {
                key: key.to_string(),
            });
            
            match client.read(request).await {
                Ok(response) => {
                    let kv = response.into_inner();
                    println!("键: {}, 值: {}", kv.key, kv.value);
                }
                Err(status) => {
                    eprintln!("读取失败: {}", status);
                }
            }
        }
        "update" => {
            let request = Request::new(KeyValue {
                key: key.to_string(),
                value: value.to_string(),
            });
            
            let response = client.update(request).await?;
            println!("更新响应: {:?}", response);
        }
        "delete" => {
            let request = Request::new(Key {
                key: key.to_string(),
            });
            
            let response = client.delete(request).await?;
            println!("删除响应: {:?}", response);
        }
        _ => {
            eprintln!("不支持的操作: {}", operation);
            eprintln!("支持的操作: create, read, update, delete");
        }
    }
    
    Ok(())
}
