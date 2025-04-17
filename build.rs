fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::compile_protos("proto/kv_service.proto")?;
    tonic_build::compile_protos("proto/consistent_hash.proto")?;
    Ok(())
}
