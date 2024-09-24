package main

type Stream[T Displayable] struct {
	items []T
}

func CreateStream[T Displayable](items []T) Stream[T] {
	return Stream[T]{items}
}

func (s Stream[T]) Filter(predicate func(T) bool) Stream[T] {
	var result []T
	for _, item := range s.items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return Stream[T]{result}
}

func (s Stream[T]) Map(transform func(T) T) Stream[T] {
	var result []T
	for _, item := range s.items {
		result = append(result, transform(item))
	}
	return Stream[T]{result}
}

func (s Stream[T]) Max(compare func(T, T) bool) T {
	maxItem := s.items[0]
	for _, item := range s.items[1:] {
		if compare(item, maxItem) {
			maxItem = item
		}
	}
	return maxItem
}

func (s Stream[T]) Reduce(accumulator func(T, T) T) T {
	result := s.items[0]
	for _, item := range s.items[1:] {
		result = accumulator(result, item)
	}
	return result
}

func (s Stream[T]) Distinct(keyExtractor func(T) string) Stream[T] {
	seen := make(map[string]bool)
	var result []T
	for _, item := range s.items {
		key := keyExtractor(item)
		if !seen[key] {
			seen[key] = true
			result = append(result, item)
		}
	}
	return Stream[T]{result}
}

func (s Stream[T]) Display() {
	for _, item := range s.items {
		println(item.Display())
	}
}
