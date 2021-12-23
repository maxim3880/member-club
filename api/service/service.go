package service

import (
	"api/persistence/model"
	"api/persistence/repo"
	"api/pkg/errors"
	"api/restapi"

	"github.com/deepmap/oapi-codegen/pkg/types"
)

type InternalService interface {
	GetUsersList() (*restapi.GetUsersPayload, error)
	CreateUserWithParams(data *restapi.CreateUserInput) (*restapi.CreateUserPayload, error)
}

type service struct {
	rp repo.Database
}

func NewService(rp repo.Database) InternalService {
	return &service{rp}
}

func (s *service) GetUsersList() (*restapi.GetUsersPayload, error) {
	list := s.rp.GetUserList()
	resp := &restapi.GetUsersPayload{
		Data: []restapi.User{},
	}
	for _, v := range list {
		u := restapi.User{
			Email:            types.Email(v.Email),
			Name:             v.Name,
			RegistrationDate: types.Date{Time: v.CreatedAt},
		}
		resp.Data = append(resp.Data, u)
	}
	return resp, nil
}

func (s *service) CreateUserWithParams(data *restapi.CreateUserInput) (*restapi.CreateUserPayload, error) {
	email := string(data.Email)
	user, err := s.rp.GetUserByEmail(email)
	if err != nil && !errors.IsNotExistsError(err) {
		return nil, err
	}
	user = &model.User{
		Email: email,
		Name:  data.Name,
	}
	user, err = s.rp.AddUser(user.Email, user)
	if err != nil {
		return nil, err
	}
	resp := &restapi.CreateUserPayload{
		Data: &restapi.User{
			Email:            types.Email(user.Email),
			Name:             user.Name,
			RegistrationDate: types.Date{Time: user.CreatedAt},
		},
	}
	return resp, nil
}
