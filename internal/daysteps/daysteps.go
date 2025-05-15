package daysteps

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
    "time"
    "../spentcalories/spentcalories"// Теперь будет работать
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	slice := strings.Fields(data)
	if len(slice) != 2 {
		return 0, 0, fmt.Errorf("expected 2 values, got")
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps format")
	}
	if steps <= 0 {
		return 0, 0, errors.New("steps must be positive")
	}
	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration format")
	}
	if duration <= 0 {
		return 0, 0, errors.New("duration must be positive")
	}

	return steps, duration, nil
}
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return " "
	}
	if steps <= 0 {
		return " "
	}
	lenght := float64(steps) * stepLength
	distance := lenght / mInKm
	calories, err := WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return " "
	}
	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.",
		steps, distance, calories)
}
