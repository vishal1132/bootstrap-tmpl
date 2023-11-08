package database

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"{{ module_name }}/postgres"
	"{{ module_name }}/users/model"
)

const (
	tableName = "users"
)

type usersDB struct {
	db *gorm.DB
}

func NewUsersDB(db *gorm.DB) *usersDB {
	return &usersDB{
		db: db,
	}
}

func (db *usersDB) CreateUser(ctx context.Context, user *model.User) error {

	return postgres.HandleError(db.db.Table(tableName).Create(user))
}

func (db *usersDB) GetUser(ctx context.Context, entityType, entityVal string) (*model.User, error) {
	var user = new(model.User)
	txn := db.db.Table(tableName).Where(fmt.Sprintf("%s = ?", entityType), entityVal).First(&user)
	return user, postgres.HandleError(txn)
}

func (db *usersDB) UpdateUser(ctx context.Context, user *model.User) error {
	return db.db.Table(tableName).Model(user).Where("id = ?", user.ID).Updates(user).Error
}
