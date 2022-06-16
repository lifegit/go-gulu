package koa

import "math"

// koa洋葱模型，gin中间件思想也如同。
// 感谢：https://blog.csdn.net/raoxiaoya/article/details/109444890

const abortIndex int8 = math.MaxInt8 / 2

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

// HandlersChain defines a HandlerFunc array.
type HandlersChain []HandlerFunc

// Context Data
type Data struct {
	Err  error       `json:"err,omitempty"`
	Data interface{} `json:"data"`
}

// Context
type Context struct {
	Result   Data
	RunCount int

	handlers HandlersChain
	index    int8
}

// New
func NewContext() *Context {
	return &Context{}
}

// Next should be used only inside middleware.
// It executes the pending handlers in the chain inside the calling handler.
// See example in GitHub.
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// Run from the first position of the context
func (c *Context) Run() *Context {
	if len(c.handlers) <= 0 {
		panic("no exist handlers")
	}

	// init
	c.index = 0
	c.RunCount++
	c.Result = Data{}

	// run
	c.handlers[c.index](c)
	c.Next()

	return c
}

// Use attaches a global middleware to the router. ie. the middleware attached though Use() will be
// For example, this is the right place for a logger or error management middleware.
func (c *Context) Use(middleware ...HandlerFunc) *Context {
	c.handlers = append(c.handlers, middleware...)

	// limit the number of handlers
	if len(c.handlers) >= int(abortIndex) {
		panic("too many handlers")
	}

	return c
}

// Abort prevents pending handlers from being called. Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers
// for this request are not called.
func (c *Context) Abort() {
	c.index = abortIndex
}

// IsAborted returns true if the current context was aborted.
func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

// Error attaches an error to the current context. The error is pushed to a list of errors.
// It's a good idea to call Error for each error that occurred during the resolution of a request.
// A middleware can be used to collect all the errors and push them to a database together,
// print a log, or append it in the HTTP response.
// Error will panic if err is nil.
func (c *Context) Error(err error) {
	c.Result.Err = err
}
