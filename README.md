# cham

A programming language based on the Chemical Abstract Machine.

This repository contains the source code for an interpreter that can run programs written in the language.   
Some additional documentation can be found in the [`docs/`](docs/) directory.

## Language

* A **guide** explaining how the language works can be found in [`language.md`](docs/language.md).
* The language **syntax** is described in [`syntax.md`](docs/syntax.md) and is accompanied by a BNF grammar in [`syntax.ebnf`](docs/syntax.ebnf).
* Some **sample programs** can be found in [`programs.md`](docs/programs.md).

## Interpreter

The interpreter is written using the [Go](https://golang.org/) programming language.

You can **compile from source** if you have the [Go toolchain](https://golang.org/doc/install) installed on your system. Just execute `go build` from the root project directory.

Alternatively, you can **download pre-built binaries** for Windows, MacOS and Linux. These can be found on GitHub under the [Actions](https://github.com/howden/cham/actions) tab. Just select the build you want, then download the desired binary from the *Artifacts* section.

#### Using the interpreter on Mac or Linux

After compiling or downloading the binary, run the following in the Terminal.

```bash
$ ./cham
```

#### Using the interpreter on Windows

After compiling or downloading the binary, run the following in the Command Prompt.

```bash
$ cham.exe
```

#### Interpreter modes

The interpreter has two modes.

In **REPL mode**, the interpreter creates a prompt into which you can continuously run programs or execute commands. Defined programs will be *stored* in memory until the process exits.

In **command-line mode**, you can perform one-off operations using the interpreter, such as evaluating a program (either from a command-line argument or read from a file) or viewing the raw output from the lexer or parser.

To view usage information, run `./cham -h` (Mac/Linux) or `cham.exe -h` (Windows).

#### Interpreter Design

The interpreter follows a fairly standard design. Program source code passes through a lexer and parser, and is then evaluated.

The evaluator is capable of executing "reaction" programs in parallel.
