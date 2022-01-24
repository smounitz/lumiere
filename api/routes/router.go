package routes

import (
	"api/config"
	"api/models"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(cfg *config.Configuration) *httprouter.Router {
	router := httprouter.New()
	gdb := models.NewGraphDb(cfg.Neo4j)
	userModel := models.NewUserModel(gdb)
	tokenModel := models.NewTokenModel(gdb, cfg.Jwt)
	router.GET("/v1/healthcheck", NewHealthCheckRoute())
	router.POST("/v1/signup", NewSignupRoute(userModel, tokenModel))
	// router.POST("/v1/login", NewLoginRoute(neo4j))
	return router
}
