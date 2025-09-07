use axum::{Router, routing::{get, post}};
use crate::handlers::hello;

pub fn hello_router() -> Router {
    Router::new()
        .route("/hello/{name}", get(hello::handle_hello))
        .route("/hello/{name}", post(hello::greet_user))
}
