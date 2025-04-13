use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
pub struct Protocol {
    pub key: String,
    pub value: String,
    pub error: Option<String>,
}
