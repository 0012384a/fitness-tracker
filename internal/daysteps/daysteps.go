package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
	mInKm      = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже

	s := strings.Split(data, ",")
	if len(s) != 2 {
		return 0, 0, errors.New("error in function parsePackage")
	}
	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, err
	}
	duration, err := time.ParseDuration(s[1])
	if err != nil {
		return 0, 0, err
	}

	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, duration, err := parsePackage(data)
	if err != nil {
		return err.Error()
	}
	if steps <= 0 {
		return " "
	}
	distance := (float64(steps) * StepLength) / float64(mInKm)
	walkingSpentCalories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	result := fmt.Sprintf("Количество шагов: %d \nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, walkingSpentCalories)
	return result
}
