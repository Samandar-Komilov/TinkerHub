#[allow(dead_code)]

// 1. App Stats
fn run_app_stats(requests: i32) -> (i32, i32) {
    let _app_name = "MyDumbApp";
    let _app_version = "1.0.0";
    let mut requests_count = 0;
    let mut errors_count = 0;

    requests_count += 1;
    errors_count += 1;

    requests_count += requests;
    (requests_count, errors_count)
}

#[test]
fn test_app_stats() {
    let (req, err) = run_app_stats(10);
    assert_eq!(req, 11); // 1 initial + 10 added
    assert_eq!(err, 1);
}

// 2. Unit Converter
#[test]
fn test_unit_converter() {
    enum Unit {
        C,
        F,
    }
    fn convert_temp(n1: i32, u1_inp: u8, u2_inp: u8) -> i32 {
        let unit1 = match u1_inp {
            1 => Unit::C,
            2 => Unit::F,
            _ => Unit::C,
        };

        let unit2 = match u2_inp {
            1 => Unit::C,
            2 => Unit::F,
            _ => Unit::C,
        };

        match (unit1, unit2) {
            (Unit::C, Unit::F) => (n1 * 9 / 5) + 32,
            (Unit::F, Unit::C) => (n1 - 32) * 5 / 9,
            _ => n1,
        }
    }

    // 0C -> 32F
    assert_eq!(convert_temp(0, 1, 2), 32);
    // 32F -> 0C
    assert_eq!(convert_temp(32, 2, 1), 0);
    // Same unit
    assert_eq!(convert_temp(100, 1, 1), 100);
}

// 3. Stats
#[test]
fn test_stats() {
    fn get_stats(arr: &[u32]) -> (u32, u32, u32) {
        let (mut max, mut min, mut sum) = (0, u32::MAX, 0);
        // Handle empty case if needed, but assuming non-empty for now
        if arr.is_empty() {
            return (0, 0, 0);
        }

        min = arr[0]; // Initialize min properly

        for &v in arr {
            if v > max {
                max = v;
            }
            if v < min {
                min = v;
            }
            sum += v;
        }

        (max, min, sum / arr.len() as u32)
    }

    let arr = [6, 8, 1, 4, 8];
    let (max, min, avg) = get_stats(&arr);
    assert_eq!(max, 8);
    assert_eq!(min, 1);
    assert_eq!(avg, 5); // 27 / 5 = 5
}

// 4. Password Checker
#[test]
fn test_password_strength() {
    #[derive(Debug, PartialEq)]
    enum StrengthLevel {
        Strong,
        Medium,
        Weak,
    }

    fn check_password_strength(s: &str) -> StrengthLevel {
        let mut upper = 0;
        let mut lower = 0;
        let mut num = 0;
        for c in s.chars() {
            match c {
                'A'..='Z' => upper += 1,
                'a'..='z' => lower += 1,
                '0'..='9' => num += 1,
                _ => (),
            }
        }
        let strength_num = upper + lower + num;
        match strength_num {
            0..=5 => StrengthLevel::Weak,
            6..=10 => StrengthLevel::Medium,
            _ => StrengthLevel::Strong,
        }
    }

    assert_eq!(check_password_strength("abc"), StrengthLevel::Weak);
    assert_eq!(check_password_strength("abcdef12"), StrengthLevel::Medium); // 8 chars
    assert_eq!(
        check_password_strength("Password123456"),
        StrengthLevel::Strong
    );
}

// 5. Menu (Logic test)
#[test]
fn test_menu() {
    fn get_menu_choice_response(choice: u8) -> &'static str {
        match choice {
            1 => "Choice 1",
            2 => "Choice 2",
            3 => "Choice 3",
            0 => "Exit",
            _ => "Invalid choice",
        }
    }
    assert_eq!(get_menu_choice_response(1), "Choice 1");
    assert_eq!(get_menu_choice_response(99), "Invalid choice");
}

// 6. Time Greeting
#[test]
fn test_greeting() {
    fn get_greeting(hour: i8) -> &'static str {
        match hour {
            0..=11 => "Good Morning",
            12..=17 => "Good Afternoon",
            18..=23 => "Good Evening",
            _ => "Invalid hour",
        }
    }
    assert_eq!(get_greeting(10), "Good Morning");
    assert_eq!(get_greeting(14), "Good Afternoon");
    assert_eq!(get_greeting(20), "Good Evening");
}

// 7. Range Analyzer
#[test]
fn test_range_analyzer() {
    fn analyze_range(mut n1: i32, n2: i32) -> (i32, i32) {
        let mut evens_cnt = 0;
        let mut odds_cnt = 0;

        while n1 < n2 {
            match n1 % 2 {
                0 => evens_cnt += 1,
                _ => odds_cnt += 1, // handles 1 and -1 etc
            }
            n1 += 1;
        }

        (evens_cnt, odds_cnt)
    }
    // 0, 1, 2, 3, 4 (end 5 exclusive)
    // Even: 0, 2, 4 (3)
    // Odd: 1, 3 (2)
    let (evens, odds) = analyze_range(0, 5);
    assert_eq!(evens, 3);
    assert_eq!(odds, 2);
}

// 9. Matrix
#[test]
fn test_matrix_creation() {
    let arr = [[0; 3]; 3];
    assert_eq!(arr.len(), 3);
    assert_eq!(arr[0].len(), 3);
    assert_eq!(arr[0][0], 0);
}
