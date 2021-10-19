package ApiHandlers

import "testing"

func TestAbletoCreateMockRiddles(t *testing.T) {
	mockRiddlesArr := mockRiddles()

	if len(mockRiddlesArr) == 0 {
		t.Errorf("Unable to create mock riddles from the mockRiddle function")
	}

	if mockRiddlesArr[0].Username != "Brendon" {
		t.Errorf("Unable to create mock riddles should container BRENDON!")
	}
}
