package user

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func BadRequestErrorMiddleware() gin.HandlerFunc {
    return func (c *gin.Context) {
        c.Next()
        detectedErrors := c.Errors

        if detectedErrors == nil {
            return
        }

        var err error = detectedErrors[0]

        if err.Error() == "EOF" {
            c.AbortWithStatusJSON(400, gin.H{"message": "provide request body", "status": 400})
            return
        }

        c.JSON(http.StatusBadRequest, gin.H{"error": detectedErrors[0].Error(), "status": http.StatusBadRequest})
        return
    }
}
