package trainings

import(
	personaldata "github.com/Yandex-Practicum/tracker/internal/personaldata"
	"time"
	"strings"
	"fmt"
    "strconv"
	spentenergy "github.com/Yandex-Practicum/tracker/internal/spentenergy"
)
type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	Personal     personaldata.Personal

	// TODO: добавить поля
}

func (t *Training) Parse(datastring string) (err error) {
	// Разделить строку на слайс строк
    parts := strings.Split(datastring, ",")

    // Проверить длину слайса
    if len(parts) != 3 {
        return fmt.Errorf("некорректный формат строки")
    }

    // Преобразовать первый элемент слайса в тип int
    parts[0] = strings.Replace(parts[0], ",", "", -1) // Удаление запятых
    parts[0] = strings.Replace(parts[0], "+", "", -1) // Удаление знака плюса
    steps, err := strconv.Atoi(parts[0])
    if err != nil {
    return err
    }

    t.Steps = steps

    if t.Steps <= 0 {
    return fmt.Errorf("количество шагов должно быть больше нуля")
    }

    // Сохранить значение типа тренировки
    t.TrainingType = parts[1]

    // Преобразовать третий элемент слайса в time.Duration
    duration, err := time.ParseDuration(parts[2])
    if err != nil {
        return err
    }
    t.Duration = duration

    if t.Duration <= 0 {
    return fmt.Errorf("продолжительность должна быть больше нуля")
    }

    return nil
// TODO: реализовать функцию
}

func (t Training) ActionInfo() (string, error) {
	// Вычислить дистанцию
    distance := spentenergy.Distance(t.Steps, t.Personal.Height)
   
    // Вычислить среднюю скорость
    meanSpeed:= spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

    // Рассчитать калории в зависимости от типа тренировки
    var calories float64
    var err error
    switch t.TrainingType {
    case "Бег":
        calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration) 
        if err != nil {
            return "", err
        }
    case "Ходьба":
        calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration) 
        if err != nil {
            return "", err
        }
    default:
        return "", fmt.Errorf("неизвестный тип тренировки")
    }

    // Сформировать строку с данными о тренировке
    result := fmt.Sprintf("Тип тренировки: %s\n"+
        "Длительность: %.2f ч.\n"+
        "Дистанция: %.2f км.\n"+
        "Скорость: %.2f км/ч\n"+
        "Сожгли калорий: %.2f\n",
        t.TrainingType,
        t.Duration.Hours(),
        distance,
        meanSpeed,
        calories)

    return result, nil
	// TODO: реализовать функцию
}
func (t Training) Print() {
    fmt.Println("Информация о тренировке:")
    fmt.Printf("Шаги: %d\n", t.Steps)
    fmt.Printf("Тип тренировки: %s\n", t.TrainingType)
    fmt.Printf("Длительность: %v\n", t.Duration)
    
}
