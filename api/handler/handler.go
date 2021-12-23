package handler

import (
	"api/config"
	"api/restapi"
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	cfg *config.ServiceConfig
	svc service.InternalService
}

func NewServerHandler(cfg *config.ServiceConfig, svc service.InternalService) restapi.ServerInterface {
	return &handler{cfg, svc}
}

// Get users list
// (GET /users)
func (h *handler) GetUsers(c *gin.Context) {
	resp, err := h.svc.GetUsersList()
	if err != nil {
		c.JSON(restapi.GetStatusCodeByErrorKind(err),
			restapi.NewGenericErrorModel(err, "failed to get user list"))
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Create user
// (POST /users)
func (h *handler) PostUsers(c *gin.Context) {
	dst := &restapi.CreateUserInput{}

	if err := c.ShouldBindJSON(dst); err != nil {
		c.JSON(http.StatusUnprocessableEntity, restapi.NewGenericErrorModel(err, "failed to parse input body"))
		return
	}
	resp, err := h.svc.CreateUserWithParams(dst)
	if err != nil {
		c.JSON(restapi.GetStatusCodeByErrorKind(err),
			restapi.NewGenericErrorModel(err, "failed to add user to db with email %v", dst.Email))
		return
	}
	c.JSON(http.StatusCreated, resp)
}
