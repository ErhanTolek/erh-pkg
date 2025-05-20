package sort

// Sort is a generic struct for sorting slices of maps.
// It uses a less function to determine the order of elements.
// K is the type of the key in the map, and V is the type of the value.
type Sort[T any] struct {
	Less func(a, b T) bool
}

func NewSort[T any](less func(a, b T) bool) *Sort[T] {
	return &Sort[T]{Less: less}
}

func (s *Sort[T]) BubbleSort(arr []T) []T {
	n := len(arr)
	sorted := make([]T, n)
	copy(sorted, arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if s.Less(sorted[j+1], sorted[j]) {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	return sorted
}
