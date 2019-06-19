# Core-utils (Go)

This is an implementation of the
[POSIX](http://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap04.html)
standard utilities written in Go.

[Decoded Core
Utils](http://www.maizure.org/projects/decoded-gnu-coreutils/) is a
helpful resource that contains explainations of each command's execution
workflow.

See [A Go Programmer’s Guide to Syscalls](https://about.sourcegraph.com/go/a-go-guide-to-syscalls)

## Completed:

- [x] true
- [x] false
- [x] uname

## Needs Implementation:

- [ ] arch
- [ ] base64
- [ ] basename
- [ ] cat
- [ ] chcon
- [ ] chgrp
- [ ] chmod
- [ ] chown
- [ ] chown
- [ ] chroot
- [ ] cksum
- [ ] comm
- [ ] cp
- [ ] csplit
- [ ] cut
- [ ] date
- [ ] dd
- [ ] df
- [ ] dir
- [ ] dircolors
- [ ] dirname
- [ ] du
- [ ] echo
- [ ] env
- [ ] expand
- [ ] expr
- [ ] factor
- [ ] false
- [ ] fmt
- [ ] fold
- [ ] groups
- [ ] head
- [ ] hostid
- [ ] hostname
- [ ] id
- [ ] install
- [ ] join
- [ ] kill
- [ ] link
- [ ] ln
- [ ] logname
- [ ] ls
- [ ] md5sum
- [ ] mkdir
- [ ] mkfifo
- [ ] mknod
- [ ] mktemp
- [ ] mv
- [ ] nice
- [ ] nl
- [ ] nohup
- [ ] nproc
- [ ] numfmt
- [ ] od
- [ ] paste
- [ ] pathchk
- [ ] pinky
- [ ] pr
- [ ] printenv
- [ ] printf
- [ ] ptx
- [ ] pwd
- [ ] readlink
- [ ] realpath
- [ ] rm
- [ ] rmdir
- [ ] runcon
- [ ] seq
- [ ] shred
- [ ] shuf
- [ ] sleep
- [ ] sort
- [ ] split
- [ ] stat
- [ ] stdbuf
- [ ] stty
- [ ] sum
- [ ] sync
- [ ] tac
- [ ] tail
- [ ] tee
- [ ] test
- [ ] timeout
- [ ] touch
- [ ] tr
- [ ] true
- [ ] truncate
- [ ] tsort
- [ ] tty
- [ ] unexpand
- [ ] uniq
- [ ] unlink
- [ ] uptime
- [ ] uptime
- [ ] users
- [ ] vdir
- [ ] wc
- [ ] who
- [ ] whoami
- [ ] xxd
- [ ] yes

### Information:

#### Behavior:

These utilities should be nearly identical to GNU's coreutils.

Since parsing the output of shell commands isn't uncommon (even if
it *is* bad behavior), most of the commands should have output that
is nearly identical to the original GNU commands.

Do note that sometimes the results could differ a little for select commands.

For example, GNU's `wc` utility relies on the current locale to determine
whether it should parse multi-byte characters or not.

The Go version, on the other hand, uses the `unicode/utf8` package
which natively detects multi-byte sequences. The trade-off is this: the
Go version is technically more correct, while the C version is faster.
