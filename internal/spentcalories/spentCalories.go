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
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// ваш код ниже
	var activity string

	s := strings.Split(data, ",")
	if len(s) != 3 {
		return 0, " ", 0, errors.New("error in function parseTraining")
	}
	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, " ", 0, err
	}
	duration, err := time.ParseDuration(s[2])
	if err != nil {
		return 0, " ", 0, err
	}
	if s[1] == "Бег" || s[1] == "Ходьба" {
		activity = s[1]
	}
	return steps, activity, duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	// ваш код ниже
	return (float64(steps) * lenStep) / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже
	if duration <= 0 {
		return 0
	}
	return distance(steps) / duration.Hours()
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// ваш код ниже
	var result string
	var spentCalories float64
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return err.Error()
	}
	switch activity {
	case "Бег":
		spentCalories = RunningSpentCalories(steps, weight, duration)
	case "Ходьба":
		spentCalories = WalkingSpentCalories(steps, weight, height, duration)
	}

	d := duration.Hours()
	distance := distance(steps)
	meanSpeed := meanSpeed(steps, duration)
	result = fmt.Sprintf("Тип тренировки: %s \nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f \n", activity, d, distance, meanSpeed, spentCalories)
	return result
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// ваш код здесь
	meanSpeed := meanSpeed(steps, duration)
	return ((runningCaloriesMeanSpeedMultiplier * meanSpeed) - runningCaloriesMeanSpeedShift) * weight
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// ваш код здесь
	meanSpeed := meanSpeed(steps, duration)

	return ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH
}
