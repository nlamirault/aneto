# Aneto

[![License Apache 2][badge-license]](LICENSE)

Aneto Glacier is the largest glacier in the Pyrenees. This tool is a personal backup
using [Amazon Glacier][]

Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/aneto/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/aneto/tree/master)
* [![GitHub version](https://badge.fury.io/gh/nlamirault%2Faneto.svg)](https://badge.fury.io/gh/nlamirault%2Faneto)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/aneto/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/aneto/tree/develop)
* https://img.shields.io/circleci/project/nlamirault/aneto/develop.svg?style=flat-square

## Usage

Setup your AWS credentials :

    $ export AWS_ACCESS_KEY_ID='AKID'
    $ export AWS_SECRET_ACCESS_KEY='SECRET'

Initialize your backup into Glacier :

    $ aneto vault create --name anetolam
    2015/08/19 02:23:49 Create vault : anetolam
    2015/08/19 02:23:50 {
        Location: "/447241706233/vaults/anetolam"
    }

Delete your backup :

    $ bin/aneto vault delete --name anetolam
    2015/08/19 02:23:59 Delete vault : anetolam
    2015/08/19 11:56:00 {
    }


## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>




[Amazon Glacier]: https://aws.amazon.com/fr/glacier/
