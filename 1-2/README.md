# golang における継承の難しさ

## Go言語で「embedded で継承ができる」と思わないほうがいいのはなぜか？
https://qiita.com/Maki-Daisuke/items/511b8989e528f7c70f80

> 本来「継承」は不必要なところに「継承」を使ってしまうと事故ることが多いので、実装の再利用には委譲を、多相性だけが必要なところにはインターフェイスを使おう、そして本当に必要なときにだけ継承を使おうね、というのがベストプラクティスとなっていったわけです。

オブジェクト志向のコード再利用の目的は、委譲だけで語れるものでもないような気はするが、 golang の継承（もどき）は癖がありすぎるため、このような棲み分けにした方が結果ラクできそうな感じはする

## オブジェクトの初期化

Java で お馴染みの new メソッドは、 golang では全く異なった定義になっているので注意

https://qiita.com/S-Masakatsu/items/6fb8e765cd443e2edd7f

```golang
type Human struct {
    name string
    age  int
}

func main() {
    // 1. {}を使い、順序にしたがって初期化する。
    sum := Human{"Sum", 22}
    // 2. field:valueの方法で初期化する。順序は任意。
    tom := Human{age: 25, name: "Tom"}
    // 3. new関数を使い、変数定義後に初期化する。
    bob := new(Human)
    bob.name, bob.age = "Bob", 28
```

↓匿名フィールドを使うことで、フィールドを継承することは可能である
https://qiita.com/S-Masakatsu/items/6fb8e765cd443e2edd7f#%E5%8C%BF%E5%90%8D%E3%83%95%E3%82%A3%E3%83%BC%E3%83%AB%E3%83%89

