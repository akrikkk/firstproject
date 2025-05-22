package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

// я простоне понимаю, что длеаю не так, по идеи вывод корректный, логика правильная, прошу прощения за мой тупизм(
const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(input string) (steps int, duration time.Duration, err error) {
	input = strings.ReplaceAll(input, " ", "")

	if input == "" {
		return 0, 0, fmt.Errorf("пустая строка")
	}

	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неверный формат данных")
	}
	steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("неверный формат шагов")
	}
	if steps < 0 {
		return 0, 0, fmt.Errorf("отрицательные шаги")
	}
	if steps == 0 {
		return 0, 0, fmt.Errorf("ноль шагов")
	}
	duration, err = time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("неверный формат продолжительности")
	}
	if duration < 0 {
		return 0, 0, fmt.Errorf("отрицательная продолжительность")
	}
	if duration == 0 {
		return 0, 0, fmt.Errorf("нулевая продолжительность")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}

	if weight <= 0 || height <= 0 {
		return fmt.Sprintf("Error: invalid parameters (weight: %.1f kg, height: %.2f m) - must be positive", weight, height)
	}

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return fmt.Sprintf("Error calculating calories for %q: %v", data, err)
	}

	distance := float64(steps) * stepLength / mInKm

	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		steps, distance, calories)
}
