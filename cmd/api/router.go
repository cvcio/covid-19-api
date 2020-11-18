package main

import (
	"net/http"
	"time"

	"github.com/cvcio/covid-19-api/cmd/api/handlers"
	"github.com/cvcio/covid-19-api/pkg/config"
	"github.com/cvcio/covid-19-api/pkg/db"
	"github.com/cvcio/covid-19-api/pkg/middleware"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"go.uber.org/zap"
)

// NewAPI Creates a new API Router using Gin
func NewAPI(cfg *config.Config, dbConn *db.DB, storeLimits limiter.Store, storeCasce *persistence.RedisStore, logger *zap.Logger) http.Handler {
	limiterMiddleware := mgin.NewMiddleware(limiter.New(storeLimits, limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  300,
	}))

	router := gin.New()

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

	router.Use(limiterMiddleware)

	// handlers
	global := handlers.NewGlobalHandler(cfg, dbConn, logger)
	greece := handlers.NewGlobalHandler(cfg, dbConn, logger)
	meta := handlers.NewGlobalHandler(cfg, dbConn, logger)

	// routes
	globalRoutes := router.Group("/global")
	{
		globalRoutes.GET("", cache.CachePage(storeCasce, 2*time.Hour, global.List))
		globalRoutes.GET("/:country", cache.CachePage(storeCasce, 2*time.Hour, global.List))
		globalRoutes.GET("/:country/:keys", cache.CachePage(storeCasce, 2*time.Hour, global.List))
		globalRoutes.GET("/:country/:keys/:from", cache.CachePage(storeCasce, 2*time.Hour, global.List))
		globalRoutes.GET("/:country/:keys/:from/:to", cache.CachePage(storeCasce, 2*time.Hour, global.List))
	}

	greeceRoutes := router.Group("/greece")
	{
		greeceRoutes.GET("", cache.CachePage(storeCasce, 2*time.Hour, greece.List))
		greeceRoutes.GET("/:region", cache.CachePage(storeCasce, 2*time.Hour, greece.List))
		greeceRoutes.GET("/:region/:keys", cache.CachePage(storeCasce, 2*time.Hour, greece.List))
		greeceRoutes.GET("/:region/:keys/:from", cache.CachePage(storeCasce, 2*time.Hour, greece.List))
		greeceRoutes.GET("/:region/:keys/:from/:to", cache.CachePage(storeCasce, 2*time.Hour, greece.List))
	}

	totalRoutes := router.Group("/total")
	{
		totalRoutes.GET("", cache.CachePage(storeCasce, 2*time.Hour, global.List))
		totalRoutes.GET("/global", cache.CachePage(storeCasce, 2*time.Hour, global.List))
		totalRoutes.GET("/global/:country", cache.CachePage(storeCasce, 2*time.Hour, global.List))
		totalRoutes.GET("/global/:country/:key", cache.CachePage(storeCasce, 2*time.Hour, global.List))
		totalRoutes.GET("/global/:country/:key/:from", cache.CachePage(storeCasce, 2*time.Hour, global.List))
		totalRoutes.GET("/global/:country/:key/:from/:to", cache.CachePage(storeCasce, 2*time.Hour, global.List))

		totalRoutes.GET("/greece", cache.CachePage(storeCasce, 2*time.Hour, greece.List))
		totalRoutes.GET("/greece/:region", cache.CachePage(storeCasce, 2*time.Hour, greece.List))
		totalRoutes.GET("/greece/:region/:key", cache.CachePage(storeCasce, 2*time.Hour, greece.List))
		totalRoutes.GET("/greece/:region/:key/:from", cache.CachePage(storeCasce, 2*time.Hour, greece.List))
		totalRoutes.GET("/greece/:region/:key/:from/:to", cache.CachePage(storeCasce, 2*time.Hour, greece.List))
	}

	metaRoutes := router.Group("/meta")
	{
		metaRoutes.GET("/countries", cache.CachePage(storeCasce, 24*time.Hour, meta.List))
		metaRoutes.GET("/routes", cache.CachePage(storeCasce, 24*time.Hour, meta.List))
	}

	// Forbid Access
	// This is usefull when you combine multiple microservices
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusForbidden, "Access Forbidden")
		c.Abort()
	})

	return router
}
