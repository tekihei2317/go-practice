# simple-bbs

以下の記事のコードを写経して練習してみます。

[ゼロからはじめるGo言語(9) 100行未満かつGo標準ライブラリだけで作る掲示板 | TECH+（テックプラス）](https://news.mynavi.jp/techplus/article/gogogo-9/)


```bash
go run bbs.go
```

## 要件

- 名前と本文を書いて、掲示板に書き込める
- 掲示板への書き込みの一覧を確認できる

## 気になったこと

### package mainでVSCode上でエラーになっているのを直す

goplsはGoのLanguage Serverです。

```text
gopls was not able to find modules in your workspace.
When outside of GOPATH, gopls needs to know which modules you are working on.
```

とりあえず`go mod init simple-bbs`すると直りました。`simple-bbs`のところは、公開するモジュールであればGitHubのパスなどを指定しますが、そうでない場合はなんでもいいようです。

Go Modulesというのがあるみたいなので、それについて調べてみようと思います。以前のGoはモジュールをグローバルの場所にインストールしていましたが、Go Modulesを使うとプロジェクトローカルにインストールするようになったみたいな認識をしています（`node_modules`みたいな感じ）。


### ホットリロードしたい

コードを書き換えるごとに`go run`を実行しているので、ホットリロードできるようにしたいです。

## メモ

### 書き込みを表示できるようにする

まずは、書き込みのデータを読み込んで掲示板に表示できるようにします。VSCodeでGo Modulesのエラーが出ていますが、よく分からないのでとりあえず先に進みます。

基本的なデータ型について。

```go
// スライスを作る構文だったはず。
return make([]Log, 0)

// byteはどういう型か、json.Unmarshalは何か
json.Unmarshal([]byte(text), &logs)
```

`fmt.Printf`の書式指定について。"%v"で値を表示、〜でGoのリテラル表記で表示、みたいな感じ。

これ、YYYY/MM/DD HH:mmとかじゃなないのが気になった。

```go
time.Unix(log.CTime, 0).Format("2006/1/2 15:04")
```

> 引数layoutに渡す時刻は、「2006年1月2日15時4分5秒 アメリカ山地標準時MST(GMT-0700)」のものを使うことになっています。

https://zenn.dev/hsaki/articles/go-time-cheatsheet

らしいです。とりあえず書き込みを表示することができました。

### 書き込みを実装する

続いて、書き込みを投稿できるように実装します。

ifが式ではなく文なので、値を書くところに書けないことがわかりました。なので、変数を宣言してif文を書きました。

```go
name := r.Form["name"][0];
if name == "" {
	name = "名無し"
}
```

[Goに三項演算子が採用されない理由](https://zenn.dev/nobonobo/articles/09d884f1f520d6)
