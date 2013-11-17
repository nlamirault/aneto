#!/usr/bin/env python

PROJECT = 'aneto'

# Change docs/sphinx/conf.py too!
VERSION = '0.1.0'

from setuptools import setup, find_packages

try:
    long_description = open('README.md', 'rt').read()
except IOError:
    long_description = ''

setup(
    name=PROJECT,
    version=VERSION,

    description='A backup tool.',
    long_description=long_description,

    author='Nicolas Lamirault',
    author_email='nicolas.lamirault@gmail.com',

    url='https://github.com/nlamirault/aneto',
    download_url='https://github.com/nlamirault/aneto/tarball/master',

    classifiers=['Development Status :: 3 - Alpha',
                 'License :: OSI Approved :: Apache Software License',
                 'Programming Language :: Python',
                 'Programming Language :: Python :: 2',
                 'Programming Language :: Python :: 2.7',
                 'Programming Language :: Python :: 3',
                 'Programming Language :: Python :: 3.2',
                 'Intended Audience :: Developers',
                 'Environment :: Console',
                 ],

    platforms=['Any'],

    scripts=[],

    provides=[],
    install_requires=['cliff'],

    namespace_packages=[],
    packages=find_packages(),
    include_package_data=True,

    entry_points={
        'console_scripts': [
            'aneto = aneto.main:main'
        ],
        'aneto': [
            'info = aneto.simple:Info',
            'config = aneto.simple:Configuration',
        ],
    },

    zip_safe=False,
)
