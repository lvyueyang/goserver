package cli

type Service struct{}

var ServiceInstance Service

func init() {
	ServiceInstance = Service{}
}

func (s *Service) GetList() []string {
	return []string{"1", "2", "3"}
}
