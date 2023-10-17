# Audit tools

## Description

This repository contains a collection of tools to audit the security of a
smart contract system. The tools are written in Golang so that they can be
distributed as a single binary.

## How to use

### Download

1. Find the latest github workflow, usually the latest commit on the `main`
   branch. Click on the check mark to see the details of the workflow. Then click
   on "Details" of the "Go / Build" step.

![Github Workflow](/docs/gh-workflow.png)

2. Download the binary from the "Summary" -> "Artifacts" section.

![Artifacts Github](/docs/gh-artifacts.png)

### Or build your own

To build the binary, run:

```bash
go build -o audit-tools
```

### Run

To run the binary, run:

```bash
./audit-tools
```

Or move it to your bin folder:

```bash
sudo mv audit-tools /usr/local/bin/audit-tools
```

Then run:

```bash
audit-tools
```

## For developers

For ease of development, the make file contains a few useful commands:

```bash
make build # Build the binary
make install # Install the binary in /usr/local/bin
```
