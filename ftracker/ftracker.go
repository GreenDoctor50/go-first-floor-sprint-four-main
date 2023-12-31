package ftracker

import (
	"fmt"
	"math"
)

const (
	// !Основные константы, необходимые для расчетов.
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
	// !Константы для расчета калорий, расходуемых при ходьбе.
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
	// !Константы для расчета калорий, расходуемых при плавании.
	swimmingLenStep                  = 1.38 // длина одного гребка.
	swimmingCaloriesMeanSpeedShift   = 1.1  // среднее количество сжигаемых колорий при плавании относительно скорости.
	swimmingCaloriesWeightMultiplier = 2    // множитель веса при плавании.
	// !Константы для расчета калорий, расходуемых при беге.
	runningCaloriesMeanSpeedMultiplier = 18   // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 1.79 // среднее количество сжигаемых калорий при беге.
)

/*
	 distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.

			Параметры:
			action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
*/
func distance(action int) float64 {
	return float64(action) * lenStep / mInKm
}

/*
	 meanSpeed возвращает значение средней скорости движения во время тренировки км./ч.

			Параметры:
			action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
			duration float64 — длительность тренировки в часах.
*/
func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return distance(action) / duration
}

/*
	 swimmingMeanSpeed возвращает среднюю скорость при плавании км./ч..

			Параметры:
			    lengthPool int — длина бассейна в метрах.
			    countPool int — сколько раз пользователь переплыл бассейн.
			    duration float64 — длительность тренировки в часах.
*/
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool) * float64(countPool) / mInKm / duration
}

/*
	 RunningSpentCalories возвращает количество потраченных колорий при беге.

			Параметры:

			action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
			weight float64 — вес пользователя в кг.
			duration float64 — длительность тренировки в часах.
*/
func RunningSpentCalories(action int, weight, duration float64) float64 {
	//((18 * СредняяСкоростьВКм/ч * 1.79) * ВесСпортсменаВКг / mInKM * ВремяТренировкиВЧасах * minInH)
	return ((runningCaloriesMeanSpeedMultiplier * meanSpeed(action, duration) * runningCaloriesMeanSpeedShift) * weight / mInKm * duration * minInH)
}

/*
WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.

	Параметры:

	action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
	duration float64 — длительность тренировки в часах.
	weight float64 — вес пользователя в кг.
	height float64 — рост пользователя в м.
*/
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	//((0.035 * ВесСпортсменаВКг + (СредняяСкоростьВМетрахВСекунду**2 / РостВМетрах) * 0.029 * ВесСпортсменаВКг) * ВремяТренировкиВЧасах * minInH)
	return ((walkingCaloriesWeightMultiplier*weight + (math.Pow(meanSpeed(action, duration) * kmhInMsec, 2)/height/cmInM)*walkingSpeedHeightMultiplier*weight) * duration * minInH)
}

/*
	 SwimmingSpentCalories возвращает количество потраченных калорий при плавании.

			Параметры:

			lengthPool int — длина бассейна в метрах.
			countPool int — сколько раз пользователь переплыл бассейн.
			duration float64 — длительность тренировки в часах.
			weight float64 — вес пользователя.
*/
func SwimmingSpentCalories(lenghtPool, countPool int, duration, weight float64) float64 {
	//(СредняяСкоростьВКм/ч * 1.1) * 2 * ВесСпортсменаВКг * ВремяТренеровкиВЧасах
	return ((swimmingMeanSpeed(lenghtPool, countPool, duration) * swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration)
}

/*
	 ShowTrainingInfo возвращает строку с информацией о тренировке.

			Параметры:

			action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
			trainingType string — вид тренировки(Бег, Ходьба, Плавание).
			duration float64 — длительность тренировки в часах.
			weight, height float64 - вес и рост пользователя.
			lengthPool, countPool int - длина бассейна в м. и количество раз переплывания бассейна.
*/
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	switch {
	case trainingType == "Бег":
		distance := distance(action)
		speed := meanSpeed(action, duration)
		calories := RunningSpentCalories(action, weight, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case trainingType == "Ходьба":
		distance := distance(action)
		speed := meanSpeed(action, duration)
		calories := WalkingSpentCalories(action, duration, weight, height/cmInM)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case trainingType == "Плавание":
		distance := distance(countPool)
		speed := swimmingMeanSpeed(lengthPool, countPool, duration)
		calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}
