package server

import (
	"api/restapi"

	ginMiddleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"
)

type server struct {
}

func NewServer(h restapi.ServerInterface) *gin.Engine {
	s := gin.Default()
	swagger, err := restapi.GetSwagger()
	if err != nil {
		log.Panic().Err(err).Msgf("")
	}
	s.Use(ginMiddleware.OapiRequestValidator(swagger))
	binding.EnableDecoderDisallowUnknownFields = true
	// r.Use(middleware.DisallowUnknownFieldsMiddlewar)
	s = restapi.RegisterHandlers(s, h)
	return s
}
