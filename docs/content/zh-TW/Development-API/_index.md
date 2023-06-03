---
title: 開發 API
weight: 700
---

本文檔提供了對公共 Eask API 的引用，您可以在您的項目和 Eask 的擴展。

{{< toc >}}

# 🚩 入口點

## 🔍 Snippet: _prepare.el

加載 `lisp/_prepare.el` 以開始使用其他 Eask API。

```elisp
(let ((dir (file-name-directory (nth 1 (member "-scriptload" command-line-args)))))
  (load (expand-file-name "_prepare.el"
                          (locate-dominating-file dir "_prepare.el"))
        nil t))
```

每個 Elisp 腳本都應該在文件的最頂部有這個片段。

## 🔍 Macro: eask-start (&rest `body`)

命令入口點。 每個命令文件都應在文件的某處包含此宏。

```elisp
(eask-start
  ;; TODO: 在這裡設計你的命令！
  )
```

# 🚩 核心

## 🔍 Variable: eask-lisp-root

從項目根目錄指向 `lisp` 目錄。

```elisp
(message "%s" eask-lisp-root)  ; path/to/eask/cli/lisp/
```

## 🔍 Function: eask-command ()

返回字符串中的當前命令。

假設命令是：

```sh
$ eask init
```

然後，

```elisp
(message "%s" (eask-command))  ; init
```

## 🔍 Function: eask-special-p ()

如果在沒有 Eask 文件存在的情況下可以運行的命令，則返回 `t`。

這允許一些命令仍然可以在不定義用戶的情況下執行目錄。 當您想在沒有的情況下進行正常操作時，這會很方便
觸摸用戶目錄。

## 🔍 Function: eask-checker-p ()

如果運行 Eask 作為檢查器，則返回 `t`。

如果沒有這個標誌，一旦發生錯誤，進程就會終止。此標誌允許您在不報告錯誤的情況下運行所有操作。

## 🔍 Function: eask-script (`script`)

返回完整的腳本文件名。

```elisp
(eask-script "extern/pacakge")  ; {project-root}/lisp/extern/package.el
```

## 🔍 Function: eask-load (`script`)

加載另一個 eask 腳本。

```elisp
(eask-load "extern/ansi")  ; load {project-root}/lisp/extern/ansi.el file
```

## 🔍 Function: eask-call (`script`)

調用另一個 eask 腳本。

```elisp
(eask-call "clean/elc")  ; call command `eask clean-elc`
```

{{< hint info >}}
💡 這個很少用！
{{< /hint >}}

## 🔍 Macro: eask-defvc< (`version` &rest `body`)

如果 Emacs 版本低於特定版本，則定義範圍。

`VERSION` 是一個整數，將與 `emacs-major-version` 進行比較。

```elisp
(eask-defvc< 28
  ;; 這在 Emacs 28 之前是缺失的； 定義它!
  (defvar package-native-compile nil))
```

{{< hint info >}}
💡 這用於 Emacs 兼容性！
{{< /hint >}}

## 🔍 Macro: eask--silent (&rest `body`)

將來自範圍內標準輸出的所有消息靜音。

```elisp
(eask--unsilent (message "你聽不到我! :("))
```

## 🔍 Macro: eask--unsilent (&rest `body`)

取消靜音來自範圍內標準輸出的所有消息。

```elisp
(eask--unsilent (message "你聽的到我! :)"))
```

## 🔍 Function: eask-dependencies ()

返回依賴項列表。

元素應該是 `(NAME . VERSION)` 或 `(NAME . RECIPE-FORMAT)`。

## 🔍 Function: eask-pkg-init (&optional `force`)

初始化包以供使用。

```elisp
(eask-start
  (eask-pkg-init)
  ;; 現在您可以使用安裝在 `package-user-dir` 中的包
  )
```

{{< hint info >}}
💡 這通常在 **eask-start** 之後調用！
{{< /hint >}}

## 🔍 Macro: eask-with-archives (`archives` &rest `body`)

臨時使存檔可用的範圍。

`ARCHIVES` 可以是字符串或字符串列表。

```elisp
(eask-with-archives "melpa"
  (eask-package-install 'package-build))  ; 安裝僅在 MELPA 中定義的包
```

{{< hint info >}}
💡 當您需要某些檔案中的某些包時，這很方便。
{{< /hint >}}

## 🔍 Function: eask-package-desc (`name` &optional `current`)

為包構建包描述符。

`CURRENT` 表示已安裝的包； 否則它將返回任何可用的來自選定包檔案的包。

## 🔍 Function: eask-argv (`index`)

通过索引返回一个命令行参数。

## 🔍 Function: eask-args ()

返回從命令行參數中提取的列表。

```sh
$ eask info --verbose 4 foo bar
```

它會忽略 `--verbose` 和 `4`，只返回 `foo` 和 `bar`。

## 🔍 Variable: eask-file

當前加載的 Eask 文件的路徑。

## 🔍 Variable: eask-file-root

當前加載的 Eask 文件的目錄。

## 🔍 Function: eask--match-file (`name`)

檢查 NAME 是否是我們的目標 Eask 文件，然後返回它。

以下輸出來自 Emacs 28.1：

```elisp
(eask--match-file "Eask")         ; t
(eask--match-file "Eask.28")      ; t
(eask--match-file "Eask.28.1")    ; t
(eask--match-file "Eask.29")      ; nil

(eask--match-file "Easkfile")     ; t
(eask--match-file "Easkfile.28")  ; t
(eask--match-file "Easkfile.29")  ; nil
```

## 🔍 Function: eask--all-files (&optional `dir`)

從 DIR 返回 Eask 文件列表。

考慮以下目錄樹：

```
. root
├── Eask
├── Eask.28
└── Eask.29
```

以下輸出來自 Emacs 28.1：

```elisp
(eask--all-files "/root/")  ; '(Eask Eask.28)
```

## 🔍 Function: eask--find-files (`start-path`)

從 START-PATH 找到 Eask 文件。

考慮以下目錄樹：

```
.project
├─ src
│ └── config.el
├── Eask
├── Eask.28
└── Eask.29
```

以下輸出來自 Emacs 28.1：

```elisp
(eask--find-files "/project/src/config.el")  ; '(/project/Eask /project/Eask.28)
```

## 🔍 Function: eask-file-try-load (`start-path`)

嘗試在 START-PATH 中加載 Eask 文件。

```elisp
(eask--find-files "/project/src/")  ; t
```

## 🔍 Function: eask-network-insecure-p ()

如果當前 Emacs 會話允許不安全的網絡連接，則返回 `t`。

# 🚩 Flags

## 🔍 Function: eask-global-p ()

如果啟用了 `global` 選項，則返回 `t`。

```elisp
(when (eask-global-p)
  user-emacs-directory)   ; ~/.eask/
```

## 🔍 Function: eask-config-p ()

如果啟用了 `config` 選項，則返回 `t`。

```elisp
(when (eask-config-p)
  user-emacs-directory)   ; ~/.emacs.d
```

{{< hint info >}}
💡 如果選項 `--config` 和 `--global` 都打開，則選擇全局空間。
{{< /hint >}}

## 🔍 Function: eask-local-p ()

這使用當前工作區，這是默認設置。

```elisp
(when (eask-local-p)
  user-emacs-directory)   ; ./.eask/{emacs-version}/
```

{{< hint info >}}
💡 此函數僅在 `(eask-global-p)` 和 `(eask-config-p)` 是 false 時返回 `t`！
{{< /hint >}}

## 🔍 Function: eask-all-p ()

如果啟用了 `all` 選項，則返回 `t`。

```elisp
(when (eask-all-p)
  ;; 運行所有測試
  ...)
```

## 🔍 Function: eask-quick-p ()

如果啟用了 `quick` 選項，則返回 `t`。

```elisp
(unless (eask-quick-p)
  (load user-init-file)
  ...)
```

## 🔍 Function: eask-force-p ()

如果啟用了 `force` 選項，則返回 `t`。

```elisp
(package-delete .. (eask-force-p))
```

## 🔍 Function: eask-dev-p ()

如果啟用了 `development` 選項，則返回 `t`。

```elisp
(when (eask-dev-p)
  (package-install 'ert-runner))  ; 安裝開發依賴
```

## 🔍 Function: eask-debug-p ()

如果啟用了 `debug` 選項，則返回 `t`。

```elisp
(when (eask-debug-p)
  (error "在調試模式下執行..."))
```

## 🔍 Function: eask-strict-p ()

如果啟用了 `strict` 選項，則返回 `t`。

```elisp
(setq byte-compile-error-on-warn (eask-strict-p))
```

## 🔍 Function: eask-timestamps-p ()

如果啟用/禁用 `timestamps` 選項，則返回 `t` / `nil`。

這些標誌不能在同一命令中共存。

```elisp
(when (eask-timestamps-p)
  (message "打印帶有時間戳的日誌！"))
```

## 🔍 Function: eask-log-level-p ()

如果啟用/禁用 `log-level` 選項，則返回 `t` / `nil`。

這些標誌不能在同一命令中共存。

```elisp
(when (eask-log-level-p)
  (message "打印帶有級別前綴的日誌！"))
```

## 🔍 Function: eask-log-file-p ()

如果啟用/禁用 `log-file` 選項，則返回 `t` / `nil`。

這些標誌不能在同一命令中共存。

```elisp
(when (eask-log-file-p)
  (message "讓我們創建一個日誌文件！"))
```

## 🔍 Function: eask-no-color-p ()

如果啟用了 `color` 選項，則返回 `t`。

```elisp
(unless (eask-no-color-p)
  (message "此字符串沒有 ANSI 代碼！"))
```

## 🔍 Function: eask-allow-error-p ()

如果啟用了 `allow-error` 選項，則返回 `t`。

```elisp
(unless (eask-allow-error-p)
  (error "停在這裡。"))
```

## 🔍 Function: eask-insecure-p ()

如果啟用了 `insecure` 選項，則返回 `t`。

```elisp
(when (eask-insecure-p)
  ;; 做一些危險的工作？
  )
```

## 🔍 Function: eask-proxy ()
## 🔍 Function: eask-http-proxy ()
## 🔍 Function: eask-https-proxy ()
## 🔍 Function: eask-no-proxy ()

返回一個 **string** 表示 `hostname` + `port number`。

```sh
$ eask [command] --proxy "localhost:1000"
$ eask [command] --http-proxy "localhost:2000"
$ eask [command] --https-proxy "localhost:3000"
$ eask [command] --no-proxy "localhost:4000"
```

## 🔍 Function: eask-destination ()

返回一個 **string** 表示目的地（輸出路徑）。

```elisp
(write-file (or (eask-destination) "./dist"))  ; 將文件寫入目標
```

## 🔍 Function: eask-depth ()

返回一個**整數**表示當前打印層級的深度。

```elisp
(setq print-level (eask-depth))
```

## 🔍 Function: eask-verbose ()

返回一個 **整數** 表示冗長級別。

```elisp
(when (= (eask-verbose) 4)
  (setq byte-compile-verbose t))
```

# 🚩 `Eask`-file

這些函數是 `Eask`-file DSL 的實際實現； 和將單詞 `eask-` 作為函數前綴。

有關詳細信息，請參閱 [DSL](https://emacs-eask.github.io/DSL/) 部分。

## 🔍 Variable: eask-package

它在 plist 中保存包的 `NAME`、`VERSION` 和 `DESCRIPTION`。

```elisp
(plist-get eask-package :name)  ; 返回包名
```

從該變量擴展的三個函數：

- `(eask-package-name)`
- `(eask-package-version)`
- `(eask-package-description)`

## 🔍 Variable: eask-package-file

指向打包主文件。

## 🔍 Variable: eask-package-desc

來自包主文件的包描述符。

```elisp
(package-desc-p eask-package-desc)  ; 返回 t
```

{{< hint warning >}}
⚠ 如果不能正確構造包描述符，這可以是 **nil**！
{{< /hint >}}

## 🔍 Variable: eask-files

持有通配符規範中的文件模式列表。

## 🔍 Variable: eask-scripts

包含可用腳本的列表，用戶可以使用 `eask run-script` 命令。

## 🔍 Variable: eask-depends-on-emacs

保存有關 Emacs 最低版本的信息。

```elisp
(depends-on "emacs" "26.1")
```

函數將返回字符串中的 Emacs 版本。

- `(eask-depends-emacs-version)` - 返回 `"26.1"`

## 🔍 Variable: eask-depends-on

持有依賴項列表。

## 🔍 Variable: eask-depends-on-dev

持有開發使用的依賴項列表。

## 🔍 Function: eask-f-package (`name` `version` `description`)

別名 `package`.

## 🔍 Function: eask-f-website-url (`url`)

別名 `website-url`.

## 🔍 Function: eask-f-keywords (&rest `keywords`)

別名 `keywords`.

## 🔍 Function: eask-f-author (`name` &optional `email`)

別名 `author`.

## 🔍 Function: eask-f-license (`name`)

別名 `license`.

## 🔍 Function: eask-f-package-file (`file`)

別名 `package-file`.

## 🔍 Function: eask-f-files (`pkg` &rest `args`)

別名 `files`.

## 🔍 Function: eask-f-script (`name` `command` &rest `args`)

別名 `script`.

## 🔍 Function: eask-f-source (`name` &optional `location`)

別名 `source`.

## 🔍 Function: eask-f-source-priority (`name` &optional `priority`)

別名 `source-priority`.

## 🔍 Function: eask-f-depends-on (`pkg` &rest `args`)

別名 `depends-on`.

## 🔍 Function: eask-f-development (&rest `dependencies`)

別名 `development`.

## 🔍 Function: eask-f-exec-paths (&rest `dirs`)

別名 `exec-paths`.

## 🔍 Function: eask-f-load-paths (&rest `dirs`)

別名 `load-paths`.

# 🚩 Logging

具有時間戳和日誌級別的記錄器實用程序。

日誌級別值在函數 `eask--verb2lvl` 中定義。

| 等級    | 描述                                               | 值 |
|:--------|:---------------------------------------------------|:---|
| `debug` | 指定對調試應用程序最有用的細粒度信息事件。         | 4  |
| `log`   | 指定普通消息。                                     | 3  |
| `info`  | 指定在粗粒度級別突出顯示應用程序進度的信息性消息。 | 2  |
| `warn`  | 指定潛在的有害情況。                               | 1  |
| `error` | 指定可能仍允許應用程序繼續運行的錯誤事件。         | 0  |

The default level is `log`.

## 🔍 Variable: eask-verbosity

詳細級別表示為整數。

```elisp
(setq eask-verbosity 4)  ; 你可以設置從 0 到 4
```

## 🔍 Variable: eask-timestamps

記錄帶有時間戳的消息。

```elisp
(setq eask-timestamps t)
```

Output:

```
2022-04-14 13:44:46 這是一條帶有時間戳的消息
```

## 🔍 Variable: eask-log-level

記錄消息級別。 （默認值：`nil`）

```elisp
(setq eask-log-level t)
```

輸出：

```
[DEBUG] 這是一條具有日誌級別的 DEBUG 消息
```

## 🔍 Variable: eask-log-file

天氣生成日誌文件。 （默認值：`nil`）

```elisp
(setq eask-log-level t)
```

使用命令 `cat` 查看日誌，

```
$ cat /.log/messages.log
```

## 🔍 Variable: eask-level-color

定義每個日誌級別顏色。

```elisp
(setq eask-level-color
      '((debug . ansi-blue)
        (log   . ansi-white)
        (info  . ansi-cyan)
        (warn  . ansi-yellow)
        (error . ansi-red)))
```

## 🔍 Macro: eask-with-verbosity (`symbol` &rest `body`)

使用詳細級別定義執行。

```elisp
(eask-with-verbosity 'debug
  ;; TODO: 在這裡執行..
  )
```

除非冗長，否則此宏範圍內的所有內容都將被靜音。 僅當您指定 `--verbose 4` 時才會打印
全局選項。

## 🔍 Function: eask-debug (`msg` &rest `args`)

```elisp
(eask-debug "這是調試信息")
```

```
2022-04-14 17:31:54 [DEBUG] 這是調試信息
```

## 🔍 Function: eask-log (`msg` &rest `args`)

```elisp
(eask-log "這是日誌消息")
```

```
2022-04-14 17:31:54 [LOG] 這是日誌消息
```

## 🔍 Function: eask-info (`msg` &rest `args`)

```elisp
(eask-info "這是信息消息")
```

```
2022-04-14 17:31:54 [INFO] 這是信息消息
```

## 🔍 Function: eask-warn (`msg` &rest `args`)

```elisp
(eask-warn "這是警告消息")
```

```
2022-04-14 17:31:54 [WARNING] 這是警告消息
```

## 🔍 Function: eask-error (`msg` &rest `args`)

```elisp
(eask-error "這是錯誤信息")
```

```
2022-04-14 17:31:54 [ERROR] 這是錯誤信息
```

## 🔍 Function: eask-msg (`msg` &rest `args`)

類似於 `message` 函數，但會用顏色替換 unicode。

```elisp
(eask-msg "用換行符打印此消息！")
```

## 🔍 Function: eask-write (`msg` &rest `args`)

類似於 eask-msg 函數，但末尾沒有換行符。

```elisp
(eask-write "不帶換行符打印此消息...")
```

## 🔍 Function: eask-report (&rest `args`)

報告錯誤/警告取決於嚴格標誌。

```elisp
(eask-report "This can be warning or error")
```

見選項 [--strict](https://emacs-eask.github.io/Getting-Started/Commands-and-options/#---strict).

# 🚩 File

## 🔍 Function: eask-guess-package-name ()

返回可能的包名稱。

## 🔍 Function: eask-package-files ()

返回包文件列表。

## 🔍 Function: eask-package-el-files ()

返回擴展名為 `.el` 的包文件列表。

## 🔍 Function: eask-package-elc-files ()

返回擴展名為 `.elc` 的包文件列表。

## 🔍 Function: eask-package-multi-p ()

如果是單個文件包，則返回 `nil`。

## 🔍 Function: eask-package-single-p ()

如果是單個文件包，則返回 `t`。

## 🔍 Function: eask-unpacked-size ()

返回當前包的大小。

{{< hint warning >}}
⚠️ 這將返回一個字符串而不是字節。
{{< /hint >}}

# 🚩 Progress

## 🔍 Macro: eask-with-progress (`msg-start` `body` `msg-end`)

使用響應消息輸出創建執行。

```elisp
(eask-with-progress 
  "檔案下載中s... "
  (eask-with-verbosity 'debug  ; 通常與 `eask-with-verbosity` 一起使用
    ;; 執行一些操作..
    )
  "完成 ✓")
```

期望輸出：

```
檔案下載中... 完成 ✓
```

## 🔍 Function: eask-print-log-buffer (&optional `buffer-or-name`)

打印緩衝區並突出顯示“錯誤和“警告”。

```elisp
(eask-print-log-buffer "*Package-Lint*")
```

{{< hint info >}}
💡 這對於創建 buffer 來顯示 **errors** 和 **warnings** 的 linters 會很方便。
{{< /hint >}}

# 🚩 Help

## 🔍 Function: eask-help (`command`)

打印位於 `lisp/help/` 目錄下的幫助手冊。

```elisp
(eask-help "core/search")
```

{{< hint info >}}
💡 這是在命令失敗時使用的，想給用戶一些提示！
{{< /hint >}}
