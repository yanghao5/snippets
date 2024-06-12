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

//3-1
#[derive(Debug)]
struct Point<T,U>{
    x:T,
    y:U,
}
fn main() {
    let u32_str=Point{ x:5,y:"wtf?" };
    let str_str=Point{x:"Bob",y:"Alice"};
    println!("{}",u32_str.return_x());
    println!("{}",str_str.return_y());
    dbg!(u32_str.mixup(str_str));
}

impl<T,U> Point<T,U> {
    fn return_x(&self)->&T{
        &self.x
    }
    fn return_y(&self)->&U{
        &self.y
    }
    fn mixup<V,W>(self,other:Point<V,W>)->Point<T,W>{
        Point{
            x:self.x,
            y:other.y,
        }
    }
}

//3-2
use std::fmt::Debug;
// print any length arrary
fn display_arr<T: Copy+Debug,const N:usize>(arr:[T;N]){
    println!("{:?}",arr)
}
fn main() {
    //let arr:[String;3]=[String::from("aaaaaaaaaaaaaaaaa"),String::from("bbbbbbbbbbbbbbbbbbbbb"),String::from("cccccccccccc")];
    let arr=[1,2,3,4,5,6];
    display_arr(arr);
    println!("{:?}",arr);
}
