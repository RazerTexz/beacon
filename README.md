# Beacon

[![Build](https://img.shields.io/github/actions/workflow/status/RazerTexz/beacon/build.yaml?style=for-the-badge)](https://github.com/RazerTexz/beacon/actions/workflows/build.yaml)
[![License](https://img.shields.io/github/license/RazerTexz/beacon?style=for-the-badge)](LICENSE)

A minimalist 24/7 endpoint monitor written in Go.

```text
$ beacon google.com github.com fakewebsite.com

PASS  69ms    https://google.com
PASS  420ms   https://github.com
FAIL  5s      https://fakewebsite.com
```

## Install

```bash
go install github.com/RazerTexz/beacon@latest
```
Or download a prebuilt binary from [Releases](https://github.com/RazerTexz/beacon/releases).

## Usage

```bash
beacon <url>...
```