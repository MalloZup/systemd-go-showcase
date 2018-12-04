# systemd golang

This is a minimal example. It is based on lib: https://github.com/coreos/go-systemd

I assume you run this in your $GOPATH and use install the `ginkgo` and `go-systemd` lib with `go get -u ...`

# BDD style

You can run the BDD `kubernetes` framework `ginkgo` with :

```bash
cd ginkgo
go test
``` 

```bash
cd gingko
ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --cover --trace --race --progress
```
# Minimal without bdd

on GOPATH `go build`. RUN!

Try other services.. :dog:
