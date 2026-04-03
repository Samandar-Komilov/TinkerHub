#[allow(dead_code)]
#[test]
fn test_ex00_user_struct() {
    #[derive(Debug, Clone)]
    struct User {
        id: u32,
        email: String,
        is_active: bool,
    }

    impl User {
        fn new(id: u32, email: String) -> Self {
            Self {
                id,
                email,
                is_active: true,
            }
        }

        fn deactivate(&mut self) {
            self.is_active = false;
        }
    }

    let user1 = User::new(1, String::from("t4jZt@example.com"));
    let mut user2 = user1.clone();
    user2.deactivate();

    // Check invariants
    assert!(user1.is_active);
    assert!(!user2.is_active);
    assert_eq!(user1.id, 1);
    assert_eq!(user2.email, "t4jZt@example.com");
}

#[test]
fn test_ex01_order_status() {
    #[derive(Debug, PartialEq, Clone, Copy)]
    enum OrderStatus {
        Created,
        Paid,
        Shipped,
        Delivered,
        Cancelled,
    }

    struct Order {
        status: OrderStatus,
    }

    impl Order {
        fn new() -> Self {
            Self {
                status: OrderStatus::Created,
            }
        }

        fn advance(&mut self) {
            match self.status {
                OrderStatus::Created => self.status = OrderStatus::Paid,
                OrderStatus::Paid => self.status = OrderStatus::Shipped,
                OrderStatus::Shipped => self.status = OrderStatus::Delivered,
                OrderStatus::Delivered => self.status = OrderStatus::Cancelled,
                OrderStatus::Cancelled => (),
            }
        }
    }

    let mut order = Order::new();
    assert_eq!(order.status, OrderStatus::Created);

    order.advance();
    assert_eq!(order.status, OrderStatus::Paid);

    order.advance();
    assert_eq!(order.status, OrderStatus::Shipped);

    order.advance();
    assert_eq!(order.status, OrderStatus::Delivered);

    order.advance();
    assert_eq!(order.status, OrderStatus::Cancelled);

    order.advance();
    assert_eq!(order.status, OrderStatus::Cancelled);
}

// New exercises

#[test]
fn test_ex1_shape() {
    enum Shape {
        Circle { radius: f64 },
        Rectangle { width: f64, height: f64 },
        Square { side: f64 },
    }

    impl Shape {
        fn area(&self) -> f64 {
            match self {
                Shape::Circle { radius } => radius.powi(2) * 3.14,
                Shape::Rectangle { width, height } => width * height,
                Shape::Square { side } => side.powi(2),
            }
        }
    }

    let c1 = Shape::Circle { radius: 5.0 };
    let r1 = Shape::Rectangle {
        width: 4.0,
        height: 3.0,
    };
    let s1 = Shape::Square { side: 7.0 };

    assert_eq!(c1.area(), 78.5);
    assert_eq!(r1.area(), 12.0);
    assert_eq!(s1.area(), 49.0);
}

#[test]
fn test_ex2_direction_navigation() {
    enum Direction {
        North,
        South,
        East,
        West,
    }
    struct Point {
        x: i32,
        y: i32,
    }

    impl Point {
        fn move_point(&mut self, dir: Direction) {
            match dir {
                Direction::East => self.x += 1,
                Direction::North => self.y += 1,
                Direction::West => self.x -= 1,
                Direction::South => self.y -= 1,
            }
        }
    }

    let mut p1 = Point { x: 0, y: 0 };
    p1.move_point(Direction::East);
    assert_eq!(p1.x, 1);
    p1.move_point(Direction::North);
    assert_eq!(p1.y, 1);
    p1.move_point(Direction::West);
    assert_eq!(p1.x, 0);
    p1.move_point(Direction::South);
    assert_eq!(p1.y, 0);
}

#[test]
fn test_ex3_media_player() {
    enum MediaState {
        Playing(String),
        Paused(String),
        Stopped,
    }

    impl MediaState {
        fn play(&self) -> String {
            match self {
                MediaState::Playing(track) => {
                    format!("Playing {}", track)
                }
                MediaState::Paused(track) => {
                    format!("Paused {}", track)
                }
                MediaState::Stopped => "Nothing to play".to_string(),
            }
        }
    }

    let m1 = MediaState::Playing("Till I Collapse".to_string());
    let m2 = MediaState::Paused("Mockingbird".to_string());
    let m3 = MediaState::Stopped;

    assert_eq!(m1.play(), "Playing Till I Collapse".to_string());
    assert_eq!(m2.play(), "Paused Mockingbird".to_string());
    assert_eq!(m3.play(), "Nothing to play".to_string());
}

#[test]
fn test_ex4_loglevels() {
    enum LogLevel {
        Info,
        Warning,
        Error,
    }
    struct LogMessage {
        level: LogLevel,
        msg: String,
    }

    fn format_log(log: LogMessage) -> String {
        match log.level {
            LogLevel::Info => format!("[INFO] {}", log.msg),
            LogLevel::Warning => format!("[WARNING] {}", log.msg),
            LogLevel::Error => format!("[ERROR] {}", log.msg),
        }
    }

    let i = LogMessage {
        level: LogLevel::Info,
        msg: "Working...".to_string(),
    };
    let w = LogMessage {
        level: LogLevel::Warning,
        msg: "This code is deprecated".to_string(),
    };
    let e = LogMessage {
        level: LogLevel::Error,
        msg: "Error occured".to_string(),
    };

    assert_eq!(format_log(i), "[INFO] Working...".to_string());
    assert_eq!(
        format_log(w),
        "[WARNING] This code is deprecated".to_string()
    );
    assert_eq!(format_log(e), "[ERROR] Error occured".to_string());
}

#[test]
fn test_ex6_safe_div() {
    fn safe_div(a: i32, b: i32) -> Option<i32> {
        if b == 0 { None } else { Some(a / b) }
    }

    assert_eq!(safe_div(6, 2), Some(3));
    assert_eq!(safe_div(5, 0), None);
}

#[test]
fn test_ex7_username() {
    fn get_username(id: u32) -> Option<String> {
        match id {
            1 => Some("Alice".to_string()),
            5 => Some("Eshmat".to_string()),
            _ => None,
        }
    }

    assert_eq!(get_username(1), Some("Alice".to_string()));
    assert_eq!(get_username(0), None);
}

#[test]
fn test_ex8_fileparser() {
    fn parse_percentage(input: &str) -> Result<u8, String> {
        match input.parse::<i32>() {
            Ok(n) => {
                if n >= 0 && n <= 100 {
                    Ok(n as u8)
                } else {
                    Err(format!("Value {} is out of range (0-100)", n))
                }
            }
            Err(_) => Err(format!("'{}' is not a valid number", input)),
        }
    }

    assert_eq!(parse_percentage("56"), Ok(56));
    assert_eq!(
        parse_percentage("hello"),
        Err("'hello' is not a valid number".to_string())
    );
    assert_eq!(
        parse_percentage("101"),
        Err("Value 101 is out of range (0-100)".to_string())
    )
}

#[test]
fn test_ex9_validate_login() {
    fn validate_login(username: &str, password: &str) -> Result<(), String> {
        // Check if username is empty
        if username.is_empty() {
            return Err("Username is empty".to_string());
        }

        // Check password length
        if password.len() < 8 {
            return Err("Password is too short".to_string());
        }

        // If both pass, return success
        Ok(())
    }

    assert_eq!(
        validate_login("", "sss"),
        Err("Username is empty".to_string())
    );
    assert_eq!(
        validate_login("user", "1234567"),
        Err("Password is too short".to_string())
    );
    assert_eq!(validate_login("user1", "q!w@e#r$"), Ok(()));
}

#[test]
fn test_ex10_array_element() {
    fn get_element(arr: &[i32], index: usize) -> Option<i32> {
        match arr.get(index) {
            Some(v) => Some(*v),
            None => None,
        }
    }

    assert_eq!(get_element(&[1, 2, 3], 1), Some(2));
    assert_eq!(get_element(&[1, 2, 3], 3), None);
}
