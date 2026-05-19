use std::env;

pub fn get_env(key: &str, default: Option<&str>) -> String {
    env::var(key).unwrap_or_else(|_| {
        default
            .map(String::from)
            .unwrap_or_else(|| panic!("env var {key} not set"))
    })
}
