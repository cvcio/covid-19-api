package handlers

import (
	"errors"
	"strings"
	"time"

	"github.com/cvcio/covid-19-api/models/greece"
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
func (h *Greece) List(c *gin.Context) {
	opts := greece.NewListOpts()

	if c.Param("region") != "" && strings.ToUpper(c.Param("region")) != "ALL" {
		opts = append(opts, greece.UID(c.Param("region")))
	}

	if c.Param("keys") != "" {
		opts = append(opts, greece.Keys(c.Param("keys")))
	}

	if c.Param("from") != "" {
		t, err := time.Parse("2006-01-02", c.Param("from"))
		if err == nil {
			opts = append(opts, greece.From(t))
		}
	}

	if c.Param("to") != "" {
		t, err := time.Parse("2006-01-02", c.Param("to"))
		if err == nil {
			opts = append(opts, greece.To(t))
		}
	}

	res, err := greece.List(h.dbConn, opts...)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	if res == nil {
		c.JSON(404, errors.New("404 Not Found"))
	} else {
		c.JSON(200, res)
	}
	return
}

// Agg Aggregate Data
func (h *Greece) Agg(c *gin.Context) {
	opts := greece.NewListOpts()

	if c.Param("region") != "" && strings.ToUpper(c.Param("region")) != "ALL" {
		opts = append(opts, greece.UID(c.Param("region")))
	}

	if c.Param("keys") != "" {
		opts = append(opts, greece.Keys(c.Param("keys")))
	}

	if c.Param("from") != "" {
		t, err := time.Parse("2006-01-02", c.Param("from"))
		if err == nil {
			opts = append(opts, greece.From(t))
		}
	}

	if c.Param("to") != "" {
		t, err := time.Parse("2006-01-02", c.Param("to"))
		if err == nil {
			opts = append(opts, greece.To(t))
		}
	}

	res, err := greece.Agg(h.dbConn, opts...)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	if res == nil {
		c.JSON(404, errors.New("404 Not Found"))
	} else {
		c.JSON(200, res)
	}

	return
}
