use tokio::net::TcpListener;
use tokio::io::{AsyncReadExt, AsyncWriteExt};
use consistent_hash::consistent_hash::ConsistentHash;
use consistent_hash::protocol::Protocol;
use serde_json;
use std::collections::HashMap;
use std::sync::Arc;
use tokio::sync::RwLock;
use std::fs::File;
use std::io::Write;

pub struct ServerState {
    pub kv_store: Arc<RwLock<HashMap<String, String>>>,
    pub hash_ring: Arc<ConsistentHash>,
}

impl ServerState {
    pub fn new(replicas: usize) -> Self {
        Self {
            kv_store: Arc::new(RwLock::new(HashMap::new())),
            hash_ring: Arc::new(ConsistentHash::new(replicas)),
        }
    }

    pub async fn add_node(&self, node: &str) {
        self.hash_ring.add_node(node);
        self.rebalance_data().await;
    }

    pub async fn remove_node(&self, node: &str) {
        self.hash_ring.remove_node(node);
        self.rebalance_data().await;
    }

    async fn rebalance_data(&self) {
        let mut kv_store = self.kv_store.write().await;
        let mut new_store = HashMap::new();

        for (key, value) in kv_store.drain() {
            if let Some(node) = self.hash_ring.get_node(&key) {
                if node == "127.0.0.1:8080" {
                    new_store.insert(key, value);
                } else {
                    // Simulate transferring data to the correct node
                    println!("Transferring key {} to node {}", key, node);
                }
            }
        }

        *kv_store = new_store;
    }
}

pub async fn log_data_distribution(server_state: Arc<ServerState>, log_file: &str) {
    let kv_store = server_state.kv_store.read().await;
    let mut file = File::create(log_file).unwrap();

    writeln!(file, "Data Distribution:").unwrap();
    for (key, value) in kv_store.iter() {
        writeln!(file, "Key: {}, Value: {}", key, value).unwrap();
    }
}

pub async fn log_data_migration(key: &str, from_node: &str, to_node: &str, log_file: &str) {
    let mut file = File::options().append(true).open(log_file).unwrap();
    writeln!(file, "Migrating Key: {} from {} to {}", key, from_node, to_node).unwrap();
}

#[tokio::main]
async fn main() {
    let listener = TcpListener::bind("127.0.0.1:8080").await.unwrap();
    let server_state = Arc::new(ServerState::new(3));

    println!("Server running on 127.0.0.1:8080");

    loop {
        let (mut socket, _) = listener.accept().await.unwrap();
        let server_state = Arc::clone(&server_state);

        tokio::spawn(async move {
            let mut buffer = vec![0; 1024];
            let n = socket.read(&mut buffer).await.unwrap();

            if n == 0 {
                return;
            }

            let request: Protocol = serde_json::from_slice(&buffer[..n]).unwrap();
            let response = match server_state.hash_ring.get_node(&request.key) {
                Some(node) => {
                    if node == "127.0.0.1:8080" {
                        let mut store = server_state.kv_store.write().await;
                        store.insert(request.key.clone(), request.value.clone());
                        Protocol {
                            key: request.key,
                            value: request.value,
                            error: None,
                        }
                    } else {
                        Protocol {
                            key: request.key,
                            value: String::new(),
                            error: Some(format!("Key belongs to node: {}", node)),
                        }
                    }
                }
                None => Protocol {
                    key: request.key,
                    value: String::new(),
                    error: Some("No nodes available".to_string()),
                },
            };

            let response = serde_json::to_vec(&response).unwrap();
            socket.write_all(&response).await.unwrap();
        });
    }
}
