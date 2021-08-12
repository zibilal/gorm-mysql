package service

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/zibilal/teepr"
	"gorm-mysql/businessrule"
	"gorm-mysql/businessrule/appid"
	"gorm-mysql/businessrule/menurule/repository"
	"reflect"
	"strings"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

type mockrepository struct {
	businessrule.BaseRepository

	inputOutputType repository.MenuEntity
	query           map[string]interface{}
}

func new_mockrepository() *mockrepository {
	return &mockrepository{
		BaseRepository: businessrule.BaseRepository{
			IdGenerator: func() appid.AppID {
				id, _ := uuid.NewUUID()
				return appid.AppID(id)
			},
		},
	}
}

func (r *mockrepository) Create(input interface{}) error {
	r.inputOutputType = repository.MenuEntity{
		Id: r.IdGenerator(),
	}

	return teepr.Teepr(input, &r.inputOutputType)
}

func (r *mockrepository) Update(input interface{}) error {
	r.inputOutputType = repository.MenuEntity{}
	err := teepr.Teepr(input, &r.inputOutputType, appid.ParseToAppID)
	if err != nil {
		return err
	}

	if teepr.IsEmpty(r.inputOutputType.Id) {
		return errors.New("the Id is empty, please provide the id for this MenuEntity")
	}

	return nil
}

func (r *mockrepository) Delete(input interface{}) error {
	r.inputOutputType = repository.MenuEntity{}
	return teepr.Teepr(input, &r.inputOutputType)
}

func (r *mockrepository) Fetch(query map[string]interface{}, output interface{}) error {
	r.query = query
	_, ok := output.(*[]repository.MenuEntity)
	if !ok {
		return errors.New("unexpected type. only accept pointer of []MenuEntity")
	}

	menu1 := repository.MenuEntity{
		Id:          *appid.NewAppID(),
		Name:        "Menu Test 1",
		Type:        "MT",
		Description: "Menu Test 1 Description",
	}
	menu2 := repository.MenuEntity{
		Id:          *appid.NewAppID(),
		Name:        "Menu Test 2",
		Type:        "MT",
		Description: "Menu Test 2 Description",
	}
	menu3 := repository.MenuEntity{
		Id:          *appid.NewAppID(),
		Name:        "Menu Test 3",
		Type:        "MT",
		Description: "Menu Test 3 Description",
	}

	oval := reflect.Indirect(reflect.ValueOf(output))
	otyp := oval.Type()
	fmt.Println("Oval:", oval, "Otyp: ", otyp)
	outSlice := reflect.MakeSlice(reflect.SliceOf(otyp.Elem()),0, 3)
	outSlice = reflect.Append(outSlice, reflect.ValueOf(menu1),
		reflect.ValueOf(menu2), reflect.ValueOf(menu3))
	oval.Set(outSlice)
	fmt.Println("Here...")

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
		repomock := new_mockrepository()
		err := svc.SetUpRepository(repomock)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		}
		t.Logf("%s expected error nil", success)

		input := struct {
			Name        string
			Type        string
			Description string
		}{}

		err = svc.Service(&input, nil)
		if err != nil {
			t.Fatalf("%s After executed Create service, expected error nil, got %s", failed, err.Error())
		}

		result := repomock.inputOutputType
		if teepr.IsEmpty(result.Id) {
			t.Fatalf("%s After executed Create service, expected Id is not empty", failed)
		}
		t.Logf("%s After executed Create service, expected Id is not empty. The Id %v", success, result.Id)
	}
}

func TestUpdateMenuService_Service(t *testing.T) {
	t.Log("Test UpdateMenuService Id type string and empty")
	{
		// strId := "1DD1B664F14E11EBACE1ACDE48001122"
		svc := NewUpdateMenuService()
		repomock := new_mockrepository()
		err := svc.SetUpRepository(repomock)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		}
		t.Logf("%s expected error nil", success)

		input := struct {
			Id          string
			Name        string
			Type        string
			Description string
		}{
			"", "Test Name", "TN", "A Test Name only",
		}

		err = svc.Service(input, nil)
		t.Logf("%s Expected error not nil, got %v", success, err)
	}

}

func TestDisplayMenuService_Service(t *testing.T) {
	t.Log("Test DisplayMenuService able to get data")
	{
		svc := NewDisplayMenuService()
		repomock := new_mockrepository()
		err := svc.SetUpRepository(repomock)
		if err != nil {
			t.Fatalf("%s expected error nil, got %v", failed, err)
		}
		t.Logf("%s expected error nil", success)

		query := map[string]interface{} {
			"type": "MT",
		}

		output := make([]repository.MenuEntity, 0)
		err = svc.Service(query, &output)
		if err != nil {
			t.Fatalf("%s expected error nil, got %v", failed, err)
		}
		if len(output) <= 0 {
			t.Fatalf("%s expected output length > 0 is true, got %v", failed, len(output) > 0)
		}

		t.Logf("%s")
		t.Logf("%s Output: %v", success, output)
	}
}

func TestUpdateAppId(t *testing.T) {
	t.Log("Test UpdateMenuService Id type string with random string value")
	{
		svc := NewUpdateMenuService()
		repomock := new_mockrepository()
		err := svc.SetUpRepository(repomock)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		}
		t.Logf("%s expected error nil", success)

		input := struct {
			Id          string
			Name        string
			Type        string
			Description string
		}{
			"12345678910", "Test Name", "TN", "A Test Name Only",
		}

		err = svc.Service(input, nil)
		if err == nil {
			t.Fatalf("%s Expected error not nil", failed)
		}
		t.Logf("%s Expected error not nil, got %v", success, err)
	}

	t.Log("Test UpdateMenuService Id type string with a valid uuid hexa string")
	{
		svc := NewUpdateMenuService()
		repomock := new_mockrepository()
		err := svc.SetUpRepository(repomock)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		}
		t.Logf("%s expected error nil", success)

		input := struct {
			Id          string
			Name        string
			Type        string
			Description string
		}{
			"1DD1B664F14E11EBACE1ACDE48001122", "Test Name", "TN", "A Test Name Only",
		}

		err = svc.Service(input, nil)
		if err != nil {
			t.Fatalf("%s Expected no error, got %v", failed, err)
		}
		t.Logf("%s Expected no error", success)
		t.Logf("%s Result: %v", success, repomock.inputOutputType)
		if teepr.IsEmpty(repomock.inputOutputType.Id) {
			t.Fatalf("%s Expected Id is not empty", failed)
		}
	}

	t.Log("Test UpdateMenuService Id type string with a valid uuid hexa string")
	{
		svc := NewUpdateMenuService()
		repomock := new_mockrepository()
		err := svc.SetUpRepository(repomock)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err)
		}
		t.Logf("%s expected error nil", success)

		strId := "1dd1b664-f14e-11eb-ace1-acde48001122"
		input := struct {
			Id          string
			Name        string
			Type        string
			Description string
		}{
			strId, "Test Name", "TN", "A Test Name Only",
		}

		err = svc.Service(input, nil)
		if err != nil {
			t.Fatalf("%s Expected no error, got %v", failed, err)
		}
		t.Logf("%s Expected no error", success)
		t.Logf("%s Result: %v", success, repomock.inputOutputType)
		if teepr.IsEmpty(repomock.inputOutputType.Id) {
			t.Fatalf("%s Expected Id is not empty", failed)
		}
	}
}

func TestTypeWithAnnotation(t *testing.T) {
	input := struct {
		Id          string
		Name        string
		Type        string
		Description string
	}{
		"1DD1B664F14E11EBACE1ACDE48001122", "Test Name", "TN", "A Test Name Only",
	}

	output := repository.MenuEntity{}
	err := teepr.Teepr(input, &output, appid.ParseToAppID)
	if err != nil {
		t.Fatalf("%s Expected error nil, got %v", failed, err)
	}

	t.Logf("%s Result: %v", success, output)
}

func TestIdBytes(t *testing.T) {
	strId := "1dd1b664-f14e-11eb-ace1-acde48001122"
	strId = strings.Replace(strId, "-", "", -1)
	b, err := hex.DecodeString(strId)
	if err != nil {
		t.Fatalf("%s Expected error nil, got %v", failed, err)
	}

	tmp, err := uuid.FromBytes(b)
	if err != nil {
		t.Logf("%s expected error nil, got %v", failed, err)
	}
	t.Logf("%s Result: %v", success, tmp)
}
