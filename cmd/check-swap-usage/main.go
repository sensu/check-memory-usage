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
			Name:     "check-swap-usage",
			Short:    "Check swap usage and provide metrics",
			Keyspace: "sensu.io/plugins/check-swap-usage/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		{
			Path:      "critical",
			Argument:  "critical",
			Shorthand: "c",
			Default:   float64(90),
			Usage:     "Critical threshold for overall swap usage",
			Value:     &plugin.Critical,
		},
		{
			Path:      "warning",
			Argument:  "warning",
			Shorthand: "w",
			Default:   float64(75),
			Usage:     "Warning threshold for overall swap usage",
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
	swapStat, err := mem.SwapMemory()
	if err != nil {
		return sensu.CheckStateCritical, fmt.Errorf("failed to get swap statistics: %v", err)
	}

	perfData := fmt.Sprintf("swap_total=%d, swap_free=%d, swap_used=%d", swapStat.Total, swapStat.Free, swapStat.Used)
	if swapStat.UsedPercent > plugin.Critical {
		fmt.Printf("%s Critical: %.2f%% swap usage | %s\n", plugin.PluginConfig.Name, swapStat.UsedPercent, perfData)

		return sensu.CheckStateCritical, nil
	} else if swapStat.UsedPercent > plugin.Warning {
		fmt.Printf("%s Warning: %.2f%% swap usage | %s\n", plugin.PluginConfig.Name, swapStat.UsedPercent, perfData)
		return sensu.CheckStateWarning, nil
	}

	fmt.Printf("%s OK: %.2f%% swap usage | %s\n", plugin.PluginConfig.Name, swapStat.UsedPercent, perfData)
	return sensu.CheckStateOK, nil
}
