package anime

import "github.com/Wessie/seasonals/anime/guessit"

type Channels struct {
	In   <-chan *Anime
	Out  chan *Anime
	Err  chan Error
	Done chan struct{}
}

type Error struct {
	Anime *Anime
	Error error
}

type Anime struct {
	// ID is an anidb identification number of this title
	ID int
	// Name as resolved by an anime database
	Name string
	// RawName is the name used before being resolved by us
	RawName string

	URL string

	// GuessIt is the raw result from the guessit API
	GuessIt *guessit.Result
}
