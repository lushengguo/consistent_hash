fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::compile_protos("proto/key_value_service.proto")?;
    tonic_build::compile_protos("proto/service_discovery.proto")?;
    Ok(())
}
