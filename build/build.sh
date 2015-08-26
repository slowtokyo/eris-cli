#!/usr/bin/env bash
# assumes golang build from source or golang >= 1.5

go get github.com/laher/goxc
goxc -c build/.goxc.json -wd cmd/eris -n eris -d $HOME/.eris/builds