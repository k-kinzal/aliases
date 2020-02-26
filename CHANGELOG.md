## v0.4.0 - 2020-02-26

New Feature:
* Changes to be able to mount docker binary without dependencies (#67)

## v0.3.1 - 2019-03-01

Bug Fixes:
* fixed use local docker binary path in docker download (https://github.com/k-kinzal/aliases/pull/65)

## v0.3.0 - 2019-02-18

New Feature:
* Parameters defined by the dependency are inherited to the parent. For example, `env` defined in dependency can be read by parent command (https://github.com/k-kinzal/aliases/pull/48)
* Define commands recursively in `dependencies`. This allows you to define the command to be called only from the command (https://github.com/k-kinzal/aliases/pull/51)
* Extend the `entrypoint` (https://github.com/k-kinzal/aliases/pull/58)
* Used shell redirection on command call (https://github.com/k-kinzal/aliases/pull/59)

Bug Fixes:
* Fixed a read waiting for input when calling command (https://github.com/k-kinzal/aliases/pull/53)

Misc:
* Refactored packages & project structure
* Changed timing for docker binary download (https://github.com/k-kinzal/aliases/pull/54)
* Changed `command` to alias of `entrypoint` (https://github.com/k-kinzal/aliases/pull/56)
* Changed execute expansion of boolean parameter shortened (https://github.com/k-kinzal/aliases/pull/57)

## v0.2.1 - 2019-01-06

Bug Fixes:
* command and arguments were not passed in `aliases run` (https://github.com/k-kinzal/aliases/pull/47)

## v0.2.0 - 2019-01-02

New Features:
* Added `aliases run` command (https://github.com/k-kinzal/aliases/pull/34, https://github.com/k-kinzal/aliases/pull/35, https://github.com/k-kinzal/aliases/pull/36)
* Added `docker` parameter to the schema. It works hosts and guests incompatible with docker binaries such as ubuntu and alpine (https://github.com/k-kinzal/aliases/pull/38)
* Supported [CircleCI Orb](https://circleci.com/orbs/registry/orb/k-kinzal/aliases) (https://github.com/k-kinzal/aliases/pull/39)

Misc:
* Refactored packages & project structure
* Added logger and the log is displayed in color.
* Changed CI setting to dogfood CircleCI Orb (https://github.com/k-kinzal/aliases/pull/40)

## v0.1.0 - 2018-12-27

First release 🎉