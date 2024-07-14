package pattern

type Handler interface {
	Handle(code int) error
	SetNext(Handler)
}

type BaseHandler struct {
	next *Handler
}

func (base *BaseHandler) SetNext(h *Handler) {
	base.next = h
}

func (base *BaseHandler) Handle(code int) error { return nil }

type ConcreteHandlerOne struct{}

func (ConcreteHandlerOne) Handle() {

}
