package adaptor

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"literank.com/rest-books/domain/model"
)

const tokenPrefix = "Bearer "

// PermCheck checks user permission
func (r *RestHandler) PermCheck(allowPerm model.UserPermission) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is required"})
			c.Abort()
			return
		}
		token := strings.Replace(authHeader, tokenPrefix, "", 1)
		hasPerm, err := r.userOperator.HasPermission(token, allowPerm)
		message := "Unauthorized"
		if err != nil {
			message = err.Error()
		}
		if !hasPerm {
			c.JSON(http.StatusUnauthorized, gin.H{"error": message})
			c.Abort()
			return
		}
		c.Next()
	}
}
