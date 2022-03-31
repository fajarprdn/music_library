package manager

import "music_library/usecase"

type UseCaseManager interface {
	TrackRegistrationUseCase() usecase.TrackRegistrationUseCase
	SearchTrackUseCase() usecase.SearchTrackUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) TrackRegistrationUseCase() usecase.TrackRegistrationUseCase {
	return usecase.NewTrackRegistrationUseCase(u.repo.TrackRepo())
}

func (u *useCaseManager) SearchTrackUseCase() usecase.SearchTrackUseCase {
	return usecase.NewSearchTrackUseCase(u.repo.TrackRepo())
}

func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: manager,
	}
}
