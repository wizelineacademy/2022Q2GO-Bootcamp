swagger:
	swagger generate spec -o ./api/swagger.yaml --scan-models

serve:
	go build cmd/main.go
	./main
