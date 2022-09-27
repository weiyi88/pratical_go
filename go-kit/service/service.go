package service

type CalculateService interface {
	Add(a, b int) int
	Reduce(a, b int) int
	Mulit(a, b int) int
	Test(msg, code int) int
}

type calculateService struct {
}

func NewService() *calculateService {
	return &calculateService{}
}

func (s *calculateService) Add(a, b int) int {
	return a + b
}
func (s *calculateService) Reduce(a, b int) int {
	return a - b
}
func (s *calculateService) Mulit(a, b int) int {
	return a * b
}

func (s *calculateService) Test(msg, code int) int {
	if code == 200 {
		return 999
	}
	return 111
}
