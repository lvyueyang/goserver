package {{.Name}}

type Service struct{}

var ServiceInstance *Service

func init() {
	ServiceInstance = &Service{}
}

func (s *Service) GetList() {
}
