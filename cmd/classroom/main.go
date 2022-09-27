package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alioygur/paribu-case/classroom"
)

func main() {
	// exit on interrupt
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	// these points will applied to students every 5 seconds
	points := []int{
		1, 3, 2, 3, // 1-2 students will get 3 points
		3, 2, 4, 2, 5, 2, 6, 2, // 3-6 students will get 2 points
		7, 1, 8, 1, 9, 1, 10, 1, // 7-10 students will get 1 point
	}

	fmt.Println("Applying points to students...")

	classroom := classroom.NewClassroom(15)
	for {
		select {
		case <-exit:
			fmt.Println("Exiting...")
			return
		case <-time.After(5 * time.Second):
			classroom.BulkReview(points...)

			// get group A students
			fmt.Println("Group A students:")
			for i, student := range classroom.GetGroupAStudents() {
				fmt.Printf("ID: %d, Score: %d\n", i+1, student.Score)
			}

			// get group B students
			fmt.Println("Group B students:")
			for i, student := range classroom.GetGroupBStudents() {
				fmt.Printf("ID: %d, Score: %d\n", i+1, student.Score)
			}
		}
	}
}
