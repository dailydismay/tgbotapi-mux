package handler

type Middleware[Update any] func(next HandleFn[Update]) HandleFn[Update]

func WithMiddleware[Update any](handler HandleFn[Update], mw ...Middleware[Update]) HandleFn[Update] {
	for i := len(mw) - 1; i >= 0; i-- {
		wrapper := mw[i]
		if wrapper != nil {
			handler = wrapper(handler)
		}
	}

	return handler
}

func CombineMiddleware[Update any](mw []Middleware[Update]) Middleware[Update] {
	return func(next HandleFn[Update]) HandleFn[Update] {
		return WithMiddleware[Update](next, mw...)
	}
}
