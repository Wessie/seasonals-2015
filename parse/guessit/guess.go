package guessit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Type      string `json:"type"`
	Title     string `json:"title"`
	Container string `json:"container"`
	Group     string `json:"releaseGroup"`

	Series  string `json:"series"`
	Season  int    `json:"season"`
	Episode int    `json:"episodeNumber"`
	Version string `json:"version"`

	Format     string `json:"format"`
	ScreenSize string `json:"screenSize"`
	VideoCodec string `json:"videoCodec"`
	VProfile   string `json:"videoProfile"`

	Channels   float64 `json:"audioChannels"`
	AudioCodec string  `json:"audioCodec"`

	CRC32 string `json:"crc32"`
	Other string `json:"other"`
}

type Error struct {
	message string
	error
}

var ResultCache = map[string]Result{}

func Guess(title string) (Result, error) {
	if Cached, ok := ResultCache[title]; ok == true {
		return Cached, nil
	}
	url := fmt.Sprintf("http://guessit.io/guess?filename=%s", title)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Result{}, Error{"guessit: failed to create request", err}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Result{}, Error{"guessit: failed to send request", err}
	}

	defer resp.Body.Close()

	result := Result{}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Result{}, Error{"guessit: failed to read response", err}
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return Result{}, Error{"guessit: failed to unmarshal response", err}
	}

	ResultCache[title] = result

	return result, nil
}
