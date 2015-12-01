package guessit

import (
	"fmt"

	"github.com/Wessie/seasonals/anime/guessit"
)

func ResolveRawName(c Channels) {
	for a := range c.In {
		if a.RawName == "" {
			c.Err <- Error{
				Anime: a,
				Error: fmt.Errorf("no raw name found"),
			}
			continue
		}

		r, err := guessit.Guess(a.RawName)
		if err != nil {
			c.Err <- Error{
				Anime: a,
				Error: err,
			}
			continue
		}

		a.GuessIt = &r
		c.Out <- res
	}
}
