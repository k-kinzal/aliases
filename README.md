# aliases

[![CircleCI](https://circleci.com/gh/k-kinzal/aliases.svg?style=svg)](https://circleci.com/gh/k-kinzal/aliases)
[![GolangCI](https://golangci.com/badges/github.com/k-kinzal/aliases.svg)](https://golangci.com/r/github.com/k-kinzal/aliases)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fk-kinzal%2Faliases.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fk-kinzal%2Faliases?ref=badge_shield)


aliases is a tool for resolving command dependencies with containers.

* YAML-based command definition
* Version of command with environment variable
* Call another dependent command from the command

## Install

```
$ go get github.com/k-kinzal/aliases
```

## Get Started

**~/.aliases/aliases.yaml**

```
/usr/local/bin/kubectl:
  image: chatwork/kubectl
  tag: 1.11.2
  volume:
  - $HOME/.kube:/root/.kube
  - $PWD:/kube
  workdir: /kube

/usr/local/bin/alpine:
  image: alpine
  tag: 3.8
  volume:
  - $PWD:/work
  workdir: /work
  dependencies:
  - /usr/local/bin/kubectl
```

```
$ eval "$(aliases gen)"
```

or 

```
$ eval "$(aliases gen --export)"
```


Using the option of `--export` can be used as a command instead of an alias.

## CLI

```
$ aliases --help
  NAME:
     aliases - Generate alias for command on container
  
  USAGE:
     aliases [global options] command [command options] [arguments...]
  
  VERSION:
     v0.2.0
  
  COMMANDS:
       gen      Generate aliases
       run      Run aliases command
       home     Get aliases home path
       help, h  Shows a list of commands or help for one command
  
  GLOBAL OPTIONS:
     --config value, -c value  Load configuration file
     --home value              Home directory for aliases [$ALIASES_HOME]
     --verbose, -v             enable verbose output
     --help                    show help
     --version                 print the version
```

## Version Environment variable

If `/usr/local/bin/kubectl` is defined, you can specify the version in the environment variable with `_VERSION` in the suffix with `kubectl` in the file name capitalized.

```
$ KUBECTL_VERSION=1.11.3 kubectl get all
```

When version environment variable is specified, command is executed by overwriting `tag` defined by YAML.


## Expand Environment Variables

```yaml
env:
  $ENV1: $ENV2
```

```yaml
volume:
- $ENV1:$ENV2:rw
```

```yaml
user: $ENV1:$ENV2
```

Environmental variables are expanded with parameters expressed like this.
`Env1` expands at the timing when aliases generates a command.
`Env2` expands at the timing when you executed a command generated by aliases

## Expand commands

```yaml
tty: $(if tty >/dev/null; then echo "true"; else echo "false"; fi)
```

```yaml
volume:
- $(helm home):/root/.helm
```

When you enclose the command with `$()`, the `STDOUT` of the command is set as the parameter.
Command expands at the timing when you executed a command generated by aliases.

Note that command is always executed on the machine (host or docker) that executed the command generated by aliases.

## Special Environment Variable

`$PWD` is a special environment variable.
The `$PWD` specified on the left always points to `$PWD` of the host.

## How to debug aliases

If you do not get what you expected, please use the `aliases run` command.

```bash
$ aliases run -it --entrypoint '' [command] sh
```

`aliases run` command overwrites the option of your defined command and executes it.

```bash
$ aliases --verbose run -it --entrypoint '' [command] sh
docker run --entrypoint "" --interactive --network "host" --rm --tty [your image] sh
```

If you want to show the docker run command, please specify option of `--verbose`.

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fk-kinzal%2Faliases.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fk-kinzal%2Faliases?ref=badge_large)
