/* 
    Exercises can be found: 
*/


// pub const MAX_REQUESTS: u32 = 1000;
// pub const TIMEOUT_SECONDS: u8 = 10;

pub fn common_main(){
    // 1
    // ex1();
    // 2
    // let res = ex2_unit_convert();
    // println!("Result: {}", res);

    // 3
    // let arr: [u32; 5] = [6,8,1,4,8];
    // println!("{:?}", ex3_get_stats(arr));

    // 4
    // let mut password = String::new();
    // io::stdin().read_line(&mut password).expect("Failed to read input");
    // ex4_password_checker(&password);

    // 5
    // ex5_simple_menu();

    // 6
    // ex6_time_based_greeting();

    // 7
    // let res = ex7_range_analyzer();
    // println!("Even: {}, Odd: {}", res.0, res.1);

    // 9
    ex9_matrix();
}


// fn ex1(){
//     // 1
//     let app_name = "MyDumbApp"; 
//     let app_version = "1.0.0";
//     let mut requests_count = 0;
//     let mut errors_count = 0;

//     requests_count += 1;
//     errors_count += 1;
//     println!("App: {}, Version: {}, Requests: {}, Errors: {}", app_name, app_version, requests_count, errors_count);
//     println!("Constants: MAX_REQUESTS: {}, TIMEOUT_SECONDS: {}", MAX_REQUESTS, TIMEOUT_SECONDS);

//     println!("Enter number of requests: ");
//     let mut buf = String::new();
//     io::stdin().read_line(&mut buf).expect("Failed to read input");
//     let input: i32 = buf.trim().parse().expect("Invalid Input");
//     requests_count += input;

//     println!("Requests count now: {}", requests_count);
// }

// fn ex2_unit_convert() -> i32{
//     enum Unit {
//         C, F
//     }

//     let mut buf = String::new();
//     io::stdin().read_line(&mut buf).expect("Failed to read input");
//     let n1: i32 = buf.trim().parse().expect("Invalid Input");

//     let mut buf2 = String::new();
//     io::stdin().read_line(&mut buf2).expect("Failed to read input");
//     let u1_inp: u8 = buf2.trim().parse().expect("Invalid Input");

//     let mut buf3 = String::new();
//     io::stdin().read_line(&mut buf3).expect("Failed to read input");
//     let u2_inp: u8 = buf3.trim().parse().expect("Invalid Input");

//     let unit1 = match u1_inp {
//         1 => Unit::C,
//         2 => Unit::F,
//         _ => Unit::C
//     };

//     let unit2 = match u2_inp {
//         1 => Unit::C,
//         2 => Unit::F,
//         _ => Unit::C
//     };

//     let res = match (unit1, unit2) {
//         (Unit::C, Unit::F) => (n1 * 9 / 5) + 32,
//         (Unit::F, Unit::C) => (n1 - 32) * 5 / 9,
//         _ => n1
//     };

//     res

// }

// fn ex3_get_stats(arr: [u32; 5]) -> (u32, u32, u32) {
//     let (mut max, mut min, mut sum) = (0, 0, 0);
//     for v in arr {
//         if v > max {
//             max = v;
//         }
//         if v < min {
//             min = v;
//         }
//         sum += v;
//     }

//     return (max, min, sum / arr.len() as u32);
// }

// fn ex4_password_checker(str: &str) -> (){
//     #[derive(Debug)]
//     enum StrengthLevel {
//         Strong, Medium, Weak
//     }
//     let mut upper = 0;
//     let mut lower = 0;
//     let mut num = 0;
//     for c in str.chars(){
//         match c {
//             'A'..='Z' => upper += 1,
//             'a'..='z' => lower += 1,
//             '0'..='9' => num += 1,
//             _ => ()
//         }
//     }
//     let strength_num = upper + lower + num;
//     let strength = match strength_num {
//         1..=5 => StrengthLevel::Weak,
//         6..=10 => StrengthLevel::Medium,
//         _ => StrengthLevel::Strong
//     };

//     println!("Password Strength: {:?}", strength);
// }

// fn ex5_simple_menu(){
//     let mut buf = String::new();
//     let mut choice: u8;
//     loop {
//         println!("Enter choice: ");
//         buf.clear();
//         io::stdin().read_line(&mut buf).expect("Failed to read choice");
//         choice = buf.trim().parse().expect("Invalid value");
//         match choice {
//             1 => println!("Choice 1"),
//             2 => println!("Choice 2"),
//             3 => println!("Choice 3"),
//             0 => break,
//             _ => println!("Invalid choice")
//         }
//     }
// }

// fn ex6_time_based_greeting(){
//     let mut buf = String::new();
//     let mut hour: i8;
//     loop {
//         println!("Enter hour: ");
//         buf.clear();
//         io::stdin().read_line(&mut buf).expect("Failed to read hour");
//         hour = buf.trim().parse().expect("Invalid value");
//         match hour {
//             0..=11 => println!("Good Morning"),
//             12..=17 => println!("Good Afternoon"),
//             18..=23 => println!("Good Evening"),
//             -1 => break,
//             _ => println!("Invalid hour")
//         }
//     }
// }

// fn ex7_range_analyzer() -> (i32, i32){
//     let mut buf = String::new();
//     io::stdin().read_line(&mut buf).expect("Failed to read input");
//     let mut n1: i32 = buf.trim().parse().expect("Invalid Input");
//     buf.clear();
//     io::stdin().read_line(&mut buf).expect("Failed to read input");
//     let n2: i32 = buf.trim().parse().expect("Invalid Input");

//     let mut evens_cnt = 0;
//     let mut odds_cnt = 0;

//     while n1 < n2{
//         match n1 % 2 {
//             0 => evens_cnt += 1,
//             1 => odds_cnt += 1,
//             _ => ()
//         }
//         n1 += 1;
//     }

//     (evens_cnt, odds_cnt)
// }

fn ex9_matrix(){
    let arr = [[0; 3]; 3];

    for i in 0..3 {
        for j in 0..3 {
            print!("{} ", arr[i][j]);
        }
        println!("");
    }
}