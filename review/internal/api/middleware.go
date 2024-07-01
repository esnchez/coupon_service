package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

		errs := c.Errors
		if len(errs) > 0 {

			// Define the default status and message
			status := http.StatusInternalServerError
			message := "Internal Server Error"
			
			err := errs[0].Err
			log.Printf("Captured error: %s\n", err.Error())

			switch err.(type) {
				case *CustomError:
					customErr := err.(*CustomError)
					status = customErr.StatusCode
					message = customErr.Message
			}

			c.JSON(status, gin.H{
				"error": message,
			})
		}
    }
}
