#!/bin/bash

glide install
cd ./cmd/traffic-api && go build && go install
cd ./cmd/traffic-collectd && go build && go install
