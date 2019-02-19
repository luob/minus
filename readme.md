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


------

## 什么是Minus

Minus就像流行的静态站点生成器jekyll, hugo那样, 使用markdown语法写文章, Minus用这些文章渲染成一个HTML站点. Minus的含义是做减法, 减去一切不需要的功能, 减去所有第三方依赖. 所有代码都是Go写的，只使用了Go的标准库, 实现了一个简单的markdown解析器, 但是有些功能比如图像调整大小看起来很难实现。

### 功能

- 快速安装, 快速生成
- 无需了解更多框架的用法（~~真的没有了~~）
- traditonal markdown支持
- 行内代码和代码块支持
- 使用Go的html/template语法自定义模板

### 不支持

- 任何更多扩展语法
- FrontMatter写法
- 图像调整大小和压缩

## todo

1. google SEO优化
2. markdown解析器优化
3. rss订阅