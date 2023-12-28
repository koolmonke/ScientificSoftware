package main

import (
	"fmt"
	"sort"
	"time"
)

type WORKER struct {
	Name        string
	title       string
	YearOfEntry int
}

var workers []WORKER

func inputData() {
	for i := 0; i < 10; i++ {
		var worker WORKER

		fmt.Printf("Введите данные для работника %d:\n", i+1)
		fmt.Print("Фамилия и инициалы: ")
		fmt.Scan(&worker.Name)

		fmt.Print("Должность: ")
		fmt.Scan(&worker.title)

		fmt.Print("Год поступления на работу: ")
		fmt.Scan(&worker.YearOfEntry)

		workers = append(workers, worker)
	}

	sort.Slice(workers, func(i, j int) bool {
		return workers[i].Name < workers[j].Name
	})
}

func printWorkersWithExperienceThreshold(thresholdYear int) {
	found := false

	fmt.Printf("\nРаботники со стажем больше %d лет:\n", thresholdYear)
	currentYear := time.Now().Year()
	for _, worker := range workers {
		experience := currentYear - worker.YearOfEntry
		if experience > thresholdYear {
			fmt.Printf("Фамилия и инициалы: %s, Должность: %s, Стаж: %d лет\n",
				worker.Name, worker.title, experience)
			found = true
		}
	}

	if !found {
		fmt.Println("Нет работников со стажем больше заданного значения.")
	}
}

func main() {
	inputData()

	var thresholdYear int
	fmt.Print("\nВведите минимальный стаж работы (в годах): ")
	fmt.Scan(&thresholdYear)

	// Вывод работников со стажем больше заданного значения
	printWorkersWithExperienceThreshold(thresholdYear)
}
