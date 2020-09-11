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
rcli ssh -U root --hosts-file hosts -d run -c "id"
```
or using `-H`
```
rcli ssh -U root -H 192.168.1.N:22,192.168.1.N:22,192.168.1.N:22 -d run -c "id"
```

#### Example

```
rcli ssh -U root --hosts-file hosts -d run -c "id"
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
rcli ssh -U root --hosts-file hosts runscript -f script-examples/for_loop.sh
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
rcli container run -c /bin/sh -i rootfs
```

#### Example
```
rcli container run -c /bin/sh -i rootfs
           ___________________ .____    .___
Welcome to \______   \_   ___ \|    |   |   |
            |       _/    \  \/|    |   |   |
            |    |   \     \___|    |___|   |
            |____|_  /\______  /_______ \___|
                  \/        \/        \/

This software comes with ABSOLUTELY NO WARRANTY.
Use at your own risk.

[root@container]# hostname
container
[root@container]# ls -la
total 2668
drwxr-xr-x   19 root     root           268 Sep 11 08:23 .
drwxr-xr-x   19 root     root           268 Sep 11 08:23 ..
-rw-rw-r--    1 root     root             0 Sep 11 05:10 ROOT_CONTAINERS
-rw-rw-r--    1 root     root       2716902 May 29 14:20 alpine-minirootfs-3.12.0-x86_64.tar.gz
drwxr-xr-x    2 root     root          4096 May 29 14:20 bin
drwxr-xr-x    2 root     root             6 May 29 14:20 dev
drwxr-xr-x   15 root     root          4096 May 29 14:20 etc
drwxr-xr-x    2 root     root             6 May 29 14:20 home
drwxr-xr-x    7 root     root           223 May 29 14:20 lib
drwxr-xr-x    5 root     root            44 May 29 14:20 media
drwxr-xr-x    2 root     root             6 May 29 14:20 mnt
drwxr-xr-x    2 root     root             6 May 29 14:20 opt
dr-xr-xr-x  336 nobody   nobody           0 Sep 11 08:23 proc
drwx------    2 root     root             6 May 29 14:20 root
drwxr-xr-x    2 root     root             6 May 29 14:20 run
drwxr-xr-x    2 root     root          4096 May 29 14:20 sbin
drwxr-xr-x    2 root     root             6 May 29 14:20 srv
drwxr-xr-x    2 root     root             6 May 29 14:20 sys
drwxrwxr-x    2 root     root             6 May 29 14:20 tmp
drwxr-xr-x    7 root     root            66 May 29 14:20 usr
drwxr-xr-x   12 root     root           137 May 29 14:20 var
[root@container]# mount
/dev/sda2 on / type xfs (rw,seclabel,relatime,attr2,inode64,noquota)
proc on /proc type proc (rw,relatime)
[root@container]# ps -elf
PID   USER     TIME  COMMAND
    1 root      0:00 /proc/self/exe container run fork -c /bin/sh -i rootfs
    6 root      0:00 /bin/sh
   10 root      0:00 ps -elf
[root@container]# ip a
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
[root@container]#
```
<a id="markdown-contact" name="contact-us"></a>
## Contact us!

This project is maintained on Github.

- Please contact us by submitting [issue](#collaborate)


