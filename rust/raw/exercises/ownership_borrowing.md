# Ownership and Borrowing Exercises

This document contains 50 exercises designed to test and solidify your understanding of Rust's ownership and borrowing rules. Given your C background, some exercises explicitly simulate scenarios where manual memory management would be error-prone.

---

## Part 1: Ownership (20 Exercises)
*Focus: Move semantics, Clone vs Copy, Scope cleanup, Ownership transfer in functions & structs.*

### 1. The Package Router
**Scenario:** You are routing network packets. Once a packet is routed to a specific handler, the main router must no longer have access to it to prevent double-processing.
**Goal:** Fix the compilation error by understanding move semantics.

```rust
struct Packet {
    content: String,
}

fn process_packet(p: Packet) {
    println!("Processing: {}", p.content);
}

fn packet_router() {
    let p = Packet { content: String::from("DATA") };
    process_packet(p);
    // process_packet(p); // Error: Use of moved value
    
    // TODO: Modify the logic (or the struct if needed, but perfer logic) so we can log it after processing? 
    // HINT: You can't use 'p' after line 11. 
    // Challenge: Is there a way to print *before* moving?
}
```

### 2. The Configuration Loader
**Scenario:** A server loads a configuration. We want to use the config to start a service, but also keep a copy for an admin dashboard.
**Goal:** Use `.clone()` explicitly where expensive copies are needed.

```rust
#[derive(Debug)] // Note: No Copy/Clone derived yet
struct Config {
    port: u16,
    domain: String,
}

fn start_service(cfg: Config) {
    println!("Starting on {}", cfg.port);
}

fn admin_dashboard() {
    let cfg = Config { port: 8080, domain: String::from("example.com") };
    
    // TODO: Fix this. We need to send 'cfg' to start_service but keep it here too.
    // start_service(cfg); 
    // println!("Admin seeing config: {:?}", cfg);
}
```

### 3. File Handle Wrapper (RAII)
**Scenario:** Simulating a file handle that closes when dropped. You cannot close it twice.
**Goal:** Ensure `consume_file` takes full ownership and the original variable cannot be used.

```rust
struct MyFile {
    fd: i32,
}
impl Drop for MyFile {
    fn drop(&mut self) {
        println!("Closing fd {}", self.fd);
    }
}

fn consume_file(f: MyFile) {
    println!("Using fd {}", f.fd);
    // f drops here
}

fn file_system() {
    let f = MyFile { fd: 10 };
    consume_file(f);
    // println!("{}", f.fd); // This must fail to compile. Why? verify it.
}
```

### 4. Partial Moves (Structs)
**Scenario:** You have a `User` struct. You want to transfer the `username` to another system but keep the `id` for logging.
**Goal:** perform a partial move.

```rust
struct User {
    id: u32,
    username: String,
}

fn register_name(name: String) {}

fn partial_move_demo() {
    let u = User { id: 1, username: String::from("alice") };
    
    // TODO: Move u.username into register_name().
    // Show that you can still access u.id afterwards.
    // Show that you CANNOT access u normally afterwards.
}
```

### 5. Array of Strings vs Array of Ints
**Scenario:** Understanding why some arrays copy and some move.
**Goal:** Explain why `[i32; 4]` implements Copy but `[String; 4]` does not.

```rust
fn array_semantics() {
    let a_int = [1, 2, 3];
    let b_int = a_int; 
    println!("{:?}", a_int); // Works

    let a_str = [String::from("a"), String::from("b")];
    // let b_str = a_str;
    // println!("{:?}", a_str); // Fix this error without cloning (impossible? or use referencing?). 
    // Actually, for this exercise: Call .clone() on the array if you want two copies. 
    // OR: Realize you can't have two owners of the strings.
}
```

### 6. Tuple Ownership
**Scenario:** A tuple `(u32, String)`.
**Goal:** Extract the String (move it out) while leaving the u32 usable?
```rust
fn tuple_extraction() {
    let t = (10, String::from("hello"));
    let (num, s) = t; // Destructuring moves 's' out.
    
    // println!("{:?}", t); // Error.
    println!("{}", num); // Is this valid? Test it.
}
```

### 7. Option::match Ownership
**Scenario:** Handling an optional resource.
**Goal:** Use `match` to consume the value inside an Option (unwrap by move).

```rust
fn process_option(o: Option<String>) {
    match o {
        Some(s) => println!("Got string: {}", s), // 's' is moved here
        None => {},
    }
    // println!("{:?}", o); // Should be invalid here.
}
```

### 8. Box<T> Moves
**Scenario:** Managing heap memory explicitly.
**Goal:** Understanding that `Box<T>` owns the heap data. Moving the box moves the pointer, not the data (cheap).

```rust
fn box_mover() {
    let b1 = Box::new(String::from("Large Data"));
    let b2 = b1; // Pointer copy, ownership move.
    
    // println!("{}", b1); // Fail
    println!("{}", b2);
}
```

### 9. Conditional Ownership
**Scenario:** A variable might be moved in one branch but not another.
**Goal:** Rust ensures variable is initialized/valid in all paths if used later.

```rust
fn conditional_move(flag: bool) {
    let s = String::from("text");
    if flag {
        let _s2 = s; // Moved here
    }
    
    // println!("{}", s); // Why is this always an error, even if flag is false?
    // Fix: If you need s later, you must handle the logic so 's' is available or re-initialized.
}
```

### 10. Loop Consumption
**Scenario:** Processing a queue of jobs.
**Goal:** Consume a vector by iterating (moving) elements out.

```rust
fn consume_vec() {
    let v = vec![String::from("job1"), String::from("job2")];
    // for job in v { // Implicit v.into_iter()
    //     println!("{}", job);
    // }
    // println!("{:?}", v); // Error.
    // Task: Rewrite loop to NOT consume v (iterate by reference) -> see Part 2.
    // For this exercise: Explicitly write `v.into_iter()` to show intent.
}
```

### 11. Returning Ownership
**Scenario:** A builder function.
**Goal:** Create a string, modify it, and return ownership.

```rust
fn build_url(domain: String) -> String {
    let mut s = String::from("https://");
    s.push_str(&domain);
    s // Return moves it out
}
```

### 12. Chain of Command
**Scenario:** Passing a token through multiple layers.
**Goal:** `layer1` passes to `layer2`, passes to `layer3`.

```rust
struct Token(u32);
fn layer3(t: Token) {}
fn layer2(t: Token) { layer3(t); }
fn layer1(t: Token) { layer2(t); }
```

### 13. The "Copy" Trait implementation
**Scenario:** A small Point struct.
**Goal:** Derive `Copy` so it acts like an integer.

```rust
#[derive(Debug, Clone, Copy)] // Try removing Copy
struct Point { x: i32, y: i32 }

fn use_point() {
    let p1 = Point { x: 1, y: 2 };
    let p2 = p1;
    println!("{:?}", p1); // Valid only if Copy
}
```

### 14. Enum with Data
**Scenario:** An event system.
**Goal:** Move data into an Enum variant.

```rust
enum Event {
    Click,
    Input(String),
}

fn handle_event() {
    let s = String::from("Hello");
    let e = Event::Input(s); // s moved into e
    // println!("{}", s); // Error
}
```

### 15. Re-assignment of Moved Variable
**Scenario:** Reusing a variable name slot.
**Goal:** You can assign a new value to a variable even after its previous value moved.

```rust
fn reassignment() {
    let mut s = String::from("v1");
    let v2 = s;
    // s is invalid now.
    s = String::from("v2"); // Now valid again!
    println!("{}", s);
}
```

### 16. Destructuring Structs completely
**Scenario:** Breaking a struct into pieces.
**Goal:** Use pattern matching to destructure and move all fields.

```rust
struct Person { name: String, age: u8 }
fn decompose() {
    let p = Person { name: String::from("Bob"), age: 30 };
    let Person { name, age } = p;
    // p is gone. name and age are new owners.
}
```

### 17. Vec::push moves
**Scenario:** Building a list.
**Goal:** Understanding that pushing a value into a Vec moves ownership into the Vec.

```rust
fn vec_push() {
    let mut list = Vec::new();
    let s = String::from("item");
    list.push(s);
    // println!("{}", s); // Error
}
```

### 18. Closures taking ownership (move keyword)
**Scenario:** A thread spawner (simulation).
**Goal:** Ensure the closure owns the data it prints.

```rust
fn closure_move() {
    let s = String::from("thread data");
    let c = move || { // 'move' keyword forces capture by value
        println!("{}", s);
    };
    c();
    // println!("{}", s); // Error
}
```

### 19. Replacing content in place
**Scenario:** Swapping a value in a mutable memory slot without leaving it empty.
**Goal:** Use `std::mem::replace`.

```rust
fn mem_replace_demo() {
    let mut s1 = String::from("A");
    let s2 = String::from("B");
    
    // We want to put s2 into s1's memory location, and get the old s1 out.
    // s1 = s2; // Error: this drops old s1 immediately, which is fine, but if we wanted it?
    let old_s1 = std::mem::replace(&mut s1, s2);
    println!("Old: {}, New in s1: {}", old_s1, s1);
}
```

### 20. Default and Take
**Scenario:** Taking a value out of a struct member that implements Default.
**Goal:** Use `std::mem::take`.

```rust
#[derive(Default)]
struct Buffer { data: String }

fn consume_buffer(b: &mut Buffer) {
    let data = std::mem::take(&mut b.data); // Swaps with String::default ("")
    println!("Consumed: {}", data);
    println!("Buffer now empty: '{}'", b.data);
}
```

---

## Part 2: Borrowing (20 Exercises)
*Focus: References `&T`, `&mut T`, Lifetimes (implicit), Slices, Data Race prevention.*

### 21. Basic Read-Only
**Scenario:** Calculate length.
**Goal:** Pass by ref `&String`.

```rust
fn borrow_len(s: &String) -> usize { s.len() }
fn usage() {
    let s = String::from("test");
    let len = borrow_len(&s);
    println!("{} {}", s, len); // s still alive
}
```

### 22. Mutable Barrow
**Scenario:** Append text.
**Goal:** Pass `&mut String`.

```rust
fn add_suffix(s: &mut String) { s.push_str(" world"); }
```

### 23. Multiple Immutable ok, One Mutable NO
**Scenario:** The Reader-Writer lock rule.
**Goal:** Trigger the error.

```rust
fn fail_borrow() {
    let mut s = String::from("data");
    let r1 = &s;
    let r2 = &s; // OK
    // let r3 = &mut s; // ERROR: cannot borrow mutably while borrowed immutably
    println!("{} {}", r1, r2);
}
```

### 24. Scope Restriction
**Scenario:** Create a scope to end a borrow early.
**Goal:** Allow mutation after reading finishes.

```rust
fn scope_tricks() {
    let mut s = String::from("data");
    {
        let r1 = &s;
        println!("{}", r1);
    } // r1 dies here
    
    let r2 = &mut s; // OK now
    r2.push('!');
}
```

### 25. Dangling Reference Prevention
**Scenario:** Returning a reference to a local variable.
**Goal:** Understand why this is forbidden (stack frame destroyed).

```rust
// fn return_ref() -> &String {
//     let s = String::from("local");
//     &s // ERROR
// }
```

### 26. Slices (String)
**Scenario:** Parsing a command. First word is command.
**Goal:** Return `&str` slice.

```rust
fn first_word(s: &String) -> &str {
    &s[0..3] // Simplified
}
```

### 27. Slices (Array)
**Scenario:** Analyzing a window of data.
**Goal:** `&[i32]`.

```rust
fn analyze_window(data: &[i32]) {
    // can pass Vec or Array
}
fn call_window() {
    let arr = [1, 2, 3, 4];
    analyze_window(&arr[1..3]);
    let vec = vec![1, 2, 3];
    analyze_window(&vec); // deref coercion to slice
}
```

### 28. Iterating by Reference
**Scenario:** Summing numbers without consuming list.
**Goal:** `for x in &vec`.

```rust
fn sum_list(v: &Vec<i32>) -> i32 {
    let mut sum = 0;
    for x in v { // x is &i32
        sum += *x;
    }
    sum
}
```

### 29. Mutating in Loop
**Scenario:** Zeroing out negative numbers.
**Goal:** `for x in &mut vec`.

```rust
fn sanitize(v: &mut Vec<i32>) {
    for x in v {
        if *x < 0 { *x = 0; }
    }
}
```

### 30. Dereferencing
**Scenario:** Following the pointer.
**Goal:** Read and write through `*`.

```rust
fn deref_basics() {
    let mut x = 10;
    let r = &mut x;
    *r += 1;
    assert_eq!(x, 11);
}
```

### 31. Ref pattern matching
**Scenario:** Match gives references.
**Goal:** `match` on `&Option`.

```rust
fn match_ref(o: &Option<String>) {
    match o {
        Some(s) => println!("Look at {}", s), // s is &String
        None => {},
    }
    // o is still valid
}
```

### 32. Ref Mut pattern matching
**Scenario:** Modify inside match.

```rust
fn modify_opt(o: &mut Option<String>) {
    match o {
        Some(s) => s.push_str(" updated"), // s is &mut String
        None => {},
    }
}
```

### 33. The "ref" keyword
**Scenario:** Legacy/explicit destructuring.
**Goal:** Bind `ref` in pattern.

```rust
fn ref_keyword() {
    let o = Some(String::from("A"));
    match o {
        Some(ref s) => println!("{}", s), // s is &String, o not moved
        None => {},
    }
}
```

### 34. Struct with References (Introduction)
**Scenario:** A View struct (like a database view).
**Note:** Requires lifetime annotations.
**Goal:** Define a struct holding a ref.

```rust
struct TextView<'a> {
    content: &'a str,
}
```

### 35. Slice of a Slice
**Scenario:** Sub-parsing.
**Goal:** Slicing is recursive.

```rust
fn sub_slice() {
    let s = String::from("abcdef");
    let slice1 = &s[0..4]; // abcd
    let slice2 = &slice1[1..3]; // bc
}
```

### 36. Implicit Re-borrowing
**Scenario:** Calling helper functions.
**Goal:** Mutable references implicitly re-borrow.

```rust
fn helper(s: &mut String) {}
fn main_job(s: &mut String) {
    helper(s); // passes transient &mut *s, effectively
    helper(s); // s still valid
}
```

### 37. Vectors and invalidation
**Scenario:** Pushing to a vector while holding a reference to an element.
**Goal:** Trigger "cannot borrow `v` as mutable because it is also borrowed as immutable".
**Why:** Reallocation moves the data.

```rust
fn invalidation() {
    let mut v = vec![1, 2, 3];
    let first = &v[0];
    // v.push(4); // Error! 
    println!("{}", first);
}
```

### 38. Split mutable borrow
**Scenario:** Modifying two parts of an array/slice.
**Goal:** `split_at_mut`. Rust doesn't allow `&mut v[0]` and `&mut v[1]` manually easily (compiler isn't smart enough to know indices don't overlap).

```rust
fn split_mut_demo() {
    let mut v = [1, 2, 3, 4];
    let (left, right) = v.split_at_mut(2);
    left[0] = 10;
    right[0] = 20;
}
```

### 39. RefCell (Runtime Borrowing) - Concept Check
**Scenario:** When compile time rules are too strict.
**Goal:** Acknowledge `RefCell<T>` checks these rules at runtime (panics if violated).

### 40. Cow (Clone on Write)
**Scenario:** Keep as reference if read-only, clone if write needed.
**Goal:** User `std::borrow::Cow`.

---

## Part 3: Mixed & Challenging (10 Exercises)
*Focus: Real world complexity, fighting the borrow checker, logic puzzles.*

### 41. The Bank Account Transaction
**Scenario:** Transfer money between two accounts.
**Challenge:** You have `&mut Bank`. You need to get mutable references to two *different* accounts inside a `HashMap` or `Vec`.
**Problem:** `bank.accounts.get_mut(a)` and `bank.accounts.get_mut(b)` might overlap in compiler's eyes (borrowing `bank.accounts` twice).
**Solution:** Use entry API or split_at_mut or indices.

### 42. Linked List (The Rust Nemesis)
**Scenario:** Try to verify a simple doubly linked list node.
**Challenge:** `struct Node { next: Option<Box<Node>>, prev: Option<&Node> }`.
**Lesson:** This is hard. Use `Rc<RefCell<Node>>` or arena indices. For this exercise, try defining just `Box` (Single linked) and see ownership work.

### 43. Iterator Invalidation Logic
**Scenario:** Removing elements from a Vec while iterating.
**Code:** `v.retain(|x| ...)` vs `for i in 0..v.len() { remove }`.
**Goal:** Efficient in-place filtering.

### 44. Self-Referential Struct
**Scenario:** A struct holding a String and a Reference to a slice of that String.
**Code:** `struct SelfRef { text: String, slice: &str }`.
**Result:** Impossible in safe Rust without crates like `ouroboros` or `Pin`. Explain why.

### 45. The Mutex Guard
**Scenario:** Understanding that a Lock Guard is a proxy that owns the lock.
**Code:** `let data = mutex.lock().unwrap();`. `data` is a RAII guard. `*data` accesses the inner content.

### 46. Thread scoped borrowing
**Scenario:** `std::thread::scope`.
**Goal:** Threads usually require `'static` (ownership). Scoped threads allow borrowing stack locals because they guarantee join before scope end.

### 47. Closure Capture Modes
**Scenario:** A closure that tries to mutate a captured variable, print another, and consume a third.
**Goal:** Determine the trait: `Fn`, `FnMut`, or `FnOnce`.

### 48. Constructor returning reference?
**Scenario:** A classic C pattern: `const char* get_name(Context* ctx)`.
**Rust:** `fn get_name(&self) -> &str`.
**Challenge:** Returning a reference tied to `self`'s lifetime.

### 49. Buffer Parser with Cursor
**Scenario:** You wrap a `Vec<u8>`. You have a `Cursor<'a>` struct that borrows the Vec.
**Invariant:** You cannot modify the Vec while the Cursor is active.

### 50. Global Singleton (lazy_static / OnceLock)
**Scenario:** Global state ownership.
**Challenge:** Who owns global data? It is `'static`. How to mutate it safe? (`Mutex`).

---

## Instructions for User
1. Copy each code block into a `.rs` file (e.g., in Rust playground or local cargo project).
2. Uncomment commented-out error lines to see the compiler message.
3. Fix the errors while maintaining the **Scenario** logic.
