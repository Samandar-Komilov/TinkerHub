// Celsius to [Farenheit, Kelvin]

pub fn temp_convert(val: f64, to_scale: &str) -> f64{
    if to_scale.eq("Farenheit"){
        let res: f64 = val * 1.8;
        return res
    }
    else if to_scale.eq("Kelvin"){
        let res: f64 = val + 315.00;
        return res
    }
    else{
        println!("You can only convert to Farenheit or Kelvin. Your measure is {to_scale}");
        return 0.0;
    }
}