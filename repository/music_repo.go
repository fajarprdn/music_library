package repository

import "music_library/model"

type TrackRepo interface {
	Insert(newTrack model.Track) (model.Track, error)
	GetByTrackId(id string) model.Track
	GetAll() []model.Track
}
