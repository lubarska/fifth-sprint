package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка на корректность входных параметров
	if steps < 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect input parameters : steps=%d, weight=%.2f, height=%.2f, duration=%v", steps, weight, height, duration)
	}
	if steps == 0 {
		return 0, fmt.Errorf("steps should be >0")
	}
	// Рассчитать среднюю скорость
	meanSpeed := MeanSpeed(steps, height, duration)

	// Перевести продолжительность в минуты
	durationInMinutes := duration.Minutes()

	// Рассчитать количество калорий
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	// Умножить на корректирующий коэффициент
	calories *= walkingCaloriesCoefficient

	return calories, nil
	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка на корректность входных параметров
	if steps < 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect input parameters: steps=%d, weight=%.2f, height=%.2f, duration=%v", steps, weight, height, duration)
	}
	if steps == 0 {
		return 0, fmt.Errorf("steps should be >0")
	}
	// Рассчитать среднюю скорость
	meanSpeed := MeanSpeed(steps, height, duration)

	// Перевести продолжительность в минуты
	durationInMinutes := duration.Minutes()

	// Рассчитать количество калорий
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil

	// TODO: реализовать функцию
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Проверка на отрицательные шаги
	if steps < 0 {
		return 0
	}
	if steps == 0 {
		return 0
	}
	// Проверка продолжительности
	if duration <= 0 {
		return 0
	}

	// Вычисляем дистанцию
	distance := Distance(steps, height)

	// Переводим продолжительность в часы
	durationHours := duration.Hours()

	// Вычисляем среднюю скорость
	meanSpeed := distance / durationHours

	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	// Приводим количество шагов к float64
	stepsFloat := float64(steps)

	// Рассчитываем длину шага
	stepLength := height * stepLengthCoefficient

	// Вычисляем дистанцию в метрах
	distanceMeters := stepsFloat * stepLength

	// Переводим дистанцию в километры
	distanceKm := distanceMeters / mInKm

	return distanceKm
}
