package Models

import "testing"

func TestAbleToInitializeAnswerObject(t *testing.T) {
	testAnswers := Answers{
		AnswerOne: "Hello",
	}

	if testAnswers.AnswerOne == "" || testAnswers.AnswerOne != "Hello" {
		t.Errorf("Unable to initialize Answer object")
	}
}

func TestAbleToInitializeRiddleObject(t *testing.T) {
	testRiddle := Riddle{
		Question: "testing",
	}
	if testRiddle.Question == "" || testRiddle.Question != "testing" {
		t.Errorf("Unable to initialize Riddle object")
	}
}
