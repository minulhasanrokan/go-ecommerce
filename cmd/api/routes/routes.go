package routes

import (
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/applicationData"
)

func Route(appData *applicationData.Application) {

	appData.Server.GET("/", appData.Handler.HealthCheck)

	api := appData.Server.Group("/api")

	UserRoutes(api, appData)
	ProductRoutes(api, appData)
	TransactionRoutes(api, appData)

}
