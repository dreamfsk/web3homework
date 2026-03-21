package user

import (
	"com.dreamfsk/blog/config"
	"com.dreamfsk/blog/services/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetProfile(c *gin.Context)
	UpdateProfile(c *gin.Context)
}

func NewHandler(db *gorm.DB, cfg *config.Config) Handler {
	return &handler{userService: user.NewService(db), jwtSecret: []byte(cfg.JWT.Secret)}
}

type handler struct {
	userService user.Service
	jwtSecret   []byte
}

func (h *handler) i() {}
