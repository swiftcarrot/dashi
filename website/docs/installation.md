---
id: installation
title: Installation
sidebar_label: Installation
---

## Requirements

Dashi requires you have following dependencies installed:

- Go working environment (Go version >=1.13)
- node
- yarn

## Installation

Dashi is a single executable binary, so its very easy to download and install as:

```sh
curl -sf https://gobinaries.com/swiftcarrot/dashi | sh
```

## Verify installation

To verify the installation, you can execute the `dashi` command in a termail:

```sh
$ dashi
dashi

Usage:
   [flags]
   [command]

Available Commands:
  generate    Generates dashboard, scaffold, migration, ...
  help        Help about any command
  new         Create new project
  pack        Pack migration files with packr2
  version     Print version

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
```
