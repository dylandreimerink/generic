package generic

// FilterMap returns a new map, only containing kv-pairs for which `filter` is true
func FilterMap[K comparable, V any](m map[K]V, filter func(k K, v V) (keep bool)) map[K]V {
	new := make(map[K]V, len(m))
	for k, v := range m {
		if filter(k, v) {
			new[k] = v
		}
	}
	return new
}

// FilterMapKeyInPlace filters the contents of the given map in-place.
// This function changes the existing map object.
func FilterMapInPlace[K comparable, V any](m map[K]V, filter func(k K, v V) (keep bool)) {
	for k, v := range m {
		if !filter(k, v) {
			delete(m, k)
		}
	}
}

// MapContains returns true if the map contains the given value
func MapContains[K comparable, V comparable](m map[K]V, value V) bool {
	_, found := MapFindKey(m, value)
	return found
}

// MapFindKey returns the key for the first instance of `value` we were able to find.
func MapFindKey[K comparable, V comparable](m map[K]V, value V) (K, bool) {
	for k, v := range m {
		if v == value {
			return k, true
		}
	}
	var ret K
	return ret, false
}
