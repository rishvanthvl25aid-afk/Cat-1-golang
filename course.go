package edtech

import (
	"encoding/json"
	"errors"
	"os"
)

type Course struct {
	CourseCode string `json:"courseCode"`
	Title      string `json:"title"`
	Instructor string `json:"instructor"`
	Credits    int    `json:"credits"`
	MaxSeats   int    `json:"maxSeats"`
}

const databaseFile = "course/courses.json"

// Load courses from JSON file
func loadCourses() ([]Course, error) {

	data, err := os.ReadFile(databaseFile)

	if err != nil {
		if os.IsNotExist(err) {
			return []Course{}, nil
		}
		return nil, err
	}

	var courses []Course

	err = json.Unmarshal(data, &courses)

	if err != nil {
		return nil, err
	}

	return courses, nil
}

// Save courses to JSON file
func saveCourses(courses []Course) error {

	data, err := json.MarshalIndent(courses, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(databaseFile, data, 0644)
}

// Create a new course
func CreateCourse(c Course) error {

	if c.Credits < 1 || c.Credits > 4 {
		return errors.New("credits must be between 1 and 4")
	}

	courses, err := loadCourses()

	if err != nil {
		return err
	}

	// Check duplicate course code
	for _, course := range courses {
		if course.CourseCode == c.CourseCode {
			return errors.New("course already exists")
		}
	}

	courses = append(courses, c)

	return saveCourses(courses)
}

// Get course by course code
func GetCourse(code string) (*Course, error) {

	courses, err := loadCourses()

	if err != nil {
		return nil, err
	}

	for i := range courses {

		if courses[i].CourseCode == code {
			return &courses[i], nil
		}
	}

	return nil, errors.New("course not found")
}

// Update instructor
func UpdateInstructor(code string, newInstructor string) error {

	courses, err := loadCourses()

	if err != nil {
		return err
	}

	for i := range courses {

		if courses[i].CourseCode == code {

			courses[i].Instructor = newInstructor

			return saveCourses(courses)
		}
	}

	return errors.New("course not found")
}

// Drop course
func DropCourse(code string) error {

	courses, err := loadCourses()

	if err != nil {
		return err
	}

	for i := range courses {

		if courses[i].CourseCode == code {

			courses = append(courses[:i], courses[i+1:]...)

			return saveCourses(courses)
		}
	}

	return errors.New("course not found")
}
