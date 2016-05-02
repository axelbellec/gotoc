![License](http://img.shields.io/:license-mit-blue.svg)

# gotoc
:page_facing_up: Markdown table of contents generator

![logo_godoc](img/gotoc.png)

`gotoc` is a tool to generate markdown table of contents inside local git repo.
Links generated refers to Github anchors.

## Installation

```sh
git clone https://github.com/Belekkk/gotoc.git gotoc
cd gotoc
go build gotoc.go
```

## Usage

```sh
./gotoc -file=README.md
```

```md
<!-- Table of Contents generated by [gotoc](https://github.com/Belekkk/gotoc) -->
**Table of Contents**
- [gotoc](#gotoc)
  - [Usage](#usage)
  - [Custom TOC title](#custom-toc-title)
  - [Max heading level for TOC entries](#max-heading-level-for-toc-entries)
```

## Custom TOC title

To specify custom TOC title like `**Repo : Table of Contents**` you can pass the argument : `-title='<yourtitle>'`.
To remove title from TOC, just use the option `-notitle`.

## Max heading level for TOC entries

To limit TOC entries to a specified level of headings, use : `-depth=3`.
