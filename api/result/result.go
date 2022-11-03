package result

// Result represents HTTP response body.
type Result[T any] struct {
	Code int    `json:"code"` // 状态码 0 为成功，负数为出错
	Msg  string `json:"msg"`  // message 消息
	Data T      `json:"data"` // 数据
}

func New[T any](i ...T) *Result[T] {
	if len(i) == 0 {
		return &Result[T]{}
	}
	return &Result[T]{Data: i[0]}
}

func With[T any](code int, msg string) *Result[T] { return &Result[T]{Code: code, Msg: msg} }

func BadRequest() *Result[any]           { return With[any](-400, "BadRequest") }
func Forbidden() *Result[any]            { return With[any](-401, "Forbidden") }
func NotFound() *Result[any]             { return With[any](-404, "NotFound") }
func ResourceConflict() *Result[any]     { return With[any](-409, "ResourceConflict") }
func UnsupportedMediaType() *Result[any] { return With[any](-415, "UnsupportedMediaType") }
func InternalServerError() *Result[any]  { return With[any](-500, "InternalServerError") }

func Status400() *Result[any] { return BadRequest() }
func Status401() *Result[any] { return Forbidden() }
func Status404() *Result[any] { return NotFound() }
func Status409() *Result[any] { return ResourceConflict() }
func Status415() *Result[any] { return UnsupportedMediaType() }
func Status500() *Result[any] { return InternalServerError() }

func (r *Result[T]) BadRequest() *Result[T]           { return r.Set(-400, "BadRequest") }
func (r *Result[T]) Forbidden() *Result[T]            { return r.Set(-401, "Forbidden") }
func (r *Result[T]) NotFound() *Result[T]             { return r.Set(-404, "NotFound") }
func (r *Result[T]) ResourceConflict() *Result[T]     { return r.Set(-409, "ResourceConflict") }
func (r *Result[T]) UnsupportedMediaType() *Result[T] { return r.Set(-415, "UnsupportedMediaType") }
func (r *Result[T]) InternalServerError() *Result[T]  { return r.Set(-500, "InternalServerError") }
func (r *Result[T]) Set(code int, msg string) *Result[T] {
	r.Code, r.Msg = code, msg
	return r
}
