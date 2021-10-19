package connectionhelper

import "testing"

func TestConstantVariables(t *testing.T) {
	if DBNAME != "escape_room_revamp" {
		t.Errorf("DB name should be escape_room-revamp")
	}

	if RIDDLE_ISSUE != "escape_riddle" {
		t.Errorf("Riddle collection name should be escape_riddle")
	}

	if SCORE_ISSUE != "escape_score" {
		t.Errorf("Score collection name should be escape_score")
	}
}
