# kins

KINS: 麒麟终端系统管控平台部署工具，支持一键部署及各组件单独部署	

## Install

```
go install github.com/283713406/kins
```

## Usage

```
» kins -h
        KINS
        麒麟终端系统管控平台部署工具，支持一键部署及各组件单独部署

Example usage:

        一键部署
        control-plane# kins kylin install

        单独部署
        worker# kins k8s install

    You can then repeat the second step on as many other machines as you like.

Usage:
  kins [command]

Available Commands:
  apps        Install apps
  basic       Install basic
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  k8s         Install kylin container cloud
  kylin       Install kylin platform
  mount       Mount glusterfs
  version     Print the version of kins

Flags:
  -h, --help   help for kins

Use "kins [command] --help" for more information about a command.
```

### Examples

```
» kins k8s
```
