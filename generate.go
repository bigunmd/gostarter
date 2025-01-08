package gostarter

//go:generate swag fmt -d internal/heroes -g internal/heroes/oas31.go
//go:generate swag init -d internal/heroes -g oas31.go --v3.1 -o docs/heroes -ot yaml
