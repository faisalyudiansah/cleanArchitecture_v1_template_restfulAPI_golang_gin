package servers

import (
	"database/sql"
	"server/controllers"
	"server/helpers"
	"server/helpers/logger"
	"server/repositories"
	"server/services"

	"github.com/gin-gonic/gin"
)

var (
	Bcrypt   *helpers.BcryptStruct
	Jwt      *helpers.JwtProviderHS256
	GetParam *helpers.GetParam
)

var (
	TransactorRepository *repositories.TransactorRepositoryImplementation
	ExampleRepository    *repositories.ExampleRepositoryImplementation
)

var (
	ExampleService *services.ExampleServiceImplementation
)

var (
	ExampleController *controllers.ExampleController
)

func SetupController(router *gin.Engine, db *sql.DB) {
	logrusLogger := logger.NewLogger()
	logger.SetLogger(logrusLogger)

	Bcrypt = helpers.NewBcryptStruct()
	Jwt = helpers.NewJWTProviderHS256()
	GetParam = helpers.NewGetParams()
	TransactorRepository = repositories.NewTransactorRepositoryImplementation(db)

	SetupExampleController(db)
	SetupYourRoute(router, ExampleController)
}

func SetupExampleController(db *sql.DB) {
	ExampleRepository = repositories.NewExampleRepositoryImplementation(db)
	ExampleService = services.NewExampleServiceImplementation(TransactorRepository)
	ExampleController = controllers.NewExampleController(ExampleService)
}
