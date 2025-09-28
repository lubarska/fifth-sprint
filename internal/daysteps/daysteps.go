package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	personaldata "github.com/Yandex-Practicum/tracker/internal/personaldata"
	spentenergy "github.com/Yandex-Practicum/tracker/internal/spentenergy"
    
)

type DaySteps struct {
    Steps int
    Duration time.Duration
    Personal personaldata.Personal
    Hours float64
    Minutes int
    StepLength float64
    Weight float64
    Height float64 
    
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// Разделить строку по символу 'h'
    parts := strings.Split(datastring, ",")
    if len(parts) != 2 {
        return fmt.Errorf("некорректный формат строки: %s", datastring)
    }
    // Извлечь количество шагов
    stepsStr := parts[0]
    stepsStr = strings.Replace(stepsStr, ",", "", -1) // Удаление запятых
    stepsStr = strings.Replace(stepsStr, "+", "", -1) // Удаление знака плюса
    steps, err := strconv.Atoi(stepsStr)
    if err != nil {
    return err
    }
    
    ds.Steps = steps 
    if steps <= 0 {
        return fmt.Errorf(" количество шагов должно быть больше 0")
      }  
    // Преобразовываем второй элемент слайса в time.Duration
    duration, err := time.ParseDuration(parts[1])
    if err != nil {
        return fmt.Errorf("ошибка при преобразовании продолжительности: %v", err)
    }

    if duration <= 0 {
        return  fmt.Errorf("продолжительность должна быть больше 0")
    }
//     // Преобразовать часы в float64
//     hoursStr := strings.Replace(parts[1], ",", ".", -1)
//     hours, err := strconv.ParseFloat(hoursStr, 64)
//     if err != nil {
//     return err
//    }
//     ds.Hours = hours

//     // Определить позицию 'm' для корректного извлечения минут
//     minutesIndex := strings.Index(parts[1], "m")
//     minutesStr := parts[1][:minutesIndex -1]

//     // Преобразовать минуты в int
//     minutes, err := strconv.Atoi(minutesStr)
//     if err != nil {
//     return err
//     }
//     ds.Minutes = minutes

//     // Установить Duration
//     totalMinutes := int(hours*60) + minutes
 //   ds.Duration = time.Duration(totalMinutes) * time.Minute
    ds.Duration = duration
    return nil
// TODO: реализовать функцию
}

func (ds DaySteps) ActionInfo() (string, error) {
	 // Вычислить дистанцию
    distance := spentenergy.Distance(ds.Steps, ds.Personal.Height) // Передаём длину шага
    

    // Вычислить количество сожжённых калорий
    var calories float64
    var err error

        calories, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
        if err != nil {
            return "", err
        }

    // Сформировать строку с данными о прогулке
    result := fmt.Sprintf("Количество шагов: %d.\n"+
        "Дистанция составила %.2f км.\n"+
        "Вы сожгли %.2f ккал.\n",
        ds.Steps,
        distance,
        calories)

    return result, nil

// TODO: реализовать функцию

}
func (ds DaySteps) Print() {
    fmt.Println("Информация о дневной активности:")
    fmt.Printf("Шаги: %d\n", ds.Steps)
    fmt.Printf("Продолжительность: %v\n", ds.Duration)
    // Возможно, вы захотите вывести и другие поля
}