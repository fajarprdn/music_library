package appreq

type TrackRequest struct {
	Id        string `json:"id" binding:"required"`
	TrackName string `json:"track_name" binding:"required"`
	Artist    string `json:"artist" binding:"required"`
}
