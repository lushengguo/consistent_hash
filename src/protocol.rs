use serde::{Deserialize, Serialize};
use std::collections::HashMap;

#[derive(Serialize, Deserialize, Clone)]
pub enum KvOp {
    Get,
    Set,
    Delete,
}

#[derive(Serialize, Deserialize)]
pub struct KvRequest {
    pub op: KvOp,
    pub key: String,
    pub value: String,
    pub timestamp: Option<String>,
}

#[derive(Serialize, Deserialize)]
pub struct KvResponse {
    pub value: Option<String>,
    pub error: Option<String>,
}

#[derive(Serialize, Deserialize, Clone)]
pub enum ServiceOp {
    Request,
    Register,
    Heartbeat,
    Shutdown,
}

#[derive(Serialize, Deserialize)]
pub struct ServiceConfigurationRequest {
    pub op: ServiceOp,
    pub id: Option<String>,
    pub timestamp: Option<String>,
}

pub struct ServerNodeConfiguration {
    pub vhost: Vec<String>,
    pub endpoint: String,
}

pub struct ServiceResponse {
    pub heartbeat_ok: Option<bool>,
    pub servers: HashMap<String, ServerNodeConfiguration>,
}
