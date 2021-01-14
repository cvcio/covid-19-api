package handlers

import (
	"errors"
	"strings"
	"time"

	"github.com/cvcio/covid-19-api/models/gr_vaccines"
	"github.com/cvcio/covid-19-api/pkg/config"
	"github.com/cvcio/covid-19-api/pkg/db"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GRVaccines Data Handlers
type GRVaccines struct {
	cfg    *config.Config
	dbConn *db.DB
	log    *zap.SugaredLogger
}

// NewGRVaccinesHandler creates the appropriate handler
func NewGRVaccinesHandler(cfg *config.Config, db *db.DB, logger *zap.Logger) *GRVaccines {
	return &GRVaccines{
		cfg:    cfg,
		dbConn: db,
		log:    logger.Sugar(),
	}
}

// List Data
func (h *GRVaccines) List(c *gin.Context) {
	opts := gr_vaccines.NewListOpts()

	if c.Param("region") != "" && strings.ToUpper(c.Param("region")) != "ALL" {
		opts = append(opts, gr_vaccines.UID(c.Param("region")))
	}

	if c.Param("keys") != "" {
		opts = append(opts, gr_vaccines.Keys(c.Param("keys")))
	}

	if c.Param("from") != "" {
		t, err := time.Parse("2006-01-02", c.Param("from"))
		if err == nil {
			opts = append(opts, gr_vaccines.From(t))
		}
	}

	if c.Param("to") != "" {
		t, err := time.Parse("2006-01-02", c.Param("to"))
		if err == nil {
			opts = append(opts, gr_vaccines.To(t))
		}
	}

	res, err := gr_vaccines.List(h.dbConn, opts...)
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
func (h *GRVaccines) Agg(c *gin.Context) {
	opts := gr_vaccines.NewListOpts()

	if c.Param("region") != "" && strings.ToUpper(c.Param("region")) != "ALL" {
		opts = append(opts, gr_vaccines.UID(c.Param("region")))
	}

	if c.Param("keys") != "" {
		opts = append(opts, gr_vaccines.Keys(c.Param("keys")))
	}

	if c.Param("from") != "" {
		t, err := time.Parse("2006-01-02", c.Param("from"))
		if err == nil {
			opts = append(opts, gr_vaccines.From(t))
		}
	}

	if c.Param("to") != "" {
		t, err := time.Parse("2006-01-02", c.Param("to"))
		if err == nil {
			opts = append(opts, gr_vaccines.To(t))
		}
	}

	res, err := gr_vaccines.Agg(h.dbConn, opts...)
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
func (h *GRVaccines) Sum(c *gin.Context) {
	opts := gr_vaccines.NewListOpts()

	if c.Param("region") != "" && strings.ToUpper(c.Param("region")) != "ALL" {
		opts = append(opts, gr_vaccines.UID(c.Param("region")))
	}

	if c.Param("from") != "" {
		t, err := time.Parse("2006-01-02", c.Param("from"))
		if err == nil {
			opts = append(opts, gr_vaccines.From(t))
		}
	}

	if c.Param("to") != "" {
		t, err := time.Parse("2006-01-02", c.Param("to"))
		if err == nil {
			opts = append(opts, gr_vaccines.To(t))
		}
	}

	res, err := gr_vaccines.Sum(h.dbConn, opts...)
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
