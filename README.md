# pxls1

あるパスに置かれた
複数のJSONファイルを解析して、
Excelファイルを作るツール

(どんなJSONかは `test/7` の下を見てください)


# もくじ

- [pxls1](#pxls1)
- [もくじ](#もくじ)
- [インストール](#インストール)
- [使い方](#使い方)
- [なんでGoで書いてあるの?](#なんでgoで書いてあるの)
- [参考](#参考)
- [TODO](#todo)


# インストール

ここの[最新リリース](https://github.com/heiwa4126/pxls1/releases)から
`pxls1_linux_x86_64.tar.gz`をダウンロードして、展開してください。

go 1.14以上がインストールされていれば
```sh
GO111MODULE=on go get -u github.com/heiwa4126/pxls1
```
でもOK(-vオプションのバージョン表記がおかしくなるけど)

# 使い方

```
pxls1 <JSONファイルのあるパス> <出力するExcelファイル> -r <rpmqa_csvのパス>
pxls1 -y <JSONファイルのあるパス> <出力するYAMLファイル>
pxls1 [-h|-v]
```

flags:
- -h    ヘルプの表示
- -v    バージョンの表示
- -y    YAMLモード
- -r    rpmqa_csvのパス


# 追加機能

## 2021-11

Excelの出力に、以前のバージョンがほしいというので追加した。
`-r`オプションで`ap rpmqa_csv`の出力パスを指定してください。



# なんでGoで書いてあるの?

PythonやPerlでいいようなコードだけど、
ライブラリとか入れる手間がめんどうだったから。
Goなら実行ファイル1個ですむ。デプロイ簡単。


# 参考

- [はじめに · Excelize ドキュメンテーション](https://xuri.me/excelize/ja/)


# TODO

- Quick fixで追加したreadjson7のユニットテストが抜けてるので、書く。(`go test ./... -v`でデバッグ用のコードしか無い)
- Rustで書き直す。
