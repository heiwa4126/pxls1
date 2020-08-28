# pxls1

あるパスに置かれたJSONファイルを解析して
Excelファイルを作るツール

(どんなJSONかは `testdata/1` の下を見てください)


# もくじ


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
