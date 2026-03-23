package user

import (
	"com.dreamfsk/blog/config"
	"com.dreamfsk/blog/services/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

func NewHandler(db *gorm.DB, cfg *config.Config, zl *zap.Logger) Handler {
	return &handler{userService: user.NewService(db), jwtSecret: []byte(cfg.JWT.Secret), zl: zl}
}

type handler struct {
	userService user.Service
	jwtSecret   []byte
	zl          *zap.Logger
}

func (h *handler) i() {}
