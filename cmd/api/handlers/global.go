package handlers

import (
	"errors"
	"strings"
	"time"

	"github.com/cvcio/covid-19-api/models/global"
	"github.com/cvcio/covid-19-api/pkg/config"
	"github.com/cvcio/covid-19-api/pkg/db"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Global Data Handlers
type Global struct {
	cfg    *config.Config
	dbConn *db.DB
	log    *zap.SugaredLogger
}

// NewGlobalHandler creates the appropriate handler
func NewGlobalHandler(cfg *config.Config, db *db.DB, logger *zap.Logger) *Global {
	return &Global{
		cfg:    cfg,
		dbConn: db,
		log:    logger.Sugar(),
	}
}

// List Data
func (h *Global) List(c *gin.Context) {
	opts := global.NewListOpts()

	if c.Param("country") != "" && strings.ToUpper(c.Param("country")) != "ALL" {
		opts = append(opts, global.ISO3(c.Param("country")))
	}

	if c.Param("keys") != "" {
		opts = append(opts, global.Keys(c.Param("keys")))
	}

	if c.Param("from") != "" {
		t, err := time.Parse("2006-01-02", c.Param("from"))
		if err == nil {
			opts = append(opts, global.From(t))
		}
	}

	if c.Param("to") != "" {
		t, err := time.Parse("2006-01-02", c.Param("to"))
		if err == nil {
			opts = append(opts, global.To(t))
		}
	}

	res, err := global.List(h.dbConn, opts...)
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
func (h *Global) Agg(c *gin.Context) {
	opts := global.NewListOpts()

	if c.Param("country") != "" && strings.ToUpper(c.Param("country")) != "ALL" {
		opts = append(opts, global.ISO3(c.Param("country")))
	}

	if c.Param("keys") != "" {
		opts = append(opts, global.Keys(c.Param("keys")))
	}

	if c.Param("from") != "" {
		t, err := time.Parse("2006-01-02", c.Param("from"))
		if err == nil {
			opts = append(opts, global.From(t))
		}
	}

	if c.Param("to") != "" {
		t, err := time.Parse("2006-01-02", c.Param("to"))
		if err == nil {
			opts = append(opts, global.To(t))
		}
	}

	res, err := global.Agg(h.dbConn, opts...)
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

// Sum Data
func (h *Global) Sum(c *gin.Context) {
	opts := global.NewListOpts()

	if c.Param("country") != "" && strings.ToUpper(c.Param("country")) != "ALL" {
		opts = append(opts, global.ISO3(c.Param("country")))
	}

	if c.Param("from") != "" {
		t, err := time.Parse("2006-01-02", c.Param("from"))
		if err == nil {
			opts = append(opts, global.From(t))
		}
	}

	if c.Param("to") != "" {
		t, err := time.Parse("2006-01-02", c.Param("to"))
		if err == nil {
			opts = append(opts, global.To(t))
		}
	}

	res, err := global.Sum(h.dbConn, opts...)
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
