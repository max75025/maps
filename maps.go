package maps

//based on https://cs.opensource.google/go/x/exp/+/master:maps/maps.go;drc=79cabaa25d7518588d46eb676385c8dff49670c3

// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// KeysFunc is like Keys, but add only values that func isAppend true
func KeysFunc[M ~map[K]V, K comparable, V any](m M, isAppend func(K) bool) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		if isAppend(k) {
			r = append(r, k)
		}

	}
	return r
}

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// ValuesFunc is like Values, but add only values that func isAppend true
func ValuesFunc[M ~map[K]V, K comparable, V any](m M, isAppend func(V) bool) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		if isAppend(v) {
			r = append(r, v)
		}

	}
	return r
}
