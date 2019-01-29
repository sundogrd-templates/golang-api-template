package user_test

import (
	"bytes"
	"context"
	"github.com/spf13/viper"
	"github.com/sundogrd-templates/golang-api-template/services/content"
	"github.com/sundogrd-templates/golang-api-template/services/user"
	"github.com/sundogrd-templates/golang-api-template/utils/config"
	"github.com/sundogrd-templates/golang-api-template/utils/db"
	"testing"
)

func initTestDB() error {
	config.Init()
	viper.SetConfigType("json") // or viper.SetConfigType("YAML")

	var jsonConfig = []byte(`{
	  	"db": {
			"type": "mysql",
			"options": {
				"user": "sundog",
				"password": "sundogPwd",
				"host": "localhost",
				"port": 3306,
				"dbname": "sundog",
				"connectTimeout": "10s"
			}
	  	}
	}`)
	viper.ReadConfig(bytes.NewBuffer(jsonConfig))
	_, err := db.Init()
	return err
}

// TestUserFindOne ...
func TestUserFindOne(t *testing.T) {
	ctx := context.Background()
	err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res := user.UserServiceInstance().FindOne(ctx, user.FindOneRequest{
		UserID: 304008154235023360,
	})
	t.Logf("FindContent: %+v", res)
}

// TestUserCreate ...
func TestUserCreate(t *testing.T) {
	var err error

	ctx := context.Background()
	err = initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := user.UserServiceInstance().Create(ctx, user.CreateRequest{
		Name:      "LWio",
		AvatarUrl: "https://avatars1.githubusercontent.com/u/9214496?v=4",
		Company:   "kk",
		Email:     "liang.peare@gmail.com",
		Extra: user.DataInfoExtra{
			GithubHome: "https://github.com/lwyj123",
		},
	})
	if err != nil {
		t.Fatalf("CreateContent err: %+v", err)
	}
	t.Logf("CreateContent: %+v", res)
}

// TestUserDelete ...
func TestUserDelete(t *testing.T) {
	var err error

	ctx := context.Background()
	err = initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := content.ContentRepositoryInstance().Delete(ctx, content.DeleteRequest{
		ContentIDs: []int64{299696847465746432},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("DeleteContent: %+v", res)
}