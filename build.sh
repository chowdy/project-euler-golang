#!/bin/bash
find solution*.go | xargs -L 1 go build -buildmode=plugin
