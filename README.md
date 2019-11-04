# Goa

Goa (*Golang Oriented to Aspects*) is a Golang library that provides us aspect-oriented programming (AOP).

## What's AOP?

> In computing, aspect-oriented programming (AOP) is a programming paradigm that aims to increase modularity by allowing the separation of cross-cutting concerns. It does so by adding additional behavior to existing code (an advice) without modifying the code itself, instead separately specifying which code is modified via a "pointcut" specification

* Extracted  from [Wikipedia](https://en.wikipedia.org/wiki/Aspect-oriented_programming)


## Goals

- Provide a handy tool that helps us to build reusable aspects.
- Code generation based in AST modifications
- A mechanism to build golang applications under AOP paradigm. 

# Getting Started

## Installation

## Usage

## Guides & Tutorials

## Roadmap

# For Collaborators

## Contributing

Please read [CONTRIBUTING.md](https://github.com/wesovilabs/goa/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Checkout the code

In order to work on a Goa enhancement or on a bug fix you just need to fork the repository.

Once you forked the repository you can checkout it locally. 

```bash
git clone https://github.com/<user>/goa.git
```

## Setup Git hooks

All the commited code must pass tests and linter checks. These are defined with golangci. To ensure
your committed code will be valid to be merged you can setup some Git hooks by running

```bash
make init
```

## Running the tests

Run the tests with the below command

```bash
make test
```

and to check the test coverage... 

```bash
make test-coverage
``` 

> Keep in mind that for approval a Pull Request the test coverage must be equal or higher than the existing one.

## Check your code

As it was mentioned on the above, the code must pass all the defined linter checks. You can check it locally

```bash
make lint
```

Linters configuration can be found [here](https://github.com/wesovilabs/goa/blob/master/.golangci.yml)

## Versioning
    
We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/wesovilabs/goa/tags).

## Authors

- **Iván Corrales Solera <ivan.corrales.solera@gmail.com>** 

See also the list of [contributors](https://github.com/wesovilabs/goa/contributors) who participated in this project.


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments
