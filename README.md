# member-club

# Swagger regenerate
oapi-codegen --package restapi -generate types -o ./api/restapi/schemas.go api/swagger.yaml
oapi-codegen --package restapi -generate spec -o -o ./api/restapi/spec.go ./api/swagger.yaml
oapi-codegen --package restapi templates ./swagger-template/ -import-mapping=fiber:github.com/gofiber/fiber/v2,cerrors:poc.bp.api/pkg/cerrors -generate server -o ./api/restapi/server.go ./api/swagger.yaml
oapi-codegen --package restapi -generate gin -o ./api/restapi/server.go ./api/swagger.yaml