use std::io;

pub fn temp_convert(val: f64, to_scale: &str) -> f64{
    let trimmed_scale:&str = to_scale.trim();
    if trimmed_scale == "Farenheit" {
        let res: f64 = val * 1.8 + 32.0;
        return res;
    } else if trimmed_scale == "Kelvin" {
        let res: f64 = val + 273.15;
        return res;
    } else {
        println!("You can only convert to Farenheit or Kelvin. Your measure is '{}'", to_scale);
        return 0.0;
    }
}

pub fn temp_convert_main() -> () {
    let mut input: String = String::new();
    io::stdin()
        .read_line(&mut input)
        .expect("Failed to read the line.");

    let inputs_lst: Vec<&str> = input.split(' ').collect();
    let val: f64 = inputs_lst[0].parse().unwrap();
    let to_scale: &str = inputs_lst[1];

    let res: f64 = temp_convert(val, to_scale);

    println!("Result: {res}");
}