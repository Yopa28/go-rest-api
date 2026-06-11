package middlewares

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strings"
)

func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        roleRaw, exists := c.Get("role")
        roleStr, ok := roleRaw.(string)

        fmt.Println("Role from token:", roleRaw)

        if !exists || !ok || strings.ToLower(roleStr) != "admin" {
            c.JSON(403, gin.H{
                "message": "Forbidden : Admin Only",
            })
            c.Abort()
            return
        }

        c.Next()
    }
}