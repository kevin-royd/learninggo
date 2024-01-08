package main

/*
接口是一组行为 的抽象
里面只能有方法，方法也不需要 func 关 键字
定义接口，同时也需要定义结构对象
*/

type Server interface {
	Route()
	Start()
}

type SdkServer struct{}

func (s *SdkServer) Route() {
	//TODO implement me
	panic("implement me")
}

func (s *SdkServer) Start() {
	//TODO implement me
	panic("implement me")
}

// 对挖暴露时根据设计设计模式进行暴露，
// 例如
type Factory func() Server

var factory Factory

func RegisterFactory(f Factory) Factory {
	factory = f
	return f
}

func NewServer() Server {
	return factory()

}

func main() {

}
