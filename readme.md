Google Authenticator Code
============

[![Build Status](https://github.com/jormin/gacode/workflows/test/badge.svg?branch=master)](https://github.com/jormin/gacode/actions?query=workflow%3Atest)
[![Codecov](https://codecov.io/gh/jormin/gacode/branch/master/graph/badge.svg)](https://codecov.io/gh/jormin/gacode)
[![Go Report Card](https://goreportcard.com/badge/github.com/jormin/gacode)](https://goreportcard.com/report/github.com/jormin/gacode)
[![](https://img.shields.io/badge/version-v1.0.0-success.svg)](https://github.com/jormin/gacode)

This is a tool to manage accounts and codes of Google Authenticator.

Support
-----

##### Account

- [x] Add an existing account
- [x] Generate a new account
- [x] Print or Export the QR code image
- [x] List all accounts configured
- [x] Remove specified or all accounts

##### Code

- [x] Show codes of specified or all accounts

Build
-----

```
# clone source code
git clone https://github.com/jormin/gacode.git

# download module
go mod download

# install
go install
```

Command
-----

```shell
NAME:
   gacode - The tool to manage accounts and codes of Google Authenticator.

USAGE:
   gacode [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

COMMANDS:
   account  Manage accounts of Google Authenticator
   code     Show codes of Google Authenticator
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

Example
-----

##### Account

- Add an existing account

```shell script
➜  ~ gacode account add test HIUKS7E5ZDQXM2HLQH5USZ7HZUQASSDQ
add account success.
name: test
secret: HIUKS7E5ZDQXM2HLQH5USZ7HZUQASSDQ
qrcode: otpauth://totp/test?secret=HIUKS7E5ZDQXM2HLQH5USZ7HZUQASSDQ
```

- Generate a new account

```shell
➜  ~ gacode account gen test2                                
generate account success.
name: test2
secret: KR2OSN2GXPQL2FARNUHTFHZWLUSSGEZ4
qrcode:otpauth://totp/test2?secret=KR2OSN2GXPQL2FARNUHTFHZWLUSSGEZ4
```

- Print the QR code image

<div align="left"><img src="https://blog.cdn.lerzen.com/img/20210904005428.png" alt="image-20210904005423753" style="zoom:30%;" /></div>

- Export the QR code image with flag `-o`

```shell
➜  ~ gacode account qr -o ~/Desktop test
export the QR code success: /Users/Jormin/Desktop/c4p583l948169o90rrhg.png
```

- List all accounts configured

```shell
➜  ~ gacode account ls
Name      Secret                               QRCode
test      HIUKS7E5ZDQXM2HLQH5USZ7HZUQASSDQ     otpauth://totp/test?secret=HIUKS7E5ZDQXM2HLQH5USZ7HZUQASSDQ
test2     KR2OSN2GXPQL2FARNUHTFHZWLUSSGEZ4     otpauth://totp/test2?secret=KR2OSN2GXPQL2FARNUHTFHZWLUSSGEZ4
```
    
- Remove specified accounts
    
```shell
➜  ~ gacode account rm test test2
remove account test success
remove account test2 success
```
        
- remove all accounts with tag `-a`
    
```shell
➜  ~ gacode account rm -a
remove all accounts success 
```

##### Code

- Show codes of specified accounts

```shell
➜  ~ gacode code test1 test2
Account     Code
test1       649669
test2       660560
```

- Show codes of all accounts

```shell
➜  ~ gacode code -a
Account     Code
test1       649669
test2       660560
test3       387946
```

License
-------

under the [MIT](./LICENSE) License