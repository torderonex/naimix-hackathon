package service

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/webapi"
	"log/slog"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ComparisonService struct {
	gpt webapi.GptApi
}

func NewComparisonService(api webapi.GptApi) *ComparisonService {
	return &ComparisonService{api}
}

func (s *ComparisonService) GetComparisonBetween2WithDescription(date1 time.Time, date2 time.Time) (entities.ComparisonResult, error) {
	taro1 := CalculateTaro(date1)
	taro2 := CalculateTaro(date2)

	prompt := fmt.Sprintf(`
Я отправлю 2 карты с номером высших арканов таро. Ты должен коротко объяснить насколько эти 2 человека смогут поладить в деловых отношениях и работе.
Пиши по структуре: краткое толкование карт, совместимость, совместимость в процентах. 
В конце своего объяснения обязательно пишешь совместимость по своему мнению в формате "Совместмость: X%%". 
1)%d %s
2)%d %s
`, taro1.Number, taro1.Name, taro2.Number, taro2.Name)
	ans, err := s.gpt.DoPrompt(prompt)
	if err != nil {
		slog.Error(err.Error())
	}
	comp, _ := ExtractCompatibility(ans)
	return entities.ComparisonResult{
		Description: RemoveCompatibility(ans),
		Percent:     comp,
	}, nil
}

func (s *ComparisonService) GetTeamComparison(part []entities.Participant, birthDate time.Time) ([]int, error) {
	var kuslievs string
	medvedev := CalculateTaro(birthDate)
	for _, p := range part {
		tmp := CalculateTaro(p.BirthDate)
		kuslievs += fmt.Sprintf("%d %s;", tmp.Number, tmp.Name)
	}
	prompt := fmt.Sprintf(`
Я отправлю %d карт с номером высших арканов таро и одну отдельно.  
Отправь в формате "Совместмость: X%%" для каждой карты из списка совместимость с картой которую я отправил отдельно.
Больше ничего не пиши. Только список совместимостей в нужном формате. Я должен получить от тебя %d совместимостей.
Карты:
1)%s
2)%d %s
`, len(part), len(part), kuslievs, medvedev.Number, medvedev.Name)
	ans, err := s.gpt.DoPrompt(prompt)
	slog.Info(ans)
	if err != nil {
		slog.Error(err.Error())
		return GenerateNumsFromParticipants(part), nil
	}
	res, err := ExtractAllCompatibilities(ans)
	if err != nil {
		return GenerateNumsFromParticipants(part), nil
	}
	return res, nil
}

func RemoveCompatibility(response string) string {
	re := regexp.MustCompile(`Совместимость:\s*\d+%`)
	return strings.TrimSpace(re.ReplaceAllString(response, ""))
}

func ExtractCompatibility(response string) (int, error) {
	re := regexp.MustCompile(`Совместимость:\s*(\d+)%`)

	matches := re.FindStringSubmatch(response)
	if len(matches) < 2 {
		return rand.Intn(101), errors.New("совместимость не найдена в ответе")
	}

	compatibility, err := strconv.Atoi(matches[1])
	if err != nil {
		return rand.Intn(101), errors.New("не удалось преобразовать совместимость в число")
	}

	return compatibility, nil
}

func GenerateNRandomNums(n int) []int {
	var res []int
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(101))
	}
	return res
}

func ExtractAllCompatibilities(response string) ([]int, error) {
	re := regexp.MustCompile(`Совместимость:\s*(\d+)%`)
	matches := re.FindAllStringSubmatch(response, -1)

	var compatibilities []int
	for _, match := range matches {
		if len(match) > 1 {
			compatibility, err := strconv.Atoi(match[1])
			if err != nil {
				return nil, err
			}
			compatibilities = append(compatibilities, compatibility)
		}
	}
	return compatibilities, nil
}

func GenerateNumsFromParticipants(part []entities.Participant) []int {
	results := make([]int, len(part))

	for i, p := range part {
		data := []byte(p.Name + p.TeamName + p.Role + p.BirthPlace + p.BirthDate.Format("2006-01-02"))
		hash := md5.Sum(data)

		results[i] = int(binary.BigEndian.Uint32(hash[:4])) % 100 // Ограничим числа диапазоном 0-99
	}

	return results
}
