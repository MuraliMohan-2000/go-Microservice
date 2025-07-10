package order

import (
	"dev.murali.go-microservice/account"
	"dev.murali.go-microservice/catalog"
)

type grpcServer struct {
	service       Service
	accountClient *account.Client
	catalogClient *catalog.Client
}
