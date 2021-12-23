package repo

import (
	"api/persistence/model"
	"time"

	pkgErrr "api/pkg/errors"

	uuid "github.com/satori/go.uuid"
)

type Database interface {
	GetUserList() map[string]*model.User
	GetUserByEmail(email string) (*model.User, error)
	AddUser(email string, user *model.User) (*model.User, error)
}

type inMemoryDB struct {
	data map[string]*model.User
}

func NewDatabase() Database {
	data := map[string]*model.User{}
	return &inMemoryDB{data: data}
}

func (db *inMemoryDB) GetUserList() map[string]*model.User {
	return db.data
}

func (db *inMemoryDB) GetUserByEmail(email string) (*model.User, error) {
	v, ok := db.data[email]
	if !ok {
		return nil, pkgErrr.NewKindedError(pkgErrr.KindNotExists, "failed to get user by email %v", email)
	}
	return v, nil
}

func (db *inMemoryDB) AddUser(email string, user *model.User) (*model.User, error) {
	_, ok := db.data[email]
	if ok {
		return nil, pkgErrr.NewKindedError(pkgErrr.KindAlreadyExists,
			"failed to add user, because user with same email %v already exists", email)
	}
	user.ID = uuid.NewV4()
	user.CreatedAt = time.Now().UTC()
	db.data[email] = user
	return user, nil
}

// func (db *inMemoryDB) dbDataToArray() []*model.User {
// 	res := make([]*model.User, 0)
// 	for _, v := range db.data {
// 		res = append(res, v)
// 	}
// 	return res
// }
