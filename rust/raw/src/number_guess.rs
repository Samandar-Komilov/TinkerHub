use std::io::{self, Read};

pub fn number_guess(user_choice: i32) -> () {
    println!("Your guess: {user_choice}");
}

pub fn number_guess_main() -> () {
    let mut input: i32;

    io::stdin().(&mut input).expect("Failed to read the line.");

}