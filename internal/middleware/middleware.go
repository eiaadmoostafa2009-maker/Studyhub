package middleware

import (
	"net/http"
	"studyhub/internal/models"
	"studyhub/internal/repository/user"
	"studyhub/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func RedirectToLogin(c *gin.Context, status int) {
	if c.GetHeader("HX-Request") == "true" {
		c.Header("HX-Redirect", "/user/login")
		c.AbortWithStatus(status)
		return
	}
	c.Redirect(http.StatusSeeOther, "/user/login")
	c.Abort()
}
func AuthMiddleWare(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authCookie, err := c.Cookie("token")
		if err != nil || authCookie == "" {
			RedirectToLogin(c, http.StatusUnauthorized)
			return
		}
		userID, userName, err := jwt.ValidateToken(authCookie, secretKey, true)
		if err != nil {
			RedirectToLogin(c, http.StatusUnauthorized)
			return
		}
		c.Set("user_name", userName)
		c.Set("user_id", userID)
		c.Next()
	}
}

// RequireRoles restricts access to specific user roles
func RequireRoles(userRepo user.UserRepository, roles ...models.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Extract user info from Gin Context (set by Auth middleware)
		val, exists := c.Get("user_id")
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := userRepo.GetUserByID(c.Request.Context(), val.(int))
        if err != nil{
			RedirectToLogin(c, http.StatusNotFound)
			return
		}
		// 2. Check if user's role matches any allowed roles
		for _, role := range roles {
			if user.Role == role {
				c.Next() // Role matched, proceed to next handler
				return
			}
		}

		// 3. Block access if no roles match
		c.AbortWithStatus(http.StatusForbidden)
	}
}
