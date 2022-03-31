package api

import (
	"github.com/gin-gonic/gin"
	"music_library/delivery/appreq"
	"music_library/usecase"
)

type TrackApi struct {
	BaseApi
	trackRegistrationUseCase usecase.TrackRegistrationUseCase
}

func (t *TrackApi) createTrack() gin.HandlerFunc {
	return func(c *gin.Context) {
		var trackReq appreq.TrackRequest
		err := t.ParseRequestBody(c, &trackReq)
		if err != nil {
			t.FailedRequest(c, "api-create-track", "02", "Can not parse body")
			return
		}
		trackRegistered, err := t.trackRegistrationUseCase.Register(trackReq.Id, trackReq.TrackName, trackReq.Artist)
		if err != nil {
			t.Failed(c, "api-createTrack", "01", "Can not create track")
			return
		}
		t.Success(c, "Track", trackRegistered)
	}
}

func NewTrackApi(trackRoute *gin.RouterGroup, trackRegistrationUseCase usecase.TrackRegistrationUseCase) {
	api := TrackApi{
		trackRegistrationUseCase: trackRegistrationUseCase,
	}
	trackRoute.POST("", api.createTrack())
}
