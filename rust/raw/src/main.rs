const MAX_REQUESTS: u32 = 1000;
const TIMEOUT_SECONDS: u8 = 10;

fn main(){
    // 1
    let app_name = "MyDumbApp"; 
    let app_version = "1.0.0";
    let mut requests_count = 0;
    let mut errors_count = 0;

    requests_count += 1;
    errors_count += 1;
    println!("App: {}, Version: {}, Requests: {}, Errors: {}", app_name, app_version, requests_count, errors_count);
    println!("Constants: MAX_REQUESTS: {}, TIMEOUT_SECONDS: {}", MAX_REQUESTS, TIMEOUT_SECONDS);

    // 3
    let arr: [u32; 5] = [6,8,1,4,8];
    println!("{:?}", get_stats(arr));
}

fn get_stats(arr: [u32; 5]) -> (u32, u32, u32) {
    let (mut max, mut min, mut sum) = (0, 0, 0);
    for v in arr {
        if v > max {
            max = v;
        }
        if v < min {
            min = v;
        }
        sum += v;
    }

    return (max, min, sum / arr.len() as u32);
}