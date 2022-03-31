package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"music_library/delivery/commonresp"
	"music_library/logger"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedError := c.Errors.Last()
		if detectedError == nil {
			return
		}
		e := detectedError.Error()
		errResp := commonresp.ErrorMessage{}
		err := json.Unmarshal([]byte(e), &errResp)
		if err != nil {
			errResp.HttpCode = http.StatusInternalServerError
			errResp.ErrorDescription = commonresp.ErrorDescription{
				Status:       "Error",
				ResponseCode: "06",
				Description:  "Convert JSON Failed",
			}
			logger.Log.Error().Err(err).Msg(errResp.Description)
		}
	}
}
