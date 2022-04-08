---
title: Domain Specific Language
permalink: dsl
---

# Domain Specific Language (DSL)

This document provides a reference on the [DSL](https://en.wikipedia.org/wiki/Domain-specific_language).

## Package metadata

#### 🔎 **package** (`name` `version` `description`)

Declare a package with the given name, version, and description:

```elisp
(package "ert-runner" "0.7.0" "Opinionated Ert testing workflow")
```

All arguments are strings. The version must be a version understood by Emacs'
built-in `version-to-list`.

#### 🔎 **package-file** (`file` `version` `description`)

Define this package and its runtime dependencies from the package headers 
of a file (used only for package development).

Example:

```elisp
(package-file "foo.el")
```

## Package contents

#### 🔎 **files** (`&rest patterns`)

Specify list of files that are included in this project.

Example:

```elisp
(files "foo.el")
(files "*.el" "core/*.el")
```

## Dependencies

#### 🔎 **depends-on** (`package-name` `&optional minimum-version`)
#### 🔎 **depends-on** (`package-name` `&rest recipe`)

Specify a dependency of this package.

Example:

```elisp
(depends-on "dash")
(depends-on "company")
(depends-on "auto-rename-tag" 
            :repo "jcs-elpa/auto-rename-tag" 
            :fetcher 'github)
(depends-on "lsp-ui" 
            :repo "emacs-lsp/lsp-ui"
            :fetcher 'github
            :files '(:defaults "lsp-ui-doc.html" "resources"))
```

#### 🔎 **development** (`&rest body`)

Scope all `depends-on` expressions in body to development.

Example:

```elisp
(development
 (depends-on "ert-runner")
 (depends-on "elsa"))
```

#### 🔎 **source** (`alias`)
#### 🔎 **source** (`name` `url`)

Add a package archive to install dependencies from.

Example:

```elisp
(source "gnu")
(source "gnu" "https://elpa.gnu.org/packages/")
```

Available aliases:

* `gnu` ([https://elpa.gnu.org/packages/](https://elpa.gnu.org/packages/))
* `nongnu` ([https://elpa.nongnu.org/nongnu/](https://elpa.nongnu.org/nongnu/))
* `celpa` ([https://celpa.conao3.com/](https://celpa.conao3.com/))
* `jcs-elpa` ([https://jcs-emacs.github.io/jcs-elpa/packages/](https://jcs-emacs.github.io/jcs-elpa/packages/))
* `marmalade` ([https://marmalade-repo.org/packages/](https://marmalade-repo.org/packages/))
* `melpa` ([https://melpa.org/packages/](https://melpa.org/packages/))
* `melpa-stable` ([https://stable.melpa.org/packages/](https://stable.melpa.org/packages/))
* `org` ([https://orgmode.org/elpa/](https://orgmode.org/elpa/))
* `shmelpa` ([https://shmelpa.commandlinesystems.com/packages/](https://shmelpa.commandlinesystems.com/packages/))

*💡 You can use `--insecure` to make `https` to `http`, but not recommended*

#### 🔎 **source-priority** (`name` `priority`)

Set archive priority.

Example:

```elisp
(source-priority "gnu" 5)
```

## Execution Order for Eask

## Example

`Eask` is the magic file that `eask` will read it as the init file in Emacs.
The syntaxes are similar to the `Cask` file, but different.

```elisp
(package "your-package" 
         "1.0.0" 
         "Your package description")

(package-file "your-package-file.el")

(source "gnu")
```

## Advanced Usage

`Eask` is just the regular Emacs Lisp file and should be read from the
Emacs itself!

```elisp
; Regular Eask file content...

(setq byte-compile-error-on-warn t)  ; Singal error if warning occurred
```

`eask` provides some hooks so you can define your own action before/after
each command. The name of the hook follows the rule of
`eask-{`before`/`after`}-{`command_name`}-hook`.

For example, to enable compile on warn on `byte-compile` command

```elisp
(add-hook 'eask-before-compile-hook 
          (lambda () (setq byte-compile-error-on-warn t)))
```
