//1-1
fn longest<'a>(x: &'a String, y: &'a String) -> &'a String {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn main() {
    let string1 = &"abdcdwe".to_string();
    let string2 = &"xyz".to_string();

    let result = longest(string1, string2);
    println!("The longest string is {}", result);
}
