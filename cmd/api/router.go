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
	"github.com/gin-contrib/gzip"
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
		router.Use(middleware.EnableCORS("*"))
	}

	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(limiterMiddleware)

	// handlers
	global := handlers.NewGlobalHandler(cfg, dbConn, logger)
	greece := handlers.NewGreeceHandler(cfg, dbConn, logger)
	// routes
	globalRoutes := router.Group("/global")
	{
		globalRoutes.GET("", cache.CachePage(storeCasce, 15*time.Minute, global.List))
		globalRoutes.GET("/:country", cache.CachePage(storeCasce, 15*time.Minute, global.List))
		globalRoutes.GET("/:country/:keys", cache.CachePage(storeCasce, 15*time.Minute, global.List))
		globalRoutes.GET("/:country/:keys/:from", cache.CachePage(storeCasce, 15*time.Minute, global.List))
		globalRoutes.GET("/:country/:keys/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, global.List))
	}

	greeceRoutes := router.Group("/greece")
	{
		greeceRoutes.GET("", cache.CachePage(storeCasce, 15*time.Minute, greece.List))
		greeceRoutes.GET("/:region", cache.CachePage(storeCasce, 15*time.Minute, greece.List))
		greeceRoutes.GET("/:region/:keys", cache.CachePage(storeCasce, 15*time.Minute, greece.List))
		greeceRoutes.GET("/:region/:keys/:from", cache.CachePage(storeCasce, 15*time.Minute, greece.List))
		greeceRoutes.GET("/:region/:keys/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, greece.List))
	}

	totalRoutes := router.Group("/agg")
	{
		totalRoutes.GET("", cache.CachePage(storeCasce, 15*time.Minute, global.Agg))
		totalRoutes.GET("/global", cache.CachePage(storeCasce, 15*time.Minute, global.Agg))
		totalRoutes.GET("/global/:country", cache.CachePage(storeCasce, 15*time.Minute, global.Agg))
		totalRoutes.GET("/global/:country/:keys", cache.CachePage(storeCasce, 15*time.Minute, global.Agg))
		totalRoutes.GET("/global/:country/:keys/:from", cache.CachePage(storeCasce, 15*time.Minute, global.Agg))
		totalRoutes.GET("/global/:country/:keys/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, global.Agg))

		totalRoutes.GET("/greece", cache.CachePage(storeCasce, 15*time.Minute, greece.Agg))
		totalRoutes.GET("/greece/:region", cache.CachePage(storeCasce, 15*time.Minute, greece.Agg))
		totalRoutes.GET("/greece/:region/:keys", cache.CachePage(storeCasce, 15*time.Minute, greece.Agg))
		totalRoutes.GET("/greece/:region/:keys/:from", cache.CachePage(storeCasce, 15*time.Minute, greece.Agg))
		totalRoutes.GET("/greece/:region/:keys/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, greece.Agg))
	}

	sumRoutes := router.Group("/total")
	{
		sumRoutes.GET("", cache.CachePage(storeCasce, 15*time.Minute, global.Sum))
		sumRoutes.GET("/global", cache.CachePage(storeCasce, 15*time.Minute, global.Sum))
		sumRoutes.GET("/global/:country", cache.CachePage(storeCasce, 15*time.Minute, global.Sum))
		sumRoutes.GET("/global/:country/:from", cache.CachePage(storeCasce, 15*time.Minute, global.Sum))
		sumRoutes.GET("/global/:country/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, global.Sum))

		sumRoutes.GET("/greece", cache.CachePage(storeCasce, 15*time.Minute, greece.Sum))
		sumRoutes.GET("/greece/:region", cache.CachePage(storeCasce, 15*time.Minute, greece.Sum))
		sumRoutes.GET("/greece/:region/:from", cache.CachePage(storeCasce, 15*time.Minute, greece.Sum))
		sumRoutes.GET("/greece/:region/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, greece.Sum))
	}

	// Forbid Access
	// This is usefull when you combine multiple microservices
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 Not Found")
		c.Abort()
	})

	return router
}
