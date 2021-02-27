ginkgo playground
----

[![CircleCI](https://circleci.com/gh/jvrmaia/ginkgo-playground/tree/master.svg?style=svg)](https://circleci.com/gh/jvrmaia/ginkgo-playground/tree/master)

# pre-setup

1. install gvm
2. install go version 1.13.1

# setup

    gvm pkgset create ginkgo-playground
    gvm pkgset use ginkgo-playground
    mkdir -p $HOME/workspace/go/src/github.com/jvrmaia
    export GOPATH=$GOPATH:$HOME/workspace/go
    cd $HOME/workspace/go/src/github.com/jvrmaia
    git clone git@github.com:jvrmaia/ginkgo-playground.git

# run

    cd ginkgo-playground/operator
    ginkgo .

Have fun!

