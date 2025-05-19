package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	stepLength = 0.65 // Длина одного шага в метрах
	mInKm      = 1000 // Количество метров в одном километре
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	slice := strings.Split(data, ",")
	if len(slice) != 2 {
		return 0, 0, errors.New("expected format 'steps,duration'")
	}

	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, err
	}

	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, err
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		return fmt.Sprintf("Error in data %q: %v", data, err)
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
		"Шаги: %d\n"+
			"Дистанция: %.2f km\n"+
			"Калорий сожжено: %.2f kcal",
		steps, distance, calories)
}
