package service

type HomeService struct {
}

var HomeServiceInstance = HomeService{}

func (s *HomeService) GetInfo() string {
	return "hello world"
}
