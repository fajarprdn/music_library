package manager

import (
	"music_library/repository"
)

type RepoManager interface {
	TrackRepo() repository.TrackRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) TrackRepo() repository.TrackRepo {
	return repository.NewTrackRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{infra: infra}
}
