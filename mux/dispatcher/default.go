package dispatcher

import (
	"github.com/dailydismay/tgbotapi-mux/mux/filter"
	"github.com/dailydismay/tgbotapi-mux/mux/handler"
)

type updateDispatcher[Update any] struct {
	handlers []handler.Handler[Update]
}

func (ud *updateDispatcher[U]) Register(f filter.Filter[U], h handler.HandleFn[U]) {
	ud.handlers = append(ud.handlers, handler.NewHandler(h, f))
}

func (ud *updateDispatcher[U]) Dispatch(u U) error {
	for _, hm := range ud.handlers {
		if !hm.Matcher()(u) {
			continue
		}

		if err := hm.Handler()(u); err != nil {
			return err
		}
		return nil
	}

	return nil
}

// default dispatcher is linear and very primitive
// you can implement more effective dispatcher following Dispatcher interface
func NewDefault[Update any]() Dispatcher[Update] {
	return &updateDispatcher[Update]{}
}
