use day_3::{find_repeat_item, find_share_item};
use std::env::var;
use std::fs;

fn main() {
    let project_dir = var("PROJECT_ROOT").unwrap();
    let path = format!("{}/days/day-3/input.txt", project_dir);
    let file = fs::read(path).unwrap();
    let file = String::from_utf8(file).unwrap();
    println!("Response(part-1): {}", find_repeat_item(&file));
    println!("Response(part-2): {}", find_share_item(&file));
}
