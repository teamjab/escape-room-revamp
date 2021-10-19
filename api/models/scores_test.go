package Models

import "testing"

func TestAbleToInitializeScoreObject(t *testing.T) {
	testAnswers := Scores{
		Username: "Hello",
		Score:    0,
	}

	if testAnswers.Username != "Hello" || testAnswers.Score != 0 {
		t.Errorf("Unable to initialize Score object")
	}
}
