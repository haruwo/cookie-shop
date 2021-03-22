## 注文

このディレクトリ（/orders）へ以下のルールに従ってファイルを追加することにより注文が完了します。
注文を行う際は注文内容を記載したファイルを追加して Pull Request を送り、マージされることにより注文が完了します。
Pull Request 作成後は CircleCI の実行を確認して下さい。

送料、商品については [CONTRIBUTING.md](../CONTRIBUTING.md) を参照してください。

### ファイル名

`yyyy-mm-dd-nnn_{username}`

例: `2021-04-01-001_haruwo`

### 中身

yaml 形式で以下の情報を登録します

```
version: 1
username: {GitHub username}
items:
  - id: {id} # ex: cookie-random-4pieces
    amount: 3 # 注文個数（3 の場合、4 pieces が 3袋、合計12個のクッキーが届きます）
```
