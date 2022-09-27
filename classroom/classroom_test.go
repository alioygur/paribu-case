package classroom

import (
	"sort"
	"testing"
)

// Note for reviewer:
// I'm also know table driven tests, but I think it's not necessary here.

func TestClassroom_BulkReview(t *testing.T) {
	classroom := NewClassroom(5)

	reviews := []int{1, 5, 2, 4, 3, 0, 4, -2, 5, -5}

	if err := classroom.BulkReview(reviews...); err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	if err := classroom.BulkReview(1, 5); err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	reviews[1] += 5

	for i := 0; i < len(reviews); i += 2 {
		if classroom.students[reviews[i]-1].Score != reviews[i+1] {
			t.Errorf("expected score to be %v, got %v", reviews[i+1], classroom.students[reviews[i]-1].Score)
		}
	}
}

func TestClassroom_BulkReview_InvalidFormat(t *testing.T) {
	classroom := NewClassroom(15)

	// invalid format so it should return an error
	// the length of rr should be even
	err := classroom.BulkReview(1, 5, 2)

	if err != ErrInvalidReviewFormat {
		t.Errorf("expected error to be %v, got %v", ErrInvalidReviewFormat, err)
	}
}

func TestClassroom_BulkReview_InvalidIndex(t *testing.T) {
	Classroom := NewClassroom(15)

	// index 16 out of range so it should return an error
	err := Classroom.BulkReview(16, 5)

	if err != ErrInvalidIndex {
		t.Errorf("expected error to be %v, got %v", ErrInvalidIndex, err)
	}
}

func TestClassroom_BulkReview_InvalidScore(t *testing.T) {
	Classroom := NewClassroom(15)

	// score 6 out of range so it should return an error
	err := Classroom.BulkReview(1, 6)

	if err != ErrInvalidScore {
		t.Errorf("expected error to be %v, got %v", ErrInvalidScore, err)
	}
}

func TestSort(t *testing.T) {

	classroom := NewClassroom(15)

	if err := classroom.BulkReview(1, 5, 2, -5, 3, 5, 4, -5, 5, 1); err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	if !sort.IsSorted(SortByScore(classroom.students)) {
		t.Errorf("expected classroom to be sorted")
	}
}

func TestClassroom_ResetScores(t *testing.T) {
	classroom := NewClassroom(15)

	if err := classroom.BulkReview(1, 5, 2, -5, 3, 5, 4, -5, 5, 1); err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	classroom.ResetScores()

	for _, student := range classroom.students {
		if student.Score != 0 {
			t.Errorf("expected score to be 0, got %v", student.Score)
		}
	}
}
