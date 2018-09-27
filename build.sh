#!/bin/bash

# I don't know if this is idiomatic (it's probably not), but I'm new to golang.
# The build artifacts should probably be in a separate folder (?)
# TODO: Make this better lol
find solution*.go | xargs -L 1 go build -buildmode=plugin