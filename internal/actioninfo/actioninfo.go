package actioninfo

import (
	"fmt"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error) // TODO: добавить методы
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Println("Ошибка парсинга:", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("Ошибка получения информации:", err, info)

		}
	}
	// TODO: реализовать функцию
}
