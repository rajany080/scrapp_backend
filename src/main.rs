use axum::{
    routing::{get, post},
    Router,
    extract::Path,
    Json,
};
use serde::Serialize;
use utoipa::{OpenApi, ToSchema};
use utoipa_swagger_ui::SwaggerUi;

/// Root endpoint - just returns your name
#[utoipa::path(
    get,
    path = "/",
    responses(
        (status = 200, description = "Returns my name as plain text")
    )
)]
async fn root() -> &'static str {
    "My name is Rajan Yadav"
}

/// Just a test hello
#[utoipa::path(
    get,
    path = "/hello",
    responses(
        (status = 200, description = "Returns hello world string")
    )
)]
async fn handle_hello() -> &'static str {
    "Hello World"
}

#[derive(Serialize, ToSchema)]
struct GreetResponse {
    message: String,
}

/// Greet a user by name
#[utoipa::path(
    post,
    path = "/hello/{name}",
    params(
        ("name" = String, Path, description = "Name of the user to greet")
    ),
    responses(
        (status = 200, description = "Successful Greeting", body = GreetResponse)
    )
)]
async fn greet_user(Path(name): Path<String>) -> Json<GreetResponse> {
    Json(GreetResponse {
        message: format!("Hello {}", name),
    })
}

#[derive(OpenApi)]
#[openapi(
    paths(
        root,
        handle_hello,
        greet_user
    ),
    components(schemas(GreetResponse))
)]
struct ApiDoc;

#[tokio::main]
async fn main() {
    let app: Router = Router::new()
        .route("/", get(root))
        .route("/hello", get(handle_hello))
        .route("/hello/{name}", post(greet_user))
        // swagger ui mount
        .merge(SwaggerUi::new("/docs").url("/api-doc/openapi.json", ApiDoc::openapi()));

    let addr = "0.0.0.0:3000";
    let listener = tokio::net::TcpListener::bind(addr).await.unwrap();

    println!("Server running at http://{}", addr);
    axum::serve(listener, app).await.unwrap();
}
