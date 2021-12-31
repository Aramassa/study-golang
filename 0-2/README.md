
# AWS SDK for go を使って AWSにアクセスしてみよう

https://aws.amazon.com/jp/sdk-for-go/

## Getting Started

https://aws.github.io/aws-sdk-go-v2/docs/getting-started/

## SDK API References

https://pkg.go.dev/github.com/aws/aws-sdk-go-v2

# build してみよう

```
go build
```

`example` というバイナリが出力された。

→ go.mod の module名がバイナリファイル名になるらしい

## その他のオプション

* -o : 出力ファイル名を指定する

とりあえず `-o` 以外は必要に応じて使えるようになる、でよさそう

# docker イメージにしてみよう

## マルチステージビルド

https://docs.docker.com/language/golang/build-images/#multi-stage-builds