package server

import (
	"api/config"
	"api/restapi"
	"net/http"

	ginMiddleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/rs/zerolog/log"
)

type server struct {
}

func NewServer(h restapi.ServerInterface, cfg *config.ServiceConfig) *gin.Engine {
	s := gin.Default()
	swagger, err := restapi.GetSwagger()
	if err != nil {
		log.Panic().Err(err).Msgf("")
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{cfg.FrontURLOrigin},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	s.Use(c)
	s.Use(ginMiddleware.OapiRequestValidator(swagger))
	binding.EnableDecoderDisallowUnknownFields = true
	s = restapi.RegisterHandlers(s, h)
	return s
}
