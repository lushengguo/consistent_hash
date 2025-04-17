pub mod protocol;
pub mod consistent_hash;

pub mod pb {
    tonic::include_proto!("kv_service");
    tonic::include_proto!("consistent_hash");
}