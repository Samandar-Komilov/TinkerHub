package basics

import (
	"fmt"
	"maps"
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
	// 15: intersection of slices: (for + contains)
	// 16:
	// 17:
	// 18:
	// 19:
	// 20:
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
