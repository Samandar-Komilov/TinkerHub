# Algebraic Data Types (ADTs): A Systems Engineering Perspective

## 1. What are Algebraic Data Types? (Concept)
Before diving into code, let's understand ADTs as a mathematical concept used to model data. The term "Algebraic" comes from the fact that we can create new types by combining existing types using algebraic operations, primarily **Sum** and **Product**.

Imagine you are designing a system state. You need to define the "universe" of all possible values your data can hold.

### The Product Type (AND)
A product type combines multiple types together. A value of a product type contains a value of type A **AND** a value of type B.
- **Mathematical Analogy**: Multiplication ($A \times B$).
- **Cardinality (Total States)**: If Type A has 2 possible states (e.g., bool) and Type B has 4 possible states, the product has $2 \times 4 = 8$ states.
- **Example**: A "Coordinate" is an Integer $X$ **AND** an Integer $Y$.

### The Sum Type (OR)
A sum type represents a choice between multiple types. A value of a sum type holds a value of type A **OR** a value of type B.
- **Mathematical Analogy**: Addition ($A + B$).
- **Cardinality (Total States)**: If Type A has 2 states and Type B has 4 states, the sum has $2 + 4 = 6$ states.
- **Example**: A "Result" is either Success **OR** Failure. It cannot be both.

### Why does this matter?
Most older languages (C, Java, Python) heavy rely on Product types (Classes, Structs) but struggle with Sum types (often using inheritance or null pointers to fake it). Rust treats both as first-class citizens, allowing you to model data precisely.

---

## 2. Structs (The Product Type)
In Rust, the Product type is realized primarily through **Structs**.

```rust
struct SensorReading {
    device_id: u8,    // 1 byte (256 states)
    value: u16,       // 2 bytes (65,536 states)
    is_active: bool   // 1 byte (2 states)
}
```

### Memory Layout & Performance
As a systems engineer, you control how bits are laid out in memory.
- **Sequential**: Fields are laid out roughly in order, but the compiler may reorder them to optimize packing (reduce gaps).
- **Alignment**: CPU reads are faster when data is aligned to its size (e.g., a 4-byte integer should start at an address divisible by 4). To achieve this, the compiler adds **Padding** (wasted bytes).

**Example Layout**:
If `SensorReading` kept the order `u8`, `u16`, `bool`:
- `device_id` at offset 0.
- `padding` at offset 1 (to align the next u16 to 2).
- `value` at offset 2.
- `is_active` at offset 4.
- `padding` at offset 5.
- Total Size: 6 bytes.

Rust automatically reorders to: `u16`, `u8`, `bool` (or similar) to remove padding, likely fitting it into 4 bytes.

**Performance**: Accessing struct fields is extremely fast (pointer arithmetic).

---

## 3. Enums (The Sum Type)
In Rust, the Sum type is the **Enum**. It is much more powerful than C-enums. A Rust enum can hold data.

```rust
enum NetworkEvent {
    Connected,                  // variant 1 (0 bytes data)
    Received(u8),               // variant 2 (1 byte data)
    Error(u32),                 // variant 3 (4 bytes data)
}
```

### Memory Layout (Tagged Union)
How does memory store "A OR B"? It needs to know *which* one is currently stored.
- **Tag (Discriminant)**: An integer (usually 1 byte) stored at the beginning to indicate the current variant.
- **Union**: A reserved space large enough to hold the *largest* variant.

**Layout of `NetworkEvent`**:
- **Tag**: 1 byte.
- **Payload**: 4 bytes (max size is `u32`).
- **Padding**: 3 bytes (to align the `u32`).
- **Total Size**: 1 + 3 (padding) + 4 = 8 bytes.

Every instance of `NetworkEvent` takes 8 bytes, regardless of whether it holds `Connected` (nothing) or `Error` (4 bytes).

---

## 4. The Option Enum
Safe nullability. Instead of a "null pointer" which crashes billion-dollar systems, Rust uses `Option<T>`.

It is literally defined as:
```rust
enum Option<T> {
    None,
    Some(T),
}
```

- **Usage**: Use it when a value might be missing.
- **Safety**: You *must* check if it is `Some` or `None` before using the data. You cannot accidentally dereference `None`.

**Example**:
```rust
fn find_user(id: i32) -> Option<String> {
    if id == 1 {
        Some("Alice".to_string())
    } else {
        None
    }
}

// Usage
let user = find_user(5);
match user {
    Some(name) => println!("Found: {}", name),
    None => println!("User not found"),
}
```

---

## 5. The Result Enum
Standardized Error Handling. Instead of throwing exceptions (which secretly break control flow), Rust returns a value indicating success or failure.

Defined as:
```rust
enum Result<T, E> {
    Ok(T),
    Err(E),
}
```

- **Usage**: When an operation can fail (I/O, parsing, logic).
- **T**: The type of value if successful.
- **E**: The type of error if failed.

**Example**:
```rust
fn divide(a: i32, b: i32) -> Result<i32, String> {
    if b == 0 {
        Err("Division by zero".to_string())
    } else {
        Ok(a / b)
    }
}

let outcome = divide(10, 2);
match outcome {
    Ok(val) => println!("Result is {}", val),
    Err(msg) => println!("Error: {}", msg),
}
```

---

## 6. Control Flow: `if let` and `let else`
`match` is great, but sometimes verbose if you only care about one case.

### `if let`
Use when you only care about one variant and ignore the rest.
```rust
let config = Some("dark_mode");

if let Some(theme) = config {
    println!("Loading theme: {}", theme);
}
// does nothing if None
```

### `let else` (Rust 1.65+)
Use when you want to handle the "bad" case immediately and return early. This reduces nesting.
```rust
fn process_input(input: Option<String>) {
    let Some(text) = input else {
        println!("No input provided!");
        return;
    };
    
    // 'text' is now available here directly
    println!("Processing: {}", text);
}
```

---

## 7. Common Collections
Since you know basics, here are the three workhorses of Rust Standard Library.

### Vector (`Vec<T>`)
A growable array.
- **Heap Allocated**: The data lives on the heap.
- **Layout**: Pointer to data, Capacity, Length.
```rust
let mut numbers = Vec::new(); // or vec![1, 2, 3];
numbers.push(10);
numbers.push(20);
let first = numbers[0]; // Access with index
```

### String (`String`)
A growable, UTF-8 encoded text buffer.
- NOT an array of characters. It handles complex unicode.
- Under the hood, it's just a `Vec<u8>`.
```rust
let mut s = String::from("Hello");
s.push_str(", World");
```

### HashMap (`HashMap<K, V>`)
Key-Value store. Fast lookups O(1).
```rust
use std::collections::HashMap;

let mut scores = HashMap::new();
scores.insert("Blue Team", 10);
scores.insert("Red Team", 50);

let blue_score = scores.get("Blue Team"); // Returns Option<&i32>
```
