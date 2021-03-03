[![Go Report Card](https://goreportcard.com/badge/github.com/zerodayz/rcli)](https://goreportcard.com/report/github.com/zerodayz/rcli)

# rcli - Multipurpose CLI
This is multipurpose tool written in Go. It can now `ssh` to thousands of hosts in around 40 seconds. It can start container without `docker` or `podman` purely using Linux namespaces. I am adding more [features](#features) as I can think of them.

**Table of contents**
<!-- TOC depthFrom:1 insertAnchor:true orderedList:true -->

- [Introduction](#introduction)
- [Releases](#releases)
- [Features](#features)
  - [ssh](#ssh)
  - [container](#container)
- [Contact us!](#contact-us)

<!-- /TOC -->

<a id="markdown-introduction" name="introduction"></a>
## Introduction

Please if you have any idea on any improvements please do not hesitate to open an issue.

<a id="markdown-releases" name="releases"></a>
## Releases
- For the latest and greatest bits, please compile from the source.
- For latest released version, please see releases page in Github.

<a id="markdown-features" name="features"></a>
## Features

<a id="markdown-features-ssh" name="ssh"></a>
### SSH
Example of SSH execution against 1,000 hosts takes around 40 seconds.
#### Syntax
Example of hosts file:
```
192.168.1.2:22
192.168.1.3:22
192.168.1.4:22
192.168.1.5:22
```
`--hosts-file`
```
$ rcli ssh -U root --hosts-file hosts exec -c id
```
or using `-H`
```
$ rcli ssh -U root -H 192.168.1.2:22,192.168.1.3:22,192.168.1.4:22 exec -c id
```
or execute script file using `-f` 
```
$ rcli ssh -U root -H 192.168.1.2:22,192.168.1.3:22,192.168.1.4:22 exec -f script-examples/for_loop.sh
```

#### Build
At the moment requires build specifically for linux due to namespaces requirements.
```
 GOOS=linux go build
```

<a id="markdown-contact" name="contact-us"></a>
## Contact us!

This project is maintained on Github.

- Please contact us by submitting [issue](#collaborate)


