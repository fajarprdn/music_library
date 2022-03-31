package usecase

import (
	"music_library/model"
	"music_library/repository"
)

type TrackRegistrationUseCase interface {
	Register(id, trackName, artist string) (model.Track, error)
}

type trackRegistrationUseCase struct {
	repo repository.TrackRepo
}

func (a *trackRegistrationUseCase) Register(id, trackName, artist string) (model.Track, error) {
	newTrack := model.NewTrack(id, trackName, artist)
	return a.repo.Insert(newTrack)

}

func NewTrackRegistrationUseCase(repo repository.TrackRepo) TrackRegistrationUseCase {
	return &trackRegistrationUseCase{
		repo: repo,
	}
}
