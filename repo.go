package main

type Repo interface {
	GetNumber() (int64, error)
	SetNumber(num int64) error
	GetParams() (Params, error)
	SetMaximumValue(maximumValue int64) error
	SetStepValue(stepValue int64) error
	SetParams(maximumValue, stepValue int64) error
}
