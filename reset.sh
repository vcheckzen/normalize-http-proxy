#!/usr/bin/env bash

[ "$(command -v git.exe)" ] && git=git.exe || git=git

$git checkout --orphan latest_branch
$git rm -rf --cached .
$git add -A
$git commit -m "$1"
$git branch -D main
$git branch -m main
$git push -f -u origin main
