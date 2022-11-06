# 概要

5-1 で作った WEBアプリケーションを `go build` した `main` を持ってきて、Dockerイメージにする

## 覚書

scratch でビルドする場合、ライブラリに動的リンクしているアプリケーションは、dockerイメージ側にもそのライブラリがないと動かない。

もしくは環境変数 CGO_ENABLED=0 を設定する必要がある。
(このライブラリを指定した場合は、クロスコンパイルはできなくなるので注意)

* 参考記事: [dockerのscratchイメージでgolangのWebアプリを動かす](https://qiita.com/4486/items/915e9e4c77f012c365ff)
