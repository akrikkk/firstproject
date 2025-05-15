package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	slice := strings.Fields(data)
	if len(slice) != 3 {
		return 0, "", 0, fmt.Errorf("expected 3 values, got")
	}
	steps, err := strconv.Atoi(slice[1])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid steps format")
	}
	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid duration format")
	}
	return steps, slice[0], duration, nil
}

func distance(steps int, height float64) float64 {
	lenghtOfStep := stepLengthCoefficient * height
	meters := float64(steps) * lenghtOfStep
	return meters / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	kilometers := distance(steps, height)
	durationInHours := duration.Hours()
	return kilometers / durationInHours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, typeOfTreaning, t, err := parseTraining(data)
	if err != nil {
		return " ", nil
	}
	switch typeOfTreaning {
	case "Ходьба":
		kilometers := distance(steps, height)
		averageSpeed := meanSpeed(steps, height, t)
		durationInMinutes := t.Minutes()
		durationHours := t.Hours()
		calories := ((weight * averageSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient
		return fmt.Sprintf(
			"Тип тренировки: %s\n"+
				"Длительность: %.2f ч.\n"+
				"Дистанция: %.2f км\n"+
				"Скорость: %.2f км/ч\n"+
				"Сожгли калорий: %.2f",
			typeOfTreaning,
			durationHours,
			kilometers,
			averageSpeed,
			calories), nil
	case "Бег":
		kilometers := distance(steps, height)
		averageSpeed := meanSpeed(steps, height, t)
		durationInMinutes := t.Minutes()
		durationHours := t.Hours()
		calories := ((weight * averageSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient
		return fmt.Sprintf(
			"Тип тренировки: %s\n"+
				"Длительность: %.2f ч.\n"+
				"Дистанция: %.2f км\n"+
				"Скорость: %.2f км/ч\n"+
				"Сожгли калорий: %.2f",
			typeOfTreaning,
			durationHours,
			kilometers,
			averageSpeed,
			calories), nil
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Incorrect parameter format")
	}
	AverageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return ((weight * AverageSpeed * durationInMinutes) / minInH), nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Incorrect parameter format")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * averageSpeed * durationInMinutes) / minInH
	return (walkingCaloriesCoefficient * calories), nil
}
