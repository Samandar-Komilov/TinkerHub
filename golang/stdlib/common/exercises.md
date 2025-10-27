# Common stdlib package Exercises

## 1. `cmp` — Generic Comparisons

*(Go ≥1.21)*

**Core Functions:**
`cmp.Compare`, `cmp.Less`, `cmp.Or`, `cmp.Ordered` constraints.

### Exercises

1. Implement a generic function `Max[T cmp.Ordered](a, b T) T` using `cmp.Compare`.
2. Sort an integer slice using `sort.Slice` + `cmp.Less`.
3. Write a generic `Between[T cmp.Ordered](x, low, high T) bool` returning true if `low ≤ x ≤ high`.
4. Compare two strings lexicographically using `cmp.Compare`.
5. Chain comparisons: `cmp.Or(cmp.Compare(len(s1), len(s2)), cmp.Compare(s1, s2))` to sort by length then lexicographically.
6. Implement a min-heap comparator using `cmp.Less`.
7. Use `cmp.Compare` with custom structs via a key field (e.g. `Person.Age`).

---

## 2. `bytes` — Immutable string-like operations for byte slices

**Core Concepts:** slice-based string ops, builders, buffers, readers.

### Exercises

1. Check if a `[]byte` starts or ends with another (`bytes.HasPrefix`, `bytes.HasSuffix`).
2. Split byte data on newlines (`bytes.Split`, `bytes.Fields`, `bytes.FieldsFunc`).
3. Join multiple byte slices with a delimiter (`bytes.Join`).
4. Replace substrings in `[]byte` (`bytes.ReplaceAll`).
5. Find index of a pattern (`bytes.Index`, `bytes.LastIndex`, `bytes.Contains`).
6. Compare slices (`bytes.EqualFold` case-insensitive).
7. Trim space and custom runes (`bytes.TrimSpace`, `bytes.Trim`, `bytes.TrimFunc`).
8. Convert between `string` and `[]byte` safely.
9. Use `bytes.NewReader` to create an `io.Reader` and read it line by line.
10. Copy a reader to stdout using `io.Copy(os.Stdout, bytes.NewReader(...))`.
11. Use `bytes.Buffer` for accumulating chunks (simulate logging).
12. Reset, Grow, Truncate a `bytes.Buffer` — observe behavior.
13. Serialize integers to a buffer (`fmt.Fprintf(&buf, "%d", x)`).
14. Compare the performance of `+` string concatenation vs. `bytes.Buffer`.
15. Create your own in-memory file abstraction using `bytes.Buffer`.

---

## 3. `strings` — High-level text manipulation

**Core Concepts:** immutable, rune-safe string utilities.

### Exercises

1. Search substring (`strings.Contains`, `strings.Index`).
2. Replace and count replacements (`strings.Replace`, `strings.ReplaceAll`).
3. Split/Join on different delimiters (`strings.Split`, `strings.Join`).
4. Trim and Pad (`strings.Trim`, `TrimSpace`, `TrimSuffix`, `TrimPrefix`).
5. Case transformations (`ToUpper`, `ToLower`, `Title`, `EqualFold`).
6. Builder-based concatenation (`strings.Builder`).
7. Count words, letters, vowels manually using `strings.Count` and loops.
8. Map transformations (`strings.Map`) to shift letters (Caesar cipher).
9. Use `strings.Cut` and `CutPrefix` for safe splitting.
10. Build a CSV parser using `strings.FieldsFunc`.
11. Implement substring search manually and verify with `strings.Contains`.
12. Benchmark `strings.Builder` vs `+` for large string creation.
13. Parse query-like strings `"key1=val1&key2=val2"` using `strings.Split`.
14. Implement string reversal using rune slice.

---

## 4. `slices` — Generic slice operations (Go ≥1.21)

**Core Functions:** `slices.Equal`, `slices.Clone`, `slices.Sort`, `slices.Delete`, `slices.Compact`, etc.

### Exercises

1. Check equality of two slices with `slices.Equal`.
2. Clone a slice and verify independence.
3. Sort ascending/descending using `slices.SortFunc` + `cmp.Compare`.
4. Delete element by index using `slices.Delete`.
5. Compact duplicates (`slices.Compact`).
6. Insert at arbitrary index using `slices.Insert`.
7. Check containment manually and using `slices.Index`.
8. Reverse in place (`slices.Reverse`).
9. Compare slices lexicographically using `slices.Compare`.
10. Use `slices.BinarySearch` on a sorted list.
11. Write generic `Min` and `Max` using `slices.Min`, `slices.Max`.
12. Deduplicate unsorted list manually using map and verify with `slices.Compact`.
13. Partition a slice into chunks of N.
14. Rotate a slice left/right by N positions.
15. Benchmark performance between `append` and `slices.Insert`.

---

## 5. `maps` — Generic map utilities (Go ≥1.21)

**Core Functions:** `maps.Copy`, `maps.Clone`, `maps.Equal`, `maps.Keys`, `maps.Values`, `maps.DeleteFunc`.

### Exercises

1. Merge two maps using `maps.Copy`.
2. Deep clone a map using `maps.Clone` and test mutability.
3. Extract all keys and values (`maps.Keys`, `maps.Values`) and sort them.
4. Compare two maps for equality with `maps.Equal`.
5. Delete entries based on predicate using `maps.DeleteFunc`.
6. Build frequency counter of words using `map[string]int`.
7. Invert a map (value→key).
8. Build a small cache system with expiration times.
9. Use `map[string][]string` to group names by first letter.
10. Benchmark map lookup vs. slice search.

---

## 6. `structs` — Reflection and Deep Equality (no stdlib “structs” package)

**Conceptual Coverage:**

* Definition
* Tag usage
* Anonymous embedding
* Field iteration via `reflect`
* Comparison via `reflect.DeepEqual`

### Exercises

1. Define a struct with nested fields; initialize using composite literal.
2. Access fields via pointers and values.
3. Embed another struct and demonstrate method promotion.
4. Tag fields with `json:"name"` and print tags via reflection.
5. Iterate all fields dynamically using `reflect.TypeOf(...).NumField()`.
6. Compare two struct instances manually and using `reflect.DeepEqual`.
7. Create a slice of structs and sort by one field using `slices.SortFunc`.
8. Marshal/unmarshal struct to JSON (using `encoding/json`).
9. Implement a `String()` method for pretty printing.
10. Add a method with pointer receiver and another with value receiver; show differences.

---

## 7. `errors` — Error Creation, Wrapping, and Inspection

**Core Functions:**
`errors.New`, `fmt.Errorf` with `%w`, `errors.Is`, `errors.As`, `errors.Join`.

### Exercises

1. Create a simple error using `errors.New`.
2. Return it from a function and handle via `if err != nil`.
3. Wrap an error using `fmt.Errorf("context: %w", err)` and inspect with `errors.Is`.
4. Extract specific error types using `errors.As`.
5. Combine multiple errors using `errors.Join` and print them.
6. Create custom error types implementing `Error() string`.
7. Build a function that retries on specific error type using `errors.Is`.
8. Simulate file parsing where each stage may fail with different wrapped errors.
9. Log stack-like context through nested `%w` wrapping.
10. Benchmark `errors.Join` vs sequential checks.
11. Implement a sentinel error (like `ErrNotFound`) and check for it.
12. Demonstrate unwrapping chain using `errors.Unwrap` in a loop.

---

## Execution Sequence

1. `cmp` → foundational for generics.
2. `bytes` and `strings` → text/data manipulation.
3. `slices`, `maps` → collection utilities.
4. `structs` → complex data modeling and reflection.
5. `errors` → robust control flow.
