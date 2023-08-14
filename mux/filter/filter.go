package filter

type Filter[Update any] func(u Update) bool

func Chain[Update any](matchers ...Filter[Update]) Filter[Update] {
	return func(u Update) bool {
		for _, m := range matchers {
			if !m(u) {
				return false
			}
		}

		return true
	}
}
