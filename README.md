[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/sensu/check-memory-usage)
![Go Test](https://github.com/sensu/check-memory-usage/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/sensu/check-memory-usage/workflows/goreleaser/badge.svg)

# Sensu memory and swap usage checks

## Table of Contents
- [Overview](#overview)
  - [Checks](#checks)
- [Usage examples](#usage-examples)
  - [check-memory-usage](#check-memory-usage)
  - [check-swap-usage](#check-swap-usage)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Check definitions](#check-definitions)
- [Installation from source](#installation-from-source)
- [Contributing](#contributing)

## Overview

The Sensu memory usage checks are a collectoin of [Sensu Checks][1] that provide
alerting and metrics for memory and swap usage.  Metrics are provided in
[nagios_perfdata][5] format.

### Checks

This collection contains the following checks:

* `check-memory-usage` - for checking memory usage
* `check-swap-usage` - for checking swap usage

## Usage examples

### check-memory-usage

```
Check memory usage and provide metrics

Usage:
  check-memory-usage [flags]
  check-memory-usage [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -c, --critical float   Critical threshold for overall memory usage (default 90)
  -w, --warning float    Warning threshold for overall memory usage (default 75)
  -h, --help             help for check-memory-usage

Use "check-memory-usage [command] --help" for more information about a command.
```

### check-swap-usage

```
Check swap usage and provide metrics

Usage:
  check-swap-usage [flags]
  check-swap-usage [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -c, --critical float   Critical threshold for overall swap usage (default 90)
  -w, --warning float    Warning threshold for overall swap usage (default 75)
  -h, --help             help for check-memory-usage

Use "check-swap-usage [command] --help" for more information about a command.
```

## Configuration

### Asset registration

[Sensu Assets][2] are the best way to make use of this plugin. If you're not
using an asset, please consider doing so! If you're using sensuctl 5.13 with
Sensu Backend 5.13 or later, you can use the following command to add the asset:

```
sensuctl asset add sensu/check-memory-usage
```

If you're using an earlier version of sensuctl, you can find the asset on the
[Bonsai Asset Index][3].

### Check definitions

#### check-memory-usage

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: check-memory-usage
  namespace: default
spec:
  command: >-
    check-memory-usage
    --critical 90
    --warning 80
  output_metric_format: nagios_perfdata
  output_metric_handlers:
    - influxdb
  subscriptions:
  - system
  runtime_assets:
  - sensu/check-memory-usage
```

#### check-swap-usage

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: check-swap-usage
  namespace: default
spec:
  command: >-
    check-swap-usage
    --critical 90
    --warning 75
  output_metric_format: nagios_perfdata
  output_metric_handlers:
    - influxdb
  subscriptions:
  - system
  runtime_assets:
  - sensu/check-memory-usage
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an
Asset. If you would like to compile and install the plugin from source or
contribute to it, download the latest version or create executables from this
source.

From the local path of the check-cpu-usage repository:

```
go build ./cmd/check-memory-usage/
go build ./cmd/check-swap-usage/
```

## Contributing

For more information about contributing to this plugin, see [Contributing][4].

[1]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[2]: https://docs.sensu.io/sensu-go/latest/reference/assets/
[3]: https://bonsai.sensu.io/assets/sensu/check-memory-usage
[4]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[5]: https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-schedule/collect-metrics-with-checks/#supported-output-metric-formats
