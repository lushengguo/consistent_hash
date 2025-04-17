use std::collections::BTreeMap;
use std::hash::{Hash, Hasher};
use std::sync::Arc;
use tokio::sync::RwLock;

#[derive(Debug, Clone)]
pub struct ConsistentHash {
    ring: Arc<RwLock<BTreeMap<u64, String>>>,
    replicas: usize,
}

impl ConsistentHash {
    pub fn new(replicas: usize) -> Self {
        Self {
            ring: Arc::new(RwLock::new(BTreeMap::new())),
            replicas,
        }
    }

    pub async fn add_node(&self, node: &str) {
        let mut ring = self.ring.write().await;
        for i in 0..self.replicas {
            let hash = Self::hash(&format!("{}-{}", node, i));
            ring.insert(hash, node.to_string());
        }
    }

    pub async fn remove_node(&self, node: &str) {
        let mut ring = self.ring.write().await;
        for i in 0..self.replicas {
            let hash = Self::hash(&format!("{}-{}", node, i));
            ring.remove(&hash);
        }
    }

    pub async fn get_node(&self, key: &str) -> Option<String> {
        let ring = self.ring.read().await;
        if ring.is_empty() {
            return None;
        }
        let hash = Self::hash(&key);
        let mut keys = ring.range(hash..);
        if let Some((&k, v)) = keys.next() {
            return Some(v.clone());
        }
        // Wrap around to the first key if necessary
        if let Some((&_, v)) = ring.iter().next() {
            return Some(v.clone());
        }
        None
    }

    pub async fn get_all_nodes(&self) -> Vec<String> {
        let ring = self.ring.read().await;
        let mut nodes = Vec::new();
        let mut seen = std::collections::HashSet::new();
        
        for (_, node) in ring.iter() {
            if !seen.contains(node) {
                nodes.push(node.clone());
                seen.insert(node);
            }
        }
        
        nodes
    }

    fn hash<T: Hash>(t: &T) -> u64 {
        let mut hasher = std::collections::hash_map::DefaultHasher::new();
        t.hash(&mut hasher);
        hasher.finish()
    }
}