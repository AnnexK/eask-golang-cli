---
title: Development API
weight: 700
---

This document provides a reference to the public Eask API, which you may use in
your projects and extensions to Eask.

{{< toc >}}

# 🚩 Entry Point

## 🔍 Snippet: _prepare.el

Load `lisp/_prepare.el` to start using other Eask API.

```elisp
(load (expand-file-name
       "../_prepare.el"
       (file-name-directory (nth 1 (member "-scriptload" command-line-args))))
      nil t)
```

Each Elisp scripts should have this snippet at the very top of the file.

## 🔍 Macro: eask-start (`&rest body`)

Command entry point. Each command file should contain this macro somewhere in the file.

```elisp
(eask-start
  ;; TODO: design your command here!
  )
```

# 🚩 Flags

## 🔍 Function: eask-global-p ()

Return `t` if the `global` option is enabled.

```elisp
(if (eask-global-p)
    user-emacs-directory   ; ~/.emacs.d
  user-emacs-directory)    ; ./.eask/{emacs-version}/
```

## 🔍 Function: eask-force-p ()

Return `t` if the `force` option is enabled.

```elisp
(package-delete .. (eask-force-p))
```

## 🔍 Function: eask-dev-p ()

Return `t` if the `development` option is enabled.

```elisp
(when (eask-dev-p)
  (package-install 'ert-runner))  ; install development dependency
```

## 🔍 Function: eask-debug-p ()

Return `t` if the `debug` option is enabled.

```elisp
(when (eask-debug-p)
  (error "Executing in debug mode..."))
```

## 🔍 Function: eask-strict-p ()

Return `t` if the `strict` option is enabled.

```elisp
(setq byte-compile-error-on-warn (eask-strict-p))
```

## 🔍 Function: eask-timestamps-p ()
## 🔍 Function: eask-no-timestamps-p ()

Return `t`/`nil` if the `timestamps` option is enabled/disabled.

These flags can't co-exist in the same command.

```elisp
(when (and (eask-timestamps-p) (eask-no-timestamps-p))
  (error "This is impossible!"))
```

## 🔍 Function: eask-log-level-p ()
## 🔍 Function: eask-no-log-level-p ()

Return `t`/`nil` if the `log-level` option is enabled/disabled.

These flags can't co-exist in the same command.

```elisp
(when (and (eask-log-level-p) (eask-no-log-level-p))
  (error "This is impossible!"))
```

## 🔍 Function: eask-no-color-p ()

Return `t` if the `color` option is enabled.

```elisp
(unless (eask-no-color-p)
  (ansi-color-filter-apply "This string has no ansi code!"))
```

## 🔍 Function: eask-allow-error-p ()

Return `t` if the `allow-error` option is enabled.

```elisp
(unless (eask-allow-error-p)
  (error "Stop here."))
```

## 🔍 Function: eask-insecure-p ()

Return `t` if the `insecure` option is enabled.

```elisp
(when (eask-insecure-p)
  ;; Do some dangerous tasks?
  )
```

## 🔍 Function: eask-proxy ()
## 🔍 Function: eask-http-proxy ()
## 🔍 Function: eask-https-proxy ()
## 🔍 Function: eask-no-proxy ()

Return a **string** represents `hostname` + `port number`.

```sh
$ eask [command] --proxy "localhost:1000"
$ eask [command] --http-proxy "localhost:2000"
$ eask [command] --https-proxy "localhost:3000"
$ eask [command] --no-proxy "localhost:4000"
```

## 🔍 Function: eask-destination ()

Return a **string** represents the destination (output path).

```elisp
(write-file (or (eask-destination) "./dist"))  ; write file to destination
```

## 🔍 Function: eask-depth ()

Return an **integer** represents the depth of the current print level.

```elisp
(setq print-level (eask-depth))
```

## 🔍 Function: eask-verbose ()

Return an **integer** represents the verbosity level.

```elisp
(when (= (eask-verbose) 4)
  (setq byte-compile-verbose t))
```

# 🚩 `Eask`-file

These functions are the actual implementation of `Eask`-file DSL; and
have the word `eask-` as the function prefix.

See [DSL](https://emacs-eask.github.io/eask/dsl) section for more information.

## 🔍 Function: eask-package (`name` `version` `description`)

Alias of `package`.

## 🔍 Function: eask-package-file (`file`)

Alias of `package-file`.

## 🔍 Function: eask-files (`pkg` `&rest args`)

Alias of `files`.

## 🔍 Function: eask-depends-on (`pkg` `&rest args`)

Alias of `depends-on`.

## 🔍 Function: eask-development (`&rest dependencies`)

Alias of `development`.

## 🔍 Function: eask-load-paths (`&rest dirs`)

Alias of `load-paths`.

## 🔍 Function: eask-exec-paths (`&rest dirs`)

Alias of `exec-paths`.

## 🔍 Function: eask-source

Alias of `source`.

## 🔍 Function: eask-source-priority

Alias of `source-priority`.

# 🚩 Logging

Logger utility with timestamps and log level.

## 🔍 Variable: eask-verbosity

Verbosity level represent as an integer.

```elisp
(setq eask-verbosity 4)  ; you could set from 0 to 4
```

## 🔍 Variable: eask-timestamps

Log messages with timestamps.

```elisp
(setq eask-timestamps t)
```

Output:

```
2022-04-14 13:44:46 This is a message with timestamps
```

## 🔍 Variable: eask-log-level

Log messages with level.

```elisp
(setq eask-log-level t)
```

Output:

```
[DEBUG] This is a DEBUG message with log level
```

## 🔍 Variable: eask-level-color

Define each log level color.

```elisp
(setq eask-level-color
      '((debug . ansi-blue)
        (log   . ansi-white)
        (info  . ansi-cyan)
        (warn  . ansi-yellow)
        (error . ansi-red)))
```

## 🔍 Macro: eask-with-verbosity (`symbol` `&rest body`)

Define executions with the verbosity level.

```elisp
(eask-with-verbosity 'debug
  (message "Hello World!~"))
```

## 🔍 Function: eask-debug (`msg` `&rest args`)

```elisp
(eask-debug "This is DEBUG message")
```

```
2022-04-14 17:31:54 [DEBUG] This is DEBUG message
```

## 🔍 Function: eask-log (`msg` `&rest args`)

```elisp
(eask-log "This is LOG message")
```

```
2022-04-14 17:31:54 [LOG] This is LOG message
```

## 🔍 Function: eask-info (`msg` `&rest args`)

```elisp
(eask-info "This is INFO message")
```

```
2022-04-14 17:31:54 [INFO] This is INFO message
```

## 🔍 Function: eask-warn (`msg` `&rest args`)

```elisp
(eask-warn "This is WARNING message")
```

```
2022-04-14 17:31:54 [WARNING] This is WARNING message
```

## 🔍 Function: eask-error (`msg` `&rest args`)

```elisp
(eask-error "This is ERROR message")
```

```
2022-04-14 17:31:54 [ERROR] This is ERROR message
```

## 🔍 Function: eask-msg (`msg` `&rest args`)

Like `message` function but will replace unicodes with color.

```elisp
(eask-msg "This is a message")
```

## 🔍 Function: eask-write (`msg` `&rest args`)

Like `eask-msg` function but without newline at the end.

```elisp
(eask-write "This is a message")
```
