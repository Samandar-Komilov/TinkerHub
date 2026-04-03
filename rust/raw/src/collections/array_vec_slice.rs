#[allow(dead_code)]
#[test]
fn test_array1() {
    let arr: [u32; 5] = [5, 10, 15, 75, 35];
    // array indexing
    assert!(arr.first().unwrap_or(&0) == &5);
    assert!(arr.get(6).unwrap_or(&0) == &0);

    // max min and finding index by value
    let max = arr.iter().max().unwrap();
    let min = arr.iter().min().unwrap();
    assert!(arr.iter().position(|x| x == max).unwrap() == 3);
    assert!(arr.iter().position(|y| y == min).unwrap() == 0);

    // convert to farenheit all with .map()
    let far = arr.map(|x| x * 9 / 5 + 32);
    assert!(far[0] == 41);

    // .contains()
    assert!(arr.contains(&74) == false);
    println!("{:?}", far);
}
