# Aneto

[![Travis](https://img.shields.io/travis/nlamirault/aneto.svg)]()

Aneto Glacier is the largest glacier in the Pyrenees. This tool is a personal backup
using [Amazon Glacier][]

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
