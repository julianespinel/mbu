package stats

type SingleStat struct {
	TotalGB         float32 `json:"totalGB"`
	UsedGB          float32 `json:"usedGB"`
	AvailableGB     float32 `json:"availableGB"`
	UsagePercentage float32 `json:"usagePercentage"`
}

type MultipleStat struct {
	AverageUsagePercentage float32   `json:"averageUsagePercentage"`
	UsagePercentagePerCore []float32 `json:"usagePercentagePerCore"`
}

type AllStat struct {
	CPU  MultipleStat `json:"cpu"`
	RAM  SingleStat   `json:"ram"`
	Disk SingleStat   `json:"disk"`
}
