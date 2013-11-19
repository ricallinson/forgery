package f

type Renderer interface {
	Render(string, ...interface{}) (string, error)
}
