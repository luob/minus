# Minus

[![GoDoc](https://godoc.org/github.com/luob/minus?status.svg)](https://godoc.org/github.com/luob/minus)
[![Go Report Card](https://goreportcard.com/badge/github.com/luob/minus)](https://goreportcard.com/report/github.com/luob/minus)
[![Code Size](https://img.shields.io/github/languages/code-size/luob/minus.svg)]()
[![Lisense](https://img.shields.io/github/license/luob/minus.svg)](LICENSE)
<!-- [![Build Status](https://travis-ci.com/luob/minus.svg?branch=master)](https://travis-ci.com/luob/minus) -->

An absolutely lightweight static site generator without any third-party dependency written in Go.

## What is Minus

Minus is like the popular static site generator jekyll, hugo. You write posts in markdown, and Minus render these posts to html. Minus means to do subtraction, Minus all unnecessary functions, Minus all third-party dependencies. All the code is written in go, only use the Go's standard library, Implemented a simple markdown parser, but some functions such as image resize are difficult to implement.

### feature

- fast install and fast generating
- no need to learn more usage of the framework(~~really no more~~)
- traditonal markdown support
- inline code and code blocks support
- custom template with Go's html/template syntax

### not support

- any more extended syntax
- front-matters
- image resize or compress

## Install


Install binary
```
// todo
```

Install from source

```shell
git clone https://github.com/luob/minus
cd minus
go install
```

## Usage

make a directory like this

```shell
├── config.json
├── posts
│   ├── post1.md
│   └── post2.md
└── template
    ├── footer.html
    ├── header.html
    ├── index.html
    ├── post.html
    └── tag.html
```

then run minus

``` shell
minus

// or
minus <path-to-your-directory>
```

Minus will generate the site in the `/public` directory, then you can deploy it to the server.

## Auto Deploy

You can use travis, circleci or netlify to automate this deployment without installing minus. This is a travis configuration file example:

```yaml

```

## Contributing

Welcome to modify some small bugs by pull request. For major changes,
There are two rules to follow:

## todo

1. 二进制发布
1. 通过目录实现文章分类功能
2. 自定义页面功能
3. 多模板支持
4. 重构markdown模块以外的部分
4. 重写markdown解释器, 一次遍历解析成ast, 时间复杂度降到O(n)


