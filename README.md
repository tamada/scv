# :balance_scale: scv

[![build](https://github.com/tamada/scv/actions/workflows/build.yml/badge.svg)](https://github.com/tamada/scv/actions/workflows/build.yml)
[![Coverage Status](https://coveralls.io/repos/github/tamada/scv/badge.svg?branch=setup_ci)](https://coveralls.io/github/tamada/scv?branch=setup_ci)
[![Go Report Card](https://goreportcard.com/badge/github.com/tamada/scv)](https://goreportcard.com/report/github.com/tamada/scv)
[![codebeat badge](https://codebeat.co/badges/5221e6ba-da64-45c1-8b13-f833f678e3b9)](https://codebeat.co/projects/github-com-tamada-scv-main)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?logo=spdx)](https://github.com/tamada/scv/blob/main/LICENSE)

Similarities and distance Calculator among Vectors.

## :speaking_head: Description

There are several algorithms to calculate the similarities of two bectors; however, no commands are exists treats them.
`scv` standardizes the interface for calculating the similarities and distances among vectors.


## :runner: Usage

```sh
scv [OPTIONS] <VECTORS...>
OPTIONS
    -a, --algorithm <ALGORITHM>    specifies the calculating algorithm.  This option is mandatory.
                                   The value of this option accepts several values separated with comma.
                                   Available values are: simpson, jaccard, dice, cosine, pearson,
                                   euclidean, and manhattan.
    -f, --format <FORMAT>          specifies the resultant format. Default is default.
                                   Available values are: default, json, and xml.
    -t, --input-type <TYPE>        specifies the type of VECTORS. Default is file.
                                   If TYPE is separated with comma, each type shows
                                   the corresponding VECTORS.
                                   Available values are: file, string, and json.
    -h, --help                     prints this message.
VECTORS
    the source of vectors for calculation.
```

## :athletic_shoe: Examples

```sh
$ scv -t string -a simpson distance similarity
simpson(distance, similarity) = 0.5000
$ scv -t string -a jaccard,dice distance similarity
jaccard(distance, similarity) = 0.3333
dice(distance, similarity) = 0.5000
```

## :smile: About

### :man_office_worker: Authors :woman_office_worker:

* Haruaki Tamada ([tamada](https://github.com/tamada))

### :scroll: License

[Apache 2.0](https://github.com/tamada/scv/blob/main/LICENSE)

### :jack_o_lantern: Icon

![Icon](https://github.com/tamada/scv/blob/main/docs/static/images/scv.png)
