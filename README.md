<p align="center"><img src="http://hprose.com/banner.@2x.png" alt="Hprose" title="Hprose" width="650" height="200" /></p>

# [Hprose gateway](https://github.com/vlorc/hprose-gateway)
[简体中文](https://github.com/vlorc/hprose-gateway/blob/master/README_CN.md)

[![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![codebeat badge](https://codebeat.co/badges/c41b426c-4121-4dc8-99c2-f1b60574be64)](https://codebeat.co/projects/github-com-vlorc-hprose-gateway-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/vlorc/hprose-gateway)](https://goreportcard.com/report/github.com/vlorc/hprose-gateway)
[![GoDoc](https://godoc.org/github.com/vlorc/hprose-gateway?status.svg)](https://godoc.org/github.com/vlorc/hprose-gateway)
[![Build Status](https://travis-ci.org/vlorc/hprose-gateway.svg?branch=master)](https://travis-ci.org/vlorc/hprose-gateway?branch=master)
[![Coverage Status](https://coveralls.io/repos/github/vlorc/hprose-gateway/badge.svg?branch=master)](https://coveralls.io/github/vlorc/hprose-gateway?branch=master)

Hprose gateway based on golang

## Features
+ service resolver
+ lazy load
+ load balancing
+ plugin

## Installing
	go get github.com/vlorc/hprose-gateway

## Library
+ discovery
	+ [etcd](https://github.com/vlorc/hprose-gateway-etcd)
	+ [dns](https://github.com/vlorc/hprose-gateway-dns)
	+ [consul](https://github.com/vlorc/hprose-gateway-consul)
+ plugin
	+ [counter](https://github.com/vlorc/hprose-gateway-plugins/counter)
	+ [hash](https://github.com/vlorc/hprose-gateway-plugins/hash)
	+ [limiter](https://github.com/vlorc/hprose-gateway-plugins/limiter)
	+ [panic](https://github.com/vlorc/hprose-gateway-plugins/panic)
	+ [session](https://github.com/vlorc/hprose-gateway-plugins/session)
+ protocol
	+ [hprose](https://github.com/vlorc/hprose-gateway-protocol/hprose)
	+ [forward](https://github.com/vlorc/hprose-gateway-protocol/forward)

## License
This project is under the apache License. See the LICENSE file for the full license text.

