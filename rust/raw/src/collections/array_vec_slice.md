# Arrays, Slices, Vectors exercises

### Arrays
##### Exercise 1: Temperature Station
Create an array of 5 temperatures. Print first/last elements, find the max and its index, find the min and its index. Convert all to Fahrenheit using .map(). Check if a specific value exists with .contains(). Use .windows(2) to compute differences between consecutive readings. Use .chunks(2) to group into pairs. Sort a mutable copy ascending, then descending.
##### Exercise 2: In-Place Array Rotation
Write a function that rotates a fixed-size array left by k positions without allocating a Vec. Hint: the three-reverse trick works here.


### Vectors
##### Exercise 3: Vec CRUD Drill
Practice every way to create a Vec: Vec::new(), vec![] macro, Vec::from() an array, .collect() from a range, repeat syntax vec![0; n], and Vec::with_capacity(). Then on a single Vec, practice in order: push, pop, insert at index, remove at index, swap_remove, extend, extend_from_slice, truncate, resize with a fill value, clear, retain (keep only even numbers), sort + dedup, drain a range, and split_off. Print the Vec after each operation to see what changed.
##### Exercise 4: Vec Search & Aggregation
Given a Vec of exam scores: use contains, binary_search (sort first!), iter().position() to find the index of the first score above 90, rposition for the last, find to get a reference to it, filter + collect to get all high scores, iter().sum() to compute the average, count to count scores above 80, and all/any to check if all are positive or any are perfect 100.
##### Exercise 5: Stack & Queue from Vec
Implement a generic Stack<T> backed by a Vec with push, pop, peek (use .last()), is_empty, and len. Then implement a Queue<T> with enqueue (push to end), dequeue (remove from front), and peek (use .first()). After it works, think about why dequeue using remove(0) is O(n) and when you'd reach for VecDeque instead.

### Slices
##### Exercise 6: Slice Method Tour
Take a Vec, create a slice of a subrange. Practice: first, last, get (safe indexing), split_at, split_first, split_last, windows(3) with a sum of each window, chunks(3), iter().zip() to pair two slices together, starts_with, ends_with, copy_from_slice into a mutable destination, fill, and chaining skip + take + step_by on an iterator.
##### Exercise 7: Generic Slice Functions
Write these functions accepting &[T] so they work on both arrays and Vecs:

sum(slice: &[i32]) -> i32
mean(slice: &[f64]) -> f64
second_largest(slice: &[i32]) -> Option<i32> — use .to_vec() to avoid mutating the original
is_sorted(slice: &[i32]) -> bool — use .windows(2)
flatten(slices: &[&[i32]]) -> Vec<i32> — use .flat_map()

Test each with an array, a Vec, and a sub-slice to prove they're generic.
##### Exercise 8: Mutable Slice Algorithms
Implement these operating on &mut [i32]:

partition(slice, pivot) -> usize — move all elements less than pivot to the left using .swap(), return the partition index
unique_len(slice) -> usize — given a sorted mutable slice, remove duplicates in-place and return the new logical length. Caller uses &slice[..returned_len] to see the result.