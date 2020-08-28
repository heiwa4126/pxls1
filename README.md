# pxls1

あるパスに置かれた
複数のJSONファイルを解析して、
Excelファイルを作るツール

(どんなJSONかは `testdata/1` の下を見てください)


# もくじ

- [pxls1](#pxls1)
- [もくじ](#もくじ)
- [インストール](#インストール)
- [使い方](#使い方)
- [参考](#参考)


# インストール

releaseから`pxls1`をダウンロードして、適当な場所に置く。

go 1.14以上がインストールされていれば
```sh
go install github.com/heiwa4126/pxls1
```
でもOK

# 使い方

```
pxls1 [-h|-v] <JSONファイルのあるパス> <出力するExcelファイル>
```

# 参考

[はじめに · Excelize ドキュメンテーション](https://xuri.me/excelize/ja/)
