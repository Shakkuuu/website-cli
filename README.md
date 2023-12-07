# website-cli

## 仕様

　CLI上で以下のコマンドを実行することで使用できる。

```:
website-cli [-search 検索したい用語][site サイト名][url URL] 
```

`-search string`

　searchフラグでは、検索したい用語を入れることで、Googleでの検索結果が開かれる。

　スペースを入れて複数用語を検索したいような場合は、クオーテーション`'',""`で囲むか、用語を`+`で繋ぐ必要がある。

`-site string`

　siteフラグでは、あらかじめwebsite-cli-setting.jsonファイルで登録したサイト名を入力することで、そのサイト名に対応したサイトが開かれる。

`-url string`

　urlフラグでは、開きたいWebサイトの[https://](https://)または[http://](http://)から始まるURLを入力することでそのURLのサイトが開かれる(タイプミスに注意)

　コマンドを実行するとデフォルトのブラウザで開かれる。

　フラグを複数選択した場合は、site,url,searchの順で全て開かれる。

## 対応OS

対応をしているOSは以下の通り。

- darwin：amd64,arm64
- linux：386,amd64,arm,arm64
- windows：386,amd64,arm,arm64

(対応しているはず...)

## 使用方法

　このURL[https://github.com/Shakkuuu/website-cli/releases](https://github.com/Shakkuuu/website-cli/releases)から最新バージョンをダウンロードする。

　ダウンロードしたzipファイルを展開し、適切な場所に配置する。(実行ファイルをMacとかであれば、`/usr/local/bin/`とかに置く)

　`website-cli-setting.json`というファイルを、実行ファイルと同じディレクトリに配置する。
　記述方法は以下の通り。

```:json
[
    {
        "url": "https://twitter.com",
        "sitename": [
            "twitter",
            "Twitter",
            "x",
            "X",
            "ツイッター",
            "ついったー"
        ]
    },
    {
        "url": "https://www.youtube.com",
        "sitename": [
            "youtube",
            "Youtube",
            "ようつべ",
            "ユーチューブ"
        ]
    }
]

```

　配列のjsonで、`url`と`sitename`の項目が必要。
　`sitename`はstringの配列となっている。
　`url`には、そのサイトのURL、`sitename`には、コマンドを使用する際に実行できるようにしたいサイト名を設定する。

`website-cli`で実行!
