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
	glCovid := handlers.NewGlobalHandler(cfg, dbConn, logger)
	grCovid := handlers.NewGreeceHandler(cfg, dbConn, logger)
	grVaccines := handlers.NewGRVaccinesHandler(cfg, dbConn, logger)

	// routes
	glCovidRoutes := router.Group("/global")
	{
		glCovidRoutes.GET("", cache.CachePage(storeCasce, 15*time.Minute, glCovid.List))
		glCovidRoutes.GET("/:country", cache.CachePage(storeCasce, 15*time.Minute, glCovid.List))
		glCovidRoutes.GET("/:country/:keys", cache.CachePage(storeCasce, 15*time.Minute, glCovid.List))
		glCovidRoutes.GET("/:country/:keys/:from", cache.CachePage(storeCasce, 15*time.Minute, glCovid.List))
		glCovidRoutes.GET("/:country/:keys/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, glCovid.List))
	}

	grCovidRoutes := router.Group("/greece")
	{
		grCovidRoutes.GET("", cache.CachePage(storeCasce, 15*time.Minute, grCovid.List))
		grCovidRoutes.GET("/:region", cache.CachePage(storeCasce, 15*time.Minute, grCovid.List))
		grCovidRoutes.GET("/:region/:keys", cache.CachePage(storeCasce, 15*time.Minute, grCovid.List))
		grCovidRoutes.GET("/:region/:keys/:from", cache.CachePage(storeCasce, 15*time.Minute, grCovid.List))
		grCovidRoutes.GET("/:region/:keys/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, grCovid.List))
	}

	grVaccinesRoutes := router.Group("/vaccines/greece")
	{
		grVaccinesRoutes.GET("", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.List))
		grVaccinesRoutes.GET("/:region", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.List))
		grVaccinesRoutes.GET("/:region/:keys", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.List))
		grVaccinesRoutes.GET("/:region/:keys/:from", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.List))
		grVaccinesRoutes.GET("/:region/:keys/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.List))
	}

	totalRoutes := router.Group("/agg")
	{
		totalRoutes.GET("", cache.CachePage(storeCasce, 15*time.Minute, glCovid.Agg))
		totalRoutes.GET("/global", cache.CachePage(storeCasce, 15*time.Minute, glCovid.Agg))
		totalRoutes.GET("/global/:country", cache.CachePage(storeCasce, 15*time.Minute, glCovid.Agg))
		totalRoutes.GET("/global/:country/:keys", cache.CachePage(storeCasce, 15*time.Minute, glCovid.Agg))
		totalRoutes.GET("/global/:country/:keys/:from", cache.CachePage(storeCasce, 15*time.Minute, glCovid.Agg))
		totalRoutes.GET("/global/:country/:keys/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, glCovid.Agg))

		totalRoutes.GET("/greece", cache.CachePage(storeCasce, 15*time.Minute, grCovid.Agg))
		totalRoutes.GET("/greece/:region", cache.CachePage(storeCasce, 15*time.Minute, grCovid.Agg))
		totalRoutes.GET("/greece/:region/:keys", cache.CachePage(storeCasce, 15*time.Minute, grCovid.Agg))
		totalRoutes.GET("/greece/:region/:keys/:from", cache.CachePage(storeCasce, 15*time.Minute, grCovid.Agg))
		totalRoutes.GET("/greece/:region/:keys/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, grCovid.Agg))

		totalRoutes.GET("/vaccines/greece", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.Agg))
		totalRoutes.GET("/vaccines/greece/:region", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.Agg))
		totalRoutes.GET("/vaccines/greece/:region/:keys", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.Agg))
		totalRoutes.GET("/vaccines/greece/:region/:keys/:from", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.Agg))
		totalRoutes.GET("/vaccines/greece/:region/:keys/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.Agg))
	}

	sumRoutes := router.Group("/total")
	{
		sumRoutes.GET("", cache.CachePage(storeCasce, 15*time.Minute, glCovid.Sum))
		sumRoutes.GET("/global", cache.CachePage(storeCasce, 15*time.Minute, glCovid.Sum))
		sumRoutes.GET("/global/:country", cache.CachePage(storeCasce, 15*time.Minute, glCovid.Sum))
		sumRoutes.GET("/global/:country/:from", cache.CachePage(storeCasce, 15*time.Minute, glCovid.Sum))
		sumRoutes.GET("/global/:country/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, glCovid.Sum))

		sumRoutes.GET("/greece", cache.CachePage(storeCasce, 15*time.Minute, grCovid.Sum))
		sumRoutes.GET("/greece/:region", cache.CachePage(storeCasce, 15*time.Minute, grCovid.Sum))
		sumRoutes.GET("/greece/:region/:from", cache.CachePage(storeCasce, 15*time.Minute, grCovid.Sum))
		sumRoutes.GET("/greece/:region/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, grCovid.Sum))

		sumRoutes.GET("/vaccines/greece", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.Sum))
		sumRoutes.GET("/vaccines/greece/:region", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.Sum))
		sumRoutes.GET("/vaccines/greece/:region/:from", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.Sum))
		sumRoutes.GET("/vaccines/greece/:region/:from/:to", cache.CachePage(storeCasce, 15*time.Minute, grVaccines.Sum))
	}

	// Forbid Access
	// This is usefull when you combine multiple microservices
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 Not Found")
		c.Abort()
	})

	return router
}
