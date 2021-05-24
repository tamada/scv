# scv

Similarities and distance Calculator among Vectors.

## Description

There are several algorithms to calculate the similarities of two bectors; however, no commands are exists treats them.
`scv` standardizes the interface for calculating the similarities and distances among vectors.


## Usage

```sh
scv [OPTIONS] <VECTORS...>
OPTIONS
    -a, --algorithm <ALGORITHM>    specifies the calculating algorithm.  This option is mandatory.
                                   The value of this option accepts several values separated with comma.
                                   Available values are: simpson, jaccard, dice, and cosine.
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

## Examples

```sh
$ scv -t string -a simpson distance similarity
simpson(distance, similarity) = 0.5000
$ scv -t string -a jaccard,dice distance similarity
jaccard(distance, similarity) = 0.3333
dice(distance, similarity) = 0.5000
```

## About

### Authors

* Haruaki Tamada ([tamada](https://github.com/tamada))

### License

[Apache 2.0](https://github.com/tamada/scv/blob/main/LICENSE)

### Icon

![Icon](https://github.com/tamada/scv/blob/main/docs/static/images/scv.png)
