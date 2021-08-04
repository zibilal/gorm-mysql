package conf

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)


func TestGetConfiguration(t *testing.T) {

	var awant = &Configuration{
		 Db: struct {
		 	 Name string `yaml:"name"`
			 Hostname string `yaml:"host"`
			 Port     int    `yaml:"port"`
			 User     string `yaml:"user"`
			 Password string `yaml:"pswd"`
		 }{
		 	 Name: "inventorydb",
			 Hostname: "localhost",
			 Port: 3306,
			 User: "root",
			 Password: "secretsample",
		 },
	}

	var str = `
dbengine:
  name: inventorydb
  host: localhost
  port: 3306
  user: root
  pswd: secretsample
`
	abuf := bytes.NewBufferString(str)
	type args struct {
		reader io.Reader
	}
	args1 := args {
		abuf,
	}
	tests := []struct {
		name    string
		args    args
		want    *Configuration
		wantErr bool
	}{
		{
			"first_test",
			args1,
			awant,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConfiguration(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfiguration() got = %v, want %v", got, tt.want)
			}
		})
	}
}
