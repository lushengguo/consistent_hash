pub mod protocol;
pub mod consistent_hash;

// 引入生成的 protobuf 代码
pub mod pb {
    tonic::include_proto!("key_value_service");
    tonic::include_proto!("service_discovery");
}