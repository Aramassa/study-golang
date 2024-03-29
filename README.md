# study-golang

## 学習ページ

### サンプルで学ぶ Go 言語
https://www.spinute.org/go-by-example/

### Go Codereview Comments
https://knsh14.github.io/translations/go-codereview-comments/#go-fmt

## 0-1 Hello World

`Hello World` を print する

## 0-2 AWS SDK Access

AWSのS3バケットから Object のリストを取得して表示する(Pagerはなし)

* build
* dockerイメージ作成

まで実施する

## 0-3 AWS SDK Access 2

https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/costexplorer#Client.GetCostAndUsage

↑のAPIを利用して、AWSの利用料を取得する

## 1-1 クラスの定義

Hoge という最もシンプルなクラスの実装

## 1-2 クラスの継承

Hoge を継承した Hoge2 というクラスを作る

## 2-1 プロジェクト構造作って使ってみよう

以下の構造を作って、2-1 の中から使ってみる

* /cmd
* /internal
* /pkg

ref: https://github.com/Aramassa/zenn_articles/wiki#golang

## 2-2 リモートの pkg を使ってみよう

2-1 の /pkg で公開したモジュールを、 2-2 ディレクトリから使ってみる

`GOMODCACHE` ディレクトリに `go get` したキャッシュが残る。一度DLしたブランチのコードは、キャッシュが優先して利用されるので注意が必要。
このキャッシュをクリアするには、 `go clean -modcache` フラグをつけて clean する必要がある（ GOMODCACHE がディレクトリごと削除される）

## 3-1 test を書いてみよう

## 4-1 goroutine を書いてみよう

## 4-1 goroutine + channel で処理をしてみよう

## 5-1 WEBアプリケーションを作ってみよう

ドキュメントルートにアクセスしたら `Hello World!!` を表示するWEBサイトを作る

* WEBアプリケーションとしてHTTPで待ち受け

## 5-2 WEBアプリケーションを作ってAPIを公開してみよう

* OpenAPIを使った RestAPIの構築

## 5-3 gRPC サイトを作ろう

## 6-1 DockerContainer を作ってみよう

## 6-2 k8s で Podを構築してみよう