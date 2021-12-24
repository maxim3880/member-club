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

	"H4sIAAAAAAAC/9RUTW/TQBD9K6uFo9WkhZNvwAFVCNFLTqiKJt6Js5X3o7NjwFT+72jHoU1Y0/QQDtxW",
	"65n35r156wfdBBeDR89J1w86NTt0IMeGEBjXfUJaWx97zpevCbe61q8WT22Lfc/ioHKsjtojDF0AkwEi",
	"hYjEFoXDAMNLYPU4VpqHiLrWYXOHjVC06JFss0aiQGsXDHYlhXzMh317YrK+FUDC+94SGl1/3ZfdzrKw",
	"qEinZVhGl16m55EIiGAoxhHEuWmkuxTpwIr2bSAHrOv9TfWn6kp7cDhjR6ZvbWICtsGvDTAe4clFdcJE",
	"wa4euUvEvyl6Cti5dT03YTlNLrd+GwTIcpe/fUa3QVJN12/Uu5trXelvSMkGr2t9ebG8WGb2ENFDtLrW",
	"b+Sq0hF4Jxpk4XJqUSQaTA3ZyBPER+SkoOuUlKnNoCK01otpKgKBS1rwJx+vzdSzEtAsLsXg02TX1XIp",
	"Lzd4Ri9cEGNnG+lc3KVM+PuNn4ppGXtx53j4L5+y+LfLy1LXykPPu0D2Jxopuro642zlw5+ZbuXxR8SG",
	"0ajpeeeS1DsHNEwu7k3vbOKcKWhTDsm0sNux0jGkmZV9kF+b9Krvlncqv1a1peDUJpihWNdNSAf7uu8x",
	"8ftcdy47yj/1eJx7ph7HIiuX/2SAZ9Iy+Wb+48gcbH4mL+M4/goAAP//TlIDGVAHAAA=",
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
