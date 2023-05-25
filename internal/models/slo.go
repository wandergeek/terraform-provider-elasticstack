package models

type Slo struct {
	ID              string
	Name            string
	Description     string
	Indicator       Indicator
	TimeWindow      TimeWindow
	BudgetingMethod string
	Objective       Objective
	Settings        Settings
	SpaceID         string
}
type Params struct {
	Index           string
	Filter          string
	Good            string
	Total           string
	TimestampField  string
	Service         string
	Environment     string
	TransactionType string
	TransactionName string
	GoodStatusCodes []string
}
type Indicator struct {
	Params Params
	Type   string
}

type TimeWindow struct {
	Duration   string
	IsRolling  bool
	IsCalendar bool
}

type Objective struct {
	Target           float64
	TimeslicesTarget float64
	TimeslicesWindow string
}
type Settings struct {
	SyncDelay string
	Frequency string
}
