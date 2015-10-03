package stats

type SingleStat struct {
	TotalGB         float64 `json:"totalGB"`
	UsedGB          float64 `json:"usedGB"`
	AvailableGB     float64 `json:"availableGB"`
	UsagePercentage float64 `json:"usagePercentage"`
}

type MultipleStat struct {
	NumberOfCores          int       `json:"numberOfCores"`
	AverageUsagePercentage float64   `json:"averageUsagePercentage"`
	UsagePercentagePerCore []float64 `json:"usagePercentagePerCore"`
}

type AllStat struct {
	CPU  MultipleStat `json:"cpu"`
	RAM  SingleStat   `json:"ram"`
	Disk SingleStat   `json:"disk"`
}
