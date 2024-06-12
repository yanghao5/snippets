// 1-1
fn largest<T:PartialOrd+Copy>(list: &[T])->T{
    let mut largest=list[0];
    for &li in list.iter() {
        if li>largest {
            largest=li;
        }
    }
    largest
}
fn main() {
    let number_list=vec![34,50,25,100,65];
    let largest_number=largest(&number_list);
    let char_list=vec!['y','m','a','q'];
    let largest_char=largest(&char_list);
    println!("{} {}",largest_number,largest_char);
}
