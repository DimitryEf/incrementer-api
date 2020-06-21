package main

import (
	"errors"
)

// Вычисляем максимальное значение типа int для данной архитектуры платформы
// ^uint(0) - максимальное значение uint побитово сдвигаем вправо для бита под знак
const MaximumInt = int64(^uint64(0) >> 1)

// Ошибка при установке отрицательно значения максимума
var (
	ErrNegativeMaximumValue     = errors.New("the maximum value of increment must be a non negative number")
	ErrNegativeStepValue        = errors.New("the step value of increment must be a non negative number")
	ErrMaximumLessThenStepValue = errors.New("the maximum value of increment must be bigger then step value")
)

// Объявляем неэкспортируемую структуру increment, которая реализует интерфейс Incrementor
type Incrementer struct {
	Repo Repo
}

// NewIncrementor - конструктор объекта increment с параметрами по умолчанию, реализующего интерфейс Incrementor.
func NewIncrementer(repo Repo) *Incrementer {
	return &Incrementer{
		Repo: repo,
	}
}

// Возвращает текущее число
func (inc *Incrementer) GetNumber() (int64, error) {
	return inc.Repo.GetNumber()
}

// Увеличивает число
func (inc *Incrementer) IncrementNumber() error {
	p, err := inc.Repo.GetParams()
	if err != nil {
		return err
	}

	// Обнуляем значение, если оно превысило или превысит максимальное
	if p.Number >= p.MaximumValue || p.Number+p.StepValue > p.MaximumValue {
		err = inc.Repo.SetNumber(0)
		if err != nil {
			return err
		}
		return nil
	}

	err = inc.Repo.SetNumber(p.Number + p.StepValue)
	if err != nil {
		return err
	}
	return nil
}

// Устанавливает максимальное число
func (inc *Incrementer) SetMaximumValue(maximumValue int64) error {
	// Проверяем, чтобы maximumValue было неотрицательным, иначе возвращаем ошибку.
	// Отсутствие проверки maximumValue > MaximumInt обусловлено тем, что сам компилятор не допустит
	// переполнения int и выдаст ошибку - overflows int
	if maximumValue < 0 {
		return ErrNegativeMaximumValue
	}
	// Используем mutex для блокировки доступа к объекту из других вызовов метода,
	// пока не будет установлено новое значение

	p, err := inc.Repo.GetParams()
	if err != nil {
		return err
	}

	if maximumValue < p.StepValue {
		return ErrMaximumLessThenStepValue
	}

	// Если при смене максимального значения число начинает превышать данное  максимальное значение, то число обнуляется
	if p.Number > p.MaximumValue {
		err = inc.Repo.SetMaximumValue(maximumValue)
		if err != nil {
			return err
		}
		err = inc.Repo.SetNumber(0)
		if err != nil {
			return err
		}
		return nil
	}

	err = inc.Repo.SetMaximumValue(maximumValue)
	if err != nil {
		return err
	}

	// Возвращаем nil вместо ошибки
	return nil
}

func (inc *Incrementer) SetStepValue(stepValue int64) error {
	if stepValue < 0 {
		return ErrNegativeStepValue
	}

	p, err := inc.Repo.GetParams()
	if err != nil {
		return err
	}

	if stepValue > p.MaximumValue {
		return ErrMaximumLessThenStepValue
	}

	err = inc.Repo.SetStepValue(stepValue)
	if err != nil {
		return err
	}
	return nil
}

func (inc *Incrementer) SetParams(maximumValue, stepValue int64) error {
	if maximumValue < stepValue {
		return ErrMaximumLessThenStepValue
	}
	if maximumValue < 0 {
		return ErrNegativeMaximumValue
	}
	if stepValue < 0 {
		return ErrNegativeStepValue
	}

	num, err := inc.Repo.GetNumber()
	if err != nil {
		return err
	}

	if num > maximumValue {
		err = inc.Repo.SetParams(maximumValue, stepValue)
		if err != nil {
			return err
		}
		err = inc.Repo.SetNumber(0)
		if err != nil {
			return err
		}
		return nil
	}
	err = inc.Repo.SetParams(maximumValue, stepValue)
	if err != nil {
		return err
	}
	return nil

}
