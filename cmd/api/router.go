package main

import (
	"net/http"
	"time"

	"github.com/cvcio/covid-19-api/pkg/config"
	"github.com/cvcio/covid-19-api/pkg/db"
	"github.com/cvcio/covid-19-api/pkg/middleware"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewAPI Creates a new API Router using Gin
func NewAPI(cfg *config.Config, dbConn *db.DB, logger *zap.Logger) http.Handler {
	router := gin.Default()

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true
	router.ForwardedByClientIP = true

	router.Use(gin.Recovery())
	// log middleware
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	// cors middleware
	if cfg.Env == "development" {
		router.Use(middleware.EnableCORS("*"))
	} else {
		gin.SetMode(gin.ReleaseMode)
		router.Use(middleware.EnableCORS(" *." + cfg.Server.DomainName))
	}

	// handlers
	// routes

	// Forbid Access
	// This is usefull when you combine multiple microservices
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusForbidden, "Access Forbidden")
		c.Abort()
	})

	return router
}
