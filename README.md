[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/nixwiz/check-memory-usage)
![Go Test](https://github.com/nixwiz/check-memory-usage/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/nixwiz/check-memory-usage/workflows/goreleaser/badge.svg)

# Sensu memory usage check

## Table of Contents
- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Check definition](#check-definition)
- [Installation from source](#installation-from-source)
- [Contributing](#contributing)

## Overview

The Sensu memory usage check is a [Sensu Check][1] that provides alerting and
metrics for memory usage.  Metrics are provided in [nagios_perfdata][5] format.

## Usage examples

```
Check memory usage and provide metrics

Usage:
  check-memory-usage [flags]
  check-memory-usage [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -c, --critical float   Critical threshold for overall CPU usage (default 90)
  -w, --warning float    Warning threshold for overall CPU usage (default 75)
  -h, --help             help for check-memory-usage

Use "check-memory-usage [command] --help" for more information about a command.
```

## Configuration

### Asset registration

[Sensu Assets][2] are the best way to make use of this plugin. If you're not
using an asset, please consider doing so! If you're using sensuctl 5.13 with
Sensu Backend 5.13 or later, you can use the following command to add the asset:

```
sensuctl asset add nixwiz/check-memory-usage
```

If you're using an earlier version of sensuctl, you can find the asset on the
[Bonsai Asset Index][3].

### Check definition

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
  - nixwiz/check-memory-usage
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an
Asset. If you would like to compile and install the plugin from source or
contribute to it, download the latest version or create an executable from this
source.

From the local path of the check-cpu-usage repository:

```
go build
```

## Contributing

For more information about contributing to this plugin, see [Contributing][4].

[1]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[2]: https://docs.sensu.io/sensu-go/latest/reference/assets/
[3]: https://bonsai.sensu.io/assets/nixwiz/check-memory-usage
[4]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[5]: https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-schedule/collect-metrics-with-checks/#supported-output-metric-formats
