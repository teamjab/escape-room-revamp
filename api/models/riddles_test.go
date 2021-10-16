package Riddles

import "testing"

func TestAbleToInitializeAnswerObject(t *testing.T) {
	testAnswers := Answers{
		AnswerOne: "Hello",
	}

	if testAnswers.AnswerOne == "" || testAnswers.AnswerOne != "Hello" {
		t.Fatalf("Unable to initialize Answer object")
	}
}

func TestAbleToInitializeRiddleObject(t *testing.T) {
	testRiddle := Riddle{
		Question: "testing",
	}
	if testRiddle.Question == "" || testRiddle.Question != "testing" {
		t.Fatalf("Unable to initialize Riddle object")
	}
}
