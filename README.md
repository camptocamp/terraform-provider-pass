# Pass Terraform Provider

[![Terraform Registry Version](https://img.shields.io/badge/dynamic/json?color=blue&label=registry&query=%24.version&url=https%3A%2F%2Fregistry.terraform.io%2Fv1%2Fproviders%2Fcamptocamp%2Fpass)](https://registry.terraform.io/providers/camptocamp/pass)
[![Go Report Card](https://goreportcard.com/badge/github.com/camptocamp/terraform-provider-pass)](https://goreportcard.com/report/github.com/camptocamp/terraform-provider-pass)
[![Build Status](https://travis-ci.org/camptocamp/terraform-provider-pass.svg?branch=master)](https://travis-ci.org/camptocamp/terraform-provider-pass)
[![By Camptocamp](https://img.shields.io/badge/by-camptocamp-fb7047.svg)](http://www.camptocamp.com)

This provider adds integration between Terraform and [Pass][] and [Gopass][] password stores.

[Pass][] is a password store using gpg to encrypt password and git to version.
[Gopass][] is a rewrite of the pass password manager in Go with the aim of making it cross-platform and adding additional features.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 1.5.x
- [Go](https://golang.org/doc/install) 1.21 and [dep](https://golang.github.io/dep/) (to build the provider plugin)

## Building The Provider

Download the provider source code

```sh
$ go get github.com/camptocamp/terraform-provider-pass
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/camptocamp/terraform-provider-pass
$ dep ensure
$ make build
```

## Installing the provider

After building the provider, install it using the Terraform instructions for [installing a third party provider](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).

## Example

```hcl
provider "pass" {
  store_dir = "/srv/password-store"    # defaults to $PASSWORD_STORE_DIR
  refresh_store = false                # do not call `git pull`
}


resource "pass_password" "test" {
  path = "secret/foo"
  password = "0123456789"
  data = {
    zip = "zap"
  }
}

data "pass_password" "test" {
  path = "${pass_password.test.path}"
}
```

## Usage

### The `pass` provider

#### Argument Reference

The provider takes no arguments.

### The `pass_password` resource

#### Argument Reference

The resource takes the following arguments:

- `path` - Full path from which a password will be read
- `password` - Secret password
- `data` - (Optional) Additional secret data (keys and values, not nested)
- `yaml` - (Optional) YAML document, can't be set together with data

#### Attribute Reference

The following attributes are exported:

- `path` - Full path from which the password was read
- `password` - Secret password
- `data` - Additional secret data
- `body` - Raw secret data, only filled if not stored as YAML
- `full` - Entire raw secret contents

### The `pass_password` data source

#### Argument Reference

The data source takes the following arguments:

- `path` - Full path from which a password will be read

#### Attribute Reference

The following attributes are exported:

- `path` - Full path from which the password was read
- `password` - Secret password
- `data` - Additional secret data
- `body` - Raw secret data, only filled if not stored as YAML
- `full` - Entire raw secret contents

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.21+ is _required_). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-$PROVIDER_NAME
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

_Note:_ Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

[pass]: https://www.passwordstore.org/
[gopass]: https://www.justwatch.com/gopass/
