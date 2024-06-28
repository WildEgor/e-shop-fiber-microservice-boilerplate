package routers

// Router generic interface
type Router[T any] interface {
	Setup(app T)
}
