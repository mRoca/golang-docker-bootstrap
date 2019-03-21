# Golang docker bootstrap

This is my docker bootstrap project for golang.

- full dockerized stack
- nothing to put in your ~/go path
- vendors management with dep
- code style lint & check 
- gitlab ci pipeline 
- go binary build after every file update

In order to use this bootstrap for another project: 

- rename the `src/github.com/mRoca/golang-docker-bootstrap` folder path
- change all the `github.com/mRoca/golang-docker-bootstrap` strings
- run a `make` command
- start coding

## Prod

Usage:

```bash
docker run --rm mroca/golang-docker-bootstrap say -text test
```

## Dev

### Dependencies with docker

- Make
- Docker
- Docker-compose

### Dependencies without docker - this is a very bad choice ;-)

- Make
- Go
- Dep
- inotifywait (dev)

### Install & start dev

```bash
make
```

### Other commands

See the [Makefile](Makefile) :-)

### Coding style

See [https://github.com/golang/go/wiki/CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)

**Check the cs**

```bash
make test-cs
```

**Fix the cs**

```bash
make fix-cs
```

### GOPATH

All the go commands are run in the `cli` docker container. If you want to use this project on your host, or if you want to be able to use autocompletion in your favorite IDE,
you can easily add it to your GOPATH variable :

```bash
export GOPATH=$HOME/go:$THIS_PROJECT_DIRECTORY
```

See [https://github.com/golang/go/wiki/SettingGOPATH](https://github.com/golang/go/wiki/SettingGOPATH) for more informations.


### Run a command in containers as current host user

```bash
bin/exec command
```

### Add a package

```bash
bin/exec dep ensure -add github.com/foo/bar github.com/baz/quux

```

### Build a prod image

```bash
make build-prod
docker run --rm golang-docker-bootstrap say -text test
```
