# Web Workshop

A gentle series of progressive workshops to incrementally build production web services in Go, with a focus on test driving a TodoBackend.

We will try to keep it as idiomatic as possible, but where a decision needs to be made, we will try to hew to the most popular choices. (e.g. relational database vs mongo). We could definitely use advice on what that is at times! (e.g. logr vs zap vs go-kit log vs logrus vs. std log)

This repository also aims to provide useful comparisons for different necessary choices (e.g. REST vs GraphQL vs gRPC).

## Why

Excellent resources like [Rosetta Code](http://rosettacode.org/wiki/Rosetta_Code) present solutions to the same task in as many different languages as possible, to demonstrate how languages are similar and different, and to aid a person with a grounding in one approach to a problem in learning another.

Similarly, [Todo Backend](https://todobackend.com/) and [TodoMVC](http://todomvc.com/) provide solutions using a variety of languages and frameworks.
Meanwhile, resources like [Ardan Labs Service Training](https://github.com/ardanlabs/service-training), [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests/), [Test With Go](https://testwithgo.com/), and [Exercism.io](http://exercism.io/) each provide other important aspects of daily tasks on the road to production.

We would like to accelerate onboarding Go developers, by providing a unified guide to the test-driven path to production.

## Requirements

This project was designed against Go 1.14. It should work for 1.13 but 1.14 is recommended.

Supporting services like the database are hosted in Docker. If you cannot install Docker on your machine you can still follow most of this material by hosting a database elsewhere and modifying the connection information to your needs.

## Setup

Clone this repository somewhere on your computer. The location does not especially matter but if it is outside of your `$GOPATH` then the Go modules features will work automatically.

In a separate folder make a directory where you will be building your API. We recommend you initialize that folder as a Git repository to track your work.


```sh
mkdir ~/training
cd ~/training
git clone https://github.com/Khan/web-workshop.git
mkdir go-todo
cd go-todo
git init .
```

---

You must also use `go mod init` to set the import path for this project. Doing
this exactly as shown will allow you to copy and paste code without a need to
modify import paths.

```sh
go mod init github.com/Khan/go-todo
git add go.mod
git commit -m "Initial commit"
```

## Postman API Client

For the class we will be building up a REST API. You may use any HTTP client you prefer to make requests but we recommend [Postman](https://www.getpostman.com/).
For convenience you may use the import button in the top left to import the included `postman_environment.json` and `postman_collection.json` files to get a client up and running quickly. Be sure to select the "Todos API" environment in the top right.

## Diffing Folders

Reviewing the differences between the successive steps helps to reinforce the ideas each change is about. This is made easier by running the following command to define a git alias called `dirdiff`:

```sh
git config --global alias.dirdiff 'diff -p --stat -w --no-index'
```

With that alias in place, run this command from the top level folder to see the differences between the `01-startup` directory and the `02-shutdown` directory.

```sh
git dirdiff 01-startup 02-shutdown
```
## Credit where credit is due

This repository draws mainly from [Ardan Labs Service Training](https://github.com/ardanlabs/service-training), with bits from [Test With Go](https://testwithgo.com/), [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests/), and other fine resources.
---
---
---
