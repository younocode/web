package web

type Router struct {
	trees map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		trees: make(map[string]Handler),
	}
}

func (r *Router) Add(method string, handler Handler) {
	r.trees[method] = handler
}

func (r *Router) Get(method string) Handler {
	return r.trees[method]
}
