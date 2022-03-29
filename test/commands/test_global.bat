@echo off

:: Copyright (C) 2022 Jen-Chieh Shen

:: This program is free software; you can redistribute it and/or modify
:: it under the terms of the GNU General Public License as published by
:: the Free Software Foundation; either version 3, or (at your option)
:: any later version.

:: This program is distributed in the hope that it will be useful,
:: but WITHOUT ANY WARRANTY; without even the implied warranty of
:: MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
:: GNU General Public License for more details.

:: You should have received a copy of the GNU General Public License
:: along with GNU Emacs; see the file COPYING.  If not, write to the
:: Free Software Foundation, Inc., 51 Franklin Street, Fifth Floor,
:: Boston, MA 02110-1301, USA.

::: Commentary:
::
:: Here we test all global (~/.emacs.d/) that the Emacser can be use daily!
::
:: Notice, to make global commands work; we need a minimum configuration
:: (mini.emacs.d), and place it under to the default Emacs directory!
::

echo "Copy test configuration"
robocopy /e "./test/mini.emacs.d/" "%UserProfile%/.emacs.d"

:: Test for global commands
eask archives -g
eask install -g spinner ivy beacon
eask list --depth=0 -g