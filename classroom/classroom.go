package classroom

import (
	"errors"
	"sync"
)

// Errors returned by BulkReview
var (
	ErrInvalidIndex        = errors.New("invalid index; index must be between 1 and 15")
	ErrInvalidScore        = errors.New("invalid score; score must be between -5 and 5")
	ErrInvalidReviewFormat = errors.New("invalid review format; review should contains index and score pairs")
)

// Student represents a student in the classroom
type Student struct {
	ID    int
	Score int
}

// NewClassroom returns a new classroom with given students
func NewClassroom(n int) *Classroom {
	return &Classroom{
		students: make([]Student, n),
	}
}

// Classroom represents a classroom
type Classroom struct {
	mu       sync.Mutex
	students []Student
}

// GetGroupAStudents returns the first 10 students in the classroom
func (c *Classroom) GetGroupAStudents() []Student {
	return c.students[:10]
}

// GetGroupBStudents returns the last 5 students in the classroom
func (c *Classroom) GetGroupBStudents() []Student {
	return c.students[10:]
}

// ResetScore resets the score of all students in the classroom to 0
func (c *Classroom) ResetScores() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := range c.students {
		c.students[i].Score = 0
	}
}

// BulkReview applies the given points to the students in the classroom
// based on the given index and score pairs
// Example: BulkReview(1,5, 2,-5, 3,5, 4,-5, 5,1) will update the score of the first 5 students
// in the classroom by 5, -5, 5, -5, and 1 respectively
func (c *Classroom) BulkReview(pp ...int) error {
	// check if the number of arguments is even
	if len(pp)%2 != 0 {
		return ErrInvalidReviewFormat
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	for i := 0; i < len(pp); i += 2 {
		index, score := pp[i], pp[i+1]

		// check if index is valid
		if index > len(c.students) || index < 0 {
			return ErrInvalidIndex
		}

		// check if score is valid
		if score > 5 || score < -5 {
			return ErrInvalidScore
		}

		// update the score of the student
		// index-1 because the index starts from 1
		c.students[index-1].Score += score
	}

	// resort the classroom after updating the score
	c.resort()

	return nil
}
