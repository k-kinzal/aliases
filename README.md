# aliases

[![CircleCI](https://circleci.com/gh/k-kinzal/aliases.svg?style=svg)](https://circleci.com/gh/k-kinzal/aliases)
[![GolangCI](https://golangci.com/badges/github.com/k-kinzal/aliases.svg)](https://golangci.com/r/github.com/k-kinzal/aliases)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fk-kinzal%2Faliases.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fk-kinzal%2Faliases?ref=badge_shield)


`aliases` is a tool for resolving command dependencies with containers.

* YAML-based command definition
* Version of command with the environment variable
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
   aliases - Generate alias for command on the container

USAGE:
   aliases [global options] command [command options] [arguments...]

VERSION:
   v0.5.0

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

When version environment variable is specified, a command is executed by overwriting `tag` defined by YAML.


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

Environmental variables expand with parameters expressed like this.
`Env1` expands at the timing when `aliases` generates a command.
`Env2` expands at the timing when you executed a command generated by `aliases`

## Expand commands

```yaml
tty: $(if tty >/dev/null; then echo "true"; else echo "false"; fi)
```

```yaml
volume:
- $(helm home):/root/.helm
```

When you enclose the command with `$(...)`, the `STDOUT` of the command is set as the parameter.
Command expanded at the timing when you executed a command generated by `aliases`.

Note: Always execute the generated command on the machine (host or docker) that executed the alias.

## Special Environment Variable

`$PWD` is a special environment variable.
The `$PWD` specified on the left always points to `$PWD` of the host.

## Extend entrypoint

```yaml
/usr/local/bin/helm:
  image: chatwork/helm
  tag: 2.12.3
  volume:
  - $HOME/.helm:/root/.helm
  - $PWD:/helm
  workdir: /helm
  dependencies:
  - /bin/bash:
      image: bash
      tag: 5.0.2
  - /bin/curl:
      image: byrnedo/alpine-curl
      tag: 0.1.7
  entrypoint: |
    #!/bin/sh
    if [ -f $(helm home)/plugins/helm-import ]; then
      helm plugin install https://github.com/k-kinzal/helm-import --version v0.2.1
    if
    helm "$@"
```

If you want to extend `entrypoint`, please define a string with shebang in `entrypoint`.

NOTE: Please understand that extend entrypoint is less reproducible. It should be included in the docker image if possible.

## Dependencies commands.

`aliases` can define commands that depend on commands.

```yaml
/usr/local/bin/sops:
  image: mozilla/sops
  tag: a2d0328e35e6e37b51f3ad468dc6f213c7b44014 
  env:
    AWS_ACCESS_KEY_ID: ${AWS_ACCESS_KEY_ID}
    AWS_SECRET_ACCESS_KEY: ${AWS_SECRET_ACCESS_KEY}
    AWS_PROFILE: ${AWS_PROFILE}
    AWS_DEFAULT_REGION: ${AWS_DEFAULT_REGION}
  volume:
  - $PWD:/sops
  workdir: /sops

/usr/local/bin/helm:
  image: chatwork/helm
  tag: 2.12.3
  volume:
  - $HOME/.helm:/root/.helm
  - $PWD:/helm
  workdir: /helm
  dependencies:
  - /usr/local/bin/sops
  - /bin/bash:
      image: bash
      tag: 5.0.2
  - /bin/curl:
      image: byrnedo/alpine-curl
      tag: 0.1.7
```

For dependency, reference to another command defined or define command.
If you want to define a command that can only be called from a command, please recursively define the dependency.

Also, parent commands inherit dependent parameters.

```bash
$ AWS_ACCESS_KEY_ID=xxx AWS_SECRET_ACCESS_KEY=xxx helm secrets enc ...
```

If the same parameter exists, the parent parameter takes precedence.

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

If you want to show the docker run command, please specify the option of `--verbose`.

## for CircleCI

`aliases` supports [CircleCI Orb](https://circleci.com/orbs/registry/orb/k-kinzal/aliases).
You can quickly execute precisely the same commands on local and CI.

```yaml
version: "2.1"

orbs:
  aliases: k-kinzal/aliases@0.2.1

jobs:
  aliases:
    machine: true
    steps:
    - checkout
    - aliases/install
    - aliases/gen
    - run: [your command]

workflows:
  version: 2

  aliases:
    jobs:
    - aliases
```

However, since A uses privileged and volume, please use [Machine Executor](https://circleci.com/docs/2.0/executor-types/#using-machine).

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fk-kinzal%2Faliases.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fk-kinzal%2Faliases?ref=badge_large)
