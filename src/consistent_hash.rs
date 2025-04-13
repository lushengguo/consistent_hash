use std::collections::BTreeMap;
use std::hash::{Hash, Hasher};
use std::sync::RwLock;

pub struct ConsistentHash {
    ring: RwLock<BTreeMap<u64, String>>,
    replicas: usize,
}

impl ConsistentHash {
    pub fn new(replicas: usize) -> Self {
        Self {
            ring: RwLock::new(BTreeMap::new()),
            replicas,
        }
    }

    pub fn add_node(&self, node: &str) {
        let mut ring = self.ring.write().unwrap();
        for i in 0..self.replicas {
            let hash = Self::hash(&format!("{}-{}", node, i));
            ring.insert(hash, node.to_string());
        }
    }

    pub fn remove_node(&self, node: &str) {
        let mut ring = self.ring.write().unwrap();
        for i in 0..self.replicas {
            let hash = Self::hash(&format!("{}-{}", node, i));
            ring.remove(&hash);
        }
    }

    pub fn get_node(&self, key: &str) -> Option<String> {
        let ring = self.ring.read().unwrap();
        if ring.is_empty() {
            return None;
        }
        let hash = Self::hash(&key);
        let mut keys = ring.keys();
        keys.find(|&&k| k >= hash).or_else(|| keys.next()).map(|&k| ring[&k].clone())
    }

    fn hash<T: Hash>(t: &T) -> u64 {
        let mut hasher = std::collections::hash_map::DefaultHasher::new();
        t.hash(&mut hasher);
        hasher.finish()
    }
}