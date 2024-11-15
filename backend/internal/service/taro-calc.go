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
		0:  "The Fool",
		1:  "The Magician",
		2:  "The High Priestess",
		3:  "The Empress",
		4:  "The Emperor",
		5:  "The Hierophant",
		6:  "The Lovers",
		7:  "The Chariot",
		8:  "Strength",
		9:  "The Hermit",
		10: "Wheel of Fortune",
		11: "Justice",
		12: "The Hanged Man",
		13: "Death",
		14: "Temperance",
		15: "The Devil",
		16: "The Tower",
		17: "The Star",
		18: "The Moon",
		19: "The Sun",
		20: "Judgement",
		21: "The World",
		22: "The Fool",
	}

	return entities.TarotCard{
		Name:   tarotCards[number],
		Number: number,
	}
}

func CalculateTaro(date time.Time) entities.TarotCard {
	return getTarotCard(calculateTarotNumber(date))
}
