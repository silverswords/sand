package controller

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand/models"
)

type UserController struct {
	db *sql.DB
}

func NewUserController(db *sql.DB) *ProductController {
	return &ProductController{
		db: db,
	}
}

func (c *UserController) RegisterRouter(r gin.IRouter) {
	if err := models.CreateUserTable(c.db); err != nil {
		log.Fatal(err)
	}
}
