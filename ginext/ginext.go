// Package ginext предоставляет расширения для веб-фреймворка Gin.
package ginext

import (
	"github.com/gin-gonic/gin"
)

// Context синоним gin.Context.
type Context = gin.Context

// HandlerFunc синоним gin.HandlerFunc.
type HandlerFunc = gin.HandlerFunc

// H синоним gin.H.
type H = gin.H

// Engine расширяет стандартный Gin Engine.
type Engine struct {
	*gin.Engine
}

// RouterGroup позволяет объединять хэндлеры в группы.
type RouterGroup struct {
	*gin.RouterGroup
}

// New создает новый экземпляр Engine.
func New() *Engine {
	return &Engine{gin.New()}
}

// Run запуск сервера.
func (e *Engine) Run(addr ...string) error {
	return e.Engine.Run(addr...)
}

// Group используется для создания группы роутов.
func (e *Engine) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{e.Engine.Group(relativePath, handlers...)}
}

// Use используется для установки middleware.
func (e *Engine) Use(middleware ...HandlerFunc) {
	e.Engine.Use(middleware...)
}

// Use используется для установки middleware.
func (g *RouterGroup) Use(middleware ...HandlerFunc) {
	g.RouterGroup.Use(middleware...)
}

// LoadHTMLGlob загружает html шаблоны.
func (e *Engine) LoadHTMLGlob(pattern string) {
	e.Engine.LoadHTMLGlob(pattern)
}

// Logger стандартный middleware.
func Logger() HandlerFunc {
	return gin.Logger()
}

// Recovery стандартный middleware.
func Recovery() HandlerFunc {
	return gin.Recovery()
}

// GET обертка для метода GET (*Engine).
func (e *Engine) GET(relativePath string, handlers ...HandlerFunc) {
	e.Engine.GET(relativePath, handlers...)
}

// POST обертка для метода POST (*Engine).
func (e *Engine) POST(relativePath string, handlers ...HandlerFunc) {
	e.Engine.POST(relativePath, handlers...)
}

// DELETE обертка для метода DELETE (*Engine).
func (e *Engine) DELETE(relativePath string, handlers ...HandlerFunc) {
	e.Engine.DELETE(relativePath, handlers...)
}

// PUT обертка для метода PUT (*Engine).
func (e *Engine) PUT(relativePath string, handlers ...HandlerFunc) {
	e.Engine.PUT(relativePath, handlers...)
}

// PATCH обертка для метода PATCH (*Engine).
func (e *Engine) PATCH(relativePath string, handlers ...HandlerFunc) {
	e.Engine.PATCH(relativePath, handlers...)
}

// OPTIONS обертка для метода OPTIONS (*Engine).
func (e *Engine) OPTIONS(relativePath string, handlers ...HandlerFunc) {
	e.Engine.OPTIONS(relativePath, handlers...)
}

// HEAD обертка для метода HEAD (*Engine).
func (e *Engine) HEAD(relativePath string, handlers ...HandlerFunc) {
	e.Engine.HEAD(relativePath, handlers...)
}

// GET обертка для метода GET (*RouterGroup).
func (g *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) {
	g.RouterGroup.GET(relativePath, handlers...)
}

// POST обертка для метода POST (*RouterGroup).
func (g *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) {
	g.RouterGroup.POST(relativePath, handlers...)
}

// DELETE обертка для метода DELETE (*RouterGroup).
func (g *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) {
	g.RouterGroup.DELETE(relativePath, handlers...)
}

// PUT обертка для метода PUT (*RouterGroup).
func (g *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) {
	g.RouterGroup.PUT(relativePath, handlers...)
}

// PATCH обертка для метода PATCH (*RouterGroup).
func (g *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) {
	g.RouterGroup.PATCH(relativePath, handlers...)
}

// OPTIONS обертка для метода OPTIONS (*RouterGroup).
func (g *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) {
	g.RouterGroup.OPTIONS(relativePath, handlers...)
}

// HEAD обертка для метода HEAD (*RouterGroup).
func (g *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) {
	g.RouterGroup.HEAD(relativePath, handlers...)
}
