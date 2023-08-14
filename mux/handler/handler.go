package handler

import (
	"github.com/dailydismay/tgbotapi-mux/mux/filter"
)

type HandleFn[Update any] func(u Update) error

type handler[Update any] struct {
	handler HandleFn[Update]
	filter  filter.Filter[Update]
}

func (h *handler[Update]) Matcher() filter.Filter[Update] {
	return h.filter
}

func (h *handler[Update]) Handler() HandleFn[Update] {
	return h.handler
}

type Handler[Update any] interface {
	Matcher() filter.Filter[Update]
	Handler() HandleFn[Update]
}

func NewHandler[Update any](handleFn HandleFn[Update], matcher filter.Filter[Update]) Handler[Update] {
	return &handler[Update]{handleFn, matcher}
}
