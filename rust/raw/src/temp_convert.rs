// Celsius to [Farenheit, Kelvin]

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