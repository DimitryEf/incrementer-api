package model

// Params - параметры инкрементора
type Params struct {
	Number       int64 `json:"-"`             // Значение инкрементора
	MaximumValue int64 `json:"maximum_value"` // Максимальное значение
	StepValue    int64 `json:"step_value"`    // Шаг инкрементора
}
