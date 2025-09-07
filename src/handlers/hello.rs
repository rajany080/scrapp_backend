use axum::extract::Path;

pub async fn handle_hello(Path(name): Path<String>) -> String {
    return format!("Hello {}", name);
}

pub async fn greet_user(Path(name): Path<String>) -> String {
    format!("Hello {}", name)
}
