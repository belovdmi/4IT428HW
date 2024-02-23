package assignment

import (
	"errors"
	"sort"
	"strings"
)

func Reverse(arr []string) []string {
	arrLen := len(arr)

	if arrLen == 0 {
		return arr
	}

	for i, j := 0, arrLen-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func Palindrome(arr []string) bool {
	arrLen := len(arr)

	if arrLen == 0 {
		return true
	}

	for i, j := 0, arrLen-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j]{ 
			return false
		}
	}

	return true
}


func Anagram(source, target string) bool {
		if len(source)!=len(target) {
			return false
		}
		
		// sort source & target arrays
		sourceArray := []rune(source)
		sort.Slice(sourceArray, func(i, j int) bool {
			return sourceArray[i] < sourceArray[j]
		})
		targetArray := []rune(target)
		sort.Slice(targetArray, func(i, j int) bool {
			return targetArray[i] < targetArray[j]
		})
	
		// Loop through the arrays and check character by character 
		for i := 0; i < len(sourceArray); i++ {
			if sourceArray[i] != targetArray[i] {
				return false
			}
		}
		return true
}

func ReplaceDigits(toBeRepaced, replacement string ) string {
	digits := []string{"0","1","2","3","4","5", "6", "7", "8", "9"}

	for i := 0; i < len(digits) - 1; i++ {
		toBeRepaced = strings.ReplaceAll(toBeRepaced, digits[i],replacement)
	}

	return toBeRepaced
}

type Student interface {
	Name() string
}

type Course interface {
	Name() string
	EnrollStudent(s Student) error
}

type DataSource interface {
	ReadStudent(studentID int) (Student, error)
	ReadCourse(courseID int) (Course, error)
}

func EnrollStudentToCourse(dataSource DataSource, studentId int, courseId int) error {
	student, errStudent := dataSource.ReadStudent(studentId)

	if errStudent != nil {
		return errors.New("invalid student id")
	}

	course, errCouse := dataSource.ReadCourse(courseId)

	if errCouse != nil {
		return errors.New("invalid course id")
	}

	return course.EnrollStudent(student)
}