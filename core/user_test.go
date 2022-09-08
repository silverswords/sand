package core

import (
	"fmt"
	"testing"

	"github.com/silverswords/sand/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		"root", "my-123456", "mqying.xyz", "3306", "mall", "utf8", true, "Local")

	config := &Config{
		Dsn: dsn,
	}

	instance := CreateApplication(config)
	db = instance.GetDefaultGormDB()
}

func TestCreateUser(t *testing.T) {
	if err := model.CreateUser(db, "1111", "1111", "12345678901"); err != nil {
		t.Error(err)
	}
}

func TestQueryByUnionId(t *testing.T) {
	id, err := model.QueryByUnionId(db, "1111")
	if err != nil {
		t.Error(err)
	}

	t.Log(id)
}

func TestQueryByMobile(t *testing.T) {
	mobile := "12345678901"

	user, err := model.QueryByMobile(db, mobile)
	if err != nil {
		t.Error(err)
		return
	}

	if user.UnionID != "1111" || user.Mobile != "12345678901" || user.OpenID != "1111" {
		t.Error("query error")
	}
}

func TestUpdateMobile(t *testing.T) {
	mobile := "98765432109"
	err := model.UpdateMobile(db, "1111", mobile)
	if err != nil {
		t.Error(err)
		return
	}

	var expected string
	db.Model(model.User{}).Select("mobile").Where(model.User{UnionID: "1111"}).Scan(&expected)
	if expected != mobile {
		t.Errorf("expected %s, got %s", expected, mobile)
	}
}
