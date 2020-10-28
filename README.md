# Todo for Terraform

* Requires (recent versions, but approximate):
  * terraform 0.13.X+
  * go 1.15.X+ (compiling)
  * docker 19.X+ (build & test scripts)
  * docker-compose 1.24.X+ (build & test scripts)

## Todo Server

### Quick Start

```shell
docker run --name todo-list -p 8080:80 spkane/todo-list-server:latest
```

### Build Server

```shell
./bin/build.sh server
```

### Usage

```shell
curl -i http://127.0.0.1:8080/
curl -i http://127.0.0.1:8080/ -d "{\"description\":\"message $RANDOM\"}" -H 'Content-Type: application/spkane.todo-list.v1+json'
curl -i http://127.0.0.1:8080/ -d "{\"description\":\"message $RANDOM\",\"completed\":false}" -H 'Content-Type: application/spkane.todo-list.v1+json'
curl -i http://127.0.0.1:8080/ -d "{\"description\":\"message $RANDOM\",\"completed\":false}" -H 'Content-Type: application/spkane.todo-list.v1+json'
curl -i http://127.0.0.1:8080/3 -X PUT -H 'Content-Type: application/spkane.todo-list.v1+json' -d '{"description":"go shopping",\"completed\":true}'
curl -i http://127.0.0.1:8080/1 -X DELETE -H 'Content-Type: application/spkane.todo-list.v1+json'
curl -i http://127.0.0.1:8080/3 -X DELETE -H 'Content-Type: application/spkane.todo-list.v1+json'
curl -i http://127.0.0.1:8080
```

## Terraform Provider

### Build & Test

```shell
./bin/build.sh provider
```

The build script runs the Integration tests. If you want to run real local terraform tests, you can run this script:

```shell
./bin/tests_manual.sh
```

### Install

In general the provider must be copied into the following directory so that terraform can find it:

`$HOME/.terraform.d/plugins/terraform.spkane.org/spkane/todo/${VERSION}/${OS}_${ARCH}/`

and in your terraform hcl you will need to refernce the prociver with something like this:

```
terraform {
  required_providers {
    todo = {
      source  = "terraform.spkane.org/spkane/todo"
      version = "1.1.0"
    }
  }
}
```

---

* *NOTE*: The todo server code for this project was directly forked from:
  * [github.com/go-swagger](https://github.com/go-swagger/go-swagger/tree/master/examples/tutorials/todo-list/server-complete)


## TODOs

* https setup
* Setup docker-compose to use pre-built binary
