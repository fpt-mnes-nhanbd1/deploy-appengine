#!/bin/sh

# VERSION=develop # AppEngineにデプロイするバージョン名
if [ -x "`which git `" ]; then
  VERSION=`git rev-parse --abbrev-ref HEAD` ||exit  # gitのブランチ名をバージョン名にセット
  VERSION=`echo $VERSION | sed "s/.*\///g"` # フォルダ(feature/)を削除する
  VERSION=`echo $VERSION | sed "s/#*_*//g"` # VERSION名で禁止文字(#と_)を削除する
fi

echo $VERSION