mod temp_convert;
mod number_guess;
use std::io;

fn main(){
    let mut input: String = String::new();
    io::stdin()
        .read_line(&mut input)
        .expect("Failed to read line.");

    let inputs_lst: Vec<&str> = input.split(' ').collect();
    let val: f64 = inputs_lst[0].parse().unwrap();
    let to_scale: &str = inputs_lst[1];

    let res: f64 = temp_convert::temp_convert(val, to_scale);

    println!("Result: {res}");
}