swagger:
	swagger generate spec -o ./api/swagger.yaml --scan-models

run:
	go build -o /out/server && ./server