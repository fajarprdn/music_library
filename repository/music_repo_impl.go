package repository

import (
	"github.com/jmoiron/sqlx"
	"music_library/logger"
	"music_library/model"
	"music_library/utils"
)

type TrackRepoImpl struct {
	trackDb *sqlx.DB
}

func (p *TrackRepoImpl) Insert(newTrack model.Track) (model.Track, error) {
	id := utils.GetUuid()
	newTrack.Id = id
	//timestamp := time.Now()
	_, err := p.trackDb.Exec("insert into music(id,name,artist) "+
		"values ($1,$2,$3)",
		newTrack.Id, newTrack.TrackName, newTrack.Artist)
	if err != nil {
		logger.Log.Error().Err(err).Str("service", "db-product-insert").Msg("Insert product failed")
		return model.Track{}, err
	}
	return newTrack, nil
}

func (p *TrackRepoImpl) GetByTrackId(id string) model.Track {
	var selectedSong model.Track
	//for _, song := range *p.trackDb {
	//	si := song.GetTrackId()
	//	if si == id {
	//		selectedSong = song
	//		break
	//	}
	//}
	return selectedSong
}

func (p *TrackRepoImpl) GetAll() []model.Track {
	return nil
}

func NewTrackRepo(trackDb *sqlx.DB) TrackRepo {
	trackRepo := TrackRepoImpl{
		trackDb: trackDb,
	}
	return &trackRepo
}
