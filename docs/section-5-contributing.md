---
title: Contributing
permalink: contributing
---

# Contributing

This document provides guidelines and information on contributing to Eask.

## Code of Conduct

Contributors to Eask should abide by the [Contributor Covenant](https://www.contributor-covenant.org/version/1/4/code-of-conduct).

## Developing Eask

### 🚩 Prerequisites

To make changes to Eask, you should have:

1. [Node.js](https://nodejs.org/en/) the development environment.
2. [npm](https://www.npmjs.com/) for the package manager.
3. [yargs](https://github.com/yargs/yargs) for parsing the command-line.
4. [Emacs](https://www.gnu.org/software/emacs/), 26.1 or above!

## 📂 State of the project

The project's barebones are pretty much done, we are currently looking for
contirbutors to give us feedback and improve our TUI/UX for this tool!

We are also looking for advice to add more Emacser often use commands and
options, so these features are prepared by default! Like command `lint` 
(package-lint) or option `--debug` refers to `debug-on-error` to `t`!

## 🏆 Goals for the project

1. Consistent sandbox testing across all systems
2. Has most commands that Emacsers often used (byte-compile, checkdoc, etc)
3. No dependencies installed by default; only install packages when it's needed!
4. Extra: Package Management Tool for your personal configuration/package

## 📂 Project Structure

The project structure are very simple, all we need is to look into 3 places.

1. `eask` file at the root of the project
2. `cmds` folder with all available commands
3. `lisp` folder with all elisp code
  - `lisp/extern` is the external modules/packages we used

`eask` is the node entry, and the main yargs definition! `cmds` and `lisp`
folders are command files that correspond to each other.

The sandbox logic is stored inside file `./lisp/_prepare.el`.
