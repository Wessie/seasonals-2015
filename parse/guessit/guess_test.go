package guessit

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestGuess(t *testing.T) {
	res, err := Guess("[UTW-Mazui]_Toaru_Kagaku_no_Railgun_S_-_14_[720p][3A139231].mkv")
	if err != nil {
		t.Error(err)
	}
	a, _ := json.MarshalIndent(res, "", " ")
	fmt.Println(string(a))
	if res.Series != "Toaru Kagaku no Railgun S" ||
		res.Group != "UTW-Mazui" ||
		res.Container != "mkv" ||
		res.Episode != 14 ||
		res.ScreenSize != "720p" {
		t.Error("incorrect data")
	}
}
