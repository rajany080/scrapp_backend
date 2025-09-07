mod routes;
mod handlers;
mod schemas;

#[tokio::main]
async fn main() {
    let app = routes::create_router();


    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    println!("Server running on localhost:3000");

    axum::serve(listener, app)
        .await
        .unwrap();
}
