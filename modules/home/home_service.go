package home

type Service struct{}

var ServiceInstance Service

func init() {
	ServiceInstance = Service{}
}
