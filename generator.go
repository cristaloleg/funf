package funf

type Generator struct {
	bound int64
	fn    func(int64) interface{}
	out   chan interface{}
}

func NewGenerator(fn func(int64) interface{}) *Generator {
	return &Generator{
		fn: fn,
	}
}

func NewBoundedGenerator(bound int64, fn func(int64) interface{}) *Generator {
	return &Generator{
		bound: bound,
		fn:    fn,
	}
}

func (g *Generator) Next() <-chan interface{} {
	g.out = make(chan interface{})
	go func() {
		for i := int64(0); g.bound == 0 || i < g.bound; i++ {
			g.out <- g.fn(i)
		}
	}()
	return g.out
}

func (g *Generator) Stop() {
	close(g.out)
}

