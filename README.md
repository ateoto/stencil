stencil
=======

Stencil is a simple go program to output templates via cli. 

Warning: Very early in development and the cli may have breaking changes.

Usage
-----

```bash
stencil - Take a go template and variables from the commandline, create output

USAGE:
   stencil [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --template               Template file
   --output                 Specify output file
   --var [--var option --var option]    Variables to be used in template
   --help, -h               show help
   --version, -v            print the version
```

If no output file is specified, output is rendered to stdout.

Examples
--------

```bash
$ echo Hello, {{ .name }}! > hello.txt.tpl
$ stencil --template hello.txt.tpl --var name=Matt
Hello, Matt!

$ stencil --template hello.txt.tpl --output hello.txt --var name=Matt
$ cat hello.txt
Hello, Matt!
```
