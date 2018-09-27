#!/bin/bash

# I don't know if this is idiomatic (it's probably not), but I'm new
# TODO: This better lol
find solution*.go | xargs -L 1 go build -buildmode=plugin