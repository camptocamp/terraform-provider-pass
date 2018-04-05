Pass Terraform Provider
=======================

[Pass](https://www.passwordstore.org/) is a password store using gpg to encrypt password and git to version.
[Gopass](https://www.justwatch.com/gopass/) is a rewrite of the pass password manager in Go with the aim of making it cross-platform and adding additional features.

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.9 and [dep](https://golang.github.io/dep/) (to build the provider plugin)

Building The Provider
---------------------

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

Using the provider
----------------------

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


Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

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

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
