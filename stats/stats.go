package stats

type SingleStat struct {
	TotalGB float32
	UsedGB float32
	AvailableGB float32
	UsagePercentage float32
}

type MultipleStat struct {
	AverageUsagePercentage float32
	UsagePercentagePerCore []float32
}
