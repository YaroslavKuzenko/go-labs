package main

// Generic структура Stream
type Stream[T Displayable] struct {
	items []T
}

// Функція для створення Stream
func CreateStream[T Displayable](items []T) Stream[T] {
	return Stream[T]{items}
}

// Функція для фільтрації
func (s Stream[T]) Filter(predicate func(T) bool) Stream[T] {
	var result []T
	for _, item := range s.items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return Stream[T]{result}
}

// Функція для перетворення (Map)
func (s Stream[T]) Map(transform func(T) T) Stream[T] {
	var result []T
	for _, item := range s.items {
		result = append(result, transform(item))
	}
	return Stream[T]{result}
}

// Функція для відображення елементів
func (s Stream[T]) Display() {
	for _, item := range s.items {
		println(item.Display())
	}
}
