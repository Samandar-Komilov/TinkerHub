# HashMap, HashSet, VecDeque exercises

### HashMap

##### Exercise 1: Word Frequency Counter
Given a `&str` paragraph, count how many times each word appears. Use `HashMap::new()`, `.entry().or_insert()` to build the map. Then: find the most frequent word using `.iter().max_by_key()`, print all words that appear exactly once using `.iter().filter()`, get the total unique word count with `.len()`, check if a word exists with `.contains_key()`, and remove a specific word with `.remove()`. Finally, `.drain()` the map and collect into a `Vec<(String, usize)>` sorted by count descending.
Hint: `.split_whitespace()` and `.to_lowercase()` for splitting/normalizing words.

##### Exercise 2: Two-Map Merge & Conflict Resolution
Create two `HashMap<String, i32>` representing item prices from two stores. Merge them into a third map keeping the lower price for each item. Practice: `.get()` vs direct indexing `[]`, `.or_insert()` vs `.or_insert_with()`, `.keys()` and `.values()` iterators, `.iter().map()` to apply a 10% discount to all values and `.collect()` into a new map. Use `.retain()` to keep only items under a price threshold. Use `.extend()` to add a batch of new entries.
Hint: `.entry()` API is key here; for merge logic, match on the entry variant.

##### Exercise 3: Grouped Aggregation
Given a `Vec<(&str, &str, u32)>` of `(department, employee, salary)` tuples, build a `HashMap<String, Vec<(String, u32)>>` grouping employees by department. Then compute per-department stats: average salary, highest earner, headcount. Use `.entry().or_insert_with(Vec::new)` to build groups. Convert the final result to a `Vec` of tuples and sort by average salary descending.
Hint: `.or_default()` is shorthand for `.or_insert_with(Default::default)`.

##### Exercise 4: Inventory System
Build a `HashMap<String, (u32, f64)>` mapping item name to `(quantity, price)`. Implement these operations as functions:
- `add_stock(map, item, qty)` -- insert or increase quantity
- `sell(map, item, qty) -> Result<f64, String>` -- decrease quantity, return total cost, error if insufficient stock
- `total_value(map) -> f64` -- sum of `qty * price` for all items
- `low_stock(map, threshold) -> Vec<&str>` -- items below threshold
- `remove_discontinued(map, items: &[&str])` -- remove multiple items

Test by simulating a sequence of stock additions, sales, and queries.
Hint: use `.get_mut()` to modify values in place.

### HashSet

##### Exercise 5: Set Operations Drill
Create two `HashSet<i32>` from arrays. Practice every set operation: `.union()`, `.intersection()`, `.difference()`, `.symmetric_difference()`, `.is_subset()`, `.is_superset()`, `.is_disjoint()`. Collect each result into a new sorted `Vec` for easy assertion. Then: `.insert()` a value, `.remove()` a value, `.take()` to remove and return, `.contains()` to check membership, `.len()` and `.is_empty()`.
Hint: set operation methods return iterators -- you need `.collect()` to get a new set or vec.

##### Exercise 6: Duplicate Detector & Unique Filter
Write a function `first_duplicate(slice: &[i32]) -> Option<i32>` that returns the first element that appears more than once, using a `HashSet` to track seen values. Then write `unique_ordered(slice: &[i32]) -> Vec<i32>` that returns elements in original order with duplicates removed (not the same as dedup -- it should handle non-adjacent duplicates). Finally write `find_missing(nums: &[i32], range: std::ops::RangeInclusive<i32>) -> Vec<i32>` that finds all numbers in the range not present in the slice.
Hint: `.insert()` returns `false` if the value already existed.

##### Exercise 7: Tag System
Build a `HashMap<String, HashSet<String>>` mapping article titles to their tags. Implement:
- `add_tags(map, article, tags: &[&str])` -- add multiple tags to an article
- `remove_tag(map, article, tag)` -- remove one tag
- `articles_with_tag(map, tag) -> Vec<&str>` -- all articles that have a specific tag
- `common_tags(map, article1, article2) -> HashSet<String>` -- intersection of two articles' tags
- `all_tags(map) -> HashSet<&str>` -- every unique tag across all articles

Hint: iterate `.values()` and use `.flat_map()` for `all_tags`.

### VecDeque

##### Exercise 8: VecDeque Method Tour
Create a `VecDeque<i32>` using `VecDeque::new()`, `VecDeque::from(vec)`, and `VecDeque::with_capacity()`. Practice both ends: `push_back`, `push_front`, `pop_back`, `pop_front`, `front()`, `back()`. Then: `insert(index, val)`, `remove(index)`, `swap(i, j)`, `contains()`, `len()`, `get(index)`, `iter()`, `make_contiguous()`, `rotate_left(n)`, `rotate_right(n)`, `truncate()`, `drain(range)`, `retain()`, `split_off(at)`. Print after each operation.
Hint: VecDeque is a ring buffer -- elements may not be contiguous in memory. `make_contiguous()` fixes that and returns a `&mut [T]`.

##### Exercise 9: Sliding Window Maximum
Given a `Vec<i32>` and a window size `k`, find the maximum value in every contiguous window of size `k`. Return a `Vec<i32>` of these maxima. First implement the naive O(n*k) version using `windows(k)`. Then implement an O(n) version using a `VecDeque` as a monotonic deque: store indices, and for each new element, pop from the back while the new element is larger, push the new index to the back, pop from the front if the index is out of window range, and `front()` is always the current max's index.
Hint: the deque stores indices, not values. Compare `nums[*deque.back().unwrap()]` with `nums[i]`.

##### Exercise 10: Recent Events Buffer
Build a `struct RecentBuffer<T> { buf: VecDeque<T>, capacity: usize }` that keeps only the N most recent items. Implement:
- `new(capacity)` -- create with max size
- `push(item)` -- add to back, if full `pop_front` first
- `recent(n) -> Vec<&T>` -- last n items (or fewer if not enough)
- `oldest() -> Option<&T>` -- peek front
- `newest() -> Option<&T>` -- peek back
- `contains(&item) -> bool` where `T: PartialEq`
- `clear()`
- `iter()` -- iterate oldest to newest

Test with a stream of log messages, verifying the buffer never exceeds capacity.
Hint: `VecDeque` is perfect for this because both `push_back` and `pop_front` are O(1).

##### Exercise 11: Task Scheduler (Round-Robin)
Simulate a round-robin task scheduler using `VecDeque<(String, u32)>` where each tuple is `(task_name, remaining_time)`. Each tick: `pop_front` a task, subtract 1 from its remaining time, print progress, and if time remains `push_back` to the end. Continue until the deque is empty. Track completion order in a `Vec<String>`. Add a `priority_insert(deque, task)` function that inserts high-priority tasks at the front with `push_front`.
Hint: this is literally how OS schedulers work -- VecDeque models a circular queue naturally.
