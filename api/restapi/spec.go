// Package restapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package restapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RUzW7bTAx8lcV+31GInbQn3doeiqAomotPRWDQEi1voP0JSbVVA717sZSb2JEa5+Ae",
	"ehNW5AxnOLsPtoo+xYBB2JYPlqsdetDPihAE1x0jrV1IneTD/wm3trT/LZ7aFvuexUHlUBy1J+jbCHUG",
	"SBQTkjhUjhoEXgNrh6Gw0ie0pY2bO6yUosGA5Ko1EkVa+1hjO6XY/0RmaDAf7GFYyIVGgQnvO0dY2/Lr",
	"s/LbWVZRVXxalhP0/Dp9j0RABP1kLEWcm0a7p6I9OPViG8mD2HJ/UjxXX9gAfs6WTN84FgJxMaxrEDzC",
	"04PihJmKXTxyTxH/pOgpcOfW9dKE02lyuQvbqEBO2vzvM/oNkqnabmPe3Vzbwn5DYheDLe3lxfJimdlj",
	"wgDJ2dK+0aPCJpCdatCF61eDKrFGrsglGSE+orCBtjVaZja9SdC4oKaZBASereKPPl7XY89KQbM4TjHw",
	"aNfVcqk3OQbBoFyQUusq7VzccSb8fedPxXQae3XnePgvn7L4t8vLqa5VgE52kdxPrLXo6uqMs00fgpnp",
	"VgF/JKwEa6OVmgbuvAfqRxf3preOJWcKGs4hGRd2OxQ2RZ5Z2Qd96rTXfHeyM/m2mi1Fbzax7ifruol8",
	"sK/7Dlne57pz2TF9uYfj3At1OEyycvlXBnghLaNv9T8cmYPNz+RlGIZfAQAA//+v6wroYAcAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
