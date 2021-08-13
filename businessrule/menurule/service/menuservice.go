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
	if input == nil {
		return errors.New("please provide the input value")
	}
	return s.Repository.Create(input)
}

type DisplayMenuService struct {
	MenuService
}

func NewDisplayMenuService() *DisplayMenuService {
	return new(DisplayMenuService)
}

func (s *DisplayMenuService) Service(input, output interface{}) error {
	if input == nil {
		return errors.New("please provide the input value")
	}
	query, ok := input.(map[string]interface{})
	if !ok {
		return errors.New("input accepted only of typed map[string]interface{}")
	}
	return s.Repository.Fetch(query, output)
}

type UpdateMenuService struct {
	MenuService
}

func NewUpdateMenuService() *UpdateMenuService{
	return new(UpdateMenuService)
}

func (s *UpdateMenuService) Service(input, _ interface{}) error {
	if input == nil {
		return errors.New("please provide the input value")
	}
	return s.Repository.Update(input)
}

type DeleteMenuService struct {
	MenuService
}

func NewDeleteMenuService() *DeleteMenuService {
	return new(DeleteMenuService)
}

func (s *DeleteMenuService) Service(input, _ interface{}) error {
	if input == nil {
		return errors.New("please provide the input value")
	}
	return s.Repository.Delete(input)
}

