package service

import (
	"github.com/google/uuid"
	"github.com/zibilal/teepr"
	"gorm-mysql/businessrule"
	"gorm-mysql/businessrule/appid"
	"gorm-mysql/businessrule/menurule/repository"
	"testing"
)

const (
	success = "\u2713"
	failed = "\u2717"
)

type mockrepository struct {
	businessrule.BaseRepository
}

func new_mockrepository() *mockrepository {
	return &mockrepository{
		businessrule.BaseRepository {
			IdGenerator: func() appid.AppID {
				id, _ := uuid.NewUUID()
				return appid.AppID(id)
			},
		},
	}
}

func (r *mockrepository) Create(input interface{}) error {
	entity := repository.MenuEntity{
		Id : r.IdGenerator(),
	}

	return teepr.Teepr(input, &entity)
}

func (r *mockrepository) Update(input interface{}) error {
	return nil
}

func (r *mockrepository) Delete(input interface{}) error {
	return nil
}

func (r *mockrepository) Fetch(query map[string] interface{}, output interface{}) error {
	return nil
}

func TestMenuService_SetUpRepository_Success(t *testing.T) {
	t.Log("Test MenuServcie.SetUpRepository")
	{
		svc := MenuService{}
		err := svc.SetUpRepository(new_mockrepository())

		if err != nil {
			t.Logf("%s Expected error nil, got %s", failed, err.Error())
		}
		t.Logf("%s Expected error nil", success)
	}
}

func TestCreateMenuService_Service(t *testing.T) {
	t.Log("Test CreateMenuService")
	{
		svc := NewCreateMenuService()
		err := svc.SetUpRepository(new_mockrepository())
		if err != nil {
			t.Logf("%s expected error nil, got %s", failed, err.Error())
		}
		t.Logf("%s expected error nil", success)

		input := struct {
			Name string
			Type string
			Description string
		}{}

		err= svc.Service(&input, nil)
		if err != nil {
			t.Logf("%s After executed Create service, expected error nil, got %s", failed, err.Error())
		}
		t.Logf("%s After executed Create service", success)
	}
}