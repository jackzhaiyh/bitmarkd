# bitmarkd - Main program

[![GoDoc](https://godoc.org/github.com/bitmark-inc/bitmarkd?status.svg)](https://godoc.org/github.com/bitmark-inc/bitmarkd)

Prerequisites

* Install the go language package for your system
* Configure environment variables for go system
* install the ZMQ4 library

For shell add the following to the shell's profile
(remark the `export CC=clang` if you wish to use gcc)
~~~~~
# check for go installation
GOPATH="${HOME}/gocode"
if [ -d "${GOPATH}" ]
then
  gobin="${GOPATH}/bin"
  export GOPATH
  export PATH="${PATH}:${gobin}"
  # needed for FreeBSD 10 and later
  export CC=clang
else
  unset GOPATH
fi
unset gobin
~~~~~

OnFreeBSD/PC-BSD

~~~~~
pkg install libzmq4
~~~~~

On a Debian like system
(as of Ubuntu 14.04 this only has V3, so need to search for PPA)

~~~~~
apt-get install libzmq4-dev
~~~~~

To compile simply:

~~~~~
go get github.com/bitmark-inc/bitmarkd
go install -v github.com/bitmark-inc/bitmarkd
~~~~~

# Set up

Create the configuration directory, copy sample configuration, edit it to
set up IPs, ports and local bitcoin testnet connection.

Generate key files and certificates.

~~~~~
bitmarkd ~/.config/bitmarkd/bitmarkd.conf rpc
bitmarkd ~/.config/bitmarkd/bitmarkd.conf peer
bitmarkd ~/.config/bitmarkd/bitmarkd.conf proof
bitmarkd ~/.config/bitmarkd/bitmarkd.conf dns-txt
~~~~~

Start the program.

~~~~~
bitmarkd
~~~~~

Bitmark:[数字环境中的产权系统_白皮书](https://github.com/jackzhaiyh/bitmarkd/blob/master/bitmark_%E4%B8%AD%E6%96%87%E7%99%BD%E7%9A%AE%E4%B9%A6.pdf)
