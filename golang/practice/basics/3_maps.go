package basics

import (
	"fmt"
	"maps"
	"slices"
)

type Set map[string]struct{}

func Main_maps() {
	// 1, 2, 3, 4
	mp1 := map[string]int{}
	mp1["one"] = 1
	mp1["two"] = 2
	delete(mp1, "two")
	mp1["three"] = 3
	fmt.Println(mp1)

	// 5: Comma-OK idiom
	v, ok := mp1["three"]
	fmt.Println("Value:", v, "Exists:", ok)
	// 6: Iterating over a map
	for k, v := range mp1 {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}
	// 7
	defaultVal := 199
	val, exists := mp1["four"]
	if !exists {
		val = defaultVal
	}
	fmt.Println("Value for 'four':", val)
	// 8
	s := "This is a test string! This string is actually string."
	num := count_words(s)
	fmt.Println(num)
	// 9
	mp2 := map[string]int{"one": 1, "hundred": 100, "thousand": 1000}
	mp2 = merge_2_maps(mp1, mp2)
	fmt.Println(mp2)
	// 10
	is_equal := maps.Equal(mp1, mp2)
	fmt.Println(is_equal, mp1, mp2)
	// 11
	// 12: clear all entries of the map
	// for i := range mp1 {
	// 	delete(mp1, i)
	// }
	mp1 = map[string]int{} // garbage collector collects the garbage
	fmt.Println(mp1)
	// 13
	s1 := make(Set)
	s1.Add("hello")
	s1.Add("world")
	fmt.Println(s1, s1.Contains("hi"))

	// 14: Contains duplicates (for + Contains)
	slc1 := []int{1, 2, 2, 3, 4, 4}
	any_duplicates := contains_duplicates(slc1)
	fmt.Println(any_duplicates)

	// 15: intersection of slices: (for + contains)
	slc2 := []int{6, 7, 8, 1}
	intersected_slc := slice_intersection(slc1, slc2)
	fmt.Println(intersected_slc)

	// 16: Group strings in a slice by their length using a map.
	slc3 := []string{"hello", "hi", "ah", "ohh", "whats"}
	grouped_map := group_strings(slc3)
	fmt.Println(grouped_map)

	// 17: Convert a map into a slice of keys.
	fmt.Println(slices.Collect(maps.Keys(grouped_map)))
	// 18: Convert a map into a slice of values.
	fmt.Println(slices.Collect(maps.Values(grouped_map)))
	// 19: Sort the keys of a map and print them in order.
	slices.Sort(slices.Collect(maps.Keys(grouped_map)))
	fmt.Println(grouped_map)
	// 20: Write a function that inverts a map (keys become values and vice versa).
	// for with a new map, simple
}

func count_words(s string) string {
	// 8
	return s
}

func merge_2_maps(m1 map[string]int, m2 map[string]int) map[string]int {
	// for i, v := range m2 {
	// 	m1[i] = v
	// }
	maps.Copy(m1, m2)

	return m1
}

func (s Set) Add(element string) {
	s[element] = struct{}{}
}

func (s Set) Remove(element string) {
	delete(s, element)
}

func (s Set) Contains(element string) bool {
	_, ok := s[element]
	return ok
}

func contains_duplicates(slice []int) bool {
	mp := make(map[int]int)

	for _, e := range slice {
		if mp[e] > 0 {
			return true
		}
		mp[e] += 1
	}

	return false
}

func slice_intersection(slc1 []int, slc2 []int) []int {
	mp := make(map[int]int)
	var res []int

	for _, e := range slc1 {
		if mp[e] == 0 {
			mp[e] += 1
		}
	}

	for _, e := range slc2 {
		mp[e] += 1
	}

	for key, val := range mp {
		if val == 2 {
			res = append(res, key)
		}
	}

	return res
}

func group_strings(slc []string) map[int][]string {
	mp := make(map[int][]string)

	for _, e := range slc {
		ln := len(e)
		mp[ln] = append(mp[ln], e)
	}

	return mp
}
