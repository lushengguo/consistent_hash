pub mod consistent_hash;
pub mod gossip;
pub mod protocol;

use serde::Deserialize;

#[derive(Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum RequestType {
    Create = 0,
    Read = 1,
    Update = 2,
    Delete = 3,
}

pub mod pb {
    tonic::include_proto!("kv_service");
    tonic::include_proto!("consistent_hash");
}
