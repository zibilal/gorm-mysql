package service

import (
	"errors"
	"gorm-mysql/businessrule"
)

type MenuService struct {
	Repository businessrule.Repository
}

func (s *MenuService) SetUpRepository(repo businessrule.Repository) error {
	if repo == nil {
		return errors.New("repo cannot be nil")
	}
	s.Repository = repo
	return nil
}

type CreateMenuService struct {
	MenuService
}

func NewCreateMenuService() *CreateMenuService {
	return new(CreateMenuService)
}

func(s *CreateMenuService) Service(input, _ interface{}) error {
	return s.Repository.Create(input)
}

type DisplayMenuService struct {
	MenuService
}

func NewDisplayMenuService() *DisplayMenuService {
	return new(DisplayMenuService)
}

func (s *DisplayMenuService) Service(input, output interface{}) error {
	return nil
}

type UpdateMenuService struct {
	MenuService
}

func NewUpdateMenuService() *UpdateMenuService{
	return new(UpdateMenuService)
}

func (s *UpdateMenuService) Service(input, _ interface{}) error {
	return s.Repository.Update(input)
}

type DeleteMenuService struct {
	MenuService
}

func NewDeleteMenuService() *DeleteMenuService {
	return new(DeleteMenuService)
}

func (s *DeleteMenuService) Service(input, output interface{}) error {

	return nil
}

