package service

import (
	"fmt"

	"github.com/flowdev/example-micro/internalacts-cmd/db"
)

type Service struct {
}

func New() *Service {
	fmt.Println("service.New()")
	return &Service{}
}

func (s *Service) ApiGetWidgets() {
	fmt.Println("service.ApiGetWidgets()")
	db.ReadData()
}
func (s *Service) ApiCreateWidget() {
	fmt.Println("service.ApiCreateWidget()")
	db.ReadData()
}
func (s *Service) ApiUpdateWidget() {
	fmt.Println("service.ApiUpdateWidget()")
	db.ReadData()
}
