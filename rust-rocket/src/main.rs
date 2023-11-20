use rocket::serde::json::{json, Value};

#[macro_use]
extern crate rocket;

#[get("/")]
fn index() -> Value {
    json!({"hello": "world"})
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![index])
}
