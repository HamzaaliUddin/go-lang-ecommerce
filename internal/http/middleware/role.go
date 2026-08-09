package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleChecker interface {
	HasAnyRole(
		userID uint,
		roles ...string,
	) (bool, error)
}

type RoleMiddleware struct {
	roleChecker RoleChecker
}

func NewRole(
	roleChecker RoleChecker,
) *RoleMiddleware {
	return &RoleMiddleware{
		roleChecker: roleChecker,
	}
}

func (m *RoleMiddleware) Require(
	roles ...string,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("userID")

		if !exists {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "unauthorized",
				},
			)
			return
		}

		userID, ok := value.(uint)

		if !ok {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "unauthorized",
				},
			)
			return
		}

		allowed, err := m.roleChecker.HasAnyRole(
			userID,
			roles...,
		)

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"message": "internal server error",
				},
			)
			return
		}

		if !allowed {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"message": "forbidden",
				},
			)
			return
		}

		c.Next()
	}
}