package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// 1. Basic map declaration
	var emptyMap map[string]int
	fmt.Println("Empty map (nil):", emptyMap, "Length:", len(emptyMap))

	// 2. Map initialization using make
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 80
	fmt.Println("Scores map:", scores)

	// 3. Map literal initialization
	people := map[string]int{
		"John":  30,
		"Maria": 25,
		"Peter": 40,
	}
	fmt.Println("People map:", people)

	// 4. Accessing map elements
	aliceScore := scores["Alice"]
	fmt.Println("Alice's score:", aliceScore)

	// 5. Accessing non-existing key (returns zero value)
	charlieScore := scores["Charlie"]
	fmt.Println("Charlie's score (doesn't exist):", charlieScore)

	// 6. Checking if a key exists
	bobScore, exists := scores["Bob"]
	if exists {
		fmt.Println("Bob's score:", bobScore)
	} else {
		fmt.Println("Bob is not in the scores map")
	}

	// 7. Adding and updating elements
	scores["Charlie"] = 85 // Add new key-value
	scores["Alice"] = 97   // Update existing value
	fmt.Println("Updated scores:", scores)

	// 8. Deleting elements
	delete(scores, "Bob")
	fmt.Println("After deleting Bob:", scores)

	// 9. Iterating over a map (order is not guaranteed)
	fmt.Println("Iterating over people:")
	for name, age := range people {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// 10. Iterating over keys only
	fmt.Println("Names of people:")
	for name := range people {
		fmt.Println(name)
	}

	// 11. Map length
	fmt.Println("Number of people:", len(people))

	// 12. Maps with more complex types
	studentScores := map[string]map[string]int{
		"Alice": {
			"Math":    95,
			"Science": 92,
			"English": 88,
		},
		"Bob": {
			"Math":    80,
			"Science": 85,
			"English": 90,
		},
	}
	fmt.Println("Alice's Math score:", studentScores["Alice"]["Math"])

	// 13. Maps with slices as values
	interests := map[string][]string{
		"Alice": {"Reading", "Swimming", "Coding"},
		"Bob":   {"Gaming", "Football", "Movies"},
	}
	fmt.Println("Alice's interests:", interests["Alice"])

	// Using helper functions
	fmt.Println("\n--- Helper Functions ---")

	// 14. Copying a map
	scoresCopy := copyMap(scores)
	scoresCopy["Alice"] = 100 // Modifying copy won't affect original
	fmt.Println("Original scores:", scores)
	fmt.Println("Scores copy:", scoresCopy)

	// 15. Merging maps
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "c": 4}
	merged := mergeMaps(map1, map2)
	fmt.Println("Merged map:", merged)

	// 16. Getting keys
	peopleKeys := getKeys(people)
	sort.Strings(peopleKeys) // Sort the keys
	fmt.Println("Sorted people keys:", peopleKeys)

	// 17. Getting values
	peopleValues := getValues(people)
	sort.Ints(peopleValues) // Sort the values
	fmt.Println("Sorted people values:", peopleValues)

	// 18. Finding keys by value
	ages := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 30}
	keysWithAge30 := findKeysByValue(ages, 30)
	fmt.Println("People who are 30:", keysWithAge30)

	// 19. Filtering maps
	adults := filterMap(ages, func(k string, v int) bool {
		return v >= 30
	})
	fmt.Println("Adults (30 or older):", adults)

	// 20. Check if maps are equal
	map3 := map[string]int{"a": 1, "b": 2}
	map4 := map[string]int{"b": 2, "a": 1}
	fmt.Println("Maps are equal:", mapsEqual(map3, map4))

	// 21. Map differences
	added, removed, modified := mapDiff(map1, map2)
	fmt.Println("Keys added:", added)
	fmt.Println("Keys removed:", removed)
	fmt.Println("Keys modified:", modified)

	// 22. Count word frequencies
	text := "the quick brown fox jumps over the lazy dog"
	wordFreq := wordFrequency(text)
	fmt.Println("Word frequencies:", wordFreq)

	// 23. Inverting a map (values become keys)
	inverted := invertMap(map[string]int{"one": 1, "two": 2, "three": 3})
	fmt.Println("Inverted map:", inverted)
}

// Helper Functions

// copyMap creates a new copy of the provided map
func copyMap[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// mergeMaps combines two maps (values from m2 override m1 if keys collide)
func mergeMaps[K comparable, V any](m1, m2 map[K]V) map[K]V {
	result := copyMap(m1)
	for k, v := range m2 {
		result[k] = v
	}
	return result
}

// getKeys returns all keys from the map as a slice
func getKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// getValues returns all values from the map as a slice
func getValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// findKeysByValue returns all keys that have the given value
func findKeysByValue[K comparable, V comparable](m map[K]V, val V) []K {
	var keys []K
	for k, v := range m {
		if v == val {
			keys = append(keys, k)
		}
	}
	return keys
}

// filterMap creates a new map with only elements that satisfy the predicate
func filterMap[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// mapsEqual checks if two maps have the same key-value pairs
func mapsEqual[K, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

// mapDiff compares two maps and returns added, removed, and modified keys
func mapDiff[K comparable, V comparable](old, new map[K]V) (added []K, removed []K, modified []K) {
	// Find added and modified keys
	for k, newVal := range new {
		if oldVal, exists := old[k]; !exists {
			added = append(added, k)
		} else if oldVal != newVal {
			modified = append(modified, k)
		}
	}

	// Find removed keys
	for k := range old {
		if _, exists := new[k]; !exists {
			removed = append(removed, k)
		}
	}

	return added, removed, modified
}

// wordFrequency counts occurrences of each word in a text
func wordFrequency(text string) map[string]int {
	words := strings.Fields(text)
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}
	return frequency
}

// invertMap swaps keys and values (note: if multiple keys have the same value,
// only one will be preserved in the result)
func invertMap[K comparable, V comparable](m map[K]V) map[V]K {
	result := make(map[V]K)
	for k, v := range m {
		result[v] = k
	}
	return result
}
