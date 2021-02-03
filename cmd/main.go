package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/mrcarromesa/fullcycle_codepix_go/application/grpc"
	"github.com/mrcarromesa/fullcycle_codepix_go/infrastructure/db"
)

var database *gorm.DB

func main() {

	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
