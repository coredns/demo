# demo

## Name

*demo* - returns `1.1.1.1` for `172.0.0.0/8` or `127.0.0.0/8` and `8.8.8.8` otherwise

## Description

This demo plugin showcases the implementation of source IP based service discovery.
It is intended as a boilerplate code, so that anyone that is interested in coredns
implementation could build their plugin on top of this demo plugin.

This demo plugin and its full code has been walked through in past KubeCon talks

Note: In order to add a new plugin, an additional step of `make gen` is needed. Therefore,
to build the coredns with demo plugin the following should be used:
```
docker run -it --rm -v $PWD:/v -w /v golang:1.16 sh -c 'make gen && make'
```

## Syntax

~~~ txt
demo
~~~

## Also See

See the [manual](https://coredns.io/manual).
