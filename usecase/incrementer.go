package usecase

import (
	"errors"
	"github.com/DimitryEf/incrementer-api/repo"
)

// Ошибки недопустимых значений
var (
	// Ошибка при попытке установить отрицательное максимальное значение
	ErrNegativeMaximumValue = errors.New("the maximum value of increment must be a non negative number")
	// Ошибка при попытке установить отрицательное значение для шага инкремента
	ErrNegativeStepValue = errors.New("the step value of increment must be a non negative number")
	// Ошибка в ситуации, когда максимальное значение меньше шага инкремента
	ErrMaximumLessThenStepValue = errors.New("the maximum value of increment must be bigger then step value")
)

// Incrementor - структура, содержащая в себе бизнес-логику работы инкрементора
type Incrementer struct {
	Repo repo.Repo
}

// NewIncrementor - конструктор объекта Incrementor, принимает в себя интерфейс репозитория.
func NewIncrementer(repo repo.Repo) *Incrementer {
	return &Incrementer{
		Repo: repo,
	}
}

// GetNumber - возвращает текущее число
func (inc *Incrementer) GetNumber() (int64, error) {
	return inc.Repo.GetNumber()
}

// IncrementNumber - увеличивает число
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

	// Устанавливаем значение
	err = inc.Repo.SetNumber(p.Number + p.StepValue)
	if err != nil {
		return err
	}
	return nil
}

// SetMaximumValue - устанавливает максимальное число
func (inc *Incrementer) SetMaximumValue(maximumValue int64) error {
	// Проверяем, чтобы maximumValue было неотрицательным, иначе возвращаем ошибку.
	if maximumValue < 0 {
		return ErrNegativeMaximumValue
	}

	// Получаем актуальные параметры инкрементора
	p, err := inc.Repo.GetParams()
	if err != nil {
		return err
	}

	// Проверяем, чтоб максимум был больше значения шага инкрементора
	if maximumValue < p.StepValue {
		return ErrMaximumLessThenStepValue
	}

	// Если при смене максимального значения число начинает превышать данное  максимальное значение, то число обнуляется
	if p.Number > p.MaximumValue {
		err = inc.Repo.SetMaximumValueAndZeroNumber(maximumValue)
		if err != nil {
			return err
		}
		return nil
	}

	// Устанавливаем максимальное значение
	err = inc.Repo.SetMaximumValue(maximumValue)
	if err != nil {
		return err
	}

	// Возвращаем nil вместо ошибки
	return nil
}

// SetStepValue - устанавливает значение шага инкрементора
func (inc *Incrementer) SetStepValue(stepValue int64) error {
	if stepValue < 0 {
		return ErrNegativeStepValue
	}

	// Получаем параметры
	p, err := inc.Repo.GetParams()
	if err != nil {
		return err
	}

	// Проверяем, чтоб максимум был больше значения шага инкрементора
	if stepValue > p.MaximumValue {
		return ErrMaximumLessThenStepValue
	}

	// Устанавливаем значение шага инкрементора
	err = inc.Repo.SetStepValue(stepValue)
	if err != nil {
		return err
	}
	return nil
}

// SetParams - установка параметров инкрементора: максимальное значение и шаг инкрементора
func (inc *Incrementer) SetParams(maximumValue, stepValue int64) error {
	// Проверка входных параметров
	if maximumValue != 0 && maximumValue < stepValue {
		return ErrMaximumLessThenStepValue
	}
	if maximumValue < 0 {
		return ErrNegativeMaximumValue
	}
	if stepValue < 0 {
		return ErrNegativeStepValue
	}

	// Получение значения
	num, err := inc.Repo.GetNumber()
	if err != nil {
		return err
	}

	// Если оба параметра ненулевые, то применяем метод SetParams
	if maximumValue != 0 && stepValue != 0 {
		// Проверка значения по отношению к максимуму
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
		// Установка параметра
		err = inc.Repo.SetParams(maximumValue, stepValue)
		if err != nil {
			return err
		}
		return nil
	}

	//Если параметр stepValue нулевой, то меняем только параметр maximumValue
	if stepValue == 0 {
		if num > maximumValue {
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
		// Установка параметра
		err = inc.Repo.SetMaximumValue(maximumValue)
		if err != nil {
			return err
		}
		return nil
	}

	//Если параметр maximumValue нулевой, то меняем только параметр stepValue
	if maximumValue == 0 {
		err = inc.Repo.SetStepValue(stepValue)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("unknown error")

}
