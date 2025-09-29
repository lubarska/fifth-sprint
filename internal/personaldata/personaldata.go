package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64
	//TrainingType string// TODO: добавить поля
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s, Вес: %.2f кг., Рост: %.2f м.\n", p.Name, p.Weight, p.Height)
	// TODO: реализовать функцию
}
