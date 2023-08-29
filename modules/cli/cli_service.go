package cli

type Service struct{}

var ServiceInstance = CreateService()

func CreateService() Service {
	return Service{}
}

func (s *Service) GetList() []string {
	return []string{"1", "2", "3"}
}
