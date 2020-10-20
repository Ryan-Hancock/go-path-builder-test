# Path Builder

## Table of contents

- [General info](#general-info)
  - [Rules](#rules)
- [Technologies](#technologies)
- [Setup](#setup)
  - [Tests](#tests)

## General info

This is an script to build paths with choice of separator.

### Rules

- Empty String and Null array items are ignored
- Leading Separator is allowed
- Trailing Separator is allowed
- The separator should not appear multiple times between path components, but inside a component is valid

## Technologies

These are the technologies used to create this project:

- Go 1.15
- [Testify](https://github.com/stretchr/testify)

## Setup

An example of how to use can be found in the main.go which can be built by:

```
$ go build -o path-builder
```

### Tests

In order to run tests:

```
$ go test -v ./...
```
