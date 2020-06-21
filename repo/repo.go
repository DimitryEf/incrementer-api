package repo

import "github.com/DimitryEf/incrementer-api/model"

// Repo - интерфейс для работы с репозиторием
type Repo interface {
	GetNumber() (int64, error)                             // Получить значение инкрементора
	SetNumber(num int64) error                             // Установить значение инкрементора
	GetParams() (model.Params, error)                      // Получить параметры
	SetMaximumValue(maximumValue int64) error              // Установить максимальное значение
	SetStepValue(stepValue int64) error                    // Установить значение шага инкрементора
	SetParams(maximumValue, stepValue int64) error         // Установить параметры
	SetMaximumValueAndZeroNumber(maximumValue int64) error // Установить максимум и обнулить значение
}
