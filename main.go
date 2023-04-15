package ptr

// P returns pointer to value
func P[T any](o T) (r *T) {
	return &o
}

// V returns value of pointer.
// If nil, returns zero-value of givern type
func V[T any](o *T) (r T) {
	if o != nil {
		r = *o
	}
	return r
}
