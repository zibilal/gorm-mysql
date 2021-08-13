package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/zibilal/teepr"
	"gorm-mysql/appctx"
	"gorm-mysql/businessrule"
	"gorm-mysql/businessrule/appid"
	"gorm-mysql/engine/dbengine"
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
	businessrule.BaseRepository
	dbEngine *dbengine.DbEngine
}

func NewMenuRepository(dbEngine *dbengine.DbEngine) *MenuRepository {
	return &MenuRepository{
		dbEngine: dbEngine,
		BaseRepository: businessrule.BaseRepository{
			IdGenerator: func() appid.AppID {
				id, err := uuid.NewUUID()
				if err != nil {
				panic(err)
				}
				return appid.AppID(id)
			},
		},
	}
}

func (r *MenuRepository) Create(input interface{}) error {
	if input == nil {
		return errors.New("input is required")
	}
	entity := &MenuEntity{
		Id: r.BaseRepository.IdGenerator(),
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
	if input == nil {
		return errors.New("input is required")
	}
	entity := &MenuEntity{}
	err := copyDataToMenuEntity(input, entity)
	if err != nil {
		return err
	}
	if entity.Id.IsEmpty() {
		return errors.New("the Id is empty, please provide the id for this MenuEntity")
	}
	result := r.dbEngine.Db.Save(entity)
	return result.Error
}

func (r *MenuRepository) Delete(input interface{}) error {
	if input == nil {
		return errors.New("input is required")
	}
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

	tmp := make([]MenuEntity, 0)
	result := r.dbEngine.Db.Where(query).Find(&tmp)
	if result.Error != nil {
		return result.Error
	}

	err := teepr.Teepr(tmp, output)
	if err != nil {
		return err
	}

	return nil
}

func copyDataToMenuEntity(input interface{}, entity *MenuEntity) error {
	err := teepr.Teepr(input, entity)
	if err != nil {
		return err
	}

	if teepr.IsEmpty(entity) {
		return errors.New("fail to extract Value Object into Entity")
	}

	return nil
}