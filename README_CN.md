<p align="center"><img src="http://hprose.com/banner.@2x.png" alt="Hprose" title="Hprose" width="650" height="200" /></p>

# [Hprose gateway](https://github.com/vlorc/hprose-gateway)
[简体中文](https://github.com/vlorc/hprose-gateway/blob/master/README_CN.md)

[![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![codebeat badge](https://codebeat.co/badges/c41b426c-4121-4dc8-99c2-f1b60574be64)](https://codebeat.co/projects/github-com-vlorc-hprose-gateway-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/vlorc/hprose-gateway)](https://goreportcard.com/report/github.com/vlorc/hprose-gateway)
[![GoDoc](https://godoc.org/github.com/vlorc/hprose-gateway?status.svg)](https://godoc.org/github.com/vlorc/hprose-gateway)
[![Build Status](https://travis-ci.org/vlorc/hprose-go-nats.svg?branch=master)](https://travis-ci.org/vlorc/hprose-gateway?branch=master)
[![Coverage Status](https://coveralls.io/repos/github/vlorc/hprose-go-nats/badge.svg?branch=master)](https://coveralls.io/github/vlorc/hprose-gateway?branch=master)

基于golang的hprose网关

## 特性
+ 服务发现
+ 惰性加载
+ 自定义插件
+ 负载均衡

## 安装
	go get github.com/vlorc/hprose-gateway

## 库
+ 发现服务
	+ [etcd](https://github.com/vlorc/hprose-gateway-etcd)
	+ [dns](https://github.com/vlorc/hprose-gateway-dns)
	+ [consul](https://github.com/vlorc/hprose-gateway-consul)
+ 插件
	+ [counter](https://github.com/vlorc/hprose-gateway-plugins/tree/master/counter)
	+ [hash](https://github.com/vlorc/hprose-gateway-plugins/tree/master/hash)
	+ [limiter](https://github.com/vlorc/hprose-gateway-plugins/tree/master/limiter)
	+ [panic](https://github.com/vlorc/hprose-gateway-plugins/tree/master/panic)
	+ [session](https://github.com/vlorc/hprose-gateway-plugins/tree/master/session)
+ 协议
	+ [hprose](https://github.com/vlorc/hprose-gateway-protocol/tree/master/hprose)
	+ [forward](https://github.com/vlorc/hprose-gateway-protocol/tree/master/forward)

## 许可证
这个项目是在Apache许可证下进行的。请参阅完整许可证文本的许可证文件。
