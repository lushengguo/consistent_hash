use std::collections::HashSet;

pub struct ServerInfo {
    pub server_id: String,
    pub address: String,
    heartbeat_timestamp: u64,
    vnode: HashSet<String>,
    is_slave_of: Option<String>,
}

pub struct Gossip {
    server_info: Vec<ServerInfo>,
}

impl Gossip {
    pub fn new() -> Self {
        Self {
            server_info: Vec::new(),
        }
    }

    pub fn merge(gossip: &Gossip) {}
}
