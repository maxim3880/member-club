rm ./restapi/schemas.go ./restapi/spec.go ./restapi/server.go
echo "[restapi] generated models."
oapi-codegen --package restapi -generate types -o ./restapi/schemas.go ./swagger.yaml
echo "[restapi] generated spec."
oapi-codegen --package restapi -generate spec -o ./restapi/spec.go ./swagger.yaml
echo "[restapi] generated server."
oapi-codegen --package restapi -generate gin -o ./restapi/server.go ./swagger.yaml