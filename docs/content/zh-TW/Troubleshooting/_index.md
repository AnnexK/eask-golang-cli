---
title: 故障排除
weight: 800
---

本文檔幫助您排除 Eask 故障。

{{< toc >}}

## 🚩 可能的錯誤變量

一些潛在的變量可能導致錯誤的Eask，請檢查:

- 用 `PATH` 安裝並設置 Emacs
- Eask安裝正確
- Node 版本應為 `14.x` 或更高版本

## ⛔️ Error when running an Eask command

如果您運行 Eask 命令並遇到錯誤，您可以嘗試一些方法：

- 確保您擁有最新的 Eask 版本。 您可以使用 `eask --version` 來確定當前的 Eask 版本。
- 如果您選擇從 `npm` 安裝，請使用 `eask upgrade-eask` 或 `npm install -g @emacs-eask/cli@latest` 升級 Eask

{{< hint warning >}}
**⚠ 警告**

如果您使用 **npm** 安裝了 Eask，那麼您應該通過 **npm** 升級它。 否則，您只需確保安裝了 **git**。
{{< /hint >}}

- 如果錯誤仍然存在，請嘗試從頭開始重新安裝 Eask。

如果Eask還是不行，請 [提交 issue](https://github.com/emacs-eask/cli/issues/new) 到 issue tracker.
請在啟用 [--verbose 4] 和 [--debug] 選項的情況下包含 Eask 輸出，為我們提供盡可能多的信息。
