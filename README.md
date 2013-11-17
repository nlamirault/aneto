Aneto
=====

Aneto Glacier is the largest glacier in the Pyrenees.
This tool is a personal backup.

## Install

* Install Python tools:

        $ apt-get install python-pip
		$ pip install virtualenvwrapper
		$ source /usr/local/bin/virtualenvwrapper.sh

* Install dependencies :

        $ mkvirtualenv aneto
		New python executable in aneto/bin/python
		Installing Setuptools..................done.
		Installing Pip.........................done.
        $ pip install -r requirements.txt

## Documentation

This documentation is written by contributors, for contributors.
The source is maintained in the *doc/source* folder using
[reStructuredText](http://docutils.sourceforge.net/rst.html)
and built by [Sphinx](http://sphinx-doc.org/)

Building Manually:

    $ sphinx-build -b html doc/source build/sphinx/html

Results are in the *build/sphinx/html* directory.

Documentation for **aneto** is hosted on readthedocs.org at:

	http://readthedocs.org/docs/aneto/en/latest/

## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>
