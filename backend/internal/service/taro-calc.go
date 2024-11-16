package service

import (
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"strconv"
	"strings"
	"time"
)

func calculateTarotNumber(birthdate time.Time) int {
	sum := 0
	birthdateStr := birthdate.Format("2006-01-02")
	for _, char := range strings.ReplaceAll(birthdateStr, "-", "") {
		num, _ := strconv.Atoi(string(char))
		sum += num
	}

	for sum > 22 {
		tempSum := 0
		for sum > 0 {
			tempSum += sum % 10
			sum /= 10
		}
		sum = tempSum
	}

	return sum
}

func getTarotCard(number int) entities.TarotCard {
	tarotCards := map[int]string{
		0:  "Шут",
		1:  "Маг",
		2:  "Жрица",
		3:  "Императрица",
		4:  "Император",
		5:  "Иерофант",
		6:  "Влюблённые",
		7:  "Колесница",
		8:  "Сила",
		9:  "Отшельник",
		10: "Колесо Фортуны",
		11: "Справедливость",
		12: "Повешенный",
		13: "Смерть",
		14: "Умеренность",
		15: "Дьявол",
		16: "Башня",
		17: "Звезда",
		18: "Луна",
		19: "Солнце",
		20: "Суд",
		21: "Мир",
		22: "Шут",
	}

	return entities.TarotCard{
		Name:   tarotCards[number],
		Number: number,
	}
}

func CalculateTaro(date time.Time) entities.TarotCard {
	return getTarotCard(calculateTarotNumber(date))
}
