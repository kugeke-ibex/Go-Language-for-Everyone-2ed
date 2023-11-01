#!/bin/sh

# git describe --tagsは最新のタグのあとにcommitがあった場合には[最新のタグ]-[そのあとのcommit回数]-[commit hash]という値になる
GIT_VER=`git describe --tags`

# buildに-ldflagsを指定することで、バイナリのビルド時に外部から変数の値を設定
go build -ldflags "-X main.version=${GIT_VER}"
