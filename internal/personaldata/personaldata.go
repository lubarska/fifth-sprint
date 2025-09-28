package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64 
	//TrainingType string// TODO: добавить поля
}

func (p Personal) Print() {
	fmt.Println("Имя:" , p.Name)
	fmt.Printf("Вес: %.2f кг.\n", p.Weight)
    fmt.Printf("Рост: %.2f м.\n", p.Height)
	// TODO: реализовать функцию
}
