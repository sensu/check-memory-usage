package main

import (
	"fmt"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
	"github.com/shirou/gopsutil/v3/mem"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	Critical float64
	Warning  float64
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "check-memory-usage",
			Short:    "Check memory usage and provide metrics",
			Keyspace: "sensu.io/plugins/check-memory-usage/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		{
			Path:      "critical",
			Argument:  "critical",
			Shorthand: "c",
			Default:   float64(90),
			Usage:     "Critical threshold for overall CPU usage",
			Value:     &plugin.Critical,
		},
		{
			Path:      "warning",
			Argument:  "warning",
			Shorthand: "w",
			Default:   float64(75),
			Usage:     "Warning threshold for overall CPU usage",
			Value:     &plugin.Warning,
		},
	}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
	if plugin.Critical == 0 {
		return sensu.CheckStateWarning, fmt.Errorf("--critical is required")
	}
	if plugin.Warning == 0 {
		return sensu.CheckStateWarning, fmt.Errorf("--warning is required")
	}
	if plugin.Warning > plugin.Critical {
		return sensu.CheckStateWarning, fmt.Errorf("--warning cannot be greater than --critical")
	}
	return sensu.CheckStateOK, nil
}

func executeCheck(event *types.Event) (int, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return sensu.CheckStateCritical, fmt.Errorf("failed to get virtual memory statistics: %v", err)
	}

	perfData := fmt.Sprintf("mem_total=%d, mem_available=%d, mem_used=%d, mem_free=%d", vmStat.Total, vmStat.Available, vmStat.Used, vmStat.Free)
	if vmStat.UsedPercent > plugin.Critical {
		fmt.Printf("%s Critical: %.2f%% memory usage | %s\n", plugin.PluginConfig.Name, vmStat.UsedPercent, perfData)

		return sensu.CheckStateCritical, nil
	} else if vmStat.UsedPercent > plugin.Warning {
		fmt.Printf("%s Warning: %.2f%% memory usage | %s\n", plugin.PluginConfig.Name, vmStat.UsedPercent, perfData)
		return sensu.CheckStateWarning, nil
	}

	fmt.Printf("%s OK: %.2f%% memory usage | %s\n", plugin.PluginConfig.Name, vmStat.UsedPercent, perfData)
	return sensu.CheckStateOK, nil
}
