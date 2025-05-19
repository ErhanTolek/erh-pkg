package sort

// Sort is a generic struct for sorting slices of maps.
// It uses a less function to determine the order of elements.
// K is the type of the key in the map, and V is the type of the value.
type Sort[K comparable, V comparable] struct {
	Less func(i, j V) bool
}

type SortingClient[K comparable, V comparable] interface {
	BubbleSort(arr []map[K]V, key K, less func(a, b V) bool) []map[K]V
}

// NewSort creates a new Sort instance with the provided less function.
// Less function should decide the order of elements.
func NewSort[K comparable, V comparable](less func(i, j V) bool) *Sort[K, V] {
	return &Sort[K, V]{
		Less: less,
	}
}

func (sort Sort[K, V])BubbleSort(arr []map[K]V, key K, less func(a, b V) bool) []map[K]V {
	n := len(arr)
	sorted := make([]map[K]V, n)
	copy(sorted, arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			v1, ok1 := sorted[j][key]
			v2, ok2 := sorted[j+1][key]

			if !ok1 || !ok2 {
				continue
			}

			if less(v2, v1) {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	return sorted
}