#[test]
fn test_owbor_ex1() {
    struct Packet {
        content: String,
    }

    fn process_packet(p: &Packet) -> String {
        format!("Processing: {}", p.content.clone())
    }

    fn packet_router() {
        let p = Packet {
            content: String::from("DATA"),
        };
        let processed = process_packet(&p);
        println!("{}", processed);
    }

    packet_router();
}
