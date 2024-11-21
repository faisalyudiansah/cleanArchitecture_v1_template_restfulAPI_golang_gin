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
	ServerRepository     *repositories.ServerRepositoryImplementation
)

var (
	ServerService *services.ServerServiceImplementation
)

var (
	ServerController *controllers.ServerController
)

func SetupController(router *gin.Engine, db *sql.DB) {
	logrusLogger := logger.NewLogger()
	logger.SetLogger(logrusLogger)

	Bcrypt = helpers.NewBcryptStruct()
	Jwt = helpers.NewJWTProviderHS256()
	GetParam = helpers.NewGetParams()
	TransactorRepository = repositories.NewTransactorRepositoryImplementation(db)

	SetupServerController(db)
	SetupYourRoute(router, ServerController)
}

func SetupServerController(db *sql.DB) {
	ServerRepository = repositories.NewServerRepositoryImplementation(db)
	ServerService = services.NewServerServiceImplementation(TransactorRepository)
	ServerController = controllers.NewServerController(ServerService)
}
