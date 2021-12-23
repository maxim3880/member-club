// Package restapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package restapi

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// CreateUserInput defines model for create_user_input.
type CreateUserInput UserInput

// CreateUserPayload defines model for create_user_payload.
type CreateUserPayload struct {
	Data *User `json:"data,omitempty"`
}

// GenericErrorModel defines model for generic_error_model.
type GenericErrorModel struct {
	ErrorMessage string `json:"error_message"`
}

// GetUsersPayload defines model for get_users_payload.
type GetUsersPayload struct {
	Data []User `json:"data"`
}

// User defines model for user.
type User struct {
	Email            openapi_types.Email `json:"email"`
	Name             string              `json:"name"`
	RegistrationDate openapi_types.Date  `json:"registration_date"`
}

// UserInput defines model for user_input.
type UserInput struct {
	Email openapi_types.Email `json:"email"`
	Name  string              `json:"name"`
}

// PostUsersJSONBody defines parameters for PostUsers.
type PostUsersJSONBody CreateUserInput

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody PostUsersJSONBody
