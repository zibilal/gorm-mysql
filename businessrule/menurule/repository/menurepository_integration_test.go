package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/zibilal/teepr"
	"gorm-mysql/appctx"
	"gorm-mysql/businessrule"
	"gorm-mysql/businessrule/appid"
	"gorm-mysql/engine/dbengine"
	"testing"
)

const (
	failed = "\u2717"
)
func TestIntegrationMenuRepository_CRUD(t *testing.T) {
	appctx.InitAppContext()
	appctx.AppContext.User = map[string]string{
		"name": "CRUD Testing",
	}

	db, err := dbengine.InitDbEngine("localhost", "inventoryusr", "dbpass", "inventorydb", 3306)
	if err != nil {
		t.Fatal(err)
	}

	repo := &MenuRepository{
		dbEngine: db,
		BaseRepository: businessrule.BaseRepository{
			IdGenerator: func() appid.AppID {
				temp, err := appid.FromString("1DD1B664F14E11EBACE1ACDE48001122")
				if err != nil {
					t.Fatal(err)
				}
				return *temp
			},
		},
	}

	type args struct {
		query       map[string]interface{}
		inputOutput interface{}
	}

	appId, err := appid.FromString("1DD1B664F14E11EBACE1ACDE48001122")
	if err != nil {
		t.Fatal(err)
	}
	queryById := map[string]interface{}{
		"id": appId,
	}

	createMenu1 := struct {
		Name        string
		Type        string
		Description string
	}{
		"Menu Test 1", "EN", "A test description",
	}

	updateMenu1 := struct {
		Id          appid.AppID
		Name        string
		Type        string
		Description string
	}{
		*appId, "Menu Test 1 Updated", "EN", "A test description Updated",
	}

	deleteMenu1 := struct {
		Id          appid.AppID
		Name        string
		Type        string
		Description string
	}{
		*appId, "Menu Test 1 Updated", "EN", "A test description Updated",
	}

	outputQuery := make([]MenuEntity, 0)
	queryArgs := args{
		queryById,
		&outputQuery,
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Test Create Menu 1",
			args{
				nil,
				&createMenu1,
			},
			false,
		},
		{
			"Test Select Menu 1",
			queryArgs,
			false,
		},
		{
			"Test Update Menu 1",
			args{
				nil,
				&updateMenu1,
			},
			false,
		},
		{
			"Test Delete Menu 1",
			args{
				nil,
				&deleteMenu1,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Test Create Menu 1":
				if err := repo.Create(tt.args.inputOutput); (err != nil) != tt.wantErr {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				}
			case "Test Select Menu 1":
				if err := repo.Fetch(tt.args.query, tt.args.inputOutput); (err != nil) != tt.wantErr {
					t.Errorf("Fetch() error = %v, wantErr = %v", err, tt.wantErr)
				} else {
					dt, ok := tt.args.inputOutput.(*[]MenuEntity)
					if !ok {
						t.Errorf("Unknown data")
					} else {
						tmp := *dt
						assert.Equal(t, 1, len(tmp))
						assert.Equal(t, tmp[0].Id, *appId)
						assert.Equal(t, tmp[0].Name, "Menu Test 1")
						assert.Equal(t, tmp[0].Type, "EN")
						assert.Equal(t, tmp[0].Description, "A test description")
					}
				}
			case "Test Update Menu 1":
				if err := repo.Update(tt.args.inputOutput); (err != nil) != tt.wantErr {
					t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				} else {
					var outputFetch []struct {
						Id          appid.AppID
						Name        string
						Type        string
						Description string
					}
					err := repo.Fetch(tt.args.query, &outputFetch)
					if err != nil {
						t.Fatal(err)
					}

					assert.Equal(t, 1, len(outputFetch))
					assert.Equal(t, outputFetch[0].Id, *appId)
					assert.Equal(t, outputFetch[0].Name, "Menu Test 1 Updated")
					assert.Equal(t, outputFetch[0].Description, "A test description Updated")
				}
			case "Test Delete Menu 1":
				if err := repo.Delete(tt.args.inputOutput); (err != nil) != tt.wantErr {
					t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				} else {
					outputFetch := make([]MenuEntity, 0)
					err := repo.Fetch(tt.args.query, &outputFetch)
					if err != nil {
						t.Fatal(err)
					}

					assert.Equal(t, 0, len(outputFetch))
				}
			}
		})
	}
}

func TestMenuRepository_Create(t *testing.T) {
	appctx.InitAppContext()
	appctx.AppContext.User = map[string]string{
		"name": "CRUD Testing",
	}

	db, err := dbengine.InitDbEngine("localhost", "inventoryusr", "dbpass", "inventorydb", 3306)
	if err != nil {
		t.Fatal(err)
	}

	repo := &MenuRepository{
		dbEngine: db,
		BaseRepository: businessrule.BaseRepository{
			IdGenerator: func() appid.AppID {
				temp, err := appid.FromString("1DD1B664F14E11EBACE1ACDE48001122")
				if err != nil {
					t.Fatal(err)
				}
				return *temp
			},
		},
	}
	createMenu1 := struct {
		Name        string
		Type        string
		Description string
	}{
		"Menu Test 1", "EN", "A test description",
	}
	t.Log("Test Create A New menu")
	{
		err := repo.Create(createMenu1)
		if err != nil {
			t.Fatalf("test failed: %v", err)
		}
	}
}

func TestMenuRepository_Delete(t *testing.T) {
	appctx.InitAppContext()
	appctx.AppContext.User = map[string]string{
		"name": "CRUD Testing",
	}

	db, err := dbengine.InitDbEngine("localhost", "inventoryusr", "dbpass", "inventorydb", 3306)
	if err != nil {
		t.Fatal(err)
	}

	repo := &MenuRepository{
		dbEngine: db,
		BaseRepository: businessrule.BaseRepository{
			IdGenerator: func() appid.AppID {
				temp, err := appid.FromString("1DD1B664F14E11EBACE1ACDE48001122")
				if err != nil {
					t.Fatal(err)
				}
				return *temp
			},
		},
	}

	t.Log("Test Delete a menu")
	{
		appId, _ := appid.FromString("1DD1B664F14E11EBACE1ACDE48001122")
		deleteMenu1 := struct {
			Id          appid.AppID
			Name        string
			Type        string
			Description string
		}{
			*appId, "Menu Test 1 Updated", "EN", "A test description Updated",
		}

		err := repo.Delete(deleteMenu1)
		if err != nil {
			t.Fatalf("%s delete failed; expected error nil, got %v", failed, err)
		}
	}
}

func TestMenuRepository_Fetch(t *testing.T) {
	appctx.InitAppContext()
	appctx.AppContext.User = map[string]string{
		"name": "CRUD Testing",
	}

	db, err := dbengine.InitDbEngine("localhost", "inventoryusr", "dbpass", "inventorydb", 3306)
	if err != nil {
		t.Fatal(err)
	}

	repo := &MenuRepository{
		dbEngine: db,
		BaseRepository: businessrule.BaseRepository{
			IdGenerator: func() appid.AppID {
				temp, err := appid.FromString("1DD1B664F14E11EBACE1ACDE48001122")
				if err != nil {
					t.Fatal(err)
				}
				return *temp
			},
		},
	}
	t.Log("Test Fetching data from a database: ")
	{
		appId, err := appid.FromString("1DD1B664F14E11EBACE1ACDE48001122")
		if err != nil {
			t.Fatal(err)
		}

		queryById := map[string]interface{}{
			"id": appId,
		}
		var output []struct {
			Id          appid.AppID
			Name        string
			Type        string
			Description string
		}
		err = repo.Fetch(queryById, &output)
		if err != nil {
			t.Fatalf("Fetch() error = %v, wantErr = %v", err, false)
		}

		tmp := output
		assert.Equal(t, 1, len(tmp))
		assert.Equal(t, tmp[0].Id, *appId)
		assert.Equal(t, tmp[0].Name, "Menu Test 1")
		assert.Equal(t, tmp[0].Type, "EN")
		assert.Equal(t, tmp[0].Description, "A test description")

		var outputType []struct {
			Id          appid.AppID
			Name        string
			Type        string
			Description string
		}

		err = teepr.Teepr(tmp, &outputType)
		if err != nil {
			t.Fatalf("Fetch() error = %v", err)
		}
	}
}

func TestHex_DecodeString(t *testing.T) {
	s := "1DD1B664F14E11EBACE1ACDE48001122"
	id, err := appid.FromString(s)
	if err != nil {
		t.Fatal("Error1", err)
	}
	t.Logf("%v\n", id)
}
