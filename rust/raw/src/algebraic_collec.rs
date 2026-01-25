#[allow(dead_code)]
#[test]
fn test_ex1_user_struct() {
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
fn test_ex2_order_status() {
    #[derive(Debug, PartialEq, Clone, Copy)]
    enum OrderStatus {
        Created,
        Paid,
        Shipped,
        Delivered,
        Cancelled,
    }

    struct Order {
        id: u32,
        status: OrderStatus,
    }

    impl Order {
        fn new(id: u32) -> Self {
            Self {
                id,
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

    let mut order = Order::new(101);
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
