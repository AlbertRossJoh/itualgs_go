package utilities

type Optional[T any] struct {
	val *T
}

func NewOptional[T any](val T) Optional[T] {
	return Optional[T]{val: &val}
}

// For extracting an optional data type
//
// Usage:
//
//	option := utilities.NewOptional(2)
//	var v int = 0
//	if option.Some(&v) {
//			//dostuff with v
//	}
func (o Optional[T]) Some(val *T) bool {
	if o.val != nil {
		*val = *o.val
		return true
	}
	return false
}

func (o Optional[T]) IsSome() bool {
	return o.val != nil
}

func (o Optional[T]) IsNone() bool {
	return o.val == nil
}
