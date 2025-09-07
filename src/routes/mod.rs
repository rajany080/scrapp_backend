use axum::Router;

mod hello;

pub fn create_router() -> Router {
    Router::new()
        .merge(hello::hello_router())
        .route("/", axum::routing::get(|| async { "My name is Rajan Yadav" }))
}
