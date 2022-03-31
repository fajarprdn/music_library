package model

type Track struct {
	Id        string
	TrackName string
	Artist    string
}

func (s *Track) SetTrackName(trackName string) {
	s.TrackName = trackName
}

func (s *Track) SetId(id string) {
	s.Id = id
}

func (s *Track) SetArtist(artist string) {
	s.Artist = artist
}

func (s *Track) GetTrackId() string {
	return s.Id
}

func (s *Track) GetTrackName() string {
	return s.TrackName
}

func (s *Track) GetArtist() string {
	return s.Artist
}

func NewTrack(id, name, artist string) Track {
	return Track{
		Id:        id,
		TrackName: name,
		Artist:    artist,
	}
}
