package appid

import (
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strings"
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

func (id AppID) String() string{
	return uuid.UUID(id).String()
}

func NewAppID() *AppID {
	tmp, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	tmpApp := AppID(tmp)
	return &tmpApp
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

func ParseToAppID(input interface{})(interface{}, error) {
	tmp, ok := input.(string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unable to parse input of type %v, only accepts type of string", reflect.TypeOf(input)))
	}
	tmp = strings.Replace(tmp, "-", "", -1)
	v, err := hex.DecodeString(tmp)
	if err != nil {
		return nil, err
	}
	result, err := uuid.FromBytes(v)
	if err != nil {
		return nil, err
	}
	theId := AppID(result)
	return theId, nil
}
