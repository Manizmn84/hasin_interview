package middleware

import (
	"fmt"
	"strings"

	"github.com/Manizmn84/hasin_interview/bootstrap"
	"github.com/Manizmn84/hasin_interview/internal/application/usecase"
	"github.com/Manizmn84/hasin_interview/internal/domain/exception"
	"github.com/Manizmn84/hasin_interview/internal/domain/logging"
	"github.com/Manizmn84/hasin_interview/internal/domain/ports"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	logger     logging.Logger
	cfg        *bootstrap.Config
	jwtService usecase.JwtService
	unitOfWork ports.UnitOfWork
}

func NewAuthMiddleWare(cfg *bootstrap.Config, jwtService usecase.JwtService, unitOfWork ports.UnitOfWork) *AuthMiddleware {
	logger := logging.NewLogger(cfg)
	return &AuthMiddleware{
		logger:     logger,
		cfg:        cfg,
		jwtService: jwtService,
		unitOfWork: unitOfWork,
	}
}

func (am *AuthMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			authError := &exception.UnauthorizedError{Message: fmt.Sprintf("%s.%s", am.cfg.Constant.ErrorField.Auth, am.cfg.Constant.ErrorTag.EmptyHeader)}
			panic(authError)
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			authError := &exception.UnauthorizedError{Message: fmt.Sprintf("%s.%s", am.cfg.Constant.ErrorField.Auth, am.cfg.Constant.ErrorTag.InvalidHeader)}
			panic(authError)
		}

		token := parts[1]
		if token == "" {
			authError := &exception.UnauthorizedError{Message: fmt.Sprintf("%s.%s", am.cfg.Constant.ErrorField.Jwt, am.cfg.Constant.ErrorTag.EmptyToken)}
			panic(authError)
		}

		claims, err := am.jwtService.ValidateToken(token)
		if err != nil {
			panic(err)
		}

		ctx.Set(am.cfg.Constant.Context.UserID, uint(claims["userID"].(float64)))
		ctx.Next()
	}
}
