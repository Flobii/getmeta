# GetMeta

A simple little script that neatly display all of the `meta` tags of a website.

## Installation

You can either download the source code and run `make install` in its directory or
download one of the binaries and add the path to it to your `PATH`.

- on Windows you can also just place the file in `C:\WINDOWS\System32`
- alternatively [here's a tutorial][tutorial] on how to edit the path on Windows
- on Linux: `export PATH=$PATH:/path/to/the/binary`
- on MacOS: no idea, sorry!

## Usage

`getmeta <url>`

**Example**

```bash
$ getmeta http://neuralnetworksanddeeplearning.com/

citation_title:             Neural Networks and Deep Learning
citation_author:            Nielsen, Michael A.
citation_publication_date:  2015
citation_fulltext_html_url: http://neuralnetworksanddeeplearning.com
citation_publisher:         Determination Press
```

[tutorial]: https://www.youtube.com/watch?v=gb9e3m98avk
