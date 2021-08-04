package service

import "gorm-mysql/businessrule"

type MenuService struct {
	Repository businessrule.Repository
}

type CreateMenuService struct {
	MenuService
}

func(s *CreateMenuService) Service(input, output interface{}) error {

	return nil
}

type DisplayMenuService struct {
	MenuService
}

func (s *DisplayMenuService) Service(input, output interface{}) error {

	return nil
}

type UpdateMenuService struct {

}

