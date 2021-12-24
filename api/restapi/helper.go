package restapi

import (
	"api/pkg/errors"
	"net/http"

	"github.com/rs/zerolog/log"
)

func NewGenericErrorModel(err error, msgf string, args ...interface{}) GenericErrorModel {
	log.Error().Err(err).Msgf(msgf, args...)
	return GenericErrorModel{
		Error: err.Error(),
	}
}

func GetStatusCodeByErrorKind(err error) int {
	kind := errors.GetErrorKind(err)
	switch kind {
	case errors.KindNotExists:
		return http.StatusNotFound
	case errors.KindAlreadyExists:
		return http.StatusConflict
	default:
		log.Debug().Msgf("failed to get status by error kind = %v", kind)
	}
	return http.StatusBadRequest
}
