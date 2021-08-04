package appid

import (
	"database/sql/driver"
	"encoding/hex"
	"github.com/google/uuid"
	"reflect"
)

type AppID uuid.UUID

func (id *AppID) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*id = AppID(parseByte)

	return err
}

func (id AppID) Value() (driver.Value, error) {
	return uuid.UUID(id).MarshalBinary()
}

func (id *AppID) GormDataType() string {
	return "binary(16)"
}

func (id AppID) IsEmpty() bool {
	return reflect.DeepEqual(id, reflect.Zero(reflect.TypeOf(id)).Interface())
}

func FromBytes(b []byte) (*AppID, error) {
	temp, err := uuid.FromBytes(b)
	if err != nil {
		return nil, err
	}

	appID := AppID(temp)
	return &appID, nil
}

func FromString(s string) (*AppID, error) {
	temp, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}

	return FromBytes(temp)
}
