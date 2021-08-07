package businessrule

import "gorm-mysql/businessrule/appid"

type Service interface {
	Serve(input, output interface{}) error
}

type BaseRepository struct {
	IdGenerator func() appid.AppID
	Repository
}

type Repository interface{
	DataCreator
	DataUpdater
	DataDeleter
	DataFetcher
}

type DataCreator interface {
	Create(input interface{}) error
}

type DataUpdater interface {
	Update(input interface{}) error
}

type DataDeleter interface {
	Delete(input interface{}) error
}

type DataFetcher interface {
	Fetch(query map[string]interface{}, output interface{}) error
}
