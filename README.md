# RCLI - multipurpose CLI
Prototype of an multipurpose tool written in Go

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
192.168.1.N:22
192.168.1.N:22
192.168.1.N:22
192.168.1.N:22
```
`--hosts-file`
```
$ rcli ssh -U root --hosts-file hosts --debug run --command id
```
or using `-H`
```
$ rcli ssh -U root -H 192.168.1.N:22,192.168.1.N:22,192.168.1.N:22 --debug run --command id
```

#### Example

```
$ rcli ssh -U root --hosts-file hosts --debug run --command id
           ___________________ .____    .___
Welcome to \______   \_   ___ \|    |   |   |
            |       _/    \  \/|    |   |   |
            |    |   \     \___|    |___|   |
            |____|_  /\______  /_______ \___|
                  \/        \/        \/

This software comes with ABSOLUTELY NO WARRANTY.
Use at your own risk.

2020/09/11 17:29:10 DEBUG: connecting to: 192.168.1.N:22
2020/09/11 17:29:10 DEBUG: connecting to: 192.168.1.N:22
2020/09/11 17:29:10 DEBUG: connecting to: 192.168.1.N:22
2020/09/11 17:29:10 DEBUG: connecting to: 192.168.1.N:22
2020/09/11 17:29:10 DEBUG: connecting to: 192.168.1.N:22
2020/09/11 17:29:10 DEBUG: connecting to: 192.168.1.N:22
2020/09/11 17:29:10 DEBUG: connecting to: 192.168.1.N:22
2020/09/11 17:29:10 DEBUG: connecting to: 192.168.1.N:22
[...SNIP...]
2020/09/11 17:29:13 DEBUG: connected to: 192.168.1.N:22
2020/09/11 17:29:13 DEBUG: connected to: 192.168.1.N:22
2020/09/11 17:29:13 DEBUG: connected to: 192.168.1.N:22
2020/09/11 17:29:13 DEBUG: connected to: 192.168.1.N:22
2020/09/11 17:29:13 DEBUG: connected to: 192.168.1.N:22
2020/09/11 17:29:13 DEBUG: connected to: 192.168.1.N:22
2020/09/11 17:29:13 DEBUG: connected to: 192.168.1.N:22
[...SNIP...]
2020/09/11 17:29:14 DEBUG: executing [id] on: 192.168.1.N:22
2020/09/11 17:29:14 DEBUG: executing [id] on: 192.168.1.N:22
2020/09/11 17:29:14 DEBUG: executing [id] on: 192.168.1.N:22
2020/09/11 17:29:14 DEBUG: executing [id] on: 192.168.1.N:22
2020/09/11 17:29:14 DEBUG: executing [id] on: 192.168.1.N:22
2020/09/11 17:29:14 DEBUG: executing [id] on: 192.168.1.N:22
2020/09/11 17:29:14 DEBUG: executing [id] on: 192.168.1.N:22
2020/09/11 17:29:14 DEBUG: executing [id] on: 192.168.1.N:22
[...SNIP...]
 --- 192.168.1.N:22 ---
 Output:
uid=0(root) gid=0(root) groups=0(root) context=unconfined_u:unconfined_r:unconfined_t:s0-s0:c0.c1023
 Error:
 --- 192.168.1.N:22 ---
 Output:
uid=0(root) gid=0(root) groups=0(root) context=unconfined_u:unconfined_r:unconfined_t:s0-s0:c0.c1023
 Error:
 --- 192.168.1.N:22 ---
 Output:
uid=0(root) gid=0(root) groups=0(root) context=unconfined_u:unconfined_r:unconfined_t:s0-s0:c0.c1023
 Error:
 --- 192.168.1.N:22 ---
 Output:
uid=0(root) gid=0(root) groups=0(root) context=unconfined_u:unconfined_r:unconfined_t:s0-s0:c0.c1023
 Error:
 --- 192.168.1.N:22 ---
 Output:
uid=0(root) gid=0(root) groups=0(root) context=unconfined_u:unconfined_r:unconfined_t:s0-s0:c0.c1023
 Error:
 --- 192.168.1.N:22 ---
 Output:
uid=0(root) gid=0(root) groups=0(root) context=unconfined_u:unconfined_r:unconfined_t:s0-s0:c0.c1023
 Error:
2020/09/11 17:29:14 38.323739114s
```
In case of large number of hosts it's possible some connection may be failing,
in such case error is printed on the screen:
```
 --- 192.168.1.N:22 ---
 Output:
 Error:
ssh: handshake failed: read tcp 192.168.1.N:50087->192.168.1.N:22: read: connection reset by peer
```

Same 1,000 hosts, this time running an example script. It takes roughly around 40 seconds to finish.
```
$ rcli ssh -U root --hosts-file hosts runscript --file script-examples/for_loop.sh
           ___________________ .____    .___
Welcome to \______   \_   ___ \|    |   |   |
            |       _/    \  \/|    |   |   |
            |    |   \     \___|    |___|   |
            |____|_  /\______  /_______ \___|
                  \/        \/        \/

This software comes with ABSOLUTELY NO WARRANTY.
Use at your own risk.
[..SNIP..]
 --- 192.168.1.N:22 ---
 Output:
Hello from localhost
for_loop-1
for_loop-2
for_loop-3
for_loop-4
for_loop-5
for_loop-6
for_loop-7
for_loop-8
for_loop-9
for_loop-10
 Error:
 --- 192.168.1.N:22 ---
 Output:
Hello from localhost
for_loop-1
for_loop-2
for_loop-3
for_loop-4
for_loop-5
for_loop-6
for_loop-7
for_loop-8
for_loop-9
for_loop-10
 Error:
 --- 192.168.1.N:22 ---
 Output:
Hello from localhost
for_loop-1
for_loop-2
for_loop-3
for_loop-4
for_loop-5
for_loop-6
for_loop-7
for_loop-8
for_loop-9
for_loop-10
 Error:
 --- 192.168.1.N:22 ---
 Output:
Hello from localhost
for_loop-1
for_loop-2
for_loop-3
for_loop-4
for_loop-5
for_loop-6
for_loop-7
for_loop-8
for_loop-9
for_loop-10
 Error:
 --- 192.168.1.N:22 ---
 Output:
Hello from localhost
for_loop-1
for_loop-2
for_loop-3
for_loop-4
for_loop-5
for_loop-6
for_loop-7
for_loop-8
for_loop-9
for_loop-10
 Error:
 --- 192.168.1.N:22 ---
 Output:
Hello from localhost
for_loop-1
for_loop-2
for_loop-3
for_loop-4
for_loop-5
for_loop-6
for_loop-7
for_loop-8
for_loop-9
for_loop-10
 Error:
2020/09/11 18:10:42 39.327087183s
```

<a id="markdown-features-container" name="container"></a>
### Container
NOTE: At the moment this is VERY EXPERIMENTAL AND DANGEROUS FEATURE.  
Please DO NOT EXECUTE!

This tool doesn't require any third party container technology such as docker, podman.
It creates it's own containers using namespaces.

#### Build
At the moment requires build specifically for linux due to namespaces requirements.
```
 GOOS=linux go build
```

#### Requirements
You will need `rootfs` directory, for example minimal alpine will do:
- rootfs (http://dl-cdn.alpinelinux.org/alpine/v3.12/releases/x86_64/alpine-minirootfs-3.12.0-x86_64.tar.gz)
#### Syntax
```
$ rcli container run --image ${IMAGE} -c ${COMMAND}
```

#### Example Fedora 32 image
```
$ IMAGE="fedora"
$ mkdir ${IMAGE} && podman export $(podman create ${IMAGE}) | tar -C ${IMAGE} -xvf -
$ rcli container run --image ${IMAGE} -c /bin/bash
           ___________________ .____    .___
Welcome to \______   \_   ___ \|    |   |   |
            |       _/    \  \/|    |   |   |
            |    |   \     \___|    |___|   |
            |____|_  /\______  /_______ \___|
                  \/        \/        \/

This software comes with ABSOLUTELY NO WARRANTY.
Use at your own risk.

[root@fpllngzieyoh]# cat /etc/*release
Fedora release 32 (Thirty Two)
NAME=Fedora
VERSION="32 (Container Image)"
ID=fedora
VERSION_ID=32
VERSION_CODENAME=""
PLATFORM_ID="platform:f32"
PRETTY_NAME="Fedora 32 (Container Image)"
ANSI_COLOR="0;34"
LOGO=fedora-logo-icon
CPE_NAME="cpe:/o:fedoraproject:fedora:32"
HOME_URL="https://fedoraproject.org/"
DOCUMENTATION_URL="https://docs.fedoraproject.org/en-US/fedora/f32/system-administrators-guide/"
SUPPORT_URL="https://fedoraproject.org/wiki/Communicating_and_getting_help"
BUG_REPORT_URL="https://bugzilla.redhat.com/"
REDHAT_BUGZILLA_PRODUCT="Fedora"
REDHAT_BUGZILLA_PRODUCT_VERSION=32
REDHAT_SUPPORT_PRODUCT="Fedora"
REDHAT_SUPPORT_PRODUCT_VERSION=32
PRIVACY_POLICY_URL="https://fedoraproject.org/wiki/Legal:PrivacyPolicy"
VARIANT="Container Image"
VARIANT_ID=container
Fedora release 32 (Thirty Two)
Fedora release 32 (Thirty Two)
```
<a id="markdown-contact" name="contact-us"></a>
## Contact us!

This project is maintained on Github.

- Please contact us by submitting [issue](#collaborate)


