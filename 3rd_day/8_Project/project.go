package main

import (
	"errors"
	"fmt"
)

type student struct {
	students map[string][]int
}

func calculateGrade(percentage float64) string {
	if percentage >= 90 {
		return "A"
	} else if percentage >= 75 {
		return "B"
	} else if percentage >= 60 {
		return "C"
	}
	return "Fail"
}
func getStudentResult(student map[string][]int, name string) (string, error) {
	marks, exists := student[name]
	if !exists {
		return "", errors.New("Student Not Found!")
	}
	total := 0
	for _, mark := range marks {
		total += mark
	}
	percentage := float64(total) / float64(len(marks))
	grade := calculateGrade(percentage)
	return grade, nil
}

func main() {
	students := make(map[string][]int)
	students["Vivek"] = []int{89, 90, 70, 56, 88, 23}
	students["Raju"] = []int{90, 89, 88, 98, 99, 100}
	students["Shivam"] = []int{20, 61, 20, 10, 1, 23}
	students["Aanand"] = []int{99, 12, 11, 25, 82, 30}
	var name string
	fmt.Println("Enter Student Name:")
	fmt.Scanln(&name)
	grade, err := getStudentResult(students, name)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Grade of Student is:", grade)
	}

}
