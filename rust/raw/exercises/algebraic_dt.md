# Exercises: Algebraic Data Types & Collections

Solve these exercises using a test-driven approach. Write a test for each exercise.

## Part 1: Basic Enums & Structs (Product & Sum Types)

### Exercise 1: The Shape
1. Define an enum `Shape` with variants:
   - `Circle` (radius: f64)
   - `Rectangle` (width: f64, height: f64)
   - `Square` (side: f64)
2. Implement a method `area(&self) -> f64`.
3. **Test**: Check area calculation for all 3 variants.

### Exercise 2: Direction Navigation
1. Define `Direction` { North, South, East, West }.
2. Define a struct `Point` { x: i32, y: i32 }.
3. Implement `move_point(&mut self, dir: Direction)` on `Point`.
   - North increments Y, South decrements Y, East increments X, West decrements X.
4. **Test**: Start at (0,0), move North then East, assert position is (1,1).

### Exercise 3: Media Player State
1. Define enum `MediaState`:
   - `Playing(String)`: holds track name.
   - `Paused(String)`: holds track name.
   - `Stopped`
2. Implement method `play(&self) -> String` that returns a message like "Playing [track]" or "Resuming [track]" or "Nothing to play".
3. **Test**: Create states and check messages.

### Exercise 4: Log Levels
1. Define `LogLevel` { Info, Warning, Error }.
2. Define `LogMessage` struct { level: LogLevel, msg: String }.
3. Write a function `format_log(log: LogMessage) -> String` that formats it "[INFO] msg" etc.
4. **Test**: Verify string format.

### Exercise 5: Web Event (Data variants)
1. Define `WebEvent`:
   - `PageLoad`
   - `KeyPress(char)`
   - `Click { x: i64, y: i64 }` (struct-like variant)
2. Implement `inspect(&self) -> String` returning a summary string.
3. **Test**: Ensure Click variant summary shows coordinates accurately.

---

## Part 2: Option & Result

### Exercise 6: Safe Division
1. Write `fn safe_div(a: i32, b: i32) -> Option<i32>`.
2. Return `None` if `b` is 0.
3. **Test**: Test 10/2 = Some(5) and 10/0 = None.

### Exercise 7: Username Extractor
1. Write `fn get_username(id: u32) -> Option<String>`.
2. Mock a database: if id is 1 return "Alice", 5 return "Bob", else None.
3. **Test**: Check valid and invalid IDs.

### Exercise 8: File Parser (Result)
1. Write `fn parse_percentage(input: &str) -> Result<u8, String>`.
2. If input is not a number, return Err.
3. If number < 0 or > 100, return Err.
4. Else Ok(n). (Use `input.parse::<i32>()` then check range).
5. **Test**: "50" -> Ok(50), "150" -> Err, "abc" -> Err.

### Exercise 9: Login Validator
1. Write `fn validate_login(username: &str, password: &str) -> Result<(), String>`.
2. Rules: Username valid if not empty. Password valid if len >= 8.
3. Return `Ok(())` on success.
4. **Test**: Check short password returns specific error.

### Exercise 10: Array Element Fetcher
1. Write `fn get_element(arr: &[i32], index: usize) -> Option<i32>`.
2. Use `.get()` or check bounds manually.
3. **Test**: Fetch valid index and out-of-bounds index.

---

## Part 3: Collections + ADTs (The Real World)

### Exercise 11: Student Database (Vec + Struct)
1. Struct `Student` { name: String, grade: u8 }.
2. Function `add_student(db: &mut Vec<Student>, name: String, grade: u8)`.
3. Function `find_student(db: &Vec<Student>, name: &str) -> Option<u8>` (returns grade).
4. **Test**: Add two students, find one existing, search for one non-existing.

### Exercise 12: In-Memory Key-Value Store (HashMap + Enum)
1. `Value` enum { Int(i32), Text(String) }.
2. `Store` struct holding a `HashMap<String, Value>`.
3. Impl methods `set(key, value)` and `get(key) -> Option<&Value>`.
4. **Test**: Store an Int, retrieve it, try to retrieve missing key.

### Exercise 13: Shopping Cart (Vec + Enum + Option)
1. `Item` enum { Book(f64), TShirt(f64), Laptop(f64) }.
2. `get_price(&self) -> f64`.
3. `fn total_price(cart: &[Item]) -> f64`.
4. **Test**: Cart with 1 Book and 1 Laptop, check sum.

### Exercise 14: Command Line Parser (Vec + Result)
1. `Command` enum { Help, Version, Echo(String) }.
2. `fn parse_args(args: Vec<String>) -> Result<Command, String>`.
   - "help" -> Help
   - "version" -> Version
   - "echo <msg>" -> Echo(msg)
   - Other -> Err("Unknown command")
3. **Test**: Parse vec!["echo", "hello"] -> Ok(Command::Echo("hello")).

### Exercise 15: Task Manager (All concepts)
1. `TaskStatus` { Todo, InProgress, Done }.
2. `Task` struct { id: u32, title: String, status: TaskStatus }.
3. Use a `Vec<Task>` as the DB.
4. `fn complete_task(db: &mut Vec<Task>, id: u32) -> Option<&Task>`.
   - Find task by ID. If found, set status to Done and return Some(task). If not, None. (Hint: might need to split into finding index first to satisfy borrow checker, or return bool).
   - Simpler: Return `bool` (true if updated).
5. **Test**: Create 2 tasks, complete 1, verify status changed.
