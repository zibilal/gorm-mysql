package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/zibilal/teepr"
	"gorm-mysql/appctx"
	"gorm-mysql/businessrule/appid"
	"gorm-mysql/engine/dbengine"
	"log"
	"reflect"
	"time"
)

type MenuEntity struct {
	Id   appid.AppID `gorm:"primary_key;"`
	Name string      `gorm:"name"`
	Type string `gorm:"type"`
	Description string `gorm:"description"`
	CreatedBy string `gorm:"created_by"`
	CreatedAt *time.Time `gorm:"created_at"`
	UpdatedBy string `gorm:"updated_by"`
	UpdatedAt *time.Time `gorm:"updated_at"`
}

func (MenuEntity) TableName() string{
	return "menu"
}

type MenuRepository struct {
	dbEngine *dbengine.DbEngine
	idGenerator func() appid.AppID
}

func NewMenuRepository(dbEngine *dbengine.DbEngine) *MenuRepository {
	return &MenuRepository{
		dbEngine: dbEngine,
		idGenerator: func() appid.AppID {
			id, err := uuid.NewUUID()
			if err != nil {
				panic(err)
			}
			return appid.AppID(id)
		},
	}
}

func (r *MenuRepository) Create(input interface{}) error {
	entity := &MenuEntity{
		Id: r.idGenerator(),
	}
	err := copyDataToMenuEntity(input, entity)
	if err != nil {
		return err
	}

	v, ok := appctx.AppContext.User["name"]
	if ok {
		atime := time.Now()
		entity.CreatedAt = &atime
		entity.CreatedBy = v
		entity.UpdatedAt = &atime
		entity.UpdatedBy = v
	}

	result := r.dbEngine.Db.Create(entity)

	return result.Error
}

func (r *MenuRepository) Update(input interface{}) error {
	entity := &MenuEntity{}
	err := copyDataToMenuEntity(input, entity)
	if err != nil {
		log.Println("SATU")
		return err
	}
	if entity.Id.IsEmpty() {
		log.Println("DUA")
		return errors.New("the Id is empty, please provide the id for this MenuEntity")
	}
	result := r.dbEngine.Db.Save(entity)
	log.Println("TIGA", result.Error)
	return result.Error
}

func (r *MenuRepository) Delete(input interface{}) error {
	entity := &MenuEntity{}
	err := copyDataToMenuEntity(input, entity)
	if err != nil {
		return err
	}
	if entity.Id.IsEmpty() {
		return errors.New("txt Id is empty, please provide the id fo this MenuEntity, for delete this object")
	}
	result := r.dbEngine.Db.Delete(entity)

	return result.Error
}

func (r *MenuRepository) Fetch(query map[string]interface{}, output interface{}) error {
	result := r.dbEngine.Db.Where(query).Find(output)
	return result.Error
}

func copyDataToMenuEntity(input interface{}, entity *MenuEntity) error {
	err := teepr.Teepr(input, entity)
	if err != nil {
		return err
	}

	if reflect.DeepEqual(entity, &MenuEntity{}) {
		return errors.New("fail to extract Value Object into Entity")
	}

	return nil
}