package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Grades []int
}

func findStudent(students map[string]Student, name string) (string, Student, bool) {
	lowerName := strings.ToLower(name)
	for key, student := range students {
		if strings.ToLower(key) == lowerName {
			return key, student, true
		}
	}
	return "", Student{}, false
}

func main() {
	students := map[string]Student{
		"Ярослав":   {Grades: []int{85, 90, 78, 92, 88}},
		"Ангеліна":  {Grades: []int{70, 75, 80, 85, 90}},
		"Олена":     {Grades: []int{95, 93, 97, 99, 94}},
		"Станіслав": {Grades: []int{60, 65, 70, 75, 80}},
		"Олексій":   {Grades: []int{88, 82, 85, 87, 90}},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nМеню:")
		fmt.Println("1. Створити студента")
		fmt.Println("2. Додати оцінку студенту")
		fmt.Println("3. Вивести студента з оцінками")
		fmt.Println("4. Вивести середню оцінку студента")
		fmt.Println("5. Вийти з програми")
		fmt.Print("Оберіть опцію: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			fmt.Print("Введіть повне ім'я студента: ")
			scanner.Scan()
			studentName := scanner.Text()

			if _, _, exists := findStudent(students, studentName); exists {
				fmt.Println("Студент з таким ім'ям вже існує.")
			} else {
				students[studentName] = Student{Grades: []int{}}
				fmt.Println("Студента створено.")
			}

		case "2":
			fmt.Print("Введіть повне ім'я студента: ")
			scanner.Scan()
			studentName := scanner.Text()

			if key, student, exists := findStudent(students, studentName); exists {
				fmt.Print("Введіть оцінку (0-100): ")
				scanner.Scan()
				gradeStr := scanner.Text()

				grade, err := strconv.Atoi(gradeStr)
				if err != nil || grade < 0 || grade > 100 {
					fmt.Println("Невірна оцінка.")
				} else {
					student.Grades = append(student.Grades, grade)
					students[key] = student
					fmt.Println("Оцінку додано.")
				}
			} else {
				fmt.Println("Студента не знайдено.")
			}

		case "3":
			fmt.Print("Введіть повне ім'я студента: ")
			scanner.Scan()
			studentName := scanner.Text()

			if key, student, exists := findStudent(students, studentName); exists {
				fmt.Printf("Оцінки студента %s: %v\n", key, student.Grades)
			} else {
				fmt.Println("Студента не знайдено.")
			}

		case "4":
			fmt.Print("Введіть повне ім'я студента: ")
			scanner.Scan()
			studentName := scanner.Text()

			if key, student, exists := findStudent(students, studentName); exists {
				if len(student.Grades) == 0 {
					fmt.Println("У студента немає оцінок.")
				} else {
					sum := 0
					for _, grade := range student.Grades {
						sum += grade
					}
					average := float64(sum) / float64(len(student.Grades))
					fmt.Printf("Середня оцінка студента %s: %.2f\n", key, average)
				}
			} else {
				fmt.Println("Студента не знайдено.")
			}

		case "5":
			fmt.Println("Вихід з програми.")
			return

		default:
			fmt.Println("Невірна опція. Спробуйте ще раз.")
		}
	}
}
