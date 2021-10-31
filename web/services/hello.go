package services

type HelloService interface {
	GetGreeting() string
}

type HelloServiceImpl struct {
}

func NewHelloService() *HelloServiceImpl {
	return &HelloServiceImpl{}
}

func (h *HelloServiceImpl) GetGreeting() string {
	// custom logic
	return "wello"
}
