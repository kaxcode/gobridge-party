# GoBridge Party

A gobridge 1-day curriculumn for learning the basics of Go.

## Getting Started

1.  Install [Go](https://golang.org/doc/install)
2.  Create your [workspace](https://golang.org/doc/code.html#Workspaces) directory & set up [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) env variables.
3.  To run the tests and start working on challenges
```
  cd </project-file-path>
  cd 01_math_challenges
  go test
  go test -v
 ```

## Using Docker
1. Install Docker
  - [Mac](https://docs.docker.com/docker-for-mac/install/)
  - [Windows](https://docs.docker.com/docker-for-windows/install/)
2. To run the tests and start working on challenges
```
  cd </project-file-path>
  cd 01_math_challenges
  docker run -v $(pwd):/go golang:1.11 go test
  docker run -v $(pwd):/go golang:1.11 go test -v
 ```