package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cvcio/covid-19-api/pkg/config"
	"github.com/cvcio/covid-19-api/pkg/db"
	"github.com/cvcio/covid-19-api/pkg/redis"
	"github.com/gin-contrib/cache/persistence"
	"github.com/kelseyhightower/envconfig"
	"github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"go.uber.org/zap"
)

// API Server
func main() {
	// ============================================================
	// Configuration & Logger
	// ============================================================
	// Create new configuration object
	cfg := config.New()

	// Create new logger using `zap`
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	log := logger.Sugar()

	// Read config from env variables and parse as configuration object
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatalf("[API] Error loading config: %s", err.Error())
	}

	// ============================================================
	// Start Mongo
	// ============================================================
	log.Debug("[SERVER] Initialize Mongo")
	dbConn, err := db.New(cfg.MongoURL(), cfg.Mongo.Path, cfg.Mongo.DialTimeout)
	if err != nil {
		log.Fatalf("[SERVER] Register DB: %v", err)
	}
	defer dbConn.Close()

	// database := client.Database(cfg.Mongo.Path)

	// ============================================================
	// Redis Client
	// ============================================================
	redisPool, err := redis.NewCachePool(cfg.RedisURL())
	if err != nil {
		log.Fatalf("[SERVER] error connecting to redis: %v", err)
	}

	redisClient, err := redis.NewLimitsClient(cfg.RedisWithPathURL())
	if err != nil {
		log.Fatalf("[SERVER] error connecting to redis: %v", err)
	}

	// ============================================================
	// Memory Storage
	// ============================================================
	// req cache
	storeCache := persistence.NewRedisCacheWithPool(redisPool.Pool, 15*time.Minute)
	storeCache.Flush()
	// limits
	storeLimits, err := sredis.NewStoreWithOptions(redisClient.Client, limiter.StoreOptions{
		Prefix:   "limiter",
		MaxRetry: 4,
	})
	if err != nil {
		log.Fatalf("[SERVER] error creating limits store clients to redis: %v", err)
	}
	// ============================================================
	// Start API Service
	// ============================================================
	log.Debug("[SERVER] Starting")
	server := http.Server{
		Addr: cfg.ServerURL(),
		Handler: NewAPI(
			cfg,
			dbConn,
			storeLimits,
			storeCache,
			logger,
		),
		ReadTimeout:    cfg.Server.ReadTimeout,
		WriteTimeout:   cfg.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// Sleep for a while before starting
	time.Sleep(100 * time.Millisecond)

	// Blocking main listening for requests
	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	log.Debug("[SERVER] Ready to start")
	go func() {
		log.Infof("[SERVER] Starting api Listening %s", cfg.ServerURL())
		serverErrors <- server.ListenAndServe()
	}()

	// ============================================================
	// Shutdown
	// ============================================================
	// Listen for os signals
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	// ============================================================
	// Stop API Service
	// ============================================================
	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		log.Fatalf("[SERVER] Error starting server: %v", err)

	case <-osSignals:
		log.Info("[SERVER] Start shutdown...")

		// Create context for Shutdown call.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		if err := server.Shutdown(ctx); err != nil {
			log.Infof("[SERVER] Graceful shutdown did not complete in %v: %v", cfg.Server.ShutdownTimeout, err)
			if err := server.Close(); err != nil {
				log.Fatalf("[SERVER] Could not stop http server: %v", err)
			}
		}
	}
}
