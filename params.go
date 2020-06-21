package main

/*type Params struct {
	Number       sql.NullInt64
	MaximumValue sql.NullInt64
	StepValue    sql.NullInt64
}*/

type Params struct {
	Number       int64 `json:"-"`
	MaximumValue int64 `json:"maximum_value"`
	StepValue    int64 `json:"step_value"`
}
