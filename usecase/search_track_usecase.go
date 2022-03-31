package usecase

import (
	"music_library/model"
	"music_library/repository"
)

type SearchTrackUseCase interface {
	Search(id string) []model.Track
}

type searchTrackUseCase struct {
	repo repository.TrackRepo
}

func (s *searchTrackUseCase) Search(id string) []model.Track {
	if len(id) == 0 {
		return s.repo.GetAll()
	}
	result := s.repo.GetByTrackId(id)
	if len(result.GetTrackId()) == 0 {
		return nil
	} else {
		return []model.Track{result}
	}
}

func NewSearchTrackUseCase(repo repository.TrackRepo) SearchTrackUseCase {
	return &searchTrackUseCase{repo: repo}
}
