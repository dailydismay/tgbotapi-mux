package dispatcher

import (
	"github.com/dailydismay/tgbotapi-mux/mux/filter"
	"github.com/dailydismay/tgbotapi-mux/mux/handler"
)

type Dispatcher[Update any] interface {
	Register(um filter.Filter[Update], h handler.HandleFn[Update])
	Dispatch(u Update) error
}
