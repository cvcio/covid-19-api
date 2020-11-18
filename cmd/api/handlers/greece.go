package handlers

import (
	"github.com/cvcio/covid-19-api/pkg/config"
	"github.com/cvcio/covid-19-api/pkg/db"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Greece Data Handlers
type Greece struct {
	cfg    *config.Config
	dbConn *db.DB
	log    *zap.SugaredLogger
}

// NewGreeceHandler creates the appropriate handler
func NewGreeceHandler(cfg *config.Config, db *db.DB, logger *zap.Logger) *Greece {
	return &Greece{
		cfg:    cfg,
		dbConn: db,
		log:    logger.Sugar(),
	}
}

// List Data
func (h *Greece) List(c *gin.Context) {}
