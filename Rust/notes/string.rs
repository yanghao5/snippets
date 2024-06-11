fn main() {
    let mut str=String::from("Bob");
    str.push_str("hello");
    str.push('c');
    str.insert_str(3,"hils");
    str.insert(5,'h');
    println!("{}",str);

    // replace()
    let str2=String::from("He He He");
    let new_str2=str2.replace("He","fuck");
    dbg!(new_str2);

    //replacen()
    let str3=String::from("Yes Yes Yes");
    let new_str3=str3.replacen("Yes","No",2);
    dbg!(new_str3);

    //replace_range()
    let mut str4=String::from("HHHHHHHHHHHHHHHHHHHHHHH");
    let new_str4=str4.replace_range(3..8,"Fuck");
    dbg!(new_str4);

    // pop()
    let mut str_pop=String::from("rust pop yx");
    let p1=str_pop.pop();
    let p2=str_pop.pop();
    dbg!(p1);
    dbg!(p2);

    // remove()
    let mut str_remove=String::from("你妈了个逼");
    println!(
        "string_remove uses {} bytes",
        std::mem::size_of_val(str_remove.as_str())
    );
    //str_remove.remove(0); // delete first kanji
    //str_remove.remove(1);//error, can't delete second kanji
    str_remove.remove(3);// delete second kanji
    dbg!(str_remove);

    // truncate()
    let mut str_truncate=String::from("风景独好");
    str_truncate.truncate(3);
    dbg!(str_truncate);

    //clear()
    let mut str_clear=String::from("fuck edge");
    str_clear.clear();
    dbg!(str_clear);

    //concatenote
    let str_con1=String::from("fuck ");
    let str_con2=String::from("HR ");
    // str_con1 和 str_con2 所有权被移走
    let result_str=str_con1+&str_con2;
    dbg!(&result_str);

    let mut result_str2= result_str + "and Boss";
    result_str2+="!!!";
    dbg!(result_str2);
}
