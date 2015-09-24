package stats_test

import (
	"math"
	"testing"
	"github.com/mbu/stats"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"github.com/stretchr/testify/assert"
)

func TestGetAverage_OK(t *testing.T) {

	float64Array := []float64{ 11.904762, 27.5, 5, 2.5, 7.692308, 31.707317, 2.4390242, 7.317073 }

	expected := float64(12.007560525)
	actual := stats.GetAverage(float64Array)
	difference := math.Abs(float64(expected - actual))

	if (difference >= 1) {
		t.Error("Expected a difference less than 1, got: ", difference)
	}
}

func TestGetAverage_NOK_emptyArray(t *testing.T) {

	float64Array := []float64{}

	expected := float64(0)
	actual := stats.GetAverage(float64Array)
	assert.Equal(t, expected, actual)
}

func TestBuildCPUMultipleStat_OK(t *testing.T) {

	float64Array := []float64{ 11.904762, 27.5, 5, 2.5, 7.692308, 31.707317, 2.4390242, 7.317073 }
	cpu := stats.BuildCPUMultipleStat(float64Array)

	assert := assert.New(t)
	assert.NotNil(cpu)
	assert.Equal(stats.GetAverage(float64Array), cpu.AverageUsagePercentage)
	assert.Equal(len(float64Array), len(cpu.UsagePercentagePerCore))
}

func TestBuildCPUMultipleStat_NOK_emptyArray(t *testing.T) {

	float64Array := []float64{}
	cpu := stats.BuildCPUMultipleStat(float64Array)

	assert := assert.New(t)
	assert.NotNil(cpu)
	assert.Equal(stats.GetAverage(float64Array), cpu.AverageUsagePercentage)
	assert.Equal(float64(0), cpu.AverageUsagePercentage)
	assert.Equal(len(float64Array), len(cpu.UsagePercentagePerCore))
	assert.Equal(0, len(cpu.UsagePercentagePerCore))
}

func TestBuildRamStat_OK(t *testing.T) {

	zero := uint64(0)
	total := uint64(100000000000)
	used := uint64(77000000000)
	available := uint64(23000000000)
	usagePercentage := float64(77.00)

	vm := mem.VirtualMemoryStat{ total, available, used, usagePercentage, zero, zero, zero, zero, zero, zero, zero }
	ram := stats.BuildRamStat(vm)

	assert := assert.New(t)
	assert.NotNil(ram)
	assert.Equal(float64(total)/stats.GB, ram.TotalGB)
	assert.Equal(float64(used)/stats.GB, ram.UsedGB)
	assert.Equal(float64(available)/stats.GB, ram.AvailableGB)
	assert.Equal(float64(usagePercentage), ram.UsagePercentage)
}

func TestBuildRamStat_NOK_zeroedRAM(t *testing.T) {

	zero := uint64(0)
	total := uint64(0)
	used := uint64(0)
	available := uint64(0)
	usagePercentage := float64(0.00)

	vm := mem.VirtualMemoryStat{ total, available, used, usagePercentage, zero, zero, zero, zero, zero, zero, zero }
	ram := stats.BuildRamStat(vm)

	assert := assert.New(t)
	assert.NotNil(ram)
	assert.Equal(float64(total)/stats.GB, ram.TotalGB)
	assert.Equal(float64(used)/stats.GB, ram.UsedGB)
	assert.Equal(float64(available)/stats.GB, ram.AvailableGB)
	assert.Equal(float64(usagePercentage), ram.UsagePercentage)
}

func TestBuildDiskStat_OK(t *testing.T) {

	stringValue := ""
	total := uint64(975979800000)
	free := uint64(931929260000)
	used := uint64(44050570000)
	usagePercentage := float64(4.5134716)
	zeroInt := uint64(0)
	zeroFloat := float64(0)

	diskUsage := disk.DiskUsageStat{ stringValue, stringValue, total, free, used, usagePercentage, zeroInt, zeroInt, zeroInt, zeroFloat }
	disk := stats.BuildDiskStat(diskUsage)

	assert := assert.New(t)
	assert.NotNil(disk)
	assert.Equal(float64(total)/stats.GB, disk.TotalGB)
	assert.Equal(float64(used)/stats.GB, disk.UsedGB)
	assert.Equal(float64(free)/stats.GB, disk.AvailableGB)
	assert.Equal(float64(usagePercentage), disk.UsagePercentage)
}

func TestBuildDiskStat_NOK_zeroedDisk(t *testing.T) {

	stringValue := ""
	total := uint64(0)
	free := uint64(0)
	used := uint64(0)
	usagePercentage := float64(0)
	zeroInt := uint64(0)
	zeroFloat := float64(0)

	diskUsage := disk.DiskUsageStat{ stringValue, stringValue, total, free, used, usagePercentage, zeroInt, zeroInt, zeroInt, zeroFloat }
	disk := stats.BuildDiskStat(diskUsage)

	assert := assert.New(t)
	assert.NotNil(disk)
	assert.Equal(float64(total)/stats.GB, disk.TotalGB)
	assert.Equal(float64(used)/stats.GB, disk.UsedGB)
	assert.Equal(float64(free)/stats.GB, disk.AvailableGB)
	assert.Equal(float64(usagePercentage), disk.UsagePercentage)
}
